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

type Params map[string]interface{}

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

func Get[T comparable](
	client *Client,
	ctx context.Context,
	destination T,
	query string,
	params Params,
) error {
	var err error

	if len(params) > 0 {
		query, err = prepareInList(query, params)
		if err != nil {
			return fmt.Errorf("could not prepare query for IN list: %w", err)
		}
	}

	names := make([]any, 0, len(params))
	for name, value := range params {
		names = append(names, sql.Named(name, value))
	}

	row := client.DB.QueryRowContext(ctx, query, names...)

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

func (c *Client) Get(
	ctx context.Context,
	destination interface{},
	query string,
	params Params,
) error {
	return Get(
		c,
		ctx,
		destination,
		query,
		params,
	)
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
