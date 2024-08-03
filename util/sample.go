package util

func ConvertToStringToInt64(str string) int64 {
	var result int64
	for _, c := range str {
		result = result*10 + int64(c-'0')
	}
	return result
}
