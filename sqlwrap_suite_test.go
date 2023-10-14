package sqlwrap_test

import (
	"context"
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

	When("selecting a single row", func() {
		var client *sqlwrap.Client

		BeforeEach(func() {
			var err error

			client, err = sqlwrap.Open("sqlite3", ":memory:")
			Expect(err).NotTo(HaveOccurred())

			_, err = client.Exec(`
				CREATE TABLE people (
					first_name TEXT,
					last_name  TEXT,
					email      TEXT
				);
				INSERT INTO people (first_name, last_name, email) VALUES ('Bob', 'Smith', 'bob@smith.com');
			`)
			Expect(err).NotTo(HaveOccurred())
		})

		It("bind with named parameters", func() {
			var firstName string

			err := client.Get(
				context.Background(),
				&firstName,
				"SELECT first_name FROM people WHERE email = :email",
				map[string]interface{}{
					"email": "bob@smith.com",
				},
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(firstName).To(Equal("Bob"))
		})

		It("can handle IN clauses", func() {
			var firstName string

			err := client.Get(
				context.Background(),
				&firstName,
				"SELECT first_name FROM people WHERE email IN (:email)",
				map[string]interface{}{
					"email": []string{"bob@smith.com"},
				},
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(firstName).To(Equal("Bob"))

			err = client.Get(
				context.Background(),
				&firstName,
				"SELECT first_name FROM people WHERE email IN (:email)",
				map[string]interface{}{},
			)
			Expect(err).To(HaveOccurred())
		})
	})

	When("selecting many rows", func() {
		It("bind with named parameters", func() {
		})

		It("can handle IN clauses", func() {
		})
	})
})
