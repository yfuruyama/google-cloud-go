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

// [START essentialcontacts_v1_generated_EssentialContactsService_SendTestMessage_sync]

package main

import (
	"context"

	essentialcontacts "cloud.google.com/go/essentialcontacts/apiv1"
	essentialcontactspb "google.golang.org/genproto/googleapis/cloud/essentialcontacts/v1"
)

func main() {
	ctx := context.Background()
	c, err := essentialcontacts.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &essentialcontactspb.SendTestMessageRequest{
		// TODO: Fill request struct fields.
	}
	err = c.SendTestMessage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END essentialcontacts_v1_generated_EssentialContactsService_SendTestMessage_sync]
