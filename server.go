package main

import (
  "net/http"
  "os"
  "log"
  "fmt"
  //"encoding/json"
  "regexp"

  // Non-Standard
  "gopkg.in/mgo.v2"
  "github.com/julienschmidt/httprouter"
  "github.com/ryandotclair/go-beer-me/controllers"
)

// type jsonObject struct {
//     UserProvided []*userProvided `json:"user-provided"`
// }
//
// type userProvided struct {
//     Credentials *credentials `json:"credentials"`
// }
//
// type credentials struct {
//     Location string `json:"MONGO_LOCATION"`
// }

func main() {
  // Instantiate a new router
  r := httprouter.New()

  // Get a BeerController instance
  bc := controllers.NewBeerController(getSession())

  // Get a beer resource
  r.GET("/api/beer/:id", bc.GetBeer)

  // Get all beers
  r.GET("/api/beer", bc.GetBeers)

  // Create a beer
  r.POST("/api/beer", bc.CreateBeer)

  // Delete a beer
  r.DELETE("/api/beer/:id", bc.RemoveBeer)

  // Delete all beers
  r.DELETE("/api/beer", bc.RemoveBeers)


  // Fire up the server
  http.ListenAndServe("0.0.0.0:" + os.Getenv("PORT"), r)
}

func getSession() *mgo.Session {

    // Check for location of mongodb or use default
    VCAP := []byte(os.Getenv("VCAP_SERVICES"))

    rx := regexp.MustCompile(`.*"MONGO_LOCATION":\s*?"(.+?)".*`)

    var temp []byte

    if m := rx.FindSubmatch(VCAP); len(m) > 0 {
        temp = m[1]
        os.Stdout.Write(m[1])
        fmt.Fprintln(os.Stdout)

    } else {
        panic("error matching MONGO_LOCATION")
    }

    loc := string(temp[:])
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

// func parseLocation(buf []byte) (string, error) {
// 	data := struct {
// 		UserProvided []struct {
// 			Credentials struct {
// 				Location string `json:"MONGO_LOCATION"`
// 			} `json:"credentials"`
// 		} `json:"user-provided"`
// 	}{}
// 	if err := json.Unmarshal(buf, &data); err != nil {
// 		return "", err
// 	}
// 	return data.UserProvided[0].Credentials.Location, nil
// }
