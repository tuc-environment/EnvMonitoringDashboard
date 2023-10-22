package service

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/store"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type DeviceSensor struct {
	Name string
	Unit string
	Code string
}

type Device struct {
	Code string
	Name string
	Data []DeviceSensor
}

var sampleMethodMap = map[string]string{
	"0": "采样值",
	"1": "平均值",
	"2": "总计值",
	"3": "最大值",
	"4": "最小值",
}

var deviceMap = map[string]Device{
	"1": {
		Code: "sm926",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Soil Moisture 土壤含水量",
				Code: "soil_vwc",
				Unit: "m³/m³",
			},
			// 2
			{
				Name: "Soil Conductivity (temp.corrected) 土壤电导率(温度补偿)",
				Code: "soil_ec_tc",
				Unit: "Siemens/Meter",
			},
			// 3
			{
				Name: "Soil Temperature(℃) 土壤温度(℃)",
				Code: "soil_temp_c",
				Unit: "℃",
			},
			// 4
			{
				Name: "Soil Temperature(℉) 土壤温度(℉)",
				Code: "soil_temp_f",
				Unit: "℉",
			},
			// 5
			{
				Name: "Soil Conductivity 土壤电导率",
				Code: "soil_ec",
				Unit: "Siemens/Meter",
			},
			// 6
			{
				Name: "Real Dielectric Permittivity 土壤介电常数实部",
				Code: "soil_dc_r",
				Unit: "",
			},
			// 7
			{
				Name: "Imaginary Dielectric Permittivity 土壤介电常数虚部",
				Code: "soil_dc_i",
				Unit: "",
			},
			// 8
			{
				Name: "Real Dielectric Permittivity (temp.corrected) 土壤介电常数实部(温度补偿)",
				Code: "soil_dc_r_tc",
				Unit: "",
			},
			// 9
			{
				Name: "Imaginary Dielectric Permittivity (temp.Corrected) 土壤介电常数虚部(温度补偿)",
				Code: "soil_dc_i_tc",
				Unit: "",
			},
		},
	},
	"2": {
		Code: "sm826",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Volumetric Water Content 土壤含水量",
				Code: "soil_vwc",
				Unit: "m³/m³",
			},
			// 2
			{
				Name: "Electrical Conductivity 土壤电导率",
				Code: "soil_ec",
				Unit: "dS/m",
			},
			// 3
			{
				Name: "Temperature 土壤温度",
				Code: "soil_temp",
				Unit: "°C",
			},
			// 4
			{
				Name: "Permittivity 介电常数",
				Code: "soil_dc",
				Unit: "",
			},
			// 5
			{
				Name: "Period",
				Code: "period",
				Unit: "",
			},
			// 6
			{
				Name: "Voltage Ratio",
				Code: "voltage_ratio",
				Unit: "",
			},
		},
	},
	"3": {
		Code: "trb3",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Temperature 空气温度",
				Code: "temp",
				Unit: "°C",
			},
			// 2
			{
				Name: "Relative Humidity 空气湿度",
				Code: "humidity",
				Unit: "%RH",
			},
			// 3
			{
				Name: "Barometric Pressure 大气压力",
				Code: "barometric_pressure",
				Unit: "hPa",
			},
		},
	},
	"4": {
		Code: "irs101",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Object Temperature 目标温度",
				Code: "object_temp",
				Unit: "°C",
			},
			// 2
			{
				Name: "Ambient Temperature 环境湿度",
				Code: "ambient_temp",
				Unit: "°C",
			},
		},
	},
	"5": {
		Code: "rs_nr",
		Name: "四分量净辐射传感器",
		Data: []DeviceSensor{
			// 1
			{
				Name: "天空短波辐射",
				Code: "a",
				Unit: "W/㎡",
			},
			// 2
			{
				Name: "地球反射短波辐射",
				Code: "b",
				Unit: "W/㎡",
			},
			// 3
			{
				Name: "天空长波辐射",
				Code: "c",
				Unit: "W/㎡",
			},
			// 4
			{
				Name: "地球长波辐射",
				Code: "d",
				Unit: "W/㎡",
			},
			// 5
			{
				Name: "传感器温度",
				Code: "temp",
				Unit: "°C",
			},
			// 6
			{
				Name: "反照率",
				Code: "albedo",
				Unit: "%",
			},
			// 7
			{
				Name: "净辐射",
				Code: "net_radiation",
				Unit: "W/㎡",
			},
		},
	},
	"6": {
		Code: "lc101",
		Name: "蒸渗仪",
		Data: []DeviceSensor{
			// 1
			{
				Name: "重量",
				Code: "weight",
				Unit: "g",
			},
			// 2
			{
				Name: "测量电压比率*1000",
				Code: "voltage_ratio_1000",
				Unit: "",
			},
			// 3
			{
				Name: "秤体温度",
				Code: "ambient_temp",
				Unit: "°C",
			},
			// 4
			{
				Name: "渗漏计数",
				Code: "leakage",
				Unit: "mm",
			},
		},
	},
	"7": {
		Code: "rsg01",
		Name: "称重雨雪量计",
		Data: []DeviceSensor{
			// 1
			{
				Name: "1#重量",
				Code: "weight_1",
				Unit: "g",
			},
			// 2
			{
				Name: "1#测量电压比率*1000",
				Code: "voltage_ratio_1000_1",
				Unit: "",
			},
			// 3
			{
				Name: "2#重量g",
				Code: "weight_2",
				Unit: "g",
			},
			// 4
			{
				Name: "2#测量电压比率*1000",
				Code: "voltage_ratio_1000_2",
				Unit: "",
			},
			// 5
			{
				Name: "3#重量(单位：g)",
				Code: "weight_3",
				Unit: "g",
			},
			// 6
			{
				Name: "3#测量电压比率*1000",
				Code: "voltage_ratio_1000_3",
				Unit: "",
			},
			// 7
			{
				Name: "温度值",
				Code: "ambient_temp",
				Unit: "°C",
			},
		},
	},
	"8": {
		Code: "tr_asc",
		Name: "自动集沙仪",
		Data: []DeviceSensor{
			// 1
			{
				Name: "重量",
				Code: "weight",
				Unit: "g",
			},
			// 2
			{
				Name: "秤体温度",
				Code: "ambient_temp",
				Unit: "°C",
			},
			// 3
			{
				Name: "重量电压比率",
				Code: "voltage_ratio_1000",
				Unit: "",
			},
		},
	},
	"9": {
		Code: "dv105",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "差分1(通道)数值",
				Code: "volt_diff_0",
				Unit: "mV",
			},
			// 2
			{
				Name: "差分2(通道)数值",
				Code: "volt_diff_1",
				Unit: "mV",
			},
			// 3
			{
				Name: "差分3(通道)数值",
				Code: "volt_diff_2",
				Unit: "mV",
			},
			// 4
			{
				Name: "差分4(通道)数值",
				Code: "volt_diff_3",
				Unit: "mV",
			},
			// 5
			{
				Name: "差分5(通道)数值",
				Code: "volt_diff_4",
				Unit: "mV",
			},
		},
	},
	"A": {
		Code: "cs65x",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Volumetric Water Content 土壤含水量",
				Code: "soil_vwc",
				Unit: "m³/m³",
			},
			// 2
			{
				Name: "Electrical Conductivity 土壤电导率",
				Code: "soil_ec",
				Unit: "dS/m",
			},
			// 3
			{
				Name: "Temperature 土壤温度",
				Code: "soil_temp",
				Unit: "°C",
			},
			// 4
			{
				Name: "Permittivity 介电常数",
				Code: "soil_dc",
				Unit: "",
			},
			// 5
			{
				Name: "Period",
				Code: "period",
				Unit: "",
			},
			// 6
			{
				Name: "Voltage Ratio",
				Code: "voltage_ratio",
				Unit: "",
			},
		},
	},
	"B": {
		Code: "soilvue10_50",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Volumetric Water Content, 5 cm",
				Code: "soil_vwc",
				Unit: "m³/m³",
			},
		},
	},
	"C": {
		Code: "soilvue10_100",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Volumetric Water Content, 5 cm(m3/m3)",
				Code: "soil_vwc",
				Unit: "m³/m³",
			},
		},
	},
	"D": {
		Code: "hydraprobe",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Soil Moisture 土壤含水量",
				Code: "soil_vwc",
				Unit: "m³/m³",
			},
			// 2
			{
				Name: "Soil Conductivity (temp.corrected) 土壤电导率(温度补偿)",
				Code: "soil_ec_tc",
				Unit: "Siemens/Meter",
			},
			// 3
			{
				Name: "Soil Temperature(℃) 土壤温度℃",
				Code: "soil_temp_c",
				Unit: "℃",
			},
			// 4
			{
				Name: "Soil Temperature(℉) 土壤温度℉",
				Code: "soil_temp_f",
				Unit: "℉",
			},
			// 5
			{
				Name: "土壤电导率，Soil Conductivity",
				Code: "soil_ec",
				Unit: "Siemens/Meter",
			},
			// 6
			{
				Name: "Real Dielectric Permittivity 土壤介电常数实部",
				Code: "soil_dc_r",
				Unit: "",
			},
			// 7
			{
				Name: "Imaginary Dielectric Permittivity 土壤介电常数虚部",
				Code: "soil_dc_i",
				Unit: "",
			},
			// 8
			{
				Name: "Real Dielectric Permittivity (temp.corrected) 土壤介电常数实部(温度补偿)",
				Code: "soil_dc_r_tc",
				Unit: "",
			},
			// 9
			{
				Name: "Imaginary Dielectric Permittivity (temp.Corrected) 土壤介电常数虚部(温度补偿)",
				Code: "soil_dc_i_tc",
				Unit: "",
			},
		},
	},
	"E": {
		Code: "hydraprobelite",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Soil moisture for inorganic & mineral soil 土壤含水量",
				Code: "soil_vwc",
				Unit: "m³/m³",
			},
			// 2
			{
				Name: "Bulk electrical conductivity 土壤电导率",
				Code: "soil_ec",
				Unit: "Siemens/Meter",
			},
			// 3
			{
				Name: "Temperature 土壤温度",
				Code: "soil_temp",
				Unit: "℃",
			},
		},
	},
	"F": {
		Code: "decagon_5tm",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "dielectric permittivity 介电常数",
				Code: "soil_dc",
				Unit: "",
			},
			// 2
			{
				Name: "temperature 土壤温度",
				Code: "soil_temp",
				Unit: "℃",
			},
			// 3
			{
				Name: "moisture 土壤含水量",
				Code: "soil_vwc",
				Unit: "m³/m³",
			},
		},
	},
	"G": {
		Code: "decagon_5te",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "dielectric permittivity 介电常数",
				Code: "soil_dc",
				Unit: "",
			},
			// 2
			{
				Name: "electrical conductivity 土壤电导率",
				Code: "soil_ec",
				Unit: "dS/m",
			},
			// 3
			{
				Name: "temperature 土壤温度",
				Code: "soil_temp",
				Unit: "℃",
			},
			// 4
			{
				Name: "moisture 土壤含水量",
				Code: "soil_vwc",
				Unit: "m³/m³",
			},
		},
	},
	"H": {
		Code: "mt10",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "dielectric permittivity 介电常数",
				Code: "soil_dc",
				Unit: "",
			},
			// 2
			{
				Name: "temperature 土壤温度",
				Code: "soil_temp",
				Unit: "℃",
			},
			// 3
			{
				Name: "moisture 土壤含水量",
				Code: "soil_vwc",
				Unit: "m³/m³",
			},
		},
	},
	"I": {
		Code: "mt10ec",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "dielectric permittivity 介电常数",
				Code: "soil_dc",
				Unit: "",
			},
			// 2
			{
				Name: "electrical conductivity 土壤电导率",
				Code: "soil_ec",
				Unit: "dS/m",
			},
			// 3
			{
				Name: "temperature 土壤温度",
				Code: "soil_temp",
				Unit: "℃",
			},
			// 4
			{
				Name: "moisture 土壤含水量",
				Code: "soil_vwc",
				Unit: "m³/m³",
			},
		},
	},
	"J": {
		Code: "si4X1",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "目标温度",
				Code: "target_temp",
				Unit: "℃",
			},
			// 2
			{
				Name: "环境温度",
				Code: "body_temp",
				Unit: "℃",
			},
		},
	},
	"K": {
		Code: "gmx",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "relative direction 相对风向",
				Code: "rd",
				Unit: "°",
			},
		},
	},
	"L": {
		Code: "cs45X",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "水深",
				Code: "water_level",
				Unit: "m",
			},
			// 2
			{
				Name: "水温",
				Code: "water_temp",
				Unit: "℃",
			},
		},
	},
	"M": {
		Code: "teros11",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "calibratedCountsVWC",
				Code: "soil_raw",
				Unit: "",
			},
			// 2
			{
				Name: "soil temperature 土壤温度",
				Code: "soil_temp",
				Unit: "℃",
			},
			// 3
			{
				Name: "mineral soils vwc 土壤含水量",
				Code: "soil_vwc",
				Unit: "m³/m³",
			},
			// 4
			{
				Name: "apparent dielectric 视在介电常数",
				Code: "soil_apparent_dc",
				Unit: "",
			},
		},
	},
	"N": {
		Code: "teros12",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "calibrated Counts VWC",
				Code: "soil_raw",
				Unit: "",
			},
			// 2
			{
				Name: "soil temperature 土壤温度",
				Code: "soil_temp",
				Unit: "℃",
			},
			// 3
			{
				Name: "electrical conductivity 土壤电导率",
				Code: "soil_ec",
				Unit: "dS/m",
			},
			// 4
			{
				Name: "mineral soils vwc 土壤含水量",
				Code: "soil_vwc",
				Unit: "m³/m³",
			},
			// 5
			{
				Name: "apparent dielectric 视在介电常数",
				Code: "soil_apparent_dc",
				Unit: "",
			},
		},
	},
	"P": {
		Code: "fss_ir10",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "目标温度",
				Code: "target_temp",
				Unit: "℃",
			},
			// 2
			{
				Name: "环境湿度",
				Code: "body_temp",
				Unit: "℃",
			},
		},
	},
	"Q": {
		Code: "psw_0p1",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "开关量计数,系数:0.1",
				Code: "psw",
				Unit: "",
			},
		},
	},
	"R": {
		Code: "psw_0p2",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "开关量计数,系数:0.2",
				Code: "psw",
				Unit: "",
			},
		},
	},
	"S": {
		Code: "psw_0p254",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "开关量计数,系数:0.254",
				Code: "psw",
				Unit: "",
			},
		},
	},
	"T": {
		Code: "psw_1",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "开关量计数,系数:1",
				Code: "psw",
				Unit: "",
			},
		},
	},
	"U": {
		Code: "ri205",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "C1~C8雨量",
				Code: "psw",
				Unit: "mm",
			},
		},
	},
	"V": {
		Code: "atmos14",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "vapor pressure 饱和水汽压",
				Code: "vapor_pressure",
				Unit: "kPa",
			},
			// 2
			{
				Name: "temperature 空气温度",
				Code: "temperature",
				Unit: "℃",
			},
			// 3
			{
				Name: "relative humidity 空气湿度",
				Code: "relative_humidity",
				Unit: "%RH",
			},
			// 4
			{
				Name: "barometric pressure 大气压力",
				Code: "barometric_pressure",
				Unit: "kPa",
			},
		},
	},
	"W": {
		Code: "tensiomark",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "soil temperature 土壤温度",
				Code: "temperature",
				Unit: "℃",
			},
			// 2
			{
				Name: "soil water potential 土壤水势",
				Code: "potential",
				Unit: "pF",
			},
		},
	},
	"X": {
		Code: "acclima_tdr3xx",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Volumetric Water Content 土壤含水量",
				Code: "vwc",
				Unit: "%",
			},
		},
	},
	"Y": {
		Code: "acclima_tdt",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Volumetric Water Content 土壤含水量",
				Code: "vwc",
				Unit: "%",
			},
		},
	},
	"Z": {
		Code: "sr50",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Distance 距离",
				Code: "distance",
				Unit: "m",
			},
			// 2
			{
				Name: "Quality number",
				Code: "qn",
				Unit: "",
			},
		},
	},
	"a": {
		Code: "sr50at",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Distance 距离",
				Code: "distance",
				Unit: "m",
			},
			// 2
			{
				Name: "Quality number",
				Code: "qn",
				Unit: "",
			},
			// 3
			{
				Name: "Temperature 温度",
				Code: "temperature",
				Unit: "℃",
			},
		},
	},
	"b": {
		Code: "multi_6p",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "风速",
				Code: "wind_speed",
				Unit: "m/s",
			},
			// 2
			{
				Name: "风向",
				Code: "wind_direction",
				Unit: "°",
			},
			// 3
			{
				Name: "空气温度",
				Code: "temperature",
				Unit: "℃",
			},
			// 4
			{
				Name: "空气湿度",
				Code: "humidity",
				Unit: "%RH",
			},
			// 5
			{
				Name: "大气压力",
				Code: "pressure",
				Unit: "hPa",
			},
			// 6
			{
				Name: "累计雨量",
				Code: "rain_accumulated",
				Unit: "mm",
			},
			// 7
			{
				Name: "分钟雨量",
				Code: "rain_minute",
				Unit: "mm",
			},
		},
	},
	"c": {
		Code: "tr6817",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "空气温度",
				Code: "temperature",
				Unit: "℃",
			},
			// 2
			{
				Name: "空气湿度",
				Code: "humidity",
				Unit: "%RH",
			},
			// 3
			{
				Name: "大气压力",
				Code: "pressure",
				Unit: "hPa",
			},
			// 4
			{
				Name: "风速",
				Code: "wind_speed",
				Unit: "m/s",
			},
			// 5
			{
				Name: "风向",
				Code: "wind_direction",
				Unit: "°",
			},
			// 6
			{
				Name: "累计雨量",
				Code: "rain_accumulated",
				Unit: "mm",
			},
			// 7
			{
				Name: "辐射",
				Code: "radiation",
				Unit: "W/m2",
			},
			// 8
			{
				Name: "日照时数",
				Code: "sunshine_time",
				Unit: "hours",
			},
			// 9
			{
				Name: "光合有效",
				Code: "par",
				Unit: "μmol/(㎡*s)",
			},
		},
	},
	"d": {
		Code: "es101",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "蒸发量",
				Code: "ec",
				Unit: "mm",
			},
			// 2
			{
				Name: "水深",
				Code: "water_level",
				Unit: "mm",
			},
		},
	},
	"e": {
		Code: "lc5060",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "重量",
				Code: "weight",
				Unit: "g",
			},
			// 2
			{
				Name: "当前重量的蒸发当量",
				Code: "weight_ec",
				Unit: "mm",
			},
			// 3
			{
				Name: "温度",
				Code: "temperature",
				Unit: "℃",
			},
			// 4
			{
				Name: "渗漏量",
				Code: "leakage",
				Unit: "mm",
			},
			// 5
			{
				Name: "1#传感器电压比率",
				Code: "rate_1",
				Unit: "",
			},
			// 6
			{
				Name: "2#传感器电压比率",
				Code: "rate_2",
				Unit: "",
			},
			// 7
			{
				Name: "3#传感器电压比率",
				Code: "rate_3",
				Unit: "",
			},
		},
	},
	"f": {
		Code: "lys3040",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "土柱重量",
				Code: "weight",
				Unit: "kg",
			},
			// 2
			{
				Name: "渗漏计翻斗的次数",
				Code: "count",
				Unit: "",
			},
			// 3
			{
				Name: "计数通道(本型号无效)",
				Code: "undefined",
				Unit: "",
			},
			// 4
			{
				Name: "温度",
				Code: "temperature",
				Unit: "℃",
			},
			// 5
			{
				Name: "保留",
				Code: "reserved",
				Unit: "",
			},
			// 6
			{
				Name: "设备电压",
				Code: "battery",
				Unit: "V",
			},
		},
	},
	"g": {
		Code: "tr219_4",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "温度1",
				Code: "temperature_1",
				Unit: "",
			},
			// 2
			{
				Name: "温度2",
				Code: "temperature_2",
				Unit: "",
			},
			// 3
			{
				Name: "温度3",
				Code: "temperature_3",
				Unit: "",
			},
			// 4
			{
				Name: "温度4",
				Code: "temperature_4",
				Unit: "",
			},
		},
	},
	"h": {
		Code: "drs26",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "circumference increment",
				Code: "increment",
				Unit: "mm",
			},
			// 2
			{
				Name: "temperature 温度",
				Code: "temperature",
				Unit: "℃",
			},
		},
	},
	"i": {
		Code: "wxa100_02s",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "风速",
				Code: "wind_speed",
				Unit: "m/s",
			},
			// 2
			{
				Name: "风向",
				Code: "wind_direction",
				Unit: "°",
			},
		},
	},
	"j": {
		Code: "general_sensor_1",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "测量值",
				Code: "value",
				Unit: "",
			},
		},
	},
	"k": {
		Code: "ptbx10",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "大气压力",
				Code: "pressure",
				Unit: "hPa",
			},
		},
	},
	"l": {
		Code: "sf4",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "Min. particle flux",
				Code: "flux_min",
				Unit: "g/㎡/s",
			},
			// 2
			{
				Name: "Avg particle flux",
				Code: "flux_avg",
				Unit: "g/㎡/s",
			},
			// 3
			{
				Name: "Max. particle flux",
				Code: "flux_max",
				Unit: "g/㎡/s",
			},
			// 4
			{
				Name: "Std particle flux",
				Code: "flux_std",
				Unit: "g/㎡/s",
			},
			// 5
			{
				Name: "Cumulative flux",
				Code: "flux_cumulative",
				Unit: "g/㎡",
			},
			// 6
			{
				Name: "Wind min",
				Code: "wind_min",
				Unit: "km/h",
			},
			// 7
			{
				Name: "Wind avg",
				Code: "wind_avg",
				Unit: "km/h",
			},
			// 8
			{
				Name: "Wind max",
				Code: "wind_max",
				Unit: "km/h",
			},
		},
	},
	"m": {
		Code: "general_sensor_2",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "",
				Code: "value_1",
				Unit: "测量值1",
			},
			// 2
			{
				Name: "",
				Code: "value_2",
				Unit: "测量值2",
			},
		},
	},
	"n": {
		Code: "general_sensor_3",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "测量值1",
				Code: "value_1",
				Unit: "",
			},
			// 2
			{
				Name: "测量值2",
				Code: "value_2",
				Unit: "",
			},
			// 3
			{
				Name: "测量值3",
				Code: "value_3",
				Unit: "",
			},
		},
	},
	"o": {
		Code: "general_sensor_4",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "测量值1",
				Code: "value_1",
				Unit: "",
			},
			// 2
			{
				Name: "测量值2",
				Code: "value_2",
				Unit: "",
			},
			// 3
			{
				Name: "测量值3",
				Code: "value_3",
				Unit: "",
			},
			// 4
			{
				Name: "测量值4",
				Code: "value_4",
				Unit: "",
			},
		},
	},
	"p": {
		Code: "general_sensor_5",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "测量值1",
				Code: "value_1",
				Unit: "",
			},
			// 2
			{
				Name: "测量值2",
				Code: "value_2",
				Unit: "",
			},
			// 3
			{
				Name: "测量值3",
				Code: "value_3",
				Unit: "",
			},
			// 4
			{
				Name: "测量值4",
				Code: "value_4",
				Unit: "",
			},
			// 5
			{
				Name: "测量值5",
				Code: "value_5",
				Unit: "",
			},
		},
	},
	"q": {
		Code: "w10",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "风速",
				Code: "wind_speed",
				Unit: "m/s",
			},
		},
	},
	"r": {
		Code: "w20",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "风向",
				Code: "wind_direction",
				Unit: "°",
			},
		},
	},
	"s": {
		Code: "ctd",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "水深",
				Code: "depth",
				Unit: "mm",
			},
			// 2
			{
				Name: "温度",
				Code: "temperature",
				Unit: "℃",
			},
			// 3
			{
				Name: "电导率",
				Code: "bulk_electrical_conductivity",
				Unit: "uS/cm",
			},
		},
	},
	"t": {
		Code: "crsm400",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "快中子计数",
				Code: "count",
				Unit: "",
			},
		},
	},
	"u": {
		Code: "crsm800",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "慢中子计数",
				Code: "count",
				Unit: "",
			},
		},
	},
	"v": {
		Code: "lc2030",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "重量",
				Code: "weight",
				Unit: "g",
			},
			// 2
			{
				Name: "当前重量的蒸发当量",
				Code: "weight_ec",
				Unit: "mm",
			},
			// 3
			{
				Name: "温度",
				Code: "temperature",
				Unit: "℃",
			},
			// 4
			{
				Name: "渗漏量",
				Code: "leakage",
				Unit: "mm",
			},
			// 5
			{
				Name: "传感器电压比率",
				Code: "rate",
				Unit: "",
			},
		},
	},
	"w": {
		Code: "yl100",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "温度",
				Code: "temperature",
				Unit: "℃",
			},
		},
	},
	"x": {
		Code: "bbs_688s",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "溶解氧",
				Code: "dissolved_oxygen",
				Unit: "mg/L",
			},
			// 2
			{
				Name: "浊度",
				Code: "turbidity",
				Unit: "NTU",
			},
			// 3
			{
				Name: "电导率",
				Code: "conductivity",
				Unit: "ms/cm",
			},
			// 4
			{
				Name: "酸碱度",
				Code: "pH",
				Unit: "",
			},
			// 5
			{
				Name: "温度",
				Code: "temperature",
				Unit: "℃",
			},
		},
	},
	"y": {
		Code: "rq_30a",
		Name: "",
		Data: []DeviceSensor{
			// 1
			{
				Name: "自校验",
				Code: "self_check",
				Unit: "",
			},
			// 2
			{
				Name: "水位",
				Code: "level",
				Unit: "m",
			},
			// 3
			{
				Name: "速率",
				Code: "velocity",
				Unit: "m/s",
			},
			// 4
			{
				Name: "品质(SNR)",
				Code: "quality_SNR",
				Unit: "",
			},
			// 5
			{
				Name: "释放量",
				Code: "discharge",
				Unit: "m^3/s",
			},
			// 6
			{
				Name: "区域",
				Code: "area",
				Unit: "m^2",
			},
			// 7
			{
				Name: "理论速率",
				Code: "learned_velocity",
				Unit: "m/s",
			},
			// 8
			{
				Name: "理论的释放量",
				Code: "learned_discharge",
				Unit: "m^3/s",
			},
			// 9
			{
				Name: "相反方向",
				Code: "opposite_direction",
				Unit: "%",
			},
			// 10
			{
				Name: "供电电压",
				Code: "supply_voltage",
				Unit: "V",
			},
		},
	},
}

type DeviceService struct {
	config        *config.Config
	db            *store.DBClient
	logger        *logger.Logger
	sensorService *SensorService
	recordService *RecordService
}

func NewDeviceService(c *config.Config, db *store.DBClient, logger *logger.Logger, sensorService *SensorService, recordService *RecordService) *DeviceService {
	log := logger.Sugar()
	defer log.Sync()
	return &DeviceService{c, db, logger, sensorService, recordService}
}

func (s *DeviceService) ReceiveData(data string) error {
	log := s.logger.Sugar()
	defer log.Sync()
	arr := strings.Split(data, ",")
	const kIndexIMEI = 0
	// const kIndexUploadTime = 1
	// const kIndexLat = 2
	// const kIndexLng = 3
	// const kIndexAltitude = 4
	const kIndexSensorStructs = 5
	const kIndexCreateTime = 6
	const kIndexRecordIndex = 7
	// const kIndexVoltage = 8
	// const kIndexTemperature = 9
	// const kIndexHumidity = 10
	const kIndexDataStart = 11

	const kValidSensorStructsLength = 30
	const kSensorDataSampleMethodIndex = 15
	if len(arr) < kIndexDataStart+1 {
		return errors.New("[device-service] invalid length of report data")
	}
	dataIMEI := strings.TrimSpace(arr[kIndexIMEI])
	// dataUploadTime := strings.TrimSpace(arr[kIndexUploadTime])
	// dataLat := strings.TrimSpace(arr[kIndexLat])
	// dataLng := strings.TrimSpace(arr[kIndexLng])
	// dataAltitude := strings.TrimSpace(arr[kIndexAltitude])
	dataSensorStructs := strings.TrimSpace(arr[kIndexSensorStructs])
	dataCreateTime := strings.TrimSpace(arr[kIndexCreateTime])
	dataRecordIndex := strings.TrimSpace(arr[kIndexRecordIndex])
	// dataVoltage := strings.TrimSpace(arr[kIndexVoltage])
	// dataTemperature := strings.TrimSpace(arr[kIndexTemperature])
	// dataHumidity := strings.TrimSpace(arr[kIndexHumidity])
	recordCreateTime, err := time.Parse("2006-01-02 15:04:05", dataCreateTime)
	if err != nil {
		return fmt.Errorf("[device-service] invalid create time format: %s", recordCreateTime)
	}

	recordIndexValue, err := strconv.ParseUint(dataRecordIndex, 10, 64)
	if err != nil {
		return fmt.Errorf("[device-service] invalid record index: %s", dataRecordIndex)
	}

	if len(dataIMEI) == 0 {
		return errors.New("[device-service] invalid imei length")
	}
	if len(dataSensorStructs) != kValidSensorStructsLength {
		return errors.New("[device-service] invalid sensorTypes length")
	}
	dataSensorTypes := string(dataSensorStructs[0:kSensorDataSampleMethodIndex])
	dataSensorSampleMethods := string(dataSensorStructs[kSensorDataSampleMethodIndex:])
	recordDataIndex := kIndexDataStart

	parsedRecords := []Record{}

	for i := 0; i < kValidSensorStructsLength/2; i++ {
		sensorType := string(dataSensorTypes[i])
		sensorSampleMethodIndex := string(dataSensorSampleMethods[i])
		sensorSampleMethod, ok := sampleMethodMap[sensorSampleMethodIndex]
		if !ok {
			log.Errorln("[device-service] received an invalid sample method index: ", sensorSampleMethodIndex)
			continue
		}
		if sensorType == "0" {
			log.Infof("[device-service] no sensor at position: %d\n", i)
			continue
		}
		device, ok := deviceMap[sensorType]
		if !ok {
			log.Errorln("[device-service] received an invalid type: ", sensorType)
			continue
		}
		sensors := device.Data
		for i := 0; i < len(sensors); i++ {
			sensorData := sensors[i]
			sensorReportCode := fmt.Sprintf("%s-%s-%s-%d", dataIMEI, sensorType, sensorData.Code, i)
			log.Infof("[device-service] add record in index: %d, sensorReportCode: %s, sensorCode: %s, sensorName: %s, sampleMethod: %s\n", recordDataIndex, sensorReportCode, sensorData.Code, sensorData.Name, sensorData.Unit, sensorSampleMethod)
			var dbSensor Sensor
			err := s.db.Where("sensor_report_code = ?", sensorReportCode).First(&dbSensor).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Infof("[device-service] sensor not found, insert new for code: %s\n", sensorReportCode)
				newSensor := Sensor{
					Name:             sensorData.Name,
					Unit:             sensorData.Unit,
					SensorCode:       sensorData.Code,
					IMEI:             dataIMEI,
					SensorReportCode: sensorReportCode,
					SampleMethod:     sensorSampleMethod,
				}
				insertedSensor, err := s.sensorService.Upsert(&newSensor)
				if err != nil {
					log.Errorln("[device-service] insert sensor with error: ", err)
				} else {
					dbSensor = *insertedSensor
				}
			}
			if len(arr) <= recordDataIndex {
				return fmt.Errorf("[device-service] parse sensor data out of range, at index:%d, sensorReportCode: %s", recordDataIndex, sensorReportCode)
			}
			recordData := arr[recordDataIndex]
			recordValue, err := strconv.ParseFloat(recordData, 64)
			if err != nil {
				log.Errorf("[device-service] failed to parse record value: %s, error: %v\n", recordData, err)
				continue
			}
			record := Record{
				SensorId:    dbSensor.ID,
				Value:       recordValue,
				Time:        recordCreateTime,
				RecordIndex: recordIndexValue,
			}
			parsedRecords = append(parsedRecords, record)
			recordDataIndex += 1
		}
	}
	s.recordService.BatchUpsert(&parsedRecords)
	return nil
}
