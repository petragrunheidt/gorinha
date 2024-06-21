package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ExecuteQuery(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := queryRows(context.Background(), query, args...)
	if err != nil {
			return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	results, err := parseRows(rows)
	if err != nil {
			return nil, fmt.Errorf("error parsing rows: %w", err)
	}

	return results, nil
}

func queryRows(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	rows, err := DBPool.Query(ctx, query, args...)
	if err != nil {
			return nil, fmt.Errorf("error querying database: %w", err)
	}
	return rows, nil
}

func parseRows(rows pgx.Rows) ([]map[string]interface{}, error) {
	defer rows.Close()

	var results []map[string]interface{}

	columns := rows.FieldDescriptions()

	colNames := make([]string, len(columns))
	for i, col := range columns {
			colNames[i] = string(col.Name)
	}

    for rows.Next() {
        scanArgs := make([]interface{}, len(columns))
        for i := range colNames {
            scanArgs[i] = new(interface{})
        }

        if err := rows.Scan(scanArgs...); err != nil {
            return nil, fmt.Errorf("error scanning row: %w", err)
        }

        row := make(map[string]interface{})
        for i, colName := range colNames {
            row[colName] = *scanArgs[i].(*interface{})
        }
        results = append(results, row)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error after iterating rows: %w", err)
    }

    return results, nil
}
