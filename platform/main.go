package main

import (
	"platform/placeholder"
	"platform/services"
)

// func writeMessage(logger logging.Logger, cfg config.Configuration) {

// 	section, ok := cfg.GetSection("main")
// 	if ok {
// 		message, ok := section.GetString("message")
// 		if ok {
// 			logger.Info(message)
// 		} else {
// 			logger.Panic("Cannot find configuration setting")
// 		}
// 	} else {
// 		logger.Panic("Config section not found")
// 	}
// }

func main() {
	println("Starting...")
	services.RegisterDefaultServices()
	placeholder.Start()
}
