package customer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"phonenumbers_backend/helper"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET ALL")
	users := GetCustomers()
	response, _ := json.Marshal(users)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	user := GetCustomer(id)
	response, _ := json.Marshal(user)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	user := DeleteCustomer(id)
	response, _ := json.Marshal(user)
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT")
	var user Customer
	_ = json.NewDecoder(r.Body).Decode(&user)
	updatedCustomer := UpdateCustomer(user)
	response, _ := json.Marshal(updatedCustomer)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	var user Customer
	_ = json.NewDecoder(r.Body).Decode(&user)
	newCustomer := AddCustomer(user)
	response, _ := json.Marshal(newCustomer)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
