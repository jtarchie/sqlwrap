package sqlwrap

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nasa9084/go-builderpool"
)

type Client struct {
	*sql.DB
}

type Params map[string]interface{}

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

func prepareInList(query string, params Params) (string, error) {
	builder, count := builderpool.Get(), 0
	defer builderpool.Release(builder)

	matches := inListRegex.FindAllStringSubmatchIndex(query, -1)
	for _, match := range matches {
		start, end := match[2], match[3]

		paramName := query[start+1 : end]
		if _, ok := params[paramName]; !ok {
			return "", fmt.Errorf("could not find param for IN list %q", paramName)
		}

		builder.WriteString(query[count:start])

		values, ok := params[paramName].([]string)
		if !ok {
			return "", fmt.Errorf("could not read param %q as array", paramName)
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

	return builder.String(), nil
}

func (c *Client) Get(
	ctx context.Context,
	destination interface{},
	query string,
	params Params,
) error {
	var err error

	query, err = prepareInList(query, params)
	if err != nil {
		return fmt.Errorf("could not prepare query for IN list: %w", err)
	}

	names := []any{}
	for name, value := range params {
		names = append(names, sql.Named(name, value))
	}

	row := c.DB.QueryRowContext(ctx, query, names...)

	err = row.Err()
	if err != nil {
		return fmt.Errorf("could not execute query for result: %w", err)
	}

	err = row.Scan(destination)
	if err != nil {
		return fmt.Errorf("could not scan result: %w", err)
	}

	return nil
}

func (c *Client) Select(
	ctx context.Context,
	destination *[]string,
	query string,
	params map[string]interface{},
) error {
	var err error

	query, err = prepareInList(query, params)
	if err != nil {
		return fmt.Errorf("could not prepare query for IN list: %w", err)
	}

	names := []any{}
	for name, value := range params {
		names = append(names, sql.Named(name, value))
	}

	rows, err := c.DB.QueryContext(ctx, query, names...)
	if err != nil {
		return fmt.Errorf("could not execute for results: %w", err)
	}

	if rows.Err() != nil {
		return fmt.Errorf("could not execute for results: %w", err)
	}

	defer rows.Close()

	var value string
	for rows.Next() {
		err := rows.Scan(&value)
		if err != nil {
			return fmt.Errorf("could not scan for result: %w", err)
		}

		*destination = append(*destination, value)
	}

	if err != nil {
		return fmt.Errorf("could not scan result: %w", err)
	}

	return nil
}
