package main

type place struct {
	Name      string   `json:"name"`
	Addr      address  `json:"address,omitempty"`
	Point     point    `json:"point"`
	FavColors []string `json:"favColors,string,omitempty"`
}

type address struct {
	StreetAddr string `json:"street_address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Zipcode    int    `json:"zipcode"`
}

type point struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
