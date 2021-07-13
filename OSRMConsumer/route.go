package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/* API exposed at http://router.project-osrm.org/route
 * Docs http://project-osrm.org/docs/v5.5.1/api/
 */
type RouteResponse struct {
	Code   string  `json:"code"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Distance float32 `json:"distance"`
	Duration float32 `json:"duration"`
}

type Loc [2]float32

func GetRoute(src Loc, dstPts []Loc) ([]Route, error) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%f,%f", src[0], src[1]))
	for i, _ := range dstPts {
		sb.WriteString(fmt.Sprintf(";%f,%f", dstPts[i][0], dstPts[i][1]))
	}

	resp, err := http.Get(fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%s?overview=false", sb.String()))
	if err != nil {
		return nil, err
	}

	respJSON, error := readRoute(resp)
	if error != nil {
		return nil, err
	}

	if respJSON.Code != "Ok" || resp.StatusCode != 200 {
		return nil, err
	}

	return respJSON.Routes, nil
}

func readRoute(r *http.Response) (*RouteResponse, error) {
	response := &RouteResponse{}
	err := json.NewDecoder(r.Body).Decode(response)
	return response, err
}
