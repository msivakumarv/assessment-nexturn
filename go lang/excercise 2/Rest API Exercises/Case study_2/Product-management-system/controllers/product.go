package controllers

import (
	"net/http"
	"product-management-system/models"
	"strconv"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func HomePage(w http.ResponseWriter, r *http.Request) {
	rows, err := models.DB.Query("SELECT * FROM products")
	if err != nil {
		http.Error(w, "Failed to load products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category)
		products = append(products, product)
	}

	tmpl.Execute(w, products)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		category := r.FormValue("category")

		_, err := models.DB.Exec("INSERT INTO products (name, description, price, category) VALUES (?, ?, ?, ?)", name, description, price, category)
		if err != nil {
			http.Error(w, "Failed to add product", http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		category := r.FormValue("category")

		_, err := models.DB.Exec("UPDATE products SET name = ?, description = ?, price = ?, category = ? WHERE id = ?", name, description, price, category, id)
		if err != nil {
			http.Error(w, "Failed to update product", http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := models.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
