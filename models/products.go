package models

import "github.com/product-manager/db"

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Amount            int
}

func GetAllProducts() []Product {

	//abre a conexão
	db := db.ConnectDB()

	allProducs, err := db.Query("select * from products;")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for allProducs.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = allProducs.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}

	//fecha a conexão
	defer db.Close()

	return products
}

func InsertProduct(name, description string, price float64, amount int) {
	db := db.ConnectDB()

	insertQuery, err := db.Prepare("insert into products(name, description, price, amount) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertQuery.Exec(name, description, price, amount)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDB()

	deleteQuery, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteQuery.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectDB()

	productData, err := db.Query("select * from products where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for productData.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = productData.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount
	}

	defer db.Close()
	return product
}

func UpdateProduct(id int, name, description string, price float64, amount int) {
	db := db.ConnectDB()

	updateQuery, err := db.Prepare("update products set name=$1, description=$2, price=$3, amount=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateQuery.Exec(name, description, price, amount, id)
	defer db.Close()
}
