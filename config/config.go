// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period      time.Duration `config:"period"`
	BluetoothID string
}

var DefaultConfig = Config{
	Period:      1 * time.Second,
	BluetoothID: "BB-FC0B",
}
