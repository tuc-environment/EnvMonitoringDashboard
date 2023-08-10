package controller

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/service"
	"encoding/csv"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type RecordAPI struct {
	config        *config.Config
	logger        *logger.Logger
	recordService *service.RecordService
}

type RecordTemplate struct {
	Name     string
	Unit     string
	Sensor   string
	Position service.RecordPosition
	Tag      string
	Group    string
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
//	@Description	csv column - "name-unit-sensor-position-tag-group"
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
	stationId := csvFile.StationId
	file, err := csvFile.CSVFile.Open()
	if err != nil {
		log.Fatalln("failed to open file: ", err)
		return
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	var templates []RecordTemplate
	var records []service.Record
	// column name - name-unit-sensor-position-tag-group
	for index, line := range lines {
		for columnIndex, column := range line {
			if index == 0 {
				log.Debugln("line: ", line)
				if columnIndex > 0 {
					if strings.Contains(column, "-") {
						template, ok := api.processTemplate(column)
						if ok {
							templates = append(templates, *template)
						} else {
							c.BadRequest(errors.New("invalid column format"))
							return
						}
					} else {
						templates = append(templates, RecordTemplate{Name: column})
					}
				}
			} else if columnIndex > 0 {
				time, err := time.Parse("2006/1/2 15:04:05", line[0])
				if err != nil {
					c.BadRequest(errors.New(fmt.Sprintf("failed to parse %s, error: %s", line[0], err.Error())))
					return
				}
				templateIndex := index - 1
				if len(templates) > templateIndex {
					template := templates[templateIndex]
					value, err := strconv.ParseFloat(column, 64)
					if err != nil {
						c.BadRequest(errors.New(fmt.Sprintf("failed to parse %s, error: %s\n", line[columnIndex], err.Error())))
						return
					}
					record := service.Record{
						StationId: stationId,
						Position:  template.Position,
						Tag:       template.Tag,
						Name:      template.Name,
						Sensor:    template.Sensor,
						Group:     template.Group,
						Value:     value,
						Unit:      template.Unit,
						Time:      time,
					}
					records = append(records, record)
				}
			}
		}
	}
	c.OK(records)
}

func (api *RecordAPI) processTemplate(str string) (*RecordTemplate, bool) {
	items := strings.Split(str, "-")
	if len(items) < 6 {
		return nil, false
	} else {
		position, _ := api.parsePosition(items[3])
		name := items[0]
		unit := items[1]
		sensor := items[2]
		tag := items[4]
		group := items[5]
		template := &RecordTemplate{
			Name:     name,
			Unit:     unit,
			Sensor:   sensor,
			Position: position,
			Tag:      tag,
			Group:    group,
		}
		return template, true
	}
}

func (api *RecordAPI) parsePosition(str string) (service.RecordPosition, bool) {
	c, ok := service.RecordPositionMap[strings.ToLower(str)]
	return c, ok
}
