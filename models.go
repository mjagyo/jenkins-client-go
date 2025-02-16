package jenkins

type Jobs struct {
	Jobs []Job `json:"jobs"`
}

type Job struct {
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
	FullName    string `json:"fullName"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Buildable   bool   `json:"buildable"`
	Color       string `json:"color"`
	InQueue     bool   `json:"inQueue"`
}

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
	Color       string             `json:"color"`
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
