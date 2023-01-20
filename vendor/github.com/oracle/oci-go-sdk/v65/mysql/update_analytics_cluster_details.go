// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAnalyticsClusterDetails DEPRECATED -- please use HeatWave API instead.
// Details about the Analytics Cluster properties to be updated.
type UpdateAnalyticsClusterDetails struct {

	// A change to the shape of the nodes in the Analytics Cluster will
	// result in the entire cluster being torn down and re-created with
	// Compute instances of the new Shape. This may result in significant
	// downtime for the analytics capability while the Analytics Cluster is
	// re-provisioned.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// A change to the number of nodes in the Analytics Cluster will result
	// in the entire cluster being torn down and re-created with the new
	// cluster of nodes. This may result in a significant downtime for the
	// analytics capability while the Analytics Cluster is
	// re-provisioned.
	ClusterSize *int `mandatory:"false" json:"clusterSize"`
}

func (m UpdateAnalyticsClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAnalyticsClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
