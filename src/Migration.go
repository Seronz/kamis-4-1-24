package src

import (
	products "github.com/seronz/api/src/models/Products"
	users "github.com/seronz/api/src/models/Users"
)

type Models struct {
	Models interface{}
}

func RegistryModels() []Models {
	return []Models{
		{Models: users.User{}},
		{Models: users.Address{}},
		{Models: products.Products{}},
		{Models: products.Image{}},
		{Models: products.Section{}},
		{Models: products.Categories{}},
	}
}
