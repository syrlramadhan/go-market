package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/syrlramadhan/go-market/app/util"
	"gorm.io/gorm"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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

func ProductHandler(db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		products, err := GetProduct(db)
		if err != nil {
			http.Error(w, "failed get products", http.StatusInternalServerError)
		}
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		isAuthenticated := util.IsLogin(r)
		RenderTemplate(w, "products.html", map[string]interface{}{
			"Title": "Products - GoMarket",
			"Products": products,
			"IsAuthenticated": isAuthenticated,
		})
	}
}

func ProductDetailHandler(db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		productSlug := ps.ByName("slug")
		printer := message.NewPrinter(language.Indonesian)

		selectedProduct, err := GetProductBySlug(productSlug, db)
		if err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		selectedPrice := printer.Sprintf("%.0f", selectedProduct.Price)

		relatedProducts, err := GetProduct(db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		isAuthenticated := util.IsLogin(r)
		RenderTemplate(w, "product-detail.html", map[string]interface{}{
			"Title":    selectedProduct.Name + " - GoMarket",
			"Product":  selectedProduct,
			"SelectedPrice": selectedPrice,
			"Products": relatedProducts,
			"IsAuthenticated": isAuthenticated,
		})
	}
}