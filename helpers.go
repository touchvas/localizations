package localizations

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	jsonFileExt = ".json"
	yamlFileExt = ".yaml"
	ymlFileExt  = ".yml"
	tomlFileExt = ".toml"
	csvFileExt  = ".csv"
)

type localizationFile map[string]string

func getLocalizationFiles(dir string) ([]string, error) {

	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		ext := filepath.Ext(path)

		if !info.IsDir() && (ext == jsonFileExt || ext == yamlFileExt) {

			files = append(files, path)

		}

		return nil

	})

	return files, err
}

func generateLocalizations(files []string, srcFolder string) (map[string]string, error) {

	localizations := map[string]string{}

	for _, file := range files {

		newLocalizations, err := getLocalizationsFromFile(file, srcFolder)
		if err != nil {
			return nil, err
		}
		for key, value := range newLocalizations {
			localizations[key] = value
		}

	}

	return localizations, nil
}

func getLocalizationsFromFile(file, srcFolder string) (map[string]string, error) {

	newLocalizations := map[string]string{}

	openFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(openFile)
	if err != nil {
		return nil, err
	}

	localizationFile := localizationFile{}

	ext := filepath.Ext(file)
	switch ext {
	case jsonFileExt:
		err = json.Unmarshal(byteValue, &localizationFile)
	case yamlFileExt, ymlFileExt:
		err = yaml.Unmarshal(byteValue, &localizationFile)
	case tomlFileExt:
		_, err = toml.Decode(string(byteValue), &localizationFile)
	case csvFileExt:
		err = parseCSV(byteValue, &localizationFile)
	default:
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	slicePath := getSlicePath(file, srcFolder)

	for key, value := range localizationFile {

		newLocalizations[strings.Join(append(slicePath, key), ".")] = value

	}

	return newLocalizations, nil
}

func parseCSV(value []byte, l *localizationFile) error {

	r := csv.NewReader(bytes.NewReader(value))
	localizations := localizationFile{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		localizations[record[0]] = record[1]
	}
	*l = localizations
	return nil
}

func getSlicePath(file string, srcFolder string) []string {

	dir, file := filepath.Split(file)

	paths := strings.Replace(dir, srcFolder, "", -1)

	pathSlice := strings.Split(paths, string(filepath.Separator))

	var strs []string
	for _, part := range pathSlice {
		part := strings.TrimSpace(part)
		part = strings.Trim(part, "/")
		if part != "" {
			strs = append(strs, part)
		}
	}

	strs = append(strs, strings.Replace(file, filepath.Ext(file), "", -1))
	return strs
}

func getAllFiles(directory string) []string {

	var files []string

	err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {

		if err != nil {

			log.Printf("failed to get transaction path %s ", err.Error())
			return err
		}

		info, err := d.Info()
		if err == nil {

			if info.Size() > 0 {

				ext := filepath.Ext(path)

				if !info.IsDir() && (ext == jsonFileExt || ext == yamlFileExt || ext == ymlFileExt || ext == csvFileExt) {

					log.Printf("got translation file at %s | %d", path, info.Size())
					files = append(files, path)

				}

			}

		}

		return nil
	})
	if err != nil {

		log.Println(err)

	}

	return files
}
