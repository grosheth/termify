package main

import (
	"fmt"
  // "github.com/rapito/go-spotify/spotify"
	"github.com/joho/godotenv"
	"os"
	"log"
)


func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  clientid := os.Getenv("CLIENT_ID")
  clientsecret := os.Getenv("CLIENT_SECRET")

  fmt.Println("vachier")	
	// call_api()
}

// func call_api() {
//     
//   spot := spotify.New(clientid,clientsecret)
//   result, _ := spot.Get("albums/%s", nil, "0sNOF9WDwhWunNAHPD3Baj")
//   
//   fmt.Println(string(result))	
// }
