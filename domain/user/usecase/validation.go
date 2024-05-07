package usecase

import (
	"regexp"
	"strings"
	"unicode"
	"user-service/domain/user/model"
	"user-service/lib/constant"
)

func isDataValid(data model.Users) (invalidMessages []string, isValid bool) {

	if len(data.FullName) < 3 || len(data.FullName) > 60 {
		invalidMessages = append(invalidMessages, constant.FullNameCharLength)
	}

	if len(data.PhoneNumber) < 10 || len(data.PhoneNumber) > 13 {
		invalidMessages = append(invalidMessages, constant.PhoneNumberCharLength)
	}

	if !strings.HasPrefix(data.PhoneNumber, "+62") {
		invalidMessages = append(invalidMessages, constant.PhoneNumberIndonesian)
	}

	if len(data.UserPassword) < 6 || len(data.UserPassword) > 64 {
		invalidMessages = append(invalidMessages, constant.PassWordCharLength)
	}

	if len(data.Email) == 0 {
		invalidMessages = append(invalidMessages, constant.EmailRequired)
	}

	if !isValidEmail(data.Email) {
		invalidMessages = append(invalidMessages, constant.EmailInvalidAdress)
	}

	if !isValidPasswordChar(data.UserPassword) {
		invalidMessages = append(invalidMessages, constant.PasswordReqChar)
	}

	if len(invalidMessages) > 0 {
		return invalidMessages, false
	} else {
		return []string{""}, true
	}
}

func isDataValidForUpdate(data model.Users) (invalidMessages []string, isValid bool) {

	if len(data.FullName) < 3 || len(data.FullName) > 60 {
		invalidMessages = append(invalidMessages, constant.FullNameCharLength)
	}

	if len(data.PhoneNumber) < 10 || len(data.PhoneNumber) > 13 {
		invalidMessages = append(invalidMessages, constant.PhoneNumberCharLength)
	}

	if !strings.HasPrefix(data.PhoneNumber, "+62") {
		invalidMessages = append(invalidMessages, constant.PhoneNumberIndonesian)
	}

	if len(data.UserPassword) < 6 || len(data.UserPassword) > 64 {
		invalidMessages = append(invalidMessages, constant.PassWordCharLength)
	}

	if !isValidPasswordChar(data.UserPassword) {
		invalidMessages = append(invalidMessages, constant.PasswordReqChar)
	}

	if len(invalidMessages) > 0 {
		return invalidMessages, false
	} else {
		return []string{""}, true
	}
}

func isLoginValid(data model.Users) (invalidMessages []string, isValid bool) {
	if len(data.Email) == 0 {
		invalidMessages = append(invalidMessages, constant.EmailRequired)
	}

	if !isValidEmail(data.Email) {
		invalidMessages = append(invalidMessages, constant.EmailInvalidAdress)
	}

	if !isValidPasswordChar(data.UserPassword) {
		invalidMessages = append(invalidMessages, constant.PasswordReqChar)
	}

	if len(invalidMessages) > 0 {
		return invalidMessages, false
	} else {
		return []string{""}, true
	}

}

func isValidPasswordChar(s string) bool {
	var hasUpperCase, hasNumber, hasSpecial bool

	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpperCase = true
		case unicode.IsNumber(char):
			hasNumber = true
		case !unicode.IsLetter(char) && !unicode.IsNumber(char):
			hasSpecial = true
		}

		if hasUpperCase && hasNumber && hasSpecial {
			return true
		}
	}

	return false
}

func isValidEmail(email string) bool {
	// Regular expression for basic email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the pattern
	regex := regexp.MustCompile(pattern)

	// Match the email against the pattern
	return regex.MatchString(email)
}
