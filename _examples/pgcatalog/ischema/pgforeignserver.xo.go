package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"

	"github.com/blakearnold/xo/_examples/pgcatalog/pgtypes"
)

// PgForeignServer represents a row from 'information_schema._pg_foreign_servers'.
type PgForeignServer struct {
	Oid                       pgtypes.NullOid  `json:"oid"`                          // oid
	Srvoptions                []sql.NullString `json:"srvoptions"`                   // srvoptions
	ForeignServerCatalog      sql.NullString   `json:"foreign_server_catalog"`       // foreign_server_catalog
	ForeignServerName         sql.NullString   `json:"foreign_server_name"`          // foreign_server_name
	ForeignDataWrapperCatalog sql.NullString   `json:"foreign_data_wrapper_catalog"` // foreign_data_wrapper_catalog
	ForeignDataWrapperName    sql.NullString   `json:"foreign_data_wrapper_name"`    // foreign_data_wrapper_name
	ForeignServerType         sql.NullString   `json:"foreign_server_type"`          // foreign_server_type
	ForeignServerVersion      sql.NullString   `json:"foreign_server_version"`       // foreign_server_version
	AuthorizationIdentifier   sql.NullString   `json:"authorization_identifier"`     // authorization_identifier
}
