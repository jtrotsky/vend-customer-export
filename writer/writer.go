package writer

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jtrotsky/govend/vend"
)

// WriteFile writes customer info to file.
func WriteFile(customers []vend.Customer, domainPrefix string) error {
	// Create blank CSV file to be written to.
	// File name will be the current time in unixtime.
	fname := fmt.Sprintf("%s_customer_export_%v.csv", domainPrefix,
		time.Now().Unix())
	f, err := os.Create(fmt.Sprintf("./%s", fname))
	if err != nil {
		log.Fatalf("Error creating CSV file: %s", err)
	}
	// Ensure file closes at end.
	defer f.Close()

	w := csv.NewWriter(f)

	var headerLine []string
	headerLine = append(headerLine, "id")
	headerLine = append(headerLine, "customer_code")
	headerLine = append(headerLine, "deleted_at")

	w.Write(headerLine)

	for _, customer := range customers {

		var deletedAt time.Time
		var id, code, deletedAtStr string
		if customer.ID != nil {
			id = *customer.ID
		}
		if customer.Code != nil {
			code = *customer.Code
		}
		if customer.DeletedAt != nil {
			deletedAt = *customer.DeletedAt
			deletedAtStr = deletedAt.String()
		}

		var record []string
		record = append(record, id)
		record = append(record, code)
		record = append(record, deletedAtStr)
		w.Write(record)
	}

	w.Flush()
	return err
}
