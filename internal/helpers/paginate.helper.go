package helpers

import "math"

func Paginate(page int, pagesize int, count int) (int, int) {
	maxPage := math.Ceil(float64(count) / float64(pagesize))
	clampedPage := Clamp(page-1, int(maxPage), 0)
	return clampedPage * pagesize, pagesize
}
