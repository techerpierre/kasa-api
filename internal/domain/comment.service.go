package domain

type CommentService struct {
	repo CommentRepo
}

func CreateCommentService(repo CommentRepo) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) Create(data Comment) (Comment, error) {
	return s.repo.Create(data)
}

func (s *CommentService) Update(id string, data Comment) (Comment, error) {
	return s.repo.Update(id, data)
}

func (s *CommentService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *CommentService) List(listing Listing) ([]Comment, int64, error) {
	return s.repo.List(listing)
}

func (s *CommentService) FindOne(id string) (Comment, error) {
	return s.repo.FindOne(id)
}
