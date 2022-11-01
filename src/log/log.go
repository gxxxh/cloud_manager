package log

import "log"

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func DEBUG(format string, v ...interface{}) {
	log.SetPrefix("DEBUG: ")
	log.Printf(format, v)
}

func INFO(format string, v ...interface{}) {
	log.SetPrefix("INFO: ")
	log.Printf(format, v)
}

func WARNING(format string, v ...interface{}) {
	log.SetPrefix("WARNING:  ")
	log.Printf(format, v)
}

func ERROR(format string, v ...interface{}) {
	log.SetPrefix("ERROR: ")
	log.Printf(format, v)
}

func CRITICAL(format string, v ...interface{}) {
	log.SetPrefix("CRITICAL: ")
	log.Panicf(format, v)
}
