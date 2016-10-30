FROM scratch
MAINTAINER Ryan Clair <ryan.clair@gmail.com>
EXPOSE 3000
COPY go-beer-me-linux /go-beer-me
ENTRYPOINT ["/go-beer-me"]
