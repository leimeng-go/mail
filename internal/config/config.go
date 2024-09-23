package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
    Cache cache.NodeConf `json:"redisconf"`
	Database MysqlConf      `json:"database"`
	// Log logx.LogConf
	// Mode string `yaml:",default=pro,options=dev|test|rt|pre|pro"`
	// MetricsUrl string `yaml:"metricsUrl"`
	// Prometheus prometheus.Config `yaml:"prometheus"`
	// Telemetry trace.Config `yaml:"telemetry"`
}

type MysqlConf struct{
	Driver string `json:"driver"`
	Host string `json:"host"`
	Port int `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Charset string `json:"charset"`
	ParseTime string `json:"parseTime"`
	Loc string `json:"loc"`
}