package main

import (
	"github.com.br/jimmmisss/api/configs"
	"github.com.br/jimmmisss/api/internal/entity"
	"github.com.br/jimmmisss/api/internal/infra/database"
	"github.com.br/jimmmisss/api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	configuration, err := configs.LoadConfig(".")
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configuration.TokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", configuration.JWTExpiresIn))

	r.Post("/users", userHandler.CreateUser)
	r.Post("/token", userHandler.GetToken)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configuration.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
