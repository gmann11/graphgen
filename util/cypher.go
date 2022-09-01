package util

import (
	"fmt"
	"strings"
)

func MapToCypher(m map[string]interface{}) string {
	var sb strings.Builder
	for k, v := range m {
		sb.WriteString(k)
		sb.WriteString(":")
		sb.WriteString(fmt.Sprintf("%v", v))
		sb.WriteString(",")
	}

	return stripLastChar(sb.String())
}

// func SliceToCypher(s []string) string {
// 	var sb strings.Builder
// 	for _, v := range s {
// 		sb.WriteString(v)
// 	}
//
// 	return stripLastChar(sb.String())
// }

func stripLastChar(s string) string {
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}
