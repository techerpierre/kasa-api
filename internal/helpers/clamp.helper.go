package helpers

func Clamp(val int, max int, min int) int {
	switch true {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}
