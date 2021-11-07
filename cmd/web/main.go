package main

import (
	"heyapple/internal/mock"
	"heyapple/pkg/api/v1"
	"heyapple/pkg/middleware"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := mock.NewDB().Prefill().Fail(false) // uncomment for testing
	//db := &storage.DB{} // uncomment for production

	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(middleware.Options)
	router.GET("/api/v1/foods", api.Foods(db))

	handler := middleware.Headers(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
