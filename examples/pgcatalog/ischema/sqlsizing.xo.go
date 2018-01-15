// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/kaddiya/xoexamples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// SQLSizing represents a row from 'information_schema.sql_sizing'.
type SQLSizing struct {
	Tableoid       pgtypes.Oid            `json:"tableoid"`        // tableoid
	Cmax           pgtypes.Cid            `json:"cmax"`            // cmax
	Xmax           pgtypes.Xid            `json:"xmax"`            // xmax
	Cmin           pgtypes.Cid            `json:"cmin"`            // cmin
	Xmin           pgtypes.Xid            `json:"xmin"`            // xmin
	Ctid           pgtypes.Tid            `json:"ctid"`            // ctid
	SizingID       pgtypes.CardinalNumber `json:"sizing_id"`       // sizing_id
	SizingName     pgtypes.CharacterData  `json:"sizing_name"`     // sizing_name
	SupportedValue pgtypes.CardinalNumber `json:"supported_value"` // supported_value
	Comments       pgtypes.CharacterData  `json:"comments"`        // comments
}
