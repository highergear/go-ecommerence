package input

type CreateProductInput struct {
	Name        string  `json:"product_name" binding:"required"`
	Description string  `json:"product_description"`
	Price       float32 `json:"price" binding:"required"`
}
