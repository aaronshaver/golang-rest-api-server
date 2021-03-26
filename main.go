package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
	"github.com/gorilla/mux"
	"strconv"
)

var numbers []int

func getNumbers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET params were:", r.URL.Query())

	function := r.URL.Query().Get("function")
	if function == "average" {
		sum := 0
		for i := 0; i < len(numbers); i++ {
			sum += (numbers[i])
		}
		average := (float64(sum) / (float64(len(numbers))))

		fmt.Fprintf(w, "Average: %v", average)
	} else {
		fmt.Fprintf(w, "%v", numbers)
	}
}

func postNumber(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	
	i, err := strconv.Atoi(string(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	numbers = append(numbers, i)
	fmt.Fprintf(w, "Number appended")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/numbers", getNumbers).Methods("GET")
	myRouter.HandleFunc("/numbers", postNumber).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	numbers = append(numbers, 1, 2, 3, 4)
	handleRequests()
}
