# go-web-app
First go web app

## Clone repository

```bash
git clone git@github.com:felixbouazza/go-web-app.git
```

## Runserver

```bash
go run .
```

## Add new Route

Create template on templates directory, and add new function on `handlers.go` like this

```go
func NewFunctionName(response http.ResponseWriter, request *http.Request){
	renderTemplate(response, "new_route_template_name")
}
```

Add function on `main.go`

```go
http.HandleFunc("/new-route-name", NewFunctionName)
```
