package entities

type Listing struct {
	Page     int64
	Pagesize int64
	Filters  []Filter
}
