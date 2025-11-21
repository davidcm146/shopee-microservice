package dto

type CreateProductInput struct {
	Name        string   `json:"name" validate:"required"`
	SellerID    string   `json:"sellerID" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Price       float64  `json:"price" validate:"required,gt=0"`
	Quantity    int      `json:"quantity" validate:"required,gte=0"`
	Category    string   `json:"category" validate:"required"`
	Features    []string `json:"features"`
}

type UpdateProductInput struct {
	ID          string    `json:"id" validate:"required"`
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Price       *float64  `json:"price" validate:"omitempty,gt=0"`
	Quantity    *int      `json:"quantity" validate:"omitempty,gte=0"`
	Category    *string   `json:"category"`
	Features    *[]string `json:"features"`
}

type ProductResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	SellerID    string   `json:"sellerID"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Quantity    int      `json:"quantity"`
	Features    []string `json:"features"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}
