package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	Port     = ":8080"
	User     = "postgres"
	Password = "postgres"
	Database = "rundb"
)

var AllSuppliers []Supplier

type Supplier struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Product struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	SKU      string `json:"sku"`
}

type SupplierJsonResponse struct {
	Type    string     `json:"type"`
	Message string     `json:"message"`
	Data    []Supplier `json:"data"`
}

type ProductJsonResponse struct {
	Type    string    `json:"type"`
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/suppliers/", GetSuppliers).Methods("GET")
	r.HandleFunc("/suppliers/", CreateSupplier).Methods("POST")
	r.HandleFunc("/suppliers/{name}", DeleteSupplier).Methods("DELETE")

	r.HandleFunc("/products", QueryProducts).Methods("GET")

	fmt.Println(r.Name("Server listening at port " + Port + " ..."))
	http.ListenAndServe(Port, r)
}

func CreateSupplier(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating supplier ...")

	w.Header().Set("Content-Type", "application/json")
	var supplier Supplier
	_ = json.NewDecoder(r.Body).Decode(&supplier)

	db := connectDB()
	if _, err := db.Exec("insert into suppliers values ($1, $2)", supplier.Name, supplier.Address); err != nil {
		panic(err)
	}

	fmt.Printf("\t%+v\n", supplier)

	response := SupplierJsonResponse{
		Type:    "success",
		Message: "the supplier was added successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func GetSuppliers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting suppliers ...")

	db := connectDB()
	rows, err := db.Query("select name, address from suppliers")
	if err != nil {
		panic(err)
	}

	var suppliers []Supplier

	for rows.Next() {
		var name string
		var address string
		err = rows.Scan(&name, &address)
		if err != nil {
			panic(err)
		}
		s := Supplier{Name: name, Address: address}
		fmt.Printf("\t%+v\n", s)
		suppliers = append(suppliers, s)
	}

	var response = SupplierJsonResponse{
		Type:    "success",
		Message: "all suppliers",
		Data:    suppliers,
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting supplier ...")

	params := mux.Vars(r)
	name := params["name"]

	var response SupplierJsonResponse
	if name == "" {
		response = SupplierJsonResponse{
			Type:    "error",
			Message: "the name is missing",
		}
	} else {
		db := connectDB()
		if _, err := db.Exec("delete from suppliers where name = $1", name); err != nil {
			panic(err)
		}

		fmt.Printf("\tname: %v\n", name)
		response = SupplierJsonResponse{
			Type:    "success",
			Message: "the supplier was deleted successfully",
		}
	}

	json.NewEncoder(w).Encode(response)
}

func QueryProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Querying products ...")

	query := "select name, category, sku from products"
	keys := []string{"name", "category", "sku"}
	found := 0
	for _, k := range keys {
		if v := r.URL.Query().Get(k); v != "" {
			query += queryPrefix(found) + k + " like '%" + v + "%'"
			found++
		}
	}
	fmt.Printf("\t%v\n", query)

	db := connectDB()
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	var products []Product

	for rows.Next() {
		var name string
		var category string
		var sku string
		err = rows.Scan(&name, &category, &sku)
		if err != nil {
			panic(err)
		}
		products = append(products, Product{Name: name, Category: category, SKU: sku})
	}

	var response = ProductJsonResponse{
		Type:    "success",
		Message: "all qualifying products",
		Data:    products,
	}
	json.NewEncoder(w).Encode(response)
}

func connectDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", User, Password, Database)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	return db
}

func queryPrefix(paramPosition int) string {
	if paramPosition == 0 {
		return " where "
	} else {
		return " and "
	}
}
