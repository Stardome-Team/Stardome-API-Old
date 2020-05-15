package utils

// CompareStrings :
func CompareStrings(first string, second string, err error) (bool, error) {
	if first != second {
		return true, nil
	}
	return false, err
}
