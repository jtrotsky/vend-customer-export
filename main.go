package main

import (
	"flag"
	"os"
	"strings"

	"github.com/jtrotsky/gov/vend"
	"github.com/jtrotsky/vend-customer-export/writer"
	log "github.com/sirupsen/logrus"
)

var (
	domainPrefix string
	authToken    string
	vendClient   *vend.Client
)

func main() {
	// Create new Vend Client.
	vendClient = vend.NewClient(authToken, domainPrefix, "")
	getAllCustomers()
}

// Run executes the process of grabbing customers then writing them to CSV.
func getAllCustomers() {
	log.Info("Retrieving customers")

	// Get customers.
	customers, err := vendClient.Customers.ListAll()
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"event": "retrieve",
			"msg": "Something went wrong whilst retrieving customers from " +
				"the store.",
			"suggestion": "See if there's anything obvious in the error " +
				"messaging, otherwise tap the nearest Gopher on the shoulder.",
		}).Warn("Failed while retrieving customers.")
	}

	log.Info("Writing customers to CSV file")

	err = writer.WriteFile(customers, domainPrefix)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"event": "write",
			"msg": "Something went wrong whilst writing the customers " +
				"to the CSV file.",
			"suggestion": "See if there's anything obvious in the error " +
				"messaging, otherwise tap the nearest Gopher on the shoulder.",
		}).Fatal("Failed while writing customers to CSV file.")
	}
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout.
	log.SetOutput(os.Stderr)

	// Get store info from command line flags.
	flag.StringVar(&domainPrefix, "d", "", "The Vend store name (prefix in "+
		"xxxx.vendhq.com)")
	flag.StringVar(&authToken, "t", "", "Personal API Access Token for the "+
		"store, generated from Setup -> Personal Tokens.")
	flag.Parse()

	if domainPrefix == "" {
		log.Fatal("Domain prefix is too short or was not specified, enter a " +
			"valid domain prefix with the -d argument.")
	}
	if authToken == "" {
		log.Fatal("Token was not specified, enter a valid authentication " +
			"token with the -t argument.")
	}

	// To save people who write DomainPrefix.vendhq.com.
	// Split DomainPrefix on the "." period character then grab the first part.
	parts := strings.Split(domainPrefix, ".")
	domainPrefix = parts[0]
}
