package writer

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/jtrotsky/gov/vend"
	log "github.com/sirupsen/logrus"
)

// WriteFile writes customer info to file.
func WriteFile(customers []vend.Customer, domainPrefix string) error {
	// Create a blank CSV file.
	// The file name will be the current time in unixtime and the store name.
	fileName := fmt.Sprintf("%s_customer_export_%v.csv", domainPrefix, time.Now().Unix())

	file, err := os.Create(fmt.Sprintf("./%s", fileName))
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"event":      "write",
			"msg":        "Failed while creating blank CSV file.",
			"suggestion": "Check permissions of the running directory.",
		}).Fatal("Failed to create CSV file.")
	}

	// Ensure the file is closed at the end.
	defer file.Close()

	// Create CSV writer on the file.
	writer := csv.NewWriter(file)

	// Write the header line.
	var header []string
	header = append(header, "id")
	header = append(header, "customer_code")
	header = append(header, "first_name")
	header = append(header, "last_name")
	header = append(header, "email")
	header = append(header, "note")
	header = append(header, "company_name")
	header = append(header, "mobile")
	header = append(header, "deleted_at")

	// Commit the header.
	writer.Write(header)

	// Now loop through each customer object and populate the CSV.
	for _, customer := range customers {

		var deletedAt time.Time
		var id, code, firstName, lastName, email, note, companyName,
			mobile, deletedAtStr string
		if customer.ID != nil {
			id = *customer.ID
		}
		if customer.Code != nil {
			code = *customer.Code
		}
		if customer.FirstName != nil {
			firstName = *customer.FirstName
		}
		if customer.LastName != nil {
			lastName = *customer.LastName
		}
		if customer.Email != nil {
			email = *customer.Email
		}
		if customer.Note != nil {
			note = *customer.Note
		}
		if customer.CompanyName != nil {
			companyName = *customer.CompanyName
		}
		if customer.Mobile != nil {
			mobile = *customer.Mobile
		}
		if customer.DeletedAt != nil {
			deletedAt = *customer.DeletedAt
			deletedAtStr = deletedAt.String()
		}

		var record []string
		record = append(record, id)
		record = append(record, code)
		record = append(record, firstName)
		record = append(record, lastName)
		record = append(record, email)
		record = append(record, note)
		record = append(record, companyName)
		record = append(record, mobile)
		record = append(record, deletedAtStr)
		writer.Write(record)
	}

	writer.Flush()
	return err
}
