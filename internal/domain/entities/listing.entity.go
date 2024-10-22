package entities

type Listing struct {
	Page     int
	Pagesize int
	Filters  []Filter
}
