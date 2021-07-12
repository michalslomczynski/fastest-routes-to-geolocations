package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	route "github.com/michalslomczynski/shortest-ways/OSRMConsumer"
)

type Response struct {
	Source route.Loc `json:"source"`
	Routes []Route   `json:"routes"`
}

type Route struct {
	Destination route.Loc `json:"destination"`
	Distance    float32   `json:"distance"`
	Duration    float32   `json:"duration"`
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	source := *readSource(w, r)
	destinations := readDestinations(w, r)

	var routes []Route

	for i, _ := range destinations {
		dest := []route.Loc{destinations[i]}
		route, err := route.GetRoute(source, dest)
		if err != nil {
			fmt.Fprintf(w, "Error occured during fetching remote API or invalid parameters provided.")
		}
		routes = append(routes, Route{destinations[i], route[0].Distance, route[0].Duration})

		sort.Slice(routes, func(i, j int) bool {
			return routes[i].Distance < routes[j].Distance
		})
	}
	response := &Response{source, routes}

	bJSON, err := json.Marshal(*response)
	if err != nil {
		fmt.Fprintf(w, "Internal error occured.")
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(bJSON))
}

func readSource(w http.ResponseWriter, r *http.Request) *route.Loc {
	src := r.URL.Query()["src"][0]
	loc, err := parseLoc(src)
	if err != nil {
		fmt.Fprintf(w, "Wrong source point provided.")
	}
	return loc
}

func readDestinations(w http.ResponseWriter, r *http.Request) []route.Loc {
	dst := r.URL.Query()["dst"]
	var dests []route.Loc
	for i, _ := range dst {
		loc, err := parseLoc(dst[i])
		if err != nil {
			fmt.Fprintf(w, "Wrong destination points provided.")
		}
		dests = append(dests, *loc)
	}
	return dests
}

func parseLoc(loc string) (*route.Loc, error) {
	split := strings.Split(loc, ",")

	lat, err := strconv.ParseFloat(split[0], 32)
	if err != nil {
		return nil, err
	}
	lon, err := strconv.ParseFloat(split[1], 32)
	if err != nil {
		return nil, err
	}
	return &route.Loc{float32(lat), float32(lon)}, nil
}
