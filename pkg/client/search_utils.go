package client

import (
	"fmt"
	"strings"
)

func BuildSearchString(keywords []string) string {
	var builder strings.Builder
	if i := len(keywords); i == 1 {
		builder.WriteString(keywords[0])
	} else {
		for i := 0; i < len(keywords); i++ {
			if i != (len(keywords) - 1) {
				builder.WriteString(fmt.Sprintf("%s ", keywords[i]))
			} else {
				builder.WriteString(keywords[i])
			}
		}
	}
	return builder.String()
}
