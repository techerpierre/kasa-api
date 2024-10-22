package helpers

func NilIfEmptyInt(nbr int) *int {
	if nbr == 0 {
		return nil
	} else {
		return &nbr
	}
}
