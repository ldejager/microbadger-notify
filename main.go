package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/kelseyhightower/envconfig"

	try "gopkg.in/matryer/try.v1"
)

type Microbadger struct {
	Repository string
	Token      string
}

func ConstructedURL() string {

	var args Microbadger
	err := envconfig.Process("mb", &args)

	if err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprintf("https://hooks.microbadger.com/images/%s/%s", args.Repository, args.Token)
}

func main() {

	if os.Getenv("MB_REPOSITORY") == "" {
		log.Fatal("You need to provide the repository for Microbadger\n\nUsage: microbadger-notify <repository> <token>")
		os.Exit(1)
	}

	if os.Getenv("MB_TOKEN") == "" {
		log.Fatal("You need to provide the repository for Microbadger\n\nUsage: microbadger-notify <repository> <token>")
		os.Exit(1)
	}

	log.Println("Notifying Microbadger...")

	endpoint := fmt.Sprintf(ConstructedURL())
	empty := url.Values{}
	values, _ := xml.Marshal(empty)
	data := bytes.NewBuffer(values)

	var resp *http.Response
	err := try.Do(func(attempt int) (bool, error) {
		var err error
		resp, err = http.Post(endpoint, "application/json", data)
		return attempt < 2, err
	})

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	if (resp.StatusCode) != 200 {
		log.Fatal("An error occured, check the repository and token provided")
		os.Exit(1)
	}

	resp.Body.Close()
	log.Println("Microbadger notified!")
}
