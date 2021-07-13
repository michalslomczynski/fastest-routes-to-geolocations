package handler

import (
	"testing"

	route "github.com/michalslomczynski/shortest-ways/OSRMConsumer"
)

var loc = route.Loc{0, 0}
var route1 = Route{loc, 2583, 208}
var route2 = Route{loc, 1837, 208}
var route3 = Route{loc, 183, 208}
var route4 = Route{loc, 30293, 98}
var route5 = Route{loc, 1687, 208}

func TestRoutesSorting(t *testing.T) {
	var routes = []Route{route1, route2, route3, route4, route5}
	sortRoutes(routes)

	for i := 0; i < len(routes)-1; i++ {

		if routes[i].Duration > routes[i+1].Duration {
			t.Fatalf("Expected that routes durations are ascending but they were not: %v", routes)
		} else if routes[i].Duration == routes[i+1].Duration {
			if routes[i].Distance > routes[i+1].Distance {
				t.Fatalf("Expected that routes with same duration will be sorted ascending by distance but they were not: %v", routes)
			}
		}
	}
}
