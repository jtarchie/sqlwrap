package sqlwrap_test

import (
	"context"
	"database/sql"
	"strings"
	"testing"
)

var err error //nolint

func BenchmarkGet(b *testing.B) {
	client, err := createClient()
	if err != nil {
		b.Errorf("could not create client: %s", err.Error())
	}

	b.Run(rxpad("get"), func(b *testing.B) {
		var firstName string

		for i := 0; i < b.N; i++ {
			err = client.Get(
				context.Background(),
				&firstName,
				"SELECT first_name FROM people WHERE email = 'bob@example.com'",
				map[string]interface{}{},
			)
		}
	})
	b.Run(rxpad("get with generic"), func(b *testing.B) {
		var firstName string

		for i := 0; i < b.N; i++ {
			err = client.Get(
				context.Background(),
				&firstName,
				"SELECT first_name FROM people WHERE email = 'bob@example.com'",
				map[string]interface{}{},
			)
		}
	})
	b.Run(rxpad("get with equals"), func(b *testing.B) {
		var firstName string

		for i := 0; i < b.N; i++ {
			err = client.Get(
				context.Background(),
				&firstName,
				"SELECT first_name FROM people WHERE email = :email",
				map[string]interface{}{
					"email": "bob@smith.com",
				},
			)
		}
	})
	b.Run(rxpad("get with IN"), func(b *testing.B) {
		var firstName string

		for i := 0; i < b.N; i++ {
			err = client.Get(
				context.Background(),
				&firstName,
				"SELECT first_name FROM people WHERE email IN (:email)",
				map[string]interface{}{
					"email": []string{"bob@smith.com"},
				},
			)
		}
	})
	b.Run(rxpad("classic with named"), func(b *testing.B) {
		var firstName string

		for i := 0; i < b.N; i++ {
			row := client.QueryRowContext(
				context.Background(),
				"SELECT first_name FROM people WHERE email = :email",
				sql.Named("email", "bob@smith.com"),
			)
			err = row.Scan(&firstName)
		}
	})
	b.Run(rxpad("classic"), func(b *testing.B) {
		var firstName string

		for i := 0; i < b.N; i++ {
			row := client.QueryRowContext(
				context.Background(),
				"SELECT first_name FROM people WHERE email = 'bob@smith.com'",
			)
			err = row.Scan(&firstName)
		}
	})
}

func rxpad(str string) string {
	lim := 50
	str += strings.Repeat(" ", lim)

	return str[:lim]
}
