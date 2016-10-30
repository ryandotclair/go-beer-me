package controllers

import (
  "encoding/json"
  "fmt"
  "net/http"

  "log"

  // Non-Standard
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "github.com/julienschmidt/httprouter"
  "github.com/ryandotclair/go-beer-me/models"
)

type (
    // BeerController represents the controller for operating on the Beer resource
    BeerController struct{
      session *mgo.Session
    }
)

func NewBeerController(s *mgo.Session) *BeerController {
    return &BeerController{s}
}

func (bc BeerController) GetBeer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // Grab id
    id := p.ByName("id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    // Grab id
    oid := bson.ObjectIdHex(id)

    // Stub beer
    b := models.Beer{}

    // Fetch user
    if err := bc.session.DB("go_beer_me").C("beer").FindId(oid).One(&b); err != nil {
        w.WriteHeader(404)
        return
    }

    // Marshal provided interface into JSON structure
    bj, _ := json.Marshal(b)

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", bj)
    log.Println("Retrieved beer " + oid)
}

func (bc BeerController) GetBeers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    type beer struct{
      Id   bson.ObjectId  `json:"id" bson:"_id"`
      Name string         `json:"name" bson:"name"`
      Type string         `json:"type" bson:"type"`
      ABV  float64        `json:"ABV" bson:"ABV"`
      Cost float64        `json:"cost" bson:"cost"`
    }

    var beers []beer

    // Fetch beers
    if err := bc.session.DB("go_beer_me").C("beer").Find(nil).All(&beers); err != nil {
        w.WriteHeader(404)
        return
    }

    // Marshal provided interface into JSON structure
    bj, _ := json.Marshal(beers)

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", bj)
    log.Println("Retrieved all beers.")
}

// CreateBeer creates a new beer resource
func (bc BeerController) CreateBeer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // Stub a beer to be populated from the body
    b := models.Beer{}

    // Populate the beer data
    json.NewDecoder(r.Body).Decode(&b)

    // Add an Id
    b.Id = bson.NewObjectId()

    // Write the user to mongo
    bc.session.DB("go_beer_me").C("beer").Insert(b)

    // Marshal provided interface into JSON structure
    bj, _ := json.Marshal(b)

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", bj)
    log.Println("Added beer " + b.Id)
}

// RemoveBeer removes an existing beer resource
func (bc BeerController) RemoveBeer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  // Grab id
    id := p.ByName("id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    // Grab id
    oid := bson.ObjectIdHex(id)

    // Remove user
    if err := bc.session.DB("go_beer_me").C("beer").RemoveId(oid); err != nil {
        w.WriteHeader(404)
        return
    }

    // Write status
    w.WriteHeader(200)
    log.Println("Removed beer " + oid)
}

func (bc BeerController) RemoveBeers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // Remove user
    bc.session.DB("go_beer_me").C("beer").RemoveAll(nil)

    // Write status
    w.WriteHeader(200)
    log.Println("Removed all beers.")
}
