package main

import (
	"github.com.br/jimmmisss/api/configs"
	"github.com.br/jimmmisss/api/internal/entity"
	"github.com.br/jimmmisss/api/internal/infra/database"
	"github.com.br/jimmmisss/api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&entity.Product{}, &entity.User{})
	if err != nil {
		panic(err)

	}

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	routes := initializeRoutes(productHandler, userHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	log.Println("Server running on port 8080")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func initializeRoutes(productHandler *handlers.ProductHandler, userHandler *handlers.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	//Products Routes
	mux.HandleFunc("GET /products", productHandler.GetProducts)
	mux.HandleFunc("GET /products/{id}", productHandler.GetProduct)
	mux.HandleFunc("POST /products", productHandler.CreateProduct)
	mux.HandleFunc("PUT /products/{id}", productHandler.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", productHandler.DeleteProduct)

	//Users Routes
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	return mux
}
