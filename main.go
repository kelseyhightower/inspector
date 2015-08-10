package main

import (
	"log"
	"net/http"
	"text/template"
)

type indexHandler struct{}

var mainTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <link rel="stylesheet" href="css/bootstrap.min.css">
  </head>
  <div class="container">
  <body>
    <h1>Inspector</h1>
    <p>Version: {{.Version}}</p>
  </body>
  </div>
</html>
`

func (h indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("env").Parse(mainTemplate))
	type Data struct {
		Version string
	}
	err := t.Execute(w, Data{Version})
	if err != nil {
		log.Println("executing template:", err)
	}
}

func main() {
	log.SetFlags(0)
	http.Handle("/", NewRequestLoggerHandler(new(indexHandler)))
	http.Handle("/env", NewRequestLoggerHandler(new(envHandler)))
	http.Handle("/net", NewRequestLoggerHandler(new(networkHandler)))
	http.Handle("/healthz", NewRequestLoggerHandler(new(healthzHandler)))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	log.Println("Starting inspector...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
