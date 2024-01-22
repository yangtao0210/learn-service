package algrithm

func replaceNumber(s string) string {
	replace := "number"
	replaceBytes := []byte(replace)
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= '0' && bytes[i] <= '9' {
			bytes = append(bytes[:i], append(replaceBytes, bytes[i+1:]...)...)
			i += len(replaceBytes) - 1
		}
	}
	return string(bytes)
}
