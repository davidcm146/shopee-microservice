package models

import "time"

type Product struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Name        string    `bson:"name" json:"name"`
	SellerID    string    `bson:"sellerID" json:"sellerID"`
	Category    string    `bson:"category" json:"category"`
	Description string    `bson:"description" json:"description"`
	Price       float64   `bson:"price" json:"price"`
	Quantity    int       `bson:"quantity" json:"quantity"`
	Features    []string  `bson:"features" json:"features"`
	IsDeleted   bool      `bson:"isDeleted" json:"isDeleted"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt" json:"updatedAt"`
}

func NewProduct(name, sellerID, category, description string, price float64, quantity int, features []string) *Product {
	return &Product{
		Name:        name,
		SellerID:    sellerID,
		Category:    category,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		Features:    features,
		IsDeleted:   false,
	}
}
