package args

// type SensorCreationArgs struct {
// 	ID        uint   `example:"1"`
// 	StationId uint   `example:"1" json:"station_id"`
// 	Position  string `example:"up"`
// 	Tag       string `example:"5cm"`
// 	Name      string `example:"temperature"`
// 	Group     string `example:"A"`
// 	Unit      string `example:"℃"`
// 	// v1.1
// 	IMEI       string `example:"A" json:"imei"`
// 	SensorCode string `example:"A" json:"sensor_code"`
// 	// imei-sensorType-sensorCode-index
// 	SensorReportCode string `example:"A" json:"sensor_report_code"`
// 	SampleMethod     string `example:"A" json:"sample_method"`
// }

type SensorGetArgs struct {
	ID        uint `example:"1"`
	StationId uint `example:"1" json:"station_id"`
}
