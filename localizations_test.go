package localizations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetup(t1 *testing.T) {

	t1.Run("test setup", func(t *testing.T) {

		localizer := Localizer{}
		localizer.SetSource("localizations_src")
		localizer.SetDefaultLocale("en", "en")
		localizer.SetupLocalization()

		assert.Greater(t, len(localizer.Localizations), 0, "Assert translations were generated successfully")

	})

	t1.Run("test translation without replacements", func(t *testing.T) {

		localizer := Localizer{}
		localizer.SetSource("localizations_src")
		localizer.SetDefaultLocale("en", "en")
		localizer.SetupLocalization()

		res := localizer.Translate("response.messages.verification_code_limit_reached")

		assert.Equal(t, "Wait for a few minutes to resend verification code again", res, "Assert test was translated successfully")

	})

	t1.Run("test translation with replacements", func(t *testing.T) {

		localizer := Localizer{}
		localizer.SetSource("localizations_src")
		localizer.SetDefaultLocale("en", "en")
		localizer.SetupLocalization()

		replacements := map[string]string{
			"code": "1234",
		}

		res := localizer.Translate("response.account.messages.password_reset_code", replacements)

		assert.Equal(t, "Your Account reset code is 1234", res, "Assert test was translated successfully")

	})

	t1.Run("test translation with invalid key", func(t *testing.T) {

		localizer := Localizer{}
		localizer.SetSource("localizations_src")
		localizer.SetDefaultLocale("en", "en")
		localizer.SetupLocalization()

		replacements := map[string]string{
			"code": "1234",
		}

		res := localizer.Translate("response.messages.password_reset_code_invalid", replacements)

		assert.Equal(t, "response.messages.password_reset_code_invalid", res, "Assert missing key does not cause panic")

	})

	t1.Run("test translation with invalid key", func(t *testing.T) {

		localizer := Localizer{}
		localizer.SetSource("localizations_src")
		localizer.SetDefaultLocale("en", "en")
		localizer.SetupLocalization()

		res := localizer.Translate("response.messages.password_reset_code_invalid")

		assert.Equal(t, "response.messages.password_reset_code_invalid", res, "Assert missing key does not cause panic")

	})
}
