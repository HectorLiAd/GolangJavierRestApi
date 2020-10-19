package product

type Service interface {
	GetProductById(param *getProductByIDRequest) (*Product, error) //EL param es una estructra del archivo endpoint.go
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetProductById(param *getProductByIDRequest) (*Product, error) { //Implementando a la estructura service la interface Service
	//Inteligencia de negocio
	return s.repo.GetProductById(param.ProductID)
}
