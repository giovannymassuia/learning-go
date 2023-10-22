package product

type UseCase struct {
	repository RepositoryInterface
}

func NewUseCase(repository RepositoryInterface) *UseCase {
	return &UseCase{repository}
}

func (u *UseCase) GetProduct(id int) (Product, error) {
	return u.repository.GetProduct(id)
}
