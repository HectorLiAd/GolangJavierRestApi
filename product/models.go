package product

type Product struct {
	Id           int     `json:"id"`
	ProductCode  string  `json:"productCode"`
	ProductName  string  `json:"productName"`
	Description  string  `json:"description"`
	StandardCost float64 `json:"standardCost"`
	ListPrice    float64 `json:"listPrice"`
	Category     string  `json:"category"`
}

//ESTRUCTTURA  EN EL CUAL SE SEVOLVERA EL RESPONCE
type ProductsList struct {
	Data         []*Product `json:"data"`
	TotalRecords int        `json:"totalRecords"`
}
