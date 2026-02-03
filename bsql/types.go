package bsql

import (
	"database/sql"
	"encoding/json"
)

// Wrapper for [sql.NullString] which has custom [NullString.MarshalJSON]
// which return the actual String value instead of the Object.
type NullString struct {
	sql.NullString
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if ns == nil || !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// Wrapper for [sql.NullInt16] which returns the actual int16 value.
type NullInt16 struct {
	sql.NullInt16
}

func (ni *NullInt16) MarshalJSON() ([]byte, error) {
	if ni == nil || !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int16)
}

// Wrapper for [sql.NullInt32] which returns the actual int32 value.
type NullInt32 struct {
	sql.NullInt32
}

func (ni *NullInt32) MarshalJSON() ([]byte, error) {
	if ni == nil || !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int32)
}

// Wrapper for [sql.NullInt64] which returns the actual int64 value.
type NullInt64 struct {
	sql.NullInt64
}

func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if ni == nil || !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// Wrapper for [sql.NullFloat64] which returns the actual float64 value.
type NullFloat64 struct {
	sql.NullFloat64
}

func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if nf == nil || !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

// Wrapper for [sql.NullBool] which returns the actual bool value.
type NullBool struct {
	sql.NullBool
}

func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if nb == nil || !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

// Wrapper for [sql.NullByte] which returns the actual byte value.
type NullByte struct {
	sql.NullByte
}

func (nb *NullByte) MarshalJSON() ([]byte, error) {
	if nb == nil || !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Byte)
}

// Wrapper for [sql.NullTime] which returns the actual time.Time value.
type NullTime struct {
	sql.NullTime
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if nt == nil || !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.Time)
}
