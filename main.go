package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

var IsVerbose bool = false

func init() {
	//Process command line flags
	IsVerbose = *(flag.Bool("v", false, "Verbose Logging to Console (Info Level)"))
	flag.Parse()

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	//Set Options based on verbose console or file based logging 
	if (IsVerbose) {
		log.SetLevel(log.TraceLevel)
		log.SetOutput(os.Stdout)
		log.WithField("Verbose", IsVerbose).Info("Verbose Logging to Console (Info Level)")
	} else {
		f, err := os.OpenFile("web-server.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
		}
		log.SetOutput(f)
		log.SetLevel(log.WarnLevel)
	}
}

func main() {
	log.Info("Initializing Web-Server using goFiberV2")
	app := fiber.New()
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
        return c.SendString("OK")
    })

    app.Listen(":3000")
}