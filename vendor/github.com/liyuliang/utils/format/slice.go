package format

func SliceDelDuplicate(slice []string) []string {
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

func InArray(items []string, item string) bool {
	for _, a := range items {
		if a == item {
			return true
		}
	}
	return false
}