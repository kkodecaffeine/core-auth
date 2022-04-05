package user

// Repository interface definition
type Repository interface {
	FindOne(ID int64) (*User, error)
}
