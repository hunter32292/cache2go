# cache2go
Very simple in memory cache written in Go

[![Build Status](https://travis-ci.org/hunter32292/cache2go.svg?branch=master)](https://travis-ci.org/hunter32292/cache2go)

Current project goals:
  - Accept request via HTTP requests
  - Emit stats via a metrics endpoint
  - Write to disk and store in memory

# New Features!
  - REST endpoints for (Create,Show,Find,Delete)

### Tech
cache2go uses:

* [Mux](https://github.com/gorilla/mux) - A powerful URL router and dispatcher for golang.
* [Viper](https://github.com/spf13/viper) - Go Configuration with Fangs
* [Cobra](https://github.com/spf13/cobra) - Go Commander for modern Go CLI Interactions

### Installation
Dillinger requires [GoLang](https://golang.org/) v1.8+ to run.
Install the dependencies and start the server.
```sh
$ go get -v -u github.com/gorilla/mux
$ go build -o cache2go .
$ ./cache2go
```

### Development
Want to contribute? Great!
Feel free to submit an issue!
Feel free to pull down and submit a PR!

