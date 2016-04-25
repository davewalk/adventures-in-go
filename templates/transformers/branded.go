package transformers

import (
	"fmt"
	"strings"
)

func Branded(title string, brand string) string {
	s := strings.TrimSpace(strings.Replace(title, brand, "", -1))
	s = fmt.Sprintf("%s - %s", brand, s)
	return s
}
