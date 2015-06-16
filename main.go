// +build !appengine

package main

import (
	"fmt"
	"log"
	"net/http"

	_ "golang.org/x/tools/present"

	_ "github.com/ebittleman/go-course/blog"
	_ "github.com/ebittleman/go-course/cart"
	"github.com/ebittleman/go-course/jsondb"
)

var tables map[string]interface{}

func init() {
	products := make([]*Product, 0, 100)
	tables = map[string]interface{}{
		"products": &products,
	}

	for name, table := range tables {
		err := jsondb.LoadTable(name, table)
		if err != nil {
			log.Println(err)
			continue
		}
	}

}

func main() {
	defer WriteTables()

	CreateProduct(&Product{
		Name:  "New Thing",
		Price: 16,
	})

	for _, ptr := range tables {
		switch table := ptr.(type) {
		case *[]*Product:
			fmt.Println(*table)
		}
	}

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func WriteTables() {
	for name, table := range tables {
		jsondb.WriteTable(name, table)
	}
}

func CreateProduct(product *Product) error {
	products, err := GetProducts()

	if err != nil {
		return err
	}

	*products = append(*products, product)

	return nil
}

func GetProducts() (*[]*Product, error) {
	inst, ok := tables["products"]

	if !ok {
		return nil, fmt.Errorf("Products Table Not Loaded")
	}

	products, ok := inst.(*[]*Product)

	if !ok {
		return nil, fmt.Errorf("Products Failed To Initialize Properly")
	}

	return products, nil
}

type Product struct {
	Name  string  `json"name"`
	Price float64 `json:"price"`
}

func (p *Product) String() string {
	return fmt.Sprintf("Name: %s, Price: %f", p.Name, p.Price)
}
