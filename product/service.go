package product

type Service interface {
	GetProductById(param *getProductByIDRequest) (*Product, error) //EL param es una estructra del archivo endpoint.go
	GetProducts(params *getProductsRequest) (*ProductsList, error)
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

//LLamando a los metodos del repository
func (s *service) GetProducts(params *getProductsRequest) (*ProductsList, error) {
	products, err := s.repo.GetProducts(params)
	if err != nil {
		panic(err)
	}
	totalProducts, err := s.repo.GetTotalProducts()
	if err != nil {
		panic(err)
	}
	return &ProductsList{Data: products, TotalRecords: totalProducts}, nil
}
