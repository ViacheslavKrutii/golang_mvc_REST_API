package models

type hryvna float32

type menuItem struct {
	name  string
	price hryvna
}

type Menu struct {
	MenuItems []menuItem
}
