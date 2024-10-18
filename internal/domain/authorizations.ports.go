package domain

type AuthorizationsRepo interface {
	Create(data Authorizations) (Authorizations, error)
	Update(id AuthStatus, data Authorizations) (Authorizations, error)
	Delete(id AuthStatus) error
	List() ([]Authorizations, int64, error)
	FindOne(id AuthStatus) (Authorizations, error)
}
