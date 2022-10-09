package models

import (
	"store-go/db"
)

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Amount            int
}

func GetProducts() []Product {
	db := db.ConnectDatabase()
	result, err := db.Query("SELECT * FROM products ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	product := Product{}
	products := []Product{}
	for result.Next() {
		var id, amount int
		var name, description string
		var price float64
		err = result.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}
		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount
		products = append(products, product)
	}
	defer db.Close()
	return products
}

func InsertProduct(name, description string, price float64, amount int) {
	db := db.ConnectDatabase()
	prepareQuery, err := db.Prepare("INSERT INTO products(name, description, price, amount) values ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	prepareQuery.Exec(name, description, price, amount)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDatabase()
	prepareQuery, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	prepareQuery.Exec(id)
	defer db.Close()
}

func GetProduct(id string) Product {
	db := db.ConnectDatabase()
	result, err := db.Query("SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	product := Product{}
	for result.Next() {
		var id, amount int
		var name, description string
		var price float64
		err = result.Scan(&id, &name, &description, &price, &amount)
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
	db := db.ConnectDatabase()
	prepareQuery, err := db.Prepare("UPDATE products set name=$1, description=$2, price=$3, amount=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}
	prepareQuery.Exec(name, description, price, amount, id)
	defer db.Close()
}
