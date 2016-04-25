package transformers

import "strings"

func SliceToString(c interface{}) string {
	switch val := c.(type) {
	case string:
		return val
	case []string:
		return strings.Join(val, ", ")
	}

	return ""
}
