// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package reservation_test

import (
	"context"

	reservation "cloud.google.com/go/bigquery/reservation/apiv1"
	"google.golang.org/api/iterator"
	reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_CreateReservation() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.CreateReservationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateReservation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListReservations() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.ListReservationsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListReservations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetReservation() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.GetReservationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetReservation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteReservation() {
	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.DeleteReservationRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteReservation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_UpdateReservation() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.UpdateReservationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateReservation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateCapacityCommitment() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.CreateCapacityCommitmentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateCapacityCommitment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListCapacityCommitments() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.ListCapacityCommitmentsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCapacityCommitments(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetCapacityCommitment() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.GetCapacityCommitmentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCapacityCommitment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteCapacityCommitment() {
	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.DeleteCapacityCommitmentRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteCapacityCommitment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_UpdateCapacityCommitment() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.UpdateCapacityCommitmentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateCapacityCommitment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_SplitCapacityCommitment() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.SplitCapacityCommitmentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SplitCapacityCommitment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_MergeCapacityCommitments() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.MergeCapacityCommitmentsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.MergeCapacityCommitments(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateAssignment() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.CreateAssignmentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateAssignment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListAssignments() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.ListAssignmentsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListAssignments(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_DeleteAssignment() {
	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.DeleteAssignmentRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteAssignment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_SearchAssignments() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.SearchAssignmentsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.SearchAssignments(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_MoveAssignment() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.MoveAssignmentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.MoveAssignment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetBiReservation() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.GetBiReservationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetBiReservation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateBiReservation() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.UpdateBiReservationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateBiReservation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}