package models

type User struct {
	Name 	string 	`json:"name" bson:"name"`
	Age 	int 	`json:"age" bson:"age"`
	Address Address
}

type Address struct {
	State 	string	`json:"state" bson:"state"`
	City 	string	`json:"city" bson:"city"`
	Pincode int		`json:"pincode" bson:"pincode"`
}