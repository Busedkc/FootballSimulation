package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Rastgele sayı üreticisini başlat

	// HTTP sunucusunu ayarla ve başlat
	setupRoutes()
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
