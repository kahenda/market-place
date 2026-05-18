package models

import "time"

type Listing struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	Gender      string    `json:"gender"`
	Size        string    `json:"size"`
	Condition   string    `json:"condition"`
	Area        string    `json:"area"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateListingRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Category    string  `json:"category" binding:"required,oneof=tops bottoms dresses suits jackets shoes accessories kids_clothing"`
	Gender      string  `json:"gender" binding:"required,oneof=male female unisex kids"`
	Size        string  `json:"size" binding:"required,oneof=XXS XS S M L XL XXL XXXL 'UK6' 'UK7' 'UK8' 'UK9' 'UK10' 'UK11' 'UK12' one_size"`
	Condition   string  `json:"condition" binding:"required,oneof=new like_new good fair"`
	Area        string  `json:"area" binding:"required,oneof=CBD Milimani Kondele Kibuye Nyalenda Mamboleo 'Tom Mboya' Otonglo Manyatta Bandani Migosi Riat Lolwe Kaloleni"`
}

var Categories = []string{
	"tops", "bottoms", "dresses", "suits",
	"jackets", "shoes", "accessories", "kids_clothing",
}

var Genders = []string{"male", "female", "unisex", "kids"}

var Sizes = []string{
	"XXS", "XS", "S", "M", "L", "XL", "XXL", "XXXL",
	"UK6", "UK7", "UK8", "UK9", "UK10", "UK11", "UK12",
	"one_size",
}

var Conditions = []string{"new", "like_new", "good", "fair"}

var KisumuAreas = []string{
	"CBD", "Milimani", "Kondele", "Kibuye", "Nyalenda",
	"Mamboleo", "Tom Mboya", "Otonglo", "Manyatta",
	"Bandani", "Migosi", "Riat", "Lolwe", "Kaloleni",
}
