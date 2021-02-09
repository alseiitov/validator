package validator

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func Validate(data interface{}) error {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldData := reflect.ValueOf(data).Field(i)

		flags, ok := field.Tag.Lookup("validator")
		if !ok {
			continue
		}

		flags = trimWhiteSpaces(flags)
		flagsArray := strings.Split(flags, ",")

		for _, flag := range flagsArray {
			keyValue := strings.Split(flag, "=")

			switch len(keyValue) {
			case 1:
				key := keyValue[0]
				switch key {
				case "required":
					if err := checkRequired(fieldData.Interface()); err != nil {
						return fmt.Errorf("%v %v", field.Name, err.Error())
					}
				case "password":
					if err := checkPassword(fieldData.String()); err != nil {
						return fmt.Errorf("%v %v", field.Name, err.Error())
					}
				case "email":
					if err := checkEmail(fieldData.String()); err != nil {
						return fmt.Errorf("%v", err.Error())
					}
				case "username":
					if err := checkUsername(fieldData.String()); err != nil {
						return fmt.Errorf("%v", err.Error())
					}
				default:
					log.Fatalf("'%v' invalid key at %v", key, t)
				}

			case 2:
				key, value := keyValue[0], keyValue[1]
				switch key {
				case "min":
					if err := checkMin(fieldData.Interface(), value); err != nil {
						return fmt.Errorf("%v %v", field.Name, err.Error())
					}
				case "max":
					if err := checkMax(fieldData.Interface(), value); err != nil {
						return fmt.Errorf("%v %v", field.Name, err.Error())
					}
				default:
					return fmt.Errorf("%v invalid key", key)
				}
			default:
				log.Fatalf("'%v' invalid format at %v", flag, t)
			}
		}
	}
	return nil
}

func checkRequired(data interface{}) error {
	switch v := data.(type) {
	case int:
		if v == 0 {
			return errors.New("is required")
		}
	case string:
		if len(strings.TrimSpace(v)) == 0 {
			return errors.New("is required")
		}
	}
	return nil
}

func checkPassword(pass string) error {
	numRegex := regexp.MustCompile(`[0-9]{1}`)
	lowercaseRegex := regexp.MustCompile(`[a-z]{1}`)
	uppercaseRegex := regexp.MustCompile(`[A-Z]{1}`)
	symbolRegex := regexp.MustCompile(`[!@#~$%^&*()+|_]{1}`)

	if !numRegex.MatchString(pass) {
		return errors.New("must contain at least one number")
	}
	if !lowercaseRegex.MatchString(pass) {
		return errors.New("must contain at least one lowercase letter")
	}
	if !uppercaseRegex.MatchString(pass) {
		return errors.New("must contain at least one uppercase letter")
	}
	if !symbolRegex.MatchString(pass) {
		return errors.New("must contain at least one symbol\n(!, @, #, ~, $, %, ^, &, *, (, ), +, |, _, )")
	}
	return nil
}

func checkEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,24}$`)

	if !emailRegex.MatchString(email) {
		return errors.New("e-mail is invalid")
	}
	return nil
}

func checkUsername(username string) error {
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)

	if !usernameRegex.MatchString(username) {
		return errors.New("username is invalid")
	}
	return nil
}
func checkMin(data interface{}, minStr string) error {
	min, err := parseInt(minStr)
	if err != nil {
		return err
	}

	switch v := data.(type) {
	case int:
		if v < min {
			return fmt.Errorf("value (%v) is lower than minimum value (%v)", v, min)
		}
	case string:
		l := len(v)
		if l < min {
			return fmt.Errorf("length (%v) is lower than minimum length (%v)", l, min)
		}
	}
	return nil
}

func checkMax(data interface{}, maxStr string) error {
	max, err := parseInt(maxStr)
	if err != nil {
		return err
	}

	switch v := data.(type) {
	case int:
		if v > max {
			return fmt.Errorf("value (%v) is higher than maximum value (%v)", v, max)
		}
	case string:
		l := len(v)
		if l > max {
			return fmt.Errorf("length (%v) length is higher than maximim length (%v)", l, max)
		}
	}
	return nil
}

func parseInt(intStr string) (int, error) {
	n, err := strconv.Atoi(intStr)
	if err != nil {
		return n, fmt.Errorf("%v is not integer", intStr)
	}
	return n, nil
}

func trimWhiteSpaces(s string) string {
	whiteSpaces := []string{" ", "\t", "\v", "\n"}
	for _, w := range whiteSpaces {
		s = strings.ReplaceAll(s, w, "")
	}
	return s
}
