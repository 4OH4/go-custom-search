// Copyright 2018 Google LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// From: https://github.com/googleapis/google-api-go-client/blob/master/examples/customsearch.go

// Request format: https://developers.google.com/custom-search/v1/reference/rest/v1/cse/list
// Response format: https://developers.google.com/custom-search/v1/reference/rest/v1/Search

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
)

// const (
// 	apiKey = "some-api-key"
// 	cx     = "some-custom-search-engine-id"
// )

func customSearchMain(query string) {

	// Get credentials from environment variables - with docker-compose, set in credentials.env
	apiKey := os.Getenv("CSE_API_KEY")
	cx := os.Getenv("CSE_ID")

	client := &http.Client{Transport: &transport.APIKey{Key: apiKey}}

	svc, err := customsearch.New(client)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := svc.Cse.List().Cx(cx).Q(query).Do()
	if err != nil {
		log.Fatal(err)
	}

	for i, result := range resp.Items {
		fmt.Printf("#%d: %s\n", i+1, result.Title)
		fmt.Printf("\t%s\n", result.Snippet)
	}
}
