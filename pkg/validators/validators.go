package validators

import "regexp"

var (
	phoneRegexp    = regexp.MustCompile(`^\+\d{11}$`)
	fullNameRegexp = regexp.MustCompile(`^[a-zA-Zа-яА-Яё]+ [a-zA-Zа-яА-Яё]+ [a-zA-Zа-яА-Яё]+$`)
)

func IsValidPhone(phone string) bool {
	return phoneRegexp.MatchString(phone)
}

func IsValidFullName(fullName string) bool {
	return fullNameRegexp.MatchString(fullName)
}
