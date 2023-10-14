package sqlwrap_test

import (
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
