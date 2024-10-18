package domain

type UserRepo interface {
	Create(data User) (User, error)
	Update(id string, data User) (User, error)
	Delete(id string) error
	List(listing Listing) ([]User, int64, error)
	FindOne(id string) (User, error)
	FindOneByEmail(email string) (User, error)
}
