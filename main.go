package main

import(
	"sync"
)

//Constants
const kingPath = "king.txt" // Path to king file
const mapPath = "map.txt"   // Path to map file
const flags = 4             // amount of flags
var isMapDeleted = false    // Initial assignment of the /api/delete check bool

func main() {

	var wg sync.WaitGroup // wg is the WaitGroup, keeps the concurrent process running, otherwise the program would exit without running the server.
	wg.Add(1)

	go func() {
		serv() // serv() is in serv.go
		wg.Done()
	}()

	wg.Wait()

}

