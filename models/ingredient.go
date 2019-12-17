package models

// Ingredient describes a ingredient.
type Ingredient struct {
	ID string `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
	Quantity int64  `yaml:"quantity" json:"quantity"`
}