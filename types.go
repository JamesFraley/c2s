package main

type place struct {
	Name      string   `json:"name"`
	Addr      address  `json:"address,omitempty"`
	FavColors []string `json:"favColors,string"`
}

type address struct {
	StreetAddr string `json:"street_address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Zipcode    int    `json:"zipcode"`
}
