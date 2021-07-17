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

// Succesfull HTTP response content
type Response struct {
	Source route.Loc `json:"source"`
	Routes []Route   `json:"routes"`
}

type Route struct {
	Destination route.Loc `json:"destination"`
	Distance    float32   `json:"distance"`
	Duration    float32   `json:"duration"`
}

// Sends back JSON with routes between source and each destination point
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	source := readSource(w, r)
	destinations := readDestinations(w, r)

	if (source == route.Loc{} || len(destinations) == 0) {
		return
	}
	var routes []Route

	for _, dest := range destinations {
		route, err := route.GetRoute(source, dest)
		if err != nil {
			http.Error(w, "Error occured during fetching remote API or invalid parameters provided.", http.StatusBadRequest)
			return
		}
		routes = append(routes, Route{dest, route.Distance, route.Duration})
	}
	sortRoutes(routes)
	response := &Response{source, routes}

	bJSON, err := json.Marshal(*response)
	if err != nil {
		http.Error(w, "Internal error occured", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", string(bJSON))
}

// Parses source point from query and returns it as location(float32, float32)
func readSource(w http.ResponseWriter, r *http.Request) route.Loc {
	if len(r.URL.Query()["src"]) == 0 {
		http.Error(w, "No source point provided.", http.StatusBadRequest)
		return route.Loc{}
	}
	src := r.URL.Query()["src"][0]

	loc, err := parseLoc(src)
	if err != nil {
		http.Error(w, "Invalid source point provided.", http.StatusBadRequest)
		return route.Loc{}
	}
	return *loc
}

// Parses destination points from query and returns []locations(float32, float32)
func readDestinations(w http.ResponseWriter, r *http.Request) []route.Loc {
	dst := r.URL.Query()["dst"]
	if len(dst) == 0 {
		http.Error(w, "No destination points provided.", http.StatusBadRequest)
		return nil
	}
	var dests []route.Loc

	for i, _ := range dst {
		loc, err := parseLoc(dst[i])
		if err != nil {
			http.Error(w, "Invalid destination points provided.", http.StatusBadRequest)
			return nil
		}
		dests = append(dests, *loc)
	}
	return dests
}

// Takes string like "12.32387,13.478623" and returns location(float32, float32)
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

// Sorts routes by duration, if its equal then by distance
func sortRoutes(routes []Route) {
	sort.Slice(routes, func(i, j int) bool {
		if routes[i].Duration == routes[j].Duration {
			return routes[i].Distance < routes[j].Distance
		}
		return routes[i].Duration < routes[j].Duration
	})
}
