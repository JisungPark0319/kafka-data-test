package generator

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	Name    string `fake:"{name}"`
	Gender  string `fake:"{gender}"`
	Hobby   string `fake:"{hobby}"`
	Email   string `fake:"{email}"`
	Phone   string `fake:"{phone}"`
	Car     string `fake:"{carmodel}"`
	CarType string `fake:"{cartype}"`
	City    string `fake:"{city}"`
	Country string `fake:"{country}"`
	State   string `fake:"{state}"`
	Street  string `fake:"{street}"`
}

type Image struct {
	Name       string `fake:"{inputname}"`
	Color      string `fake:"{color}"`
	Image      []byte
	CreateDate time.Time `fake:"{date}"`
	UpdateDate time.Time `fake:"{date}"`
}

func NewUser() *User {
	var user User
	gofakeit.Struct(&user)

	return &user
}

func NewImage(width int, height int) *Image {
	var image Image
	gofakeit.Struct(&image)
	image.Image = gofakeit.ImageJpeg(width, height)
	return &image
}
