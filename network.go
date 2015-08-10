package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"text/template"
)

type networkHandler struct{}

var networkTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <link rel="stylesheet" href="css/bootstrap.min.css">
  </head>
  <div class="container">
  <body>
  <h1>Network Info</h1>
  <h3>Interfaces</h3>
  <table class="table table-striped">
    <tr>
      <th>Name</th><th>HardwareAddr</th><th>IP</th>
    </tr>
    {{range $iface := .Interfaces}}<tr>
      <td>{{$iface.Name}}</td>
      <td>{{$iface.HardwareAddr}}</td>
      <td><ul class="list-unstyled">{{range $addr := $iface.Addrs}}
            <li>{{$addr.String}}</li>{{end}}
          </ul>
      </td>
    </tr>{{end}}
  <table>
  <h3>/etc/resolv.conf</h3>
  <pre>{{.ResolvConf}}</pre>
  </body>
  </div>
</html>

`

func (h networkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		ResolvConf string
		Interfaces []net.Interface
	}

	resolvConf, err := ioutil.ReadFile("/etc/resolv.conf")
	if err != nil {
		log.Println(err)
		return
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Println(err)
		return
	}

	data := Data{
		ResolvConf: string(resolvConf),
		Interfaces: ifaces,
	}

	t := template.Must(template.New("env").Parse(networkTemplate))
	err = t.Execute(w, data)
	if err != nil {
		log.Println("executing template:", err)
	}
}
