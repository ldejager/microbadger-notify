package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/kelseyhightower/envconfig"

	try "gopkg.in/matryer/try.v1"
)

type microbadger struct {
	Repository string
	Token      string
}

func constructedURL() string {

	var args microbadger
	err := envconfig.Process("mb", &args)

	if err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprintf("https://hooks.microbadger.com/images/%s/%s", args.Repository, args.Token)
}

func usage() string {
	return fmt.Sprintf("\nUsage: microbadger-notify <repository> <token>")
}

func main() {

	if os.Getenv("MB_REPOSITORY") == "" {
		log.Println("You need to provide the repository for Microbadger", usage())
		os.Exit(1)
	}

	if os.Getenv("MB_TOKEN") == "" {
		log.Println("You need to provide the token for Microbadger", usage())
		os.Exit(1)
	}

	endpoint := fmt.Sprintf(constructedURL())
	empty := url.Values{}
	values, _ := json.Marshal(empty)
	data := bytes.NewBuffer(values)

	var resp *http.Response
	err := try.Do(func(attempt int) (bool, error) {
		var err error
		resp, err = http.Post(endpoint, "application/json", data)
		return attempt < 2, err
	})

	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	} else if (resp.StatusCode) != 200 {
		log.Println("Received unexpected response code:", resp.StatusCode)
		log.Println("Attempted POST URL: ", endpoint)
		os.Exit(1)
	} else {
		resp.Body.Close()
		log.Println("Microbadger successfully notified.")
	}
}
