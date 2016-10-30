package models

import "gopkg.in/mgo.v2/bson"

type (
  // Beer represents the structure of the resource
  Beer struct {
    Id   bson.ObjectId  `json:"id" bson:"_id"`
    Name string         `json:"name" bson:"name"`
    Type string         `json:"type" bson:"type"`
    ABV  float64        `json:"ABV" bson:"ABV"`
    Cost float64        `json:"cost" bson:"cost"`
  }
)
