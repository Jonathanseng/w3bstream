package models

import (
	"github.com/machinefi/w3bstream/pkg/depends/base/types"
	"github.com/machinefi/w3bstream/pkg/depends/kit/sqlx/datatypes"
	"github.com/machinefi/w3bstream/pkg/enums"
)

// Project schema for project information
// @def primary                    ID
// @def unique_index UI_project_id ProjectID
// @def unique_index UI_name       Name
//
//go:generate toolkit gen model Project --database DB
type Project struct {
	datatypes.PrimaryID
	RelProject
	RelAccount
	ProjectInfo
	datatypes.OperationTimesWithDeleted
}

type RelProject struct {
	ProjectID types.SFID `db:"f_project_id" json:"projectID"`
}

type ProjectInfo struct {
	Name    string         `db:"f_name"               json:"name"`               // Name project name
	Version string         `db:"f_version,default=''" json:"version,omitempty"`  // Version project version
	Proto   enums.Protocol `db:"f_proto,default='0'"  json:"protocol,omitempty"` // Proto project protocol for event publisher
}
