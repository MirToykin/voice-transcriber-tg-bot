package sqlc

import "database/sql"

func stringToSQLNullString(s string) sql.NullString {
	return sql.NullString{String: s}
}

func intToSQLNullInt(i int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(i)}
}
