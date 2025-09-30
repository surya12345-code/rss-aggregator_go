package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/surya123/RSSaggregator/internal/database"
)

type apiconfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("hello world")
	godotenv.Load(".env")
	portstring := os.Getenv("PORT")
	if portstring == "" {
		log.Fatal("port is not found in the environment")

	}

	fmt.Println("port:", portstring)
	dburl := os.Getenv("DB_URL")
	if dburl == "" {
		log.Fatal("DB_URL is not found in the environment")
	}
	conn, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal("There is an error in connecting to database")
	}
	queries := database.New(conn)
	apicfg := apiconfig{
		DB: queries,
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/responderror", handlererror)
	v1Router.Post("/users", apicfg.handlerCreateUser)
	v1Router.Get("/users", apicfg.middlewareAuth(apicfg.handlerGetUserByApiKey))
	v1Router.Post("/feed", apicfg.middlewareAuth(apicfg.handlerCreatefeed))
	v1Router.Get("/feeds", apicfg.handlerGetFeeds)
	v1Router.Post("/feedfollows", apicfg.middlewareAuth(apicfg.handlerFeedfollows))
	v1Router.Get("/allusers", apicfg.Getallusers)
	v1Router.Get("/userfeed", apicfg.middlewareAuth(apicfg.handlerGetFeedfollows))
	v1Router.Delete("/deleteuserfeed/{feedfollowid}", apicfg.middlewareAuth(apicfg.handlerDeletefeed))
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
