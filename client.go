package sqlwrap

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	*sql.DB
}

func Open(driver string, connection string) (*Client, error) {
	database, err := sql.Open(driver, connection)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	return &Client{
		DB: database,
	}, nil
}

var inListRegex = regexp.MustCompile(`(?im)\s+IN\s+\(\s*([@$:]\p{L}+)\s*\)`)

func (c *Client) Get(
	ctx context.Context,
	destination interface{},
	query string,
	params map[string]interface{},
) error {
	builder, count := &strings.Builder{}, 0

	matches := inListRegex.FindAllStringSubmatchIndex(query, -1)
	for _, match := range matches {
		start, end := match[2], match[3]

		paramName := query[start+1 : end]
		if _, ok := params[paramName]; !ok {
			return fmt.Errorf("could not find param for IN list %q", paramName)
		}

		builder.WriteString(query[count:start])

		values, ok := params[paramName].([]string)
		if !ok {
			return fmt.Errorf("could not read param %q as array", paramName)
		}

		for index, value := range values {
			indexParamName := paramName + strconv.Itoa(index)

			builder.WriteByte(query[start])
			builder.WriteString(indexParamName)

			if index < len(values)-1 {
				builder.WriteByte(',')
			}

			params[indexParamName] = value
		}

		delete(params, paramName)

		count = end
	}

	builder.WriteString(query[count:])
	query = builder.String()

	names := []any{}
	for name, value := range params {
		names = append(names, sql.Named(name, value))
	}

	row := c.DB.QueryRowContext(ctx, query, names...)

	err := row.Err()
	if err != nil {
		return fmt.Errorf("could not execute query: %w", err)
	}

	err = row.Scan(destination)
	if err != nil {
		return fmt.Errorf("could not scan result: %w", err)
	}

	return nil
}
