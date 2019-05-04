package beater

import (
	"fmt"
	"time"

	"github.com/mrubelmann/bb8beat/bb8"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/mrubelmann/bb8beat/config"
)

// Bb8beat configuration.
type Bb8beat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
	bb8    bb8.BB8
}

// New creates an instance of bb8beat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	// Instantiate a BB-8.
	robot := bb8.NewBB8(c.BluetoothID)
	robot.AddCollisionEventHandler(bb8.OnCollision)

	bt := &Bb8beat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts bb8beat.
func (bt *Bb8beat) Run(b *beat.Beat) error {
	logp.Info("bb8beat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":    b.Info.Name,
				"counter": counter,
			},
		}
		bt.client.Publish(event)
		logp.Info("Event sent")
		counter++
	}
}

// Stop stops bb8beat.
func (bt *Bb8beat) Stop() {
	bt.client.Close()
	close(bt.done)
}
