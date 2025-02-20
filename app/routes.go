package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/syrlramadhan/go-market/app/controller"
	"gorm.io/gorm"
)

func Routes(router *httprouter.Router, db *gorm.DB) *httprouter.Router {
	router.GET("/", controller.HomeHandler(db))
	router.GET("/products", controller.ProductHandler(db))
	router.GET("/products/product-detail/:id", controller.ProductDetailHandler(db))
	router.GET("/register", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		data := map[string]interface{}{
			"Title": "Register - GoMarket",
		}
		controller.RenderTemplate(w, "register.html", data)
	})
	router.GET("/login", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		data := map[string]interface{}{
			"Title": "Login - GoMarket",
		}
		controller.RenderTemplate(w, "login.html", data)
	})

	router.ServeFiles("/assets/*filepath", http.Dir("assets"))

	return router
}