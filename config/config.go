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
	Database           string        `json:"database"`
}

type CollectionsConfig struct {
	Customer string `json:"customer"`
	Supplier string `json:"supplier"`
}

type NoSqlDataConfig struct {
	Mongo       MongoDataConfig   `json:"mongo"`
	Collections CollectionsConfig `json:"collections"`
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

type TopicConfig struct {
	Topic    string `json:"topic"`
	Consumer string `json:"consumer"`
}

type SubsConfig struct {
	User       TopicConfig `json:"user"`
	Enterprise TopicConfig `json:"enterprise"`
}

type PubsConfig struct {
	Customer TopicConfig `json:"customer"`
	Supplier TopicConfig `json:"supplier"`
}

type AmqIntegrationConfig struct {
	Subs SubsConfig `json:"subs"`
	Pubs PubsConfig `json:"pubs"`
}

type IntegrationConfig struct {
	Amqp AmqIntegrationConfig `json:"amqp"`
}

type Config struct {
	Data         DataConfig         `json:"data"`
	Presentation PresentationConfig `json:"presentation"`
	Integration  IntegrationConfig  `json:"integration"`
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
