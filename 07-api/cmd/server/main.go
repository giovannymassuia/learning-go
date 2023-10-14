package main

import (
	"github.com/giovannymassuia/learning-go/07-api/configs"
	"github.com/giovannymassuia/learning-go/07-api/internal/entity"
	"github.com/giovannymassuia/learning-go/07-api/internal/infra/database"
	"github.com/giovannymassuia/learning-go/07-api/internal/infra/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"

	_ "github.com/giovannymassuia/learning-go/07-api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Example API
// @version v1
// @description This is a sample server Petstore server.
// @TermsOfService http://swagger.io/terms/

// @contact.name Giovanny
// @contact.url giovannymassuia.io

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	r := chi.NewRouter()
	//r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.TokenAuth))
	r.Use(LogRequest)

	//r.Use(LogResponse)

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Post("/", productHandler.CreateProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, config.JWTExpiresIn)

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/jwt", userHandler.GetJwt)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":"+config.WebServerPort, r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("My middleware ->> Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func LogResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("My middleware ->> Response: %s %s", r.Method, r.URL.Path)
		log.Printf("My middleware ->> Response: %s", w.Header().Get("Content-Type"))
		next.ServeHTTP(w, r)
	})
}
