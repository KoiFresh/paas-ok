package materials

type Material struct {
	Id      int     `json:"id"`
	Price   float64 `json:"price"`
	Density float64 `json:"density"`
	Color   string  `json:"color"`
	Type    string  `json:"type"`
}
