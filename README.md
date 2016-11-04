## Purpose
A simple Go webapp that uses mongoDB and supports GET/PUT/DELETE of beers. Created for my own learning.

## Special Thanks
Huge shout out to [akutz](https://github.com/akutz) from the [Dell EMC {Code} community](https://community.codedellemc.com/). This would have taken a lot longer if it wasn't for his golang expertise. Especially the pesky JSON parsing problem. I mean come on, this [gist](https://gist.github.com/akutz/8cd297d7f8d7bf5b4c43384e98a7f00b) is a work of art.

## Requires

The only thing this requires is mongoDB up and running already and PCF.

To run this on your laptop, I recommend using docker to spin up mongo (`docker run -dP mongo`), and PCF Dev (https://github.com/pivotal-cf/pcfdev).

It doesn't use user/password (so don't add one... this is not meant to be used in prod). Also, go-beer-me assumes that you attached mongodb as a user provided service. Specifically it looks for the env var `MONGO_LOCATION` which contains the `IP:Port` of where mongo runs. You can add as a user provided service (`cf cups mongo -p '{"MONGO_LOCATION":"192.168.99.100:32774"}'`).

## Running it

cd into the main directory of this repository, and run `cf push`

Note: the included `manifest.yml` file is read during the push, so the default name of this is `go-beer-me`.

Now you can use `curl` to use it

+ PUT beers in
  + `curl -XPOST -H 'Content-Type: application/json' -d '{"name": "Corona", "type": "Mexican", "ABV": 3.2, "cost":1.99}' go-beer-me.local.pcfdev.io/api/beer`
+ GET the beer you put in
  + `curl go-beer-me.local.pcfdev.io/api/beer/<id_of_beer>`
+ GET all beers
  + `curl go-beer-me.local.pcfdev.io/api/beer`
+ DELETE a beer from the app
  + `curl -XDELETE go-beer-me.local.pcfdev.io/api/beer/<id_of_beer>`
+ DELETE all beers from the app
  + `curl -XDELETE go-beer-me.local.pcfdev.io/api/beer`
