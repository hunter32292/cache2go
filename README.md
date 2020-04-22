# cachego
Very simple in memory cache written in Go

[![Build Status](https://travis-ci.org/hunter32292/cachego.svg?branch=master)](https://travis-ci.org/hunter32292/cachego)

Current project goals:
  - Accept request via HTTP requests
  - Emmit stats via a metrics endpoint
  - write to disk and store in memory

# New Features!
  - REST endpoints for (Create,Show,Find,Delete)

### Tech
cache2go uses:

* [Mux](https://github.com/gorilla/mux) - A powerful URL router and dispatcher for golang.

### Installation
Dillinger requires [GoLang](https://golang.org/) v1.8+ to run.
Install the dependencies and start the server.
```sh
$ go get -v -u github.com/gorilla/mux
$ go build -o cachego src/hunter32292.github.com/cachego/*.go
$ ./cachego
```

### Development
Want to contribute? Great!
Feel free to submit an issue!
Feel free to pull down and submit a PR!

