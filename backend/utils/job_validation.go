package utils

// Helper functions for enum validation
func IsValidJobType(jobType string) bool {
	validTypes := []string{"full_time", "part_time", "contract", "internship"}
	for _, t := range validTypes {
		if t == jobType {
			return true
		}
	}
	return false
}

func IsValidCategory(category string) bool {
	validCategories := []string{"construction", "repairing", "house_keeping", "management"}
	for _, c := range validCategories {
		if c == category {
			return true
		}
	}
	return false
}