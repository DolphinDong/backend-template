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
func RemoveDuplicateElement(slice []string) []string {
	result := make([]string, 0, len(slice))
	temp := map[string]struct{}{}
	for _, item := range slice {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
