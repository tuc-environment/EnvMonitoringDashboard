package args

type SensorCreationArgs struct {
	ID        uint   `example:"1"`
	StationId uint   `example:"1" json:"station_id"`
	Position  string `example:"up"`
	Tag       string `example:"5cm"`
	Name      string `example:"temperature"`
	Group     string `example:"A"`
	Unit      string `example:"℃"`
}

type SensorGetArgs struct {
	ID        uint `example:"1"`
	StationId uint `example:"1" json:"station_id"`
}
