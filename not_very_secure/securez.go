package not_very_secure

import (
	"regexp"
	"strings"
)

func alphanumeric(str string) bool {
	sourceArray := strings.Split(str, "")
	re := regexp.MustCompile("([0-9]|[A-Z]|[a-z])+")

	if len(sourceArray) == 0 {
		return false
	}

	for _, v := range sourceArray {
		temp := re.MatchString(v)
		if !temp {
			return false
		}
	}

	return true
}
