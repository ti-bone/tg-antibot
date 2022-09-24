/*
 * config.go
 * Copyright (c) ti-bone 2022.
 */

package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// CurrentConfiguration - Current loaded configuration, contains Config struct.
var CurrentConfiguration Config

// LoadConfig - Load Configuration into CurrentConfiguration variable
func LoadConfig(path string) {
	// Read config file
	cfgFile, err := os.ReadFile(path)

	// Catch errors
	if err != nil {
		panic(fmt.Errorf("Cannot load config file!\nerr: %w", err))
	}

	// Parse config
	errP := json.Unmarshal(cfgFile, &CurrentConfiguration)

	// Catch errors
	if errP != nil {
		panic(fmt.Errorf("Cannot parse config file!\nerr: %w", errP))
	}
}

// Config struct, used for parsing json
type Config struct {
	Bot struct {
		Token   string `json:"Token"`   // Bot token
		AdminID int64  `json:"AdminID"` // Admin ID, where will come errors
	}
	//Mysql struct {
	//	Host string `json:"Host"`
	//	Port int64  `json:"Port"`
	//	User string `json:"User"`
	//	Pass string `json:"Pass"`
	//} TODO: MySQL support
}
