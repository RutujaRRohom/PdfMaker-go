package main

import (
	"fmt"
	"strconv"

	"github.com/RohomRutuja/Employee/PDF/route"

	"github.com/urfave/negroni"
)

func main() {
	dep := route.InitDependencies()
	router := route.InitRouter(dep)
	//go CheckStatus.GetStatus()

	service := negroni.Classic()
	service.UseHandler(router)
	port := 8080 // This can be changed to the service port number via environment variable.
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	service.Run(addr)

}
