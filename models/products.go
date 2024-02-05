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
