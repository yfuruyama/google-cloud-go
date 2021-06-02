// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START binaryauthorization_v1beta1_generated_BinauthzManagementServiceV1Beta1_UpdateAttestor_sync]

package main

import (
	"context"

	binaryauthorization "cloud.google.com/go/binaryauthorization/apiv1beta1"
	binaryauthorizationpb "google.golang.org/genproto/googleapis/cloud/binaryauthorization/v1beta1"
)

func main() {
	ctx := context.Background()
	c, err := binaryauthorization.NewBinauthzManagementServiceV1Beta1Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &binaryauthorizationpb.UpdateAttestorRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateAttestor(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END binaryauthorization_v1beta1_generated_BinauthzManagementServiceV1Beta1_UpdateAttestor_sync]
