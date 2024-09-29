package utils

import "github.com/jackc/pgx/v5/pgtype"

func ConvertNumericToFloat(n pgtype.Numeric) float64 {
	if !n.Valid {
		return 0
	}
	f, _ := n.Float64Value()
	return f.Float64
}
