package timev2api

import "encoding/json"

// To save client data
type ClientsReturn struct {
	Total   int `json:"total"`
	Clients []ClientDataReturn
}

type ClientDataReturn struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Rate    string `json:"rate"`
	SevDesk string `json:"sevdesk,omitempty"`
}

// Get a list of all clients
func Clients(domainkey, token string) (ClientsReturn, error) {

	// Response to timev2
	response, err := response("clients", domainkey, token)
	if err != nil {
		return ClientsReturn{}, err
	}

	// Save clients
	var decode ClientsReturn

	// Decode json code in struct
	err = json.NewDecoder(response).Decode(&decode)
	if err != nil {
		return ClientsReturn{}, err
	}

	// Return data
	return decode, nil

}
