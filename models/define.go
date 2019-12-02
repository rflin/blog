package models

type User struct{
	Id int `gorm:"primary_key"`
	Name string
}

type Category struct{
	Id int `gorm:"primary_key"`
	Name string
	Articles []Article
}

type Article struct{
	Id int `gorm:"primary_key"`
	Title string
	Body string
	CategoryID int
}