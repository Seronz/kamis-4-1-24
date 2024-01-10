package src

import (
	order "github.com/seronz/api/src/models/Order"
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
		{Models: order.Order{}},
		{Models: order.Items{}},
		{Models: order.Customer{}},
		{Models: order.Payment{}},
		{Models: order.Shiment{}},
	}
}
