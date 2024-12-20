package main

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type SpindlePosition struct {
	ZUp   float64 `yaml:"zup"`
	ZDown float64 `yaml:"zdown"`
}

var defSpindlePosition SpindlePosition

func parseConfig() {
	var err error

	f, err := os.ReadFile("./main.cfg")

	if err != nil {
		sugar.Errorln("CAN'T READ CONFIG FILE.")
	}

	defSpindlePosition = SpindlePosition{}

	err = yaml.Unmarshal(f, &defSpindlePosition)

	if err != nil {
		sugar.Errorln("CONFIG ERROR. CAN'T UNMARSHAL FILE.")
	}
}

func parseFlags() {
	flag.Float64Var(&defSpindlePosition.ZUp, "zup", 0, "точка в которой размещается сверло при перемещении")
	flag.Float64Var(&defSpindlePosition.ZDown, "zdown", 0, "нижняя точка в которую опускается сверло при сверлении")
}

func parseParams() {
	parseConfig()
	parseFlags()
	flag.Parse()
}
