package models

import "products/database"

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func GetAllProducts() []Product {
	db := database.ConnectDB()

	productsSelect, err := db.Query("SELECT * FROM products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for productsSelect.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productsSelect.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateProduct(name, description string, price float64, quantity int) {
	db := database.ConnectDB()

	insertData, err := db.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, quantity)
	defer db.Close()
}

func EditProduct(productId string) Product {
	db := database.ConnectDB()

	product, err := db.Query("select * from products where id=$1", productId)
	if err != nil {
		panic(err.Error())
	}

	productUpdate := Product{}

	for product.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = product.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		productUpdate.Id = id
		productUpdate.Name = name
		productUpdate.Description = description
		productUpdate.Price = price
		productUpdate.Quantity = quantity
	}

	defer db.Close()
	return productUpdate
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := database.ConnectDB()

	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := database.ConnectDB()

	deleteProduct, err := db.Prepare("delete from products where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}
