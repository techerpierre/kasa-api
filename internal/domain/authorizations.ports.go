package domain

type AuthorizationPorts interface {
	Create(data Authorizations) (Authorizations, error)
	Update(id AuthStatus, data Authorizations) (Authorizations, error)
	Delete(id AuthStatus) error
	List() ([]Authorizations, int32, error)
	FindOne(id AuthStatus) (Authorizations, error)
}
