package hashicups

// Order -
type Order struct {
	ID    int         `json:"id,omitempty"`
	Items []OrderItem `json:"items,omitempty"`
}

// OrderItem -
type OrderItem struct {
	Coffee   Coffee `json:"coffee"`
	Quantity int    `json:"quantity"`
}

// Coffee -
type Coffee struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Teaser      string             `json:"teaser"`
	Collection  string             `json:"collection"`
	Origin      string             `json:"origin"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Image       string             `json:"image"`
	Ingredient  []CoffeeIngredient `json:"ingredients"`
}

// Ingredient -
type CoffeeIngredient struct {
	ID       int    `json:"ingredient_id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

// Ingredient -
type Ingredient struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}
