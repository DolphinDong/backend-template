package tools

// 判断元素是否在切片中
func ElementInSlice(element string, sli []string) bool {
	for _, str := range sli {
		if element == str {
			return true
		}
	}
	return false
}
