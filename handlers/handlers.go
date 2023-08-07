package handlers

import (
	"net/http"
	"github.com/gwoodwa1/vlan-db/types"
	"github.com/gwoodwa1/vlan-db/db"
	"encoding/json"
)

func GetVlanHandler(w http.ResponseWriter, r *http.Request) {
	// Prepare the SQL query
	rows, err := db.Db.Query("SELECT id, name, description FROM vlan")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Iterate through the result set and put the values into a slice of Vlan
	var vlans []types.Vlan
	for rows.Next() {
		var v types.Vlan
		err := rows.Scan(&v.ID, &v.Name, &v.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		vlans = append(vlans, v)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the slice to JSON
	vlanJson, err := json.Marshal(vlans)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON to the response body
	w.Header().Set("Content-Type", "application/json")
	w.Write(vlanJson)
}

func AddNewVlan(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var newVlan types.Vlan
	err := json.NewDecoder(r.Body).Decode(&newVlan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the VLAN ID is reserved
	if isReserved(newVlan.ID) {
		http.Error(w, "VLAN ID is reserved", http.StatusForbidden)
		return
	}

	// Insert the new VLAN into the database
	_, err = db.Db.Exec(`INSERT INTO vlan (id, name, description) VALUES (?, ?, ?)`,
		newVlan.ID, newVlan.Name, newVlan.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.Write([]byte("New VLAN added successfully"))
}

func isReserved(id int) bool {
	for _, v := range types.Reserved.IDs {
		if v == id {
			return true
		}
	}
	return false
}
