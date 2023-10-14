package sqlwrap_test

import (
	"context"

	"github.com/jtarchie/sqlwrap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Select", func() {
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
				INSERT INTO people (first_name, last_name, email) VALUES ('Jane', 'Smith', 'jane@smith.com');
			`)
		Expect(err).NotTo(HaveOccurred())
	})

	It("bind with named parameters", func() {
		var firstNames []string

		err := client.Select(
			context.Background(),
			&firstNames,
			"SELECT first_name FROM people WHERE last_name = :last_name",
			map[string]interface{}{
				"last_name": "Smith",
			},
		)
		Expect(err).NotTo(HaveOccurred())
		Expect(firstNames).To(Equal([]string{"Bob", "Jane"}))
	})

	It("errors when binding is missing", func() {
		var firstNames []string

		err := client.Select(
			context.Background(),
			&firstNames,
			"SELECT first_name FROM people WHERE last_name = :last_name",
			map[string]interface{}{},
		)
		Expect(err).To(HaveOccurred())
	})

	It("can handle IN clauses", func() {
		var firstNames []string

		err := client.Select(
			context.Background(),
			&firstNames,
			"SELECT first_name FROM people WHERE email IN (:email)",
			map[string]interface{}{
				"email": []string{"bob@smith.com"},
			},
		)
		Expect(err).NotTo(HaveOccurred())
		Expect(firstNames).To(Equal([]string{"Bob"}))
	})

	It("errors with missing params of IN clauses", func() {
		var firstNames []string

		err := client.Select(
			context.Background(),
			&firstNames,
			"SELECT first_name FROM people WHERE email IN (:email)",
			map[string]interface{}{},
		)
		Expect(err).To(HaveOccurred())
	})
})