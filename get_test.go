package sqlwrap_test

import (
	"context"

	"github.com/jtarchie/sqlwrap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get", func() {
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

	It("errors on missing values for bind params", func() {
		var firstName string

		err := client.Get(
			context.Background(),
			&firstName,
			"SELECT first_name FROM people WHERE email = :email",
			map[string]interface{}{},
		)
		Expect(err).To(HaveOccurred())
	})

	It("errors when binding is missing", func() {
		var firstName string

		err := client.Get(
			context.Background(),
			&firstName,
			"SELECT first_name FROM people WHERE email = :email",
			map[string]interface{}{},
		)
		Expect(err).To(HaveOccurred())
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
	})

	It("errors with missing params of IN clauses", func() {
		var firstName string

		err := client.Get(
			context.Background(),
			&firstName,
			"SELECT first_name FROM people WHERE email IN (:email)",
			map[string]interface{}{},
		)
		Expect(err).To(HaveOccurred())
	})
})
