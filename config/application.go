package config

import "log"

// Application ...
type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
