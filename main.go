package main

import log "github.com/sirupsen/logrus"

func main() {
	log.Println("Hello world!")

	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"foo": "foo",
			"bar": "bar",
		},
	).Debug("Something happened")
}
