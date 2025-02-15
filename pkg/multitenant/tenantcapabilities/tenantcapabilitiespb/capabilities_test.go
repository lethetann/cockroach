// Copyright 2023 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package tenantcapabilitiespb

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/multitenant/tenantcapabilities"
)

func TestCapabilityGetSet(t *testing.T) {
	capabilities := TenantCapabilities{}
	for _, capID := range tenantcapabilities.CapabilityIDs {
		capability := capabilities.Cap(capID)
		value := capability.Get().Unwrap()
		capability.Set(value)
	}
}
