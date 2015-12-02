package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/jtrotsky/govend/vend"
	"github.com/jtrotsky/vend-customer-export/writer"
)

var (
	domainPrefix string
	token        string
)

func main() {
	// Third argument is timezone, not useful here.
	v := vend.NewClient(token, domainPrefix, "")
	manager := NewManager(v)

	manager.Run()
}

func init() {
	// Get store info from command line flags.
	flag.StringVar(&domainPrefix, "d", "",
		"The Vend store name (prefix of xxxx.vendhq.com)")
	flag.StringVar(&token, "t", "",
		"Personal API Access Token for the store, generated from Setup -> API Access.")
	flag.Parse()

	// To save people who write DomainPrefix.vendhq.com.
	// Split DomainPrefix on the "." period character then grab the first part.
	parts := strings.Split(domainPrefix, ".")
	domainPrefix = parts[0]
}

// Manager contains the Vend client.
type Manager struct {
	vend vend.Client
}

// NewManager creates an instance of manager.
func NewManager(vend vend.Client) *Manager {
	return &Manager{vend}
}

// Run executes the process of grabbing customers then writing them to CSV.
func (manager *Manager) Run() {
	fmt.Println("Getting customers")
	// Get customers.
	customers, err := manager.vend.Customers()
	if err != nil {
		log.Fatalf("Failed to get customers: %s", err)
	}

	fmt.Println("Writing to CSV")
	err = writer.WriteFile(*customers, domainPrefix)
}
