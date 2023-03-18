package service

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetEmployeeHandler(emp Employe) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//var num string
		num := req.URL.Query().Get("numEmp")

		numEmployee, err := strconv.Atoi(num)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		employee, err := emp.GenerateEmployee(req.Context(), numEmployee)
		if err != nil {
			http.Error(w, "error in generating random employees", http.StatusInternalServerError)
		}
		json_response, err := json.Marshal(employee)
		if err != nil {
			http.Error(w, "error in marshelling", http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(json_response)

	})
}
