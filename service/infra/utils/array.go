package utils

func Filter(source []string, target []string) ([]string, []string) {
	included := make([]string, len(source))
	excluded := make([]string, len(source))
	includedIndex := 0
	excludedIndex := 0

	for _, content := range source {
		if !Contains(content, target) {
			excluded[excludedIndex] = content
			excludedIndex++
		} else {
			included[includedIndex] = content
			includedIndex++
		}
	}

	return included[0:includedIndex], excluded[0:excludedIndex]
}

func Contains(obj string, target []string) bool {
	for _, item := range target {
		if obj == item {
			return true
		}
	}
	return false
}
