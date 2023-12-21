package domain

type User struct {
	Name string
}
type Popularity struct {
	User
	Value int
}
