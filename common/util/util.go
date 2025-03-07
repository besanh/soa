package util

import (
	"strconv"
	"strings"
)

func GetLimitOffset(limitQuery, offsetQuery string) (int, int) {
	limit := 0
	offset := 0
	limit = ParseLimit(limitQuery)
	offset = ParseOffset(offsetQuery)
	return limit, offset
}

func ParseOffset(offset string) int {
	offset = ParseString(offset)
	i, err := strconv.Atoi(offset)
	if err != nil {
		i = 0
	}
	return i
}

func ParseLimit(limit string) int {
	limit = ParseString(limit)
	i, err := strconv.Atoi(limit)
	if err != nil {
		i = 10
	}
	return i
}

func ParseString(value any) string {
	str, ok := value.(string)
	if !ok {
		return str
	}
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Trim(str, "\r\n")
	str = strings.TrimSpace(str)
	return str
}

func ParseQueryArray(slice []string) []string {
	result := make([]string, 0)
	for _, v := range slice {
		if len(v) > 0 {
			result = append(result, v)
		}
	}
	return result
}
