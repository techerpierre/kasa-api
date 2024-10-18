package domain

type UserRepo interface {
	Create(data User) (User, *HttpException)
	Update(id string, data User) (User, *HttpException)
	Delete(id string) *HttpException
	List(listing Listing) ([]User, int64, *HttpException)
	FindOne(id string) (User, *HttpException)
	FindOneByEmail(email string) (User, *HttpException)
}
