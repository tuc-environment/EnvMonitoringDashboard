package args

import "time"

type RecordsGetArgs struct {
	SensorIds *[]uint    `json:"sensor_ids"`
	Offset    *int       `json:"offset"`
	Limit     *int       `json:"limit"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	// v1.1
	RecordIndex uint64 `json:"record_index"`
}
