package route

import "github.com/RohomRutuja/Employee/PDF/service"

type dependencies struct {
	//httpchecker CheckStatus.WebsiteChecker
	emp service.Employe
}

// InitDependencies initializes the router with the dependencies

func InitDependencies() *dependencies {
	return &dependencies{
		emp: service.NewFunc(),
	}
}
