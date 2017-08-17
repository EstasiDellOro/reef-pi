package utils

import (
	"github.com/reef-pi/adafruitio"
	"log"
)

type AdafruitIO struct {
	Enable bool   `yaml:"enable"`
	Token  string `yaml:"token"`
	User   string `yaml:"user"`
	Feed   string `yaml:"feed"`
}

type Telemetry struct {
	client *adafruitio.Client
	config AdafruitIO
}

func NewTelemetry(config AdafruitIO) *Telemetry {
	return &Telemetry{
		client: adafruitio.NewClient(config.Token),
		config: config,
	}
}

func (t *Telemetry) EmitMetric(feed string, v interface{}) {
	d := adafruitio.Data{
		Value: v,
	}

	if !t.config.Enable {
		log.Println("Telemetry disabled. Skipping emitting", v, "on", feed)
		return
	}
	if err := t.client.SubmitData(t.config.User, feed, d); err != nil {
		log.Println("ERROR: Failed to submit data to adafruit.io. User: ", t.config.User, "Feed:", feed, "Error:", err)
	}
}