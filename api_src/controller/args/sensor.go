package args

type SensorGetArgs struct {
	ID                 uint `example:"1"`
	StationId          uint `example:"1" json:"station_id"`
	VisibleInDashboard bool `example:"true"`
}
