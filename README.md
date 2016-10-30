## Purpose
A simple Go webapp that uses mongoDB and supports GET/PUT/DELETE of beers. Created for my own learning.

## Dependencies (If you wanted to build it yourself)

+ github.com/julienschmidt/httprouter
+ gopkg.in/mgo.v2
+ gopkg.in/mgo.v2/bson

## Running it

Easiest way is to use docker.

Grab the mongo image:

`docker pull mongo`

Run it:

`docker run -dP mongo`

Check the port number it's running on (I assume you know how to find the IP of the container host):

`docker ps`

Run this webapp:

`docker run -dP -e MONGO_LOCATION=<ip:port_of_mongo> ryandotclair/go-beer-me`

Then you can use Curl to use it

+ PUT beers in
  + `curl -XPOST -H 'Content-Type: application/json' -d '{"name": "Corona", "type": "Mexican", "ABV": 3.2, "cost":1.99}' <ip:port_of_go-beer-me>/beer`
+ GET the beer you put in
  + `curl <ip:port_of_go-beer-me>/beer/<id_of_beer>`
+ GET all beers
  + `curl <ip:port_of_go-beer-me>/beer`
+ DELETE a beer from the app
  + `curl -XDELETE <ip:port_of_go-beer-me>/beer/<id_of_beer>`
+ DELETE all beers from the app
  + `curl -XDELETE <ip:port_of_go-beer-me>/beer`
