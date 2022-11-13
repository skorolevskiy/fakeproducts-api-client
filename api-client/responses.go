package api_client

import (
	"fmt"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rating  `json:"rating"`
}

type Rating struct {
	Rate  float32 `json:"rate"`
	Count int     `json:"count"`
}

type Cart struct {
	ID       int           `json:"id"`
	UserID   int           `json:"userId"`
	Date     time.Time     `json:"date"`
	Products []CartProduct `json:"products"`
}

type CartProduct struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

func (p Product) Info() string {
	return fmt.Sprintf("[ID] %d \n[Title] %s \n[Price] %f\n[Description] %s\n[Category] %s\n[Image] %s\n[Rating] %.2f\n",
		p.ID, p.Title, p.Price, p.Description, p.Category, p.Image, p.Rating.Rate)
}
