package controller

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/service"
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type RecordAPI struct {
	config        *config.Config
	logger        *logger.Logger
	recordService *service.RecordService
}

type FileUpload struct {
	CSVFile   *multipart.FileHeader `form:"file"`
	StationId string                `form:"station_id"`
}

func NewRecordAPI(c *config.Config, l *logger.Logger, s *service.RecordService) *RecordAPI {
	return &RecordAPI{c, l, s}
}

// Register godoc
//
//	@Summary		upload csv records
//	@Description	csv column - "sensorId"
//	@Tags			records
//	@Accept			json
//	@Produce		json
//	@Success		200	"return uploaded csv"
//	@Router			/records/upload [post]
func (api *RecordAPI) UploadRecords(g *gin.Context) {
	log := api.logger.Sugar()
	defer log.Sync()
	c := WrapContext(g)
	var csvFile FileUpload
	if err := c.ShouldBind(&csvFile); err != nil {
		log.Fatalln("failed to binding csv file: ", err)
		return
	}
	if csvFile.CSVFile == nil {
		log.Fatalln("csv file is missing")
		return
	}
	file, err := csvFile.CSVFile.Open()
	if err != nil {
		log.Fatalln("failed to open file: ", err)
		return
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	var sensorIds []string
	var records []service.Record
	// column name - name-unit-sensor-position-tag-group
	for index, line := range lines {
		for columnIndex, column := range line {
			if index == 0 {
				if columnIndex > 0 {
					sensorIds = append(sensorIds, column)
				}
			} else if columnIndex > 0 {
				time, err := time.Parse("2006/1/2 15:04:05", line[0])
				if err != nil {
					c.BadRequest(errors.New(fmt.Sprintf("failed to parse %s, error: %s", line[0], err.Error())))
					return
				}
				sensorIndex := columnIndex - 1
				if len(sensorIds) > sensorIndex {
					sensorId, err := strconv.ParseUint(sensorIds[sensorIndex], 10, 64)
					if err != nil {
						c.BadRequest(errors.New(fmt.Sprintf("failed to parse sensor id: %s, error: %s\n", sensorIds[sensorIndex], err.Error())))
						return
					}
					value, err := strconv.ParseFloat(column, 64)
					if err != nil {
						c.BadRequest(errors.New(fmt.Sprintf("failed to parse %s, error: %s\n", column, err.Error())))
						return
					}
					record := service.Record{
						SensorId: uint(sensorId),
						Value:    value,
						Time:     time,
					}
					records = append(records, record)
				}
			}
		}
	}
	updatedRecords, insertErr := api.recordService.BatchUpsert(&records)
	if insertErr == nil {
		c.OK(updatedRecords)
	} else {
		c.BadRequest(insertErr)
	}
}

// Register godoc
//
//	@Summary		download csv records
//	@Description	generate record upload csv template
//	@Tags			records
//	@Accept			json
//	@Produce		json
//	@Success		200	"export csv template"
//	@Router			/records/template [get]
func (api *RecordAPI) ExportCSVTemplate(g *gin.Context) {
	log := api.logger.Sugar()
	defer log.Sync()
	c := WrapContext(g)
	csvData := [][]string{
		{"数据时间(ignored)", "$Sensor_ID_1 (can be get from 'admin page / Sensors tab')", "$Sensor_ID_2 (can be get from 'admin page / Sensors tab')"},
		{"2023/4/1 15:30:00 (standard datetime format)", "(number only)", "(number only)"},
	}
	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	w.WriteAll(csvData)
	c.Writer.Write(b.Bytes())
}