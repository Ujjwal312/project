package models

type User struct {
	Email string `json:"email" bson:"email"`
	Name  string `json:"name" bson:"name"`
	Age   int    `json:"age" bson:"age"`
}

type Product struct {
	Name  string `json:"name" bson:"product_name"`
	Price int    `json:"price" bson:"price"`
}
