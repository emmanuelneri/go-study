package structs

// Order struct
type Order struct {
	ID       int
	Customer string
	Total    float32
	Items    []Item
}
