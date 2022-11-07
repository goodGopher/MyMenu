package core

type Dish struct {
	Id          int
	Name        string
	Description string
	Categoryes  []Category
}

type User struct {
	Id       int
	Login    string
	Password string
}

type Category struct {
	Id   string
	Name string
}
