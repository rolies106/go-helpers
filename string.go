package gohelpers

import (
	"regexp"
	"strconv"
	"strings"
)

type String struct{}

/**
 * Convert kebab-case into CamelCase
 */
func (this *String) KebabToCamelCase(kebab string) (camelCase string) {

	isToUpper := true
	for _, runeValue := range kebab {
		if isToUpper {
			camelCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '-' {
				isToUpper = true
			} else {
				camelCase += string(runeValue)
			}
		}
	}
	return
}

/**
 * Extract integer from string
 */
func (this *String) ExtractIntFromString(s string) (result int) {

	resultString := this.ExtractNumberFromString(s)

	if result, err := strconv.Atoi(resultString); err == nil {
		return result
	}

	result = 0

	return
}

func (this *String) ExtractNumberFromString(s string) (result string) {

	re := regexp.MustCompile("[0-9]+")
	array := re.FindAllString(s, -1)

	result = strings.Join(array, "")

	return
}
