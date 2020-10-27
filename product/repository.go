package product

import "database/sql"

type Repository interface { //Conformada por firma de metodos
	GetProductById(productId int) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(dataBaseConnection *sql.DB) Repository {
	return &repository{db: dataBaseConnection}
}

func (repo *repository) GetProductById(productId int) (*Product, error) { //Consulta sobre la tabla productos
	const sql = `select id, product_code, product_name, COALESCE(description, ''), 
					standard_cost, list_price, 
					category 
					FROM products
					WHERE id=?`
	row := repo.db.QueryRow(sql, productId)
	product := &Product{}

	err := row.Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.Description, //Mapenado los datos
		&product.StandardCost, &product.ListPrice, &product.Category) //Mapendo a la estructura los resultados del select
	if err != nil {
		panic(err)
	}
	return product, err
}

func (repo *repository) GetProducts(params *getProductsRequest) ([]*Product, error) {
	const sql = `SELECT 
					id, product_code, 
					product_name, COALESCE(description, ''),
					standard_cost, list_price,
					category 
				FROM products
				order by id
				LIMIT ? OFFSET ?`
	result, err := repo.db.Query(sql, params.Limit, params.Offset)
	if err != nil {
		panic(err)
	}
	var products []*Product
	for result.Next() {
		product := &Product{} //Creando instancia vacia del tipo Product

		err = result.Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.Description,
			&product.StandardCost, &product.ListPrice, &product.Category)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	return products, nil
}

func (repo *repository) GetTotalProducts() (int, error) {
	const sql = `select count(*) FROM products`
	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	if err != nil {
		panic(err)
	}
	return total, nil
}
