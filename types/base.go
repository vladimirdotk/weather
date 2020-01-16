package types

type Weather struct {
	ID       int     `json:"id,omitempty"`
	Lat      string  `json:"lat,omitempty"`
	Long     string  `json:"long,omitempty"`
	Temp     float32 `json:"temp,omitempty"`
	Humidity float32 `json:"humidity,omitempty"`
	Pressure float32 `json:"pressure,omitempty"`
}

type ApiResponse struct {
	Count int            `json:"count"`
	Data  []WeaherObject `json:"data"`
}

type WeaherObject struct {
	Temp     float32 `json:"temp"`
	Humidity float32 `json:"rh"`
	Pressure float32 `json:"pres"`
}
