package domain

type AuthorizationsRepo interface {
	Create(data Authorizations) (Authorizations, *HttpException)
	Update(id AuthStatus, data Authorizations) (Authorizations, *HttpException)
	Delete(id AuthStatus) *HttpException
	List() ([]Authorizations, int64, *HttpException)
	FindOne(id AuthStatus) (Authorizations, *HttpException)
}
