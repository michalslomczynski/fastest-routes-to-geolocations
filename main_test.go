package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	route "github.com/michalslomczynski/shortest-ways/OSRMConsumer"
)

var service string = "http://localhost:8086/routes"

var src = route.Loc{12, 12.5}
var dst1 = route.Loc{12.35823, 13.4721}
var dst2 = route.Loc{12.12394, 12.9756}
var dst3 = route.Loc{15.32763, 11.2376}

// Tests status code and content type
func TestValidQuery(t *testing.T) {
	resp, err := http.Get(
		fmt.Sprintf(
			"%v?src=%v,%v&dst=%v,%v&dst=%v,%v&dst=%v,%v",
			service, src[0], src[1], dst1[0], dst1[1], dst2[0], dst2[1], dst3[0], dst3[1],
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Want status code to be 200, got %v", resp.StatusCode)
	}

	if resp.Header["Content-Type"][0] != "application/json" {
		t.Fatalf("Want content-type to be application/json, got %v", resp.Header["Content-Type"][0])
	}
}

func TestNoDestinations(t *testing.T) {
	resp, err := http.Get(
		fmt.Sprintf(
			"%v?src=%v,%v",
			service, src[0], src[1],
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 400 {
		t.Fatalf("Want status code to be 400, got %v", resp.StatusCode)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	if string(bodyBytes) != "No destination points provided.\n" {
		t.Fatalf("Want response to be 'No destination points provided.', but was %v.", string(bodyBytes))
	}
}

func TestNoSourcePoint(t *testing.T) {
	resp, err := http.Get(
		fmt.Sprintf(
			"%v?dst=%v,%v",
			service, dst1[0], dst1[1],
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 400 {
		t.Fatalf("Want status code to be 400, got %v", resp.StatusCode)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	if string(bodyBytes) != "No source point provided.\n" {
		t.Fatalf("Want response to be 'No source point provided.', but was '%v'", string(bodyBytes))
	}
}
