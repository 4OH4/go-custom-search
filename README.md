# go-custom-search
Docker container for running Google Custom Search queries - written in Go

This repository is a template Go API running in a Docker container, built using the Gin web framework. The demo code integrates with Google Cloud to request data from the Custom Search API, parse the results and return the search result URLs.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites


### Installing/building from scratch

Running a development environment - native Golang installation, tested on Windows 10:

```
$ git clone https://github.com/4OH4/go-custom-search
$ cd go-custom-search
$ go build
$ go-custom-search
```

## Running the tests

No tests have been written yet.

## Deployment

Production use, or at least use inside the Docker container:

```
docker build . -t go-custom-search
docker run -p 3000:3000 go-custom-search
```

## Built With

* [Gin](https://github.com/gin-gonic/gin) - Golang API web framework

## Contributing

Pull requests accepted.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Afdol Riski: [Complete Guide to Create Docker Container for Your Golang Application](https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e)
* https://github.com/afdolriski/golang-docker
