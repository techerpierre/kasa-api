package helpers

func ArrayContains(arr []any, vals ...any) bool {
	for _, val := range vals {
		for _, el := range arr {
			if el == val {
				return true
			}
		}
	}
	return false
}
