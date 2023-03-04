package database

import "database/sql"

type Scanner[T any] interface {
	Scan(*sql.Rows) (T, error)
}

// GenericDBCall is a way to wrap all of the manual mess that is needed to
// parse Database queries
//
// # Arguments
//
// - query :: SQL Query String
//
// - returnValue :: default value to return if error
//
// - dbValue :: Scanner to use to parse values from Database
//
// - args ... :: position args to pass to Query. Referenced by $ in SQL Query
func GenericDBCall[ReturnValue any](query string, returnValue ReturnValue, dbValue Scanner[ReturnValue], args ...any) (ReturnValue, error) {
	db, err := GetConnection()

	defaultValue := returnValue

	if err != nil {
		return defaultValue, err
	}

	rows, err := db.Query(query, args...)

	if err != nil {
		return defaultValue, err
	}

	defer rows.Close()

	holder := dbValue
	result, err := holder.Scan(rows)

	if err != nil {
		return defaultValue, err
	}

	return result, nil
}
