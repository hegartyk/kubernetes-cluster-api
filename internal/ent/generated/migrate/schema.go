// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ClustersColumns holds the columns for the "clusters" table.
	ClustersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "tenant_id", Type: field.TypeString},
	}
	// ClustersTable holds the schema information for the "clusters" table.
	ClustersTable = &schema.Table{
		Name:       "clusters",
		Columns:    ClustersColumns,
		PrimaryKey: []*schema.Column{ClustersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "cluster_created_at",
				Unique:  false,
				Columns: []*schema.Column{ClustersColumns[1]},
			},
			{
				Name:    "cluster_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ClustersColumns[2]},
			},
			{
				Name:    "cluster_tenant_id",
				Unique:  false,
				Columns: []*schema.Column{ClustersColumns[5]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ClustersTable,
	}
)

func init() {
}
