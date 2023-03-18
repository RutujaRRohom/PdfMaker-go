package route

import (
	"net/http"

	"github.com/RohomRutuja/Employee/PDF/service"
	"github.com/gorilla/mux"
)

func InitRouter(dp *dependencies) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/getEmployee", service.GetEmployeeHandler(dp.emp)).Methods(http.MethodGet)

	// router.HandleFunc("/POST/websites", CheckStatus.AddWebsitesHandler(dp.httpchecker)).Methods(http.MethodPost)
	// router.HandleFunc("/GET/websites", CheckStatus.GetWebsitesHandler(dp.httpchecker)).Methods(http.MethodGet)

	return router
}
