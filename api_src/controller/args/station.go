package args

type UpsertStationArgs struct {
	ID       uint    `json:"id" example:"1"`
	Name     string  `json:"name" example:"Station Name"`
	Lat      float64 `json:"lat" example:"10.13"`
	Lng      float64 `json:"lng" example:"11.11"`
	Altitude float64 `json:"altitude" example:"11100"`
}
