package util

import (
	"gopkg.in/ini.v1"
	"log"
)

var configFile = "config.ini"

func GetConfig(section string, key string) string {
	return GetConfigByPath(configFile, section, key)
}

func GetConfigBool(path, section, key string) bool {
	temp_bool, err := GetConfigBase(path, section, key).Bool()
	if err != nil {
		log.Fatal(err)
	}
	return temp_bool
}

func GetConfigInt(path, section, key string) int {
	temp_int, err := GetConfigBase(path, section, key).Int()
	if err != nil {
		log.Fatal(err)
	}
	return temp_int
}

func GetConfigByPath(path, section, key string) string {
	return GetConfigBase(path, section, key).String()
}

func GetConfigBase(path, section, key string) *ini.Key {
	cfg, err := ini.Load(path)
	if err != nil {
		log.Fatal(err)
	}
	return cfg.Section(section).Key(key)
}