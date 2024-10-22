package helpers

func NilIfEmptyString(str string) *string {
	if str == "" {
		return nil
	} else {
		return &str
	}
}
