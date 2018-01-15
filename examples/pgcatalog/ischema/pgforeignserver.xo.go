// Package ischema contains the types for schema 'information_schema'.
package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"

	"github.com/kaddiya/xoexamples/pgcatalog/pgtypes"
)

// PgForeignServer represents a row from 'information_schema._pg_foreign_servers'.
type PgForeignServer struct {
	Oid                       pgtypes.Oid           `json:"oid"`                          // oid
	Srvoptions                []sql.NullString      `json:"srvoptions"`                   // srvoptions
	ForeignServerCatalog      pgtypes.SQLIdentifier `json:"foreign_server_catalog"`       // foreign_server_catalog
	ForeignServerName         pgtypes.SQLIdentifier `json:"foreign_server_name"`          // foreign_server_name
	ForeignDataWrapperCatalog pgtypes.SQLIdentifier `json:"foreign_data_wrapper_catalog"` // foreign_data_wrapper_catalog
	ForeignDataWrapperName    pgtypes.SQLIdentifier `json:"foreign_data_wrapper_name"`    // foreign_data_wrapper_name
	ForeignServerType         pgtypes.CharacterData `json:"foreign_server_type"`          // foreign_server_type
	ForeignServerVersion      pgtypes.CharacterData `json:"foreign_server_version"`       // foreign_server_version
	AuthorizationIdentifier   pgtypes.SQLIdentifier `json:"authorization_identifier"`     // authorization_identifier
}
