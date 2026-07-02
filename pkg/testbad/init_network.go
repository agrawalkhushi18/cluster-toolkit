package testbad

import (
	"fmt"
	"net/http"
)

// Global client initialized eagerly in init(), violating Sage Go guidelines.
var globalResponse *http.Response

func init() {
	// Bad: Executing network calls during package initialization blocks CLI startup.
	resp, err := http.Get("https://example.com/config.json")
	if err != nil {
		fmt.Println("Error fetching config:", err)
	}
	globalResponse = resp
}

func ProcessData() {
	if globalResponse != nil {
		// Bad: Deferring Body.Close() without verifying response error or checking nil properly
		defer globalResponse.Body.Close()
	}
}
