package domain

type AuthorizationsService struct {
	repo AuthorizationsRepo
}

func CreateAuthorizationsService(repo AuthorizationsRepo) *AuthorizationsService {
	return &AuthorizationsService{
		repo: repo,
	}
}

func (s *AuthorizationsService) Create(data Authorizations) (Authorizations, *HttpException) {
	return s.repo.Create(data)
}

func (s *AuthorizationsService) Update(id AuthStatus, data Authorizations) (Authorizations, *HttpException) {
	return s.repo.Update(id, data)
}

func (s *AuthorizationsService) Delete(id AuthStatus) *HttpException {
	return s.repo.Delete(id)
}

func (s *AuthorizationsService) List() ([]Authorizations, int64, *HttpException) {
	return s.repo.List()
}

func (s *AuthorizationsService) FindOne(id AuthStatus) (Authorizations, *HttpException) {
	return s.repo.FindOne(id)
}
