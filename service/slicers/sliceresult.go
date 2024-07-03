package slicers

type SliceResult struct {
	Price float64    `json:"price"`
	Files []Metadata `json:"files"`
}
