## Purpose
A simple Go webapp that uses mongoDB and supports GET/PUT/DELETE of beers. Created for my own learning.

## Requires

The only thing this requires is mongoDB. It doesn't use user/password (so don't add one... this is not meant to be used in prod). By default, go-beer-me assumes mongoDB is running locally (`localhost`). If you've got an external mongoDB instance, you can point it to it using the MONGO_LOCATION environment variable (example: `export MONGO_LOCATION=192.168.100.50`, optionally you can add a port if it's non-standard).

## Running it

Easiest way to run this is using docker.

Steps:

+ Grab the mongo image
  + `docker pull mongo`
+ Run it
  + `docker run -dP mongo`
+ Check the port number it's running on (I assume you know how to find the IP of the container host)
  + `docker ps`
+ Run this webapp
  + `docker run -dP -e MONGO_LOCATION=<ip:port_of_mongo> ryandotclair/go-beer-me`
+ Check the port number go-beer-me is running on
  + `docker ps`

Now you can use `curl` to use it

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
