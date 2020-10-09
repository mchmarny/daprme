package format

import "regexp"

// CodeSafeString removes non-alpha characters
func CodeSafeString(val string) string {
	reg := regexp.MustCompile("[^a-zA-Z]+")
	return reg.ReplaceAllString(val, "")
}
