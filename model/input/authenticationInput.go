package input

type RegisterBuyerInput struct {
	Email            string `json:"email" binding:"required"`
	Name             string `json:"name" binding:"required"`
	Password         string `json:"password" binding:"required"`
	AlamatPengiriman string `json:"alamat_pengiriman"`
}

type RegisterSellerInput struct {
	Email        string `json:"email" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Password     string `json:"password" binding:"required"`
	AlamatPickup string `json:"alamat_pickup"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
