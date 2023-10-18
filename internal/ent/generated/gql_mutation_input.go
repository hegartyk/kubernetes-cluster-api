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

package generated

import (
	"go.infratographer.com/x/gidx"
)

// CreateClusterInput represents a mutation input for creating clusters.
type CreateClusterInput struct {
	Name        string
	Description *string
	TenantID    gidx.PrefixedID
}

// Mutate applies the CreateClusterInput on the ClusterMutation builder.
func (i *CreateClusterInput) Mutate(m *ClusterMutation) {
	m.SetName(i.Name)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	m.SetTenantID(i.TenantID)
}

// SetInput applies the change-set in the CreateClusterInput on the ClusterCreate builder.
func (c *ClusterCreate) SetInput(i CreateClusterInput) *ClusterCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateClusterInput represents a mutation input for updating clusters.
type UpdateClusterInput struct {
	Name             *string
	ClearDescription bool
	Description      *string
}

// Mutate applies the UpdateClusterInput on the ClusterMutation builder.
func (i *UpdateClusterInput) Mutate(m *ClusterMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearDescription {
		m.ClearDescription()
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
}

// SetInput applies the change-set in the UpdateClusterInput on the ClusterUpdate builder.
func (c *ClusterUpdate) SetInput(i UpdateClusterInput) *ClusterUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateClusterInput on the ClusterUpdateOne builder.
func (c *ClusterUpdateOne) SetInput(i UpdateClusterInput) *ClusterUpdateOne {
	i.Mutate(c.Mutation())
	return c
}