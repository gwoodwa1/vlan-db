package main

import (
	"log"
	"net/http"
	"github.com/gwoodwa1/vlan-db/handlers"
	"github.com/gwoodwa1/vlan-db/db"
)




func main() {
	// Initialize the sqlite database
	err := db.InitDB("./vlan.db")
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	// Create the VLAN table if it doesn't already exist
	_, err = db.Db.Exec(`CREATE TABLE IF NOT EXISTS vlan (
		id INTEGER NOT NULL PRIMARY KEY, 
		name TEXT, 
		description TEXT);`)
	if err != nil {
		log.Fatalf("Could not create vlan table: %v", err)
	}

	// Define the routes for fetching the list of VLANs and adding new ones
	http.HandleFunc("/vlans", handlers.GetVlanHandler)
	http.HandleFunc("/addnew", handlers.AddNewVlan)
	// Serve the static files from the "static" folder
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Start the HTTP server on port 8080
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

