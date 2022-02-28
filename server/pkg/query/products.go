package query

import (
	"log"
	dbmanager "team2-real-world-app/server/pkg/database"
	"team2-real-world-app/server/pkg/helpers"
)

type Product struct {
	ProductID   int    `db:"product_id"   json:"product_id"`
	ProductName string `db:"product_name" json:"product_name"`
}

// AllProducts - query that return all product in the Database
func AllProducts() ([]byte, error) {

	// create Database connection
	var db = dbmanager.NewDBManager()
	log.Printf("** Try to connected\n")

	dbx, err := db.GetConnection()
	if err != nil {
		return nil, err
	}

	defer dbx.Close()

	var product []Product
	err = dbx.Select(&product, "SELECT product_id, product_name FROM product")
	if err != nil {
		return nil, err
	}

	products, err := helpers.StructToJSON(product)

	// to read the JSON as key: value
	//fmt.Println(string(products))
	return products, nil

}
