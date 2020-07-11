# go-custom-search
Docker container for running Google Custom Search queries - written in Go

This repository is a template Go API running in a Docker container, built using the Gin web framework. The demo code integrates with Google Cloud to request data from the Custom Search API, parse the results and return the search result URLs.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

The API is querying [Google Custom Search API](https://developers.google.com/custom-search/v1/introduction). To use it, you must have a Google Cloud account and create a Custom Search Engine using the [control panel](https://cse.google.com/cse/all). You can use this to create a custom search engine that only references a certain group of websites, or to search the [entire web](https://support.google.com/programmable-search/answer/4513886?hl=en&visit_id=637301049785357027-1764750062&rd=1) with high-throuput (paid-for) access to Google Search.

Your API key and Custom Search Engine ID must be configured in the `.env` credentials file before use.

### Installing/building from scratch

Running a development environment - native Golang installation, tested on Windows 10:

```
$ git clone https://github.com/4OH4/go-custom-search
$ cd go-custom-search
$ go mod download
$ go build
$ go-custom-search
```

## Running the tests

No tests have been written yet.

## Deployment

For production use, or at least use inside a Docker container. To run via Docker Compose:

```
docker-compose up
```

Or to build and run the image directly:

```
docker build . -t go-custom-search
docker run -p 3000:3000 go-custom-search
```

## API details

**go-custom-search**

* **URL**

  `/search`

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `query=[string]`

   **Optional:**
 
   None

* **Success Response:**
  
  * **Code:** 200 <br />
    **Content:** `{ links: ["https://www.domain1.com/url1", "https://www.domain2.com/url2"] }`
 
* **Error Response:**

  * **Code:** 403 FORBIDDEN <br />
    **Content:** `{ error : "The request is missing a valid API key." }`
	
    Check the API key is configured correctly in the `.env` file.

  OR

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "Requested entity was not found." }`

    Check the CSE ID is configured correctly in the `.env` file.
	
* **Sample Call:**

    `curl localhost:3000/search?query=golang`

## Built With

* [Gin](https://github.com/gin-gonic/gin) - Golang API web framework

## Contributing

Pull requests accepted.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Afdol Riski: [Complete Guide to Create Docker Container for Your Golang Application](https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e)
* https://github.com/afdolriski/golang-docker
