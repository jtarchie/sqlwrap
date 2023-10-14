package sqlwrap

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	*sql.DB
}

type (
	Params map[string]any
	Values []any
)

//go:generate ragel -e -G2 -Z prepare_in_list.rl

func Open(driver string, connection string) (*Client, error) {
	database, err := sql.Open(driver, connection)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	return &Client{
		DB: database,
	}, nil
}

func In(query string, params Params) (string, error) {
	if len(params) == 0 {
		return query, nil
	}

	for _, value := range params {
		if _, ok := value.(Values); ok {
			query, err := prepareInList(query, params)
			if err != nil {
				return "", fmt.Errorf("could not prepare query for IN list: %w", err)
			}

			return query, nil
		}
	}

	return query, nil
}

func (c *Client) Get(
	ctx context.Context,
	destination interface{},
	query string,
	params Params,
) error {
	var err error

	query, err = In(query, params)
	if err != nil {
		return fmt.Errorf("could not setup IN list for GET: %w", err)
	}

	names := make([]any, 0, len(params))
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
