package helper

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ToIntFromSlice(target []interface{}, index int) int {
	pos := Min(index, len(target))
	num, _ := ToInt(target[pos])
	return num
}

func ToString(target interface{}) string {
	if target == nil {
		return ""
	}
	return fmt.Sprintf("%v", target)
}

func FormatNumber(target interface{}) (string, error) {
	if number, ok := target.(float64); ok {
		return strconv.FormatFloat(number, 'f', -1, 64), nil
	}

	if number, ok := target.(int64); ok {
		return strconv.FormatInt(number, 10), nil
	}

	v, err := ToInt(target)
	return fmt.Sprintf("%d", v), err
}

func ToSafeInt(value interface{}) int {
	v, _ := ToInt(value)
	return v
}

func ToInt(value interface{}) (int, error) {
	if value == nil {
		return 0, nil
	}

	i, err := strconv.Atoi(fmt.Sprintf("%v", value))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func LeftPad(s string, padStr string, overallLen int) string {
	padCountInt := 1 + ((overallLen - len(padStr)) / len(padStr))
	retStr := strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func RightPad(s string, padStr string, overallLen int) string {
	padCountInt := 1 + ((overallLen - len(padStr)) / len(padStr))
	retStr := s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}
