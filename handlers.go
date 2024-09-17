package main 

import (
	"encoding/json"
	"net/http"
)

func handleClientProfile(w http.ResponseWriter, r *http.Request){
	// Router to another function depending on the request method
	// Will allow get and patch requests(patch will allow the user profile to update)

	switch r.Method{
	case http.MethodGet:
		GetClientProfile(w,r)
	case http.MethodPatch:
		UpdateClientProfile(w,r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetClientProfile(w http.ResponseWriter, r *http.Request){
	clientProfile := r.Context().Value("clientProfile").(ClientProfile) // the last part is casting to the correct struct type

	response := ClientProfile{
		Email:clientProfile.Email,
		Name:clientProfile.Name,
		Id:clientProfile.Id,
	}
	json.NewEncoder(w).Encode(response)
	
}

func UpdateClientProfile(w http.ResponseWriter, r *http.Request){
	clientProfile := r.Context().Value("clientProfile").(ClientProfile) 

	// Decode the JSON payload directly into the struct
	var payloadData ClientProfile 
	if err := json.NewDecoder(r.Body).Decode(&payloadData); 
	err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Close the connection between the client and the server
	// defer = execute before this function closes no matter what
	defer r.Body.Close() 

	clientProfile.Email = payloadData.Email
	clientProfile.Name = payloadData.Name
	database[clientProfile.Id] = clientProfile

	w.WriteHeader(http.StatusOK)
}