package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func HomeHandler(db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		products, err := GetProduct(db)
		if err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		RenderTemplate(w, "home.html", map[string]interface{}{
			"Title": "Home - GoMarket",
			"Products": products,
		})
	}
}

func ProductHandler(db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		products, err := GetProduct(db)
		if err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		RenderTemplate(w, "products.html", map[string]interface{}{
			"Title": "Products - GoMarket",
			"Products": products,
		})
	}
}

func ProductDetailHandler(db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		productID := ps.ByName("id")

		product, err := GetProductByID(productID, db)
		if err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		products, err := GetProduct(db)
		if err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		RenderTemplate(w, "product-detail.html", map[string]interface{}{
			"Title": product.Name+" - GoMarket",
			"Product": product,
			"Products": products,
		})		
	}
}
