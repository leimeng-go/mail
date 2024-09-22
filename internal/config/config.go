package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/prometheus"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
    Cache cache.NodeConf `yaml:"cache"`
	MysqlConf MysqlConf      `yaml:"database"`
	Log logx.LogConf
	Mode string `yaml:",default=pro,options=dev|test|rt|pre|pro"`
	MetricsUrl string `yaml:"metricsUrl"`
	Prometheus prometheus.Config `yaml:"prometheus"`
	Telemetry trace.Config `yaml:"telemetry"`
}

type MysqlConf struct{
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Charset string `yaml:"charset"`
	ParseTime string `yaml:"parseTime"`
	Loc string `yaml:"loc"`
}