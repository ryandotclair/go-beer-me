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

    // Check for location of mongodb or use default
    loc := getenv("MONGO_LOCATION", "localhost")

    // TODO: Get USER/PASS overrides working
    // // Check for custom mongo username or use default
    // muser := getenv("MONGO_USER","")
    //
    // // Check for custom mongo password or use default
    // mpass := getenv("MONGO_PASS","")


    // Connect to mongo
    log.Println("Connecting to database on", loc)

    s, err := mgo.Dial(loc)

    // Check if connection error, is mongo running?
    if err != nil {
        log.Println("Can't connect to mongo. Make sure MONGO_LOCATION environment variable is set if it's not local.")
        log.Println(err)
        panic(err)
    }
    log.Println("Connected to database.")

    return s
}

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        log.Println(key, "environment variable wasn't set. Using default.")
        return fallback
    }
    log.Println(key, "was set. Overriding default.")
    return value
}
