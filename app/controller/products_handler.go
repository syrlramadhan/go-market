package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/syrlramadhan/go-market/app/model"
	"github.com/syrlramadhan/go-market/app/util"
	"gorm.io/gorm"
)

func HomeHandler(db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		products, err := GetProduct(db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		isAuthenticated := util.IsLogin(r)
		RenderTemplate(w, "home.html", map[string]interface{}{
			"Title":    "Home - GoMarket",
			"Products": products,
			"IsAuthenticated": isAuthenticated,
		})
	}
}

func ProductHandler(db *gorm.DB) ([]model.MstProducts, error) {
	products, err := GetProduct(db)
	if err != nil {
		return products, err
	}

	return products, nil
}

func ProductDetailHandler(db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		productSlug := ps.ByName("slug")

		product, err := GetProductBySlug(productSlug, db)
		if err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		products, err := GetProduct(db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		isAuthenticated := util.IsLogin(r)
		RenderTemplate(w, "product-detail.html", map[string]interface{}{
			"Title":    product.Name + " - GoMarket",
			"Product":  product,
			"Products": products,
			"IsAuthenticated": isAuthenticated,
		})
	}
}
