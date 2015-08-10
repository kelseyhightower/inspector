package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type envHandler struct{}

var envTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <link rel="stylesheet" href="css/bootstrap.min.css">
  </head>
  <div class="container">
  <body>
  <h1>Environment Info</h1>
  <h3>Environment Variables</h3>
  <table class="table table-striped">
    <tr>
      <th>Key</th><th>Value</th>
    </tr>
    {{range $key, $value := .Environment}}<tr>
      <td>{{$key}}</td><td>{{$value}}</td>
    </tr>{{end}}
  <table>
  </body>
</html>
`

func (h envHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	envMap := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		envMap[pair[0]] = pair[1]
	}
	type EnvData struct {
		Environment map[string]string
	}
	data := EnvData{envMap}
	t := template.Must(template.New("env").Parse(envTemplate))
	err := t.Execute(w, data)
	if err != nil {
		log.Println("executing template:", err)
	}
}
