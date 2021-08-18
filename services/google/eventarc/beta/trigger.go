// Copyright 2021 Google LLC. All Rights Reserved.
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
package beta

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Trigger struct {
	Name             *string                   `json:"name"`
	Uid              *string                   `json:"uid"`
	CreateTime       *string                   `json:"createTime"`
	UpdateTime       *string                   `json:"updateTime"`
	MatchingCriteria []TriggerMatchingCriteria `json:"matchingCriteria"`
	ServiceAccount   *string                   `json:"serviceAccount"`
	Destination      *TriggerDestination       `json:"destination"`
	Transport        *TriggerTransport         `json:"transport"`
	Labels           map[string]string         `json:"labels"`
	Etag             *string                   `json:"etag"`
	Project          *string                   `json:"project"`
	Location         *string                   `json:"location"`
}

func (r *Trigger) String() string {
	return dcl.SprintResource(r)
}

type TriggerMatchingCriteria struct {
	empty     bool    `json:"-"`
	Attribute *string `json:"attribute"`
	Value     *string `json:"value"`
}

type jsonTriggerMatchingCriteria TriggerMatchingCriteria

func (r *TriggerMatchingCriteria) UnmarshalJSON(data []byte) error {
	var res jsonTriggerMatchingCriteria
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTriggerMatchingCriteria
	} else {

		r.Attribute = res.Attribute

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this TriggerMatchingCriteria is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTriggerMatchingCriteria *TriggerMatchingCriteria = &TriggerMatchingCriteria{empty: true}

func (r *TriggerMatchingCriteria) Empty() bool {
	return r.empty
}

func (r *TriggerMatchingCriteria) String() string {
	return dcl.SprintResource(r)
}

func (r *TriggerMatchingCriteria) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TriggerDestination struct {
	empty           bool                               `json:"-"`
	CloudRunService *TriggerDestinationCloudRunService `json:"cloudRunService"`
	CloudFunction   *string                            `json:"cloudFunction"`
}

type jsonTriggerDestination TriggerDestination

func (r *TriggerDestination) UnmarshalJSON(data []byte) error {
	var res jsonTriggerDestination
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTriggerDestination
	} else {

		r.CloudRunService = res.CloudRunService

		r.CloudFunction = res.CloudFunction

	}
	return nil
}

// This object is used to assert a desired state where this TriggerDestination is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTriggerDestination *TriggerDestination = &TriggerDestination{empty: true}

func (r *TriggerDestination) Empty() bool {
	return r.empty
}

func (r *TriggerDestination) String() string {
	return dcl.SprintResource(r)
}

func (r *TriggerDestination) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TriggerDestinationCloudRunService struct {
	empty   bool    `json:"-"`
	Service *string `json:"service"`
	Path    *string `json:"path"`
	Region  *string `json:"region"`
}

type jsonTriggerDestinationCloudRunService TriggerDestinationCloudRunService

func (r *TriggerDestinationCloudRunService) UnmarshalJSON(data []byte) error {
	var res jsonTriggerDestinationCloudRunService
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTriggerDestinationCloudRunService
	} else {

		r.Service = res.Service

		r.Path = res.Path

		r.Region = res.Region

	}
	return nil
}

// This object is used to assert a desired state where this TriggerDestinationCloudRunService is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTriggerDestinationCloudRunService *TriggerDestinationCloudRunService = &TriggerDestinationCloudRunService{empty: true}

func (r *TriggerDestinationCloudRunService) Empty() bool {
	return r.empty
}

func (r *TriggerDestinationCloudRunService) String() string {
	return dcl.SprintResource(r)
}

func (r *TriggerDestinationCloudRunService) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TriggerTransport struct {
	empty  bool                    `json:"-"`
	Pubsub *TriggerTransportPubsub `json:"pubsub"`
}

type jsonTriggerTransport TriggerTransport

func (r *TriggerTransport) UnmarshalJSON(data []byte) error {
	var res jsonTriggerTransport
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTriggerTransport
	} else {

		r.Pubsub = res.Pubsub

	}
	return nil
}

// This object is used to assert a desired state where this TriggerTransport is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTriggerTransport *TriggerTransport = &TriggerTransport{empty: true}

func (r *TriggerTransport) Empty() bool {
	return r.empty
}

func (r *TriggerTransport) String() string {
	return dcl.SprintResource(r)
}

func (r *TriggerTransport) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type TriggerTransportPubsub struct {
	empty        bool    `json:"-"`
	Topic        *string `json:"topic"`
	Subscription *string `json:"subscription"`
}

type jsonTriggerTransportPubsub TriggerTransportPubsub

func (r *TriggerTransportPubsub) UnmarshalJSON(data []byte) error {
	var res jsonTriggerTransportPubsub
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyTriggerTransportPubsub
	} else {

		r.Topic = res.Topic

		r.Subscription = res.Subscription

	}
	return nil
}

// This object is used to assert a desired state where this TriggerTransportPubsub is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyTriggerTransportPubsub *TriggerTransportPubsub = &TriggerTransportPubsub{empty: true}

func (r *TriggerTransportPubsub) Empty() bool {
	return r.empty
}

func (r *TriggerTransportPubsub) String() string {
	return dcl.SprintResource(r)
}

func (r *TriggerTransportPubsub) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Trigger) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "eventarc",
		Type:    "Trigger",
		Version: "beta",
	}
}

func (r *Trigger) ID() (string, error) {
	if err := extractTriggerFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":             dcl.ValueOrEmptyString(nr.Name),
		"uid":              dcl.ValueOrEmptyString(nr.Uid),
		"createTime":       dcl.ValueOrEmptyString(nr.CreateTime),
		"updateTime":       dcl.ValueOrEmptyString(nr.UpdateTime),
		"matchingCriteria": dcl.ValueOrEmptyString(nr.MatchingCriteria),
		"serviceAccount":   dcl.ValueOrEmptyString(nr.ServiceAccount),
		"destination":      dcl.ValueOrEmptyString(nr.Destination),
		"transport":        dcl.ValueOrEmptyString(nr.Transport),
		"labels":           dcl.ValueOrEmptyString(nr.Labels),
		"etag":             dcl.ValueOrEmptyString(nr.Etag),
		"project":          dcl.ValueOrEmptyString(nr.Project),
		"location":         dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/triggers/{{name}}", params), nil
}

const TriggerMaxPage = -1

type TriggerList struct {
	Items []*Trigger

	nextToken string

	pageSize int32

	resource *Trigger
}

func (l *TriggerList) HasNext() bool {
	return l.nextToken != ""
}

func (l *TriggerList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listTrigger(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListTrigger(ctx context.Context, r *Trigger) (*TriggerList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListTriggerWithMaxResults(ctx, r, TriggerMaxPage)

}

func (c *Client) ListTriggerWithMaxResults(ctx context.Context, r *Trigger, pageSize int32) (*TriggerList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listTrigger(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &TriggerList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetTrigger(ctx context.Context, r *Trigger) (*Trigger, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getTriggerRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalTrigger(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeTriggerNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteTrigger(ctx context.Context, r *Trigger) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Trigger resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Trigger...")
	deleteOp := deleteTriggerOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllTrigger deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllTrigger(ctx context.Context, project, location string, filter func(*Trigger) bool) error {
	r := &Trigger{
		Project:  &project,
		Location: &location,
	}
	listObj, err := c.ListTrigger(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllTrigger(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllTrigger(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyTrigger(ctx context.Context, rawDesired *Trigger, opts ...dcl.ApplyOption) (*Trigger, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Trigger
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyTriggerHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyTriggerHelper(c *Client, ctx context.Context, rawDesired *Trigger, opts ...dcl.ApplyOption) (*Trigger, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyTrigger...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractTriggerFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.triggerDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToTriggerDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	var recreate bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				if dcl.HasLifecycleParam(lp, dcl.BlockDestruction) || dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
					return nil, dcl.ApplyInfeasibleError{
						Message: fmt.Sprintf("Infeasible update: (%v) would require recreation.", d),
					}
				}
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []triggerApiOperation
	if create {
		ops = append(ops, &createTriggerOperation{})
	} else if recreate {
		ops = append(ops, &deleteTriggerOperation{})
		ops = append(ops, &createTriggerOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeTriggerDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetTrigger(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createTriggerOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapTrigger(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeTriggerNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeTriggerNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeTriggerDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffTrigger(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
