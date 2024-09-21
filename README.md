# localizations

[![GoDoc](https://godoc.org/github.com/touchvas/localizations?status.svg)](https://godoc.org/github.com/m1/go-localize)
[![Build Status](https://travis-ci.org/touchvas/localizations?branch=master)](https://travis-ci.org/m1/go-localize)
[![Release](https://img.shields.io/github/release/m1/go-localize.svg)](https://github.com/touchvas/localizations/releases/latest)

__Simple and easy to use i18n (Internationalization and localization) engine written in Go, used for translating locale strings. 
Use with [go generate](#go-generate) or on the [CLI](#cli). Currently supports JSON, YAML, TOML and CSV translation files__

## Get Library

```go
go get -u github.com/touchvas/localizations
```

## Why another i18n library?

This package aims to be as simple and easy to use as possible. It also
takes inspiration from popular localization libraries/packages in other languages - 
so makes it easier to reason about coming from other languages and frameworks.

## Usage

### Go generate

The suggested way to use go-localize is to use `go generate`. For example, take the following directory structure:

```
goapp
└── localizations_src
    ├── en
    │   └── messages.yaml
    └── es
        ├── customer
        │   └── messages.json
        └── messages.json
```

Example of JSON translation file:

```json
{
  "hello": "Hola",
  "how_are_you": "¿Cómo estás?",
  "whats_your_name": "¿Cuál es tu nombre?",
  "hello_my_name_is": "Hola, mi nombre es {{.name}}"
}
```

Example of YAML translation file:
```yaml
hello: hello
how_are_you: How are you?
whats_your_name: "What's your name?"
hello_my_name_is: Hello my name is {{.name}}
hello_firstname_lastname: Hello {{.firstname}} {{.lastname}}
```

Example of CSV translation file:
```csv
hello, hello
how_are_you, How are you?
```

Example of TOML translation file:
```toml
hello = "hello"
how_are_you = "How are you?"
```

Now you'll be able to use the localization like so:
```go
localizer := Localizer{}
localizer.SetSource("localizations_src")
localizer.SetDefaultLocale("en", "en")
localizer.SetupLocalization()

println(localizer.Translate("response.messages.verification_code_limit_reached")) // Wait for a few minutes to resend verification code again

```

With `en` being the locale and `es` being the fallback. The localization keys are worked out using folder structure, eg:

`en/customer/messages.json` with the contents being:
```json
{
  "hello": "hello customer!"
}
```
You'll be able to access this using the key: `customer.messages.hello`.

#### Suggestions

It is suggested to instead of using hardcoded locale keys i.e. `en` to use the language keys included in key, i.e: `language.BritishEnglish.String()` 
which is `en-GB`


#### Replacements

Take this replacement string for example:
```yaml
password_reset_code: "Your Account verification code is {{.code}}"
```

To then replace `code`  variable, you can use 
something like this:


```go
localizer := Localizer{}
localizer.SetSource("localizations_src")
localizer.SetDefaultLocale("en", "en")
localizer.SetupLocalization()

replacements := map[string]string{
			"code": "1234",
		}

println(localizer.Translate("password_reset_code", replacements))
```


#### Locale defining and localization fallbacks

You can define the locale and fallbacks using:
```go
		localizer.SetDefaultLocale("en", "en")
```

Where `en` is the locale and `es` is the fallback. If no translation key-value is
found then the key will be returned. For example

```go
println(l.Get("key_doesnt_exist")) //"key_doesnt_exist" will be printed
```

#### Translation file support

We currently support JSON and YAML translation files. Please suggest
missing file type using issues or pull requests.
