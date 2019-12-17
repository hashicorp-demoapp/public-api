package models

// User describes a user.
type User struct {
	ID string `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
}