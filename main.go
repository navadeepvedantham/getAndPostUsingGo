package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var numbers []int

func main() {
	http.HandleFunc("/add", addNumber)
	http.HandleFunc("/get", getNumbers)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// This method supports both post and put methods.
func addNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	var number int
	err := json.NewDecoder(r.Body).Decode(&number)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	numbers = append(numbers, number)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Number %d added successfully", number)
}

func getNumbers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	// Using sortAscendingWithoutComparators for sorting the numbers
	sortedNumbers := sortAscendingWithoutComparators(numbers)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sortedNumbers)
}

func sortAscendingWithoutComparators(arr []int) []int {
	// Find the maximum element in the array
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	count := make([]int, max+1) // Used to store the count of each element.

	for _, val := range arr {
		count[val]++ // Update count of every element.
	}

	for i := 1; i <= max; i++ {
		count[i] += count[i-1] // Modify the count array to store the sum of counts
	}

	// Create a result array to store the sorted elements
	sorted := make([]int, len(arr))

	// Build the result array
	for _, val := range arr {
		sorted[count[val]-1] = val
		count[val]--
	}

	return sorted
}
