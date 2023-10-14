package sqlwrap_test

import (
	"fmt"
	"testing"

	"github.com/jtarchie/sqlwrap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSqlwrap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sqlwrap Suite")
}

var _ = Describe("SQLWrap", func() {
	When("connecting to a database", func() {
		It("connects successfully", func() {
			_, err := sqlwrap.Open("sqlite3", ":memory:")
			Expect(err).NotTo(HaveOccurred())
		})

		It("errors with invalid driver", func() {
			_, err := sqlwrap.Open("no-db", ":memory:")
			Expect(err).To(HaveOccurred())
		})
	})
})

func createClient() (*sqlwrap.Client, error) {
	client, err := sqlwrap.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("could not create client: %w", err)
	}

	_, err = client.Exec(`
		CREATE TABLE people (
			first_name TEXT,
			last_name  TEXT,
			email      TEXT
		);
		INSERT INTO people (first_name, last_name, email) VALUES ('Bob', 'Smith', 'bob@smith.com');
		INSERT INTO people (first_name, last_name, email) VALUES ('Jane', 'Smith', 'jane@smith.com');
	`)
	if err != nil {
		return nil, fmt.Errorf("could not execute: %w", err)
	}

	return client, nil
}
