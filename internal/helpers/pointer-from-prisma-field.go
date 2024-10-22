package helpers

func PointerFromPrismaField[T any](field T, filled bool) *T {
	if filled {
		return &field
	} else {
		return nil
	}
}
