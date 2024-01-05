package models

import (
	"github.com/behatris-fiorentini/curso_alura_go_web/database"
	"database/sql"
)

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       float64
	Quantity    int64
}

func GetProduct() ([]Product, error) {
	db, err := database.Connection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	productList, err := db.Query("select * from product")
	if err != nil {
		panic(err.Error())
	}

	products, err := modelProduct(productList)
	if err != nil {
		panic(err.Error())
	}

	return products, nil
}

func Create(name string, description string, price float64, quantity int64) {
	db, err := database.Connection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	queryProduct, err := db.Prepare("insert into product (name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = queryProduct.Exec(name, description, price, quantity)
	if err != nil {
		panic(err.Error())
	}
}

func Delete(id string) {
	db, err := database.Connection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query, err := db.Prepare("delete from product where id = $1")
	if err != nil {
		panic(err.Error())
	}

	query.Exec(id)
}

func GetByID(ID string) Product {
	db, err := database.Connection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query, err := db.Query("select * from product where id = $1", ID)
	if err != nil {
		panic(err.Error())
	}

	product, err := modelProduct(query)
	if err != nil {
		panic(err.Error())
	}

	return product[0]
}

func Update(ID, name, description string, price float64, quantity int64) {
	db, err := database.Connection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	GetByID(ID)

	query, err := db.Prepare("update product set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	query.Exec(name, description, price, quantity, ID)
}

func modelProduct(productList *sql.Rows) ([]Product, error) {
	p := Product{}
	products := []Product{}

	for productList.Next() {
		var id, quantity int64
		var name, description string
		var price float64

		err := productList.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			return nil, err
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
		p.ID = id

		products = append(products, p)
	}

	return products, nil
}
