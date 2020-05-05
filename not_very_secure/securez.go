package not_very_secure

import (
	"regexp"
	"strings"
)

func alphanumeric(str string) bool {
	sourceArray := strings.Split(str, "")
	re := regexp.MustCompile("0-9A-Za-z_")

	for _, v := range sourceArray {
		if !re.MatchString(sourceArray) {
			return false
		}
	}

	return true
}
