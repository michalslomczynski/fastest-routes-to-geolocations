# shortest-ways-to-geolocations
## How to construct an API call
 Make a http GET request with the following format
`http://localhost:8086/routes?{src}&{dst}&{dst}...`
for example
`http://localhost:8086/routes?src=13.388860,52.517037&dst=13.428555,52.523219&dst=13.397634,52.529407`
response should be in JSON format, sorted from shortest to furthest location
`{"source":[13.38886,52.517036],"routes":[{"destination":[13.397634,52.529408],"distance":1884.9,"duration":251.4},{"destination":[13.428555,52.52322],"distance":3795.1,"duration":384.5}]}`
## How to run
### Docker
With docker compose `docker-compose build` and `docker-compose up`\
or\
with docker image `docker build -t shortest-ways-to-geolocations .` then `docker run -p 8086:8086 shortest-ways-to-geolocations`\
### Local build
cd to project and resolve gorilla dependency with `go get github.com/gorilla/mux`\
then run server with `go run main.go`
