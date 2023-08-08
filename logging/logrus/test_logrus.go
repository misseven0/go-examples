package main

import log "github.com/sirupsen/logrus"

func main() {
	users := map[string]string{
		"ztz": "aaa",
		"zzz": "bbb",
	}

	// log.SetFormatter(&log.TextFormatter{
	// 	DisableColors: false,
	// 	FullTimestamp: true,
	// })

	// log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Trace("trace", users)
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Debug("debug", users)
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("info-->>>>", users)
}
