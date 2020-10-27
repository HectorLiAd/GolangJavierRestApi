//Aqui van
package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

//CAPTURAR LOS PARAMETROS DE ENTRADA DEL REQUEST
type getProductsRequest struct {
	Limit  int
	Offset int
}

type getProductByIDRequest struct {
	ProductID int //Se guardara el paramatro que está llendo en la URL -> localhost:3000/products/1
}

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) /* Nos servira para hacer concurrente */ {
		req := request.(getProductByIDRequest) //Parsenado el request y se le envia a service
		product, err := s.GetProductById(&req)

		if err != nil {
			panic(err)
		}
		return product, nil
	}

	return getProductByIdEndPoint
}

//METODO PARA INVOCAR AL METODO GETPRODUCTS
func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	//Convertir el tipo "request" al tipo  getProductsRequest
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getProductsEndPoint
}
