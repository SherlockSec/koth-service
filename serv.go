package main

import (
	"net/http"
)

func serv() {

	mux := http.NewServeMux() // http router, handles the different endpoints
	mux.HandleFunc("/", returnKing)	// mux.HandleFunc("endpoint", function)
	mux.HandleFunc("/api/get", returnFlags)	// Neither API endpoint should respond to post requests unless /api/delete hasn't been used yet
	mux.HandleFunc("/api/delete", handleDelete) // ^^
	http.ListenAndServe(":9999", mux) // Serves on port 9999

}

func returnKing(w http.ResponseWriter, r *http.Request) { 

	w.Write(readKing())	// Standard port 9999 get request response credit to NinjaJc01 for that code.

}

func returnFlags(w http.ResponseWriter, r *http.Request) {

	flagArrayChan := make(chan []string) // Create callback channel to get array of flags according to the flag map

	go getFlagArray(flagArrayChan) // Get the flags, wait for the response

	flagArray := <-flagArrayChan // Read results from channel

	switch r.Method{ // Different results for different Request Methods
	case "GET":	// GET - returns flag.txt
		w.Write(readKing())
	case "POST": // POST - If /api/delete hasn't been run yet, supply flags as JSON, if /api/delete has been run, "Status Not Implemented" response.
		if isMapDeleted == true {

			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(http.StatusText(http.StatusNotImplemented) + "\n"))
			break

		}

		json := make(chan []byte) // Open callback channel

		go packFlags(&flagArray, json) // Pack the flags into JSON (packFlags is in json.go)

		response := <-json // Read the json from the channel

		w.Write(response)
	default: // Literally any other request type - "Status Not Implemented" response.
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented) + "\n"))

	}

}

func handleDelete(w http.ResponseWriter, r *http.Request) {

	switch r.Method{ // See above function for an explanation on how the Switch:Case works
	case "GET":
		w.Write(readKing())
	case "POST":
		if isMapDeleted == true {
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(http.StatusText(http.StatusNotImplemented) + "\n"))
			break
		} else {

			statusChan := make(chan bool) // Opens response channel

			go deleteMap(statusChan) // Delete the flag map from the box (king.go)

			status := <- statusChan	// Read status of file deletion from channel

			json := make(chan []byte) // Make JSON return channel

			go packDelete(status, json) // Pack map deletion status into JSON for response

			response := <-json // Read response JSON from return channel
			w.Write(response)
		}	
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented) + "\n"))
	}

}