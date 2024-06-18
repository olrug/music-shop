package models

type Users struct {
	UserId      int
	Name        string
	SecondName  string
	MiddleName  string
	Email       string
	PhoneNumber string
	Password    string
}

type Categories struct {
	CategoryId int
	Name       string
}

type SubCategories struct {
	SubCategoryId int
	CategoryId    int
	Name          string
}

type Products struct {
	ProductId     int
	SubCategoryId int
	CategoryId    int
	Price         int
	Description   string
}

type ShopingCard struct {
	UserId    int
	ProductId int
}
