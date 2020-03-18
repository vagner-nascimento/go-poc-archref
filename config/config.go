package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

type MongoDataConfig struct {
	ConnStr            string        `json:"connStr"`
	ClientTimeOut      time.Duration `json:"clientTimeOut"`
	InsertTimeout      time.Duration `json:"insertTimeout"`
	FindTimeout        time.Duration `json:"findTimeout"`
	MaxPaginatedSearch int64         `json:"maxPaginatedSearch"`
}

type NoSqlDataConfig struct {
	Mongo MongoDataConfig `json:"mongo"`
}

type AmqpDataConfig struct {
	ConnStr string `json:"connStr"`
}

type DataConfig struct {
	NoSql NoSqlDataConfig `json:"noSql"`
	Amqp  AmqpDataConfig  `json:"amqp"`
}

type PresentationWebConfig struct {
	Port int16 `json:"port"`
}

type PresentationConfig struct {
	Web PresentationWebConfig `json:"web"`
}

type Config struct {
	Data         DataConfig         `json:"data"`
	Presentation PresentationConfig `json:"presentation"`
}

var config *Config

func Load(environment string) error {
	if config != nil {
		return errors.New("config is already loaded")
	}

	path, _ := filepath.Abs(fmt.Sprintf("config/%s.json", strings.ToLower(environment)))
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, &config); err != nil {
		return err
	}
	return nil
}

func Get() Config {
	if config == nil {
		panic("config not loaded")
	}
	return *config
}
