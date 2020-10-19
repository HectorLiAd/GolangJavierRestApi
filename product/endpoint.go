package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int //Se guardara el paramatro que está llendo en la URL -> localhost:3000/products/1
}

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) /* Nos servira para hacer concurrente */ {
		req := request.(getProductByIDRequest)
		product, err := s.GetProductById(&req)

		if err != nil {
			panic(err)
		}
		return product, nil
	}

	return getProductByIdEndPoint
}
