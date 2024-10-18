package helpers

func NillableString(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}
