package main

import (
  "net/http"
  "os"
  "log"

  // Non-Standard
  "gopkg.in/mgo.v2"
  "github.com/julienschmidt/httprouter"
  "github.com/ryandotclair/go-beer-me/controllers"
)

func main() {
  // Instantiate a new router
  r := httprouter.New()

  // Get a BeerController instance
  bc := controllers.NewBeerController(getSession())

  // Get a beer resource
  r.GET("/beer/:id", bc.GetBeer)

  // Get all beers
  r.GET("/beer", bc.GetBeers)

  // Create a beer
  r.POST("/beer", bc.CreateBeer)

  // Delete a beer
  r.DELETE("/beer/:id", bc.RemoveBeer)

  // Delete all beers
  r.DELETE("/beer", bc.RemoveBeers)


  // Fire up the server
  http.ListenAndServe("0.0.0.0:3000", r)
}

func getSession() *mgo.Session {
    // Get location of mongodb from environment
    ml := "mongodb://" + os.Getenv("MONGO_LOCATION")

    // Connect to mongo
    log.Println("Connecting to database at", os.Getenv("MONGO_LOCATION"))

    s, err := mgo.Dial(ml)

    // Check if connection error, is mongo running?
    if err != nil {
        log.Println("Can't connect to mongo. Make sure MONGO_LOCATION environment variable is set.")
        log.Println(err)
        panic(err)
    }
    log.Println("Connected to database.")

    return s
}
