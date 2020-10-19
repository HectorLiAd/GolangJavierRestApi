package product

import "database/sql"

type Repository interface { //Conformada por firma de metodos
	GetProductById(productId int) (*Product, error)
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
