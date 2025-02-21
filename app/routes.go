package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/syrlramadhan/go-market/app/controller"
	"github.com/syrlramadhan/go-market/app/middleware"
	"github.com/syrlramadhan/go-market/app/util"
	"gorm.io/gorm"
)

func Routes(router *httprouter.Router, db *gorm.DB) *httprouter.Router {
	router.GET("/", controller.HomeHandler(db))

	router.GET("/products", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		products, err := controller.ProductHandler(db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		isAuthenticated := util.IsLogin(r)
		controller.RenderTemplate(w, "products.html", map[string]interface{}{
			"Title": "Products - GoMarket",
			"Products": products,
			"IsAuthenticated": isAuthenticated,
		})
	}))
	router.GET("/products/product-detail/:slug", controller.ProductDetailHandler(db))

	router.GET("/register", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		data := map[string]interface{}{
			"Title": "Register - GoMarket",
		}
		controller.RenderTemplate(w, "register.html", data)
	})
	router.POST("/register/add", controller.CreateUserHandler(db))

	router.GET("/login", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		data := map[string]interface{}{
			"Title": "Login - GoMarket",
		}
		controller.RenderTemplate(w, "login.html", data)
	})
	router.POST("/login", controller.LoginUserHandler(db))
	router.GET("/logout", controller.LogoutUserHandler(db))

	router.ServeFiles("/assets/*filepath", http.Dir("assets"))

	return router
}