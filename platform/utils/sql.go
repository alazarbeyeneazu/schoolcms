package utils

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
)

// IntOrNull returns properly configured sql.NullInt64
func IntOrNull(n int64) sql.NullInt64 {
	return sql.NullInt64{Int64: n, Valid: true}
}

// Int32OrNull returns properly configured sql.NullInt64
func Int32OrNull(n int32) sql.NullInt32 {
	return sql.NullInt32{Int32: n, Valid: true}
}

// PositiveInt32OrNull returns properly configured sql.NullInt64
func PositiveInt32OrNull(n int32) sql.NullInt32 {
	if n <= 0 {
		return sql.NullInt32{Int32: 0, Valid: false}
	}
	return sql.NullInt32{Int32: n, Valid: true}
}

// Int64OrNull returns properly configured sql.NullInt64
func Int64OrNull(n int64) sql.NullInt64 {
	return sql.NullInt64{Int64: n, Valid: true}
}

// Int64OrNull returns properly configured sql.NullInt64
func Int64OrNullFromPntr(n *int64) sql.NullInt64 {
	if n != nil {
		return sql.NullInt64{Int64: *n, Valid: true}
	}
	return sql.NullInt64{Int64: 0, Valid: false}
}

// PositiveInt64OrNull returns properly configured sql.NullInt64
func PositiveInt64OrNull(n int64) sql.NullInt64 {
	if n <= 0 {
		return sql.NullInt64{Int64: 0, Valid: false}
	}
	return sql.NullInt64{Int64: n, Valid: true}
}

// PositiveIntOrNull returns properly configured sql.NullInt64 for a positive number
func PositiveIntOrNull(n int64) sql.NullInt64 {
	if n <= 0 {
		return sql.NullInt64{Int64: 0, Valid: false}
	}
	return sql.NullInt64{Int64: n, Valid: true}
}

// FloatOrNull returns properly configured sql.NullFloat64
func FloatOrNull(n float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: n, Valid: true}
}

// PositiveFloatOrNull returns properly configured sql.NullFloat64 for a positive number
func PositiveFloatOrNull(n float64) sql.NullFloat64 {
	if n <= 0 {
		return sql.NullFloat64{Float64: 0.0, Valid: false}
	}
	return sql.NullFloat64{Float64: n, Valid: true}
}

// StringOrNull returns properly configured sql.NullString
func StringOrNull(str string) sql.NullString {
	if str == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: str, Valid: true}
}

// StringOrNull returns properly configured sql.NullString
func StringOrNullFromPntr(str *string) sql.NullString {
	if str != nil {
		if *str != "" {
			return sql.NullString{String: *str, Valid: true}
		}
	}
	return sql.NullString{String: "", Valid: false}
}

// DecimalOrNull returns properly configured sql.NullString
func DecimalOrNull(n decimal.Decimal) decimal.NullDecimal {
	if n.IsZero() {
		return decimal.NullDecimal{
			Decimal: decimal.NewFromInt(0),
			Valid:   false}
	}
	return decimal.NullDecimal{
		Decimal: n,
		Valid:   true,
	}
}

// DecimalOrNull returns properly configured sql.NullString
func DecimalOrNullFromPntr(n *decimal.Decimal) decimal.NullDecimal {
	if n != nil {
		return decimal.NullDecimal{
			Decimal: *n,
			Valid:   true}
	}
	return decimal.NullDecimal{
		Decimal: decimal.Decimal{},
		Valid:   false,
	}
}

// PositiveDecimalOrNull returns properly configured sql.NullString
func PositiveDecimalOrNull(n decimal.Decimal) decimal.NullDecimal {
	if n.IsZero() {
		return decimal.NullDecimal{
			Decimal: decimal.NewFromInt(0),
			Valid:   false}
	}
	return decimal.NullDecimal{
		Decimal: n,
		Valid:   true,
	}
}

// TimeOrNull returns properly configured pq.TimeNull
func TimeOrNull(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Time: time.Time{}, Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

// UUIDOrNull returns properly confiigured uuid.NullUUID
func UUIDOrNull(t uuid.UUID) uuid.NullUUID {
	if t == uuid.Nil {
		return uuid.NullUUID{UUID: uuid.UUID{}, Valid: false}
	}

	return uuid.NullUUID{UUID: t, Valid: true}
}

// BoolOrNull returns properly confiigured uuid.NullUUID
func BoolOrNull(t bool) sql.NullBool {
	if !t {
		return sql.NullBool{Bool: t, Valid: false}
	}

	return sql.NullBool{Bool: t, Valid: true}
}

// BoolOrNull returns properly confiigured uuid.NullUUID
func BoolOrNullFromPntr(t *bool) sql.NullBool {
	if t != nil {
		return sql.NullBool{Bool: *t, Valid: true}
	}
	return sql.NullBool{Bool: false, Valid: false}
}

func MapJSONOrNull(t []byte) pgtype.JSON {
	if t == nil {
		return pgtype.JSON{
			Status: pgtype.Null,
		}
	}

	return pgtype.JSON{
		Bytes:  t,
		Status: pgtype.Present,
	}
}
