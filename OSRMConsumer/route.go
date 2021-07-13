package route

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* API exposed at http://router.project-osrm.org/route
 * Docs http://project-osrm.org/docs/v5.5.1/api/
 */
type RouteResponse struct {
	Code  string  `json:"code"`
	Route []Route `json:"routes"`
}

type Route struct {
	Distance float32 `json:"distance"`
	Duration float32 `json:"duration"`
}

type Loc [2]float32

// Returns fastest route between source and destination
func GetRoute(src Loc, dst Loc) (Route, error) {
	resp, err := http.Get(fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=false", src[0], src[1], dst[0], dst[1]))
	if err != nil {
		return Route{}, err
	}

	respJSON, error := readRoute(resp)
	if error != nil {
		return Route{}, err
	}

	if respJSON.Code != "Ok" || resp.StatusCode != 200 {
		return Route{}, err
	}

	return respJSON.Route[0], nil
}

// Deserializes response body
func readRoute(r *http.Response) (*RouteResponse, error) {
	response := &RouteResponse{}
	err := json.NewDecoder(r.Body).Decode(response)
	return response, err
}
