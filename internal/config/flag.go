package config

import "flag"

var Loglevel string

// ParseFlag парсит аргументы команды запуска
func ParseFlag() {
	Loglevel = *flag.String("loglevel", "INFO", "Set log level: DEBUG, INFO, WARN, ERROR")
	flag.Parse()
}
