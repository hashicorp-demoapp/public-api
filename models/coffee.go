package models

// Coffee describes a coffee.
type Coffee struct {
	ID string `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
	Price float64  `yaml:"price" json:"price"`
	Ingredients []Ingredient  `yaml:"ingredients" json:"ingredients"`
}