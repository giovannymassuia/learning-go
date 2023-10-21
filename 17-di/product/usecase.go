package product

type UseCase struct {
	repository *Repository
}

func NewUseCase(repository *Repository) *UseCase {
	return &UseCase{repository}
}

func (u *UseCase) GetProduct(id int) (Product, error) {
	return u.repository.GetProduct(id)
}
