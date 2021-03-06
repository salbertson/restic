package resourcehealth

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AvailabilityStateValues enumerates the values for availability state values.
type AvailabilityStateValues string

const (
	// Available ...
	Available AvailabilityStateValues = "Available"
	// Unavailable ...
	Unavailable AvailabilityStateValues = "Unavailable"
	// Unknown ...
	Unknown AvailabilityStateValues = "Unknown"
)

// PossibleAvailabilityStateValuesValues returns an array of possible values for the AvailabilityStateValues const type.
func PossibleAvailabilityStateValuesValues() []AvailabilityStateValues {
	return []AvailabilityStateValues{Available, Unavailable, Unknown}
}

// ReasonChronicityTypes enumerates the values for reason chronicity types.
type ReasonChronicityTypes string

const (
	// Persistent ...
	Persistent ReasonChronicityTypes = "Persistent"
	// Transient ...
	Transient ReasonChronicityTypes = "Transient"
)

// PossibleReasonChronicityTypesValues returns an array of possible values for the ReasonChronicityTypes const type.
func PossibleReasonChronicityTypesValues() []ReasonChronicityTypes {
	return []ReasonChronicityTypes{Persistent, Transient}
}

// AvailabilityStatus availabilityStatus of a resource.
type AvailabilityStatus struct {
	autorest.Response `json:"-"`
	// ID - Azure Resource Manager Identity for the availabilityStatuses resource.
	ID *string `json:"id,omitempty"`
	// Name - current.
	Name *string `json:"name,omitempty"`
	// Type - Microsoft.ResourceHealth/AvailabilityStatuses.
	Type *string `json:"type,omitempty"`
	// Location - Azure Resource Manager geo location of the resource.
	Location *string `json:"location,omitempty"`
	// Properties - Properties of availability state.
	Properties *AvailabilityStatusProperties `json:"properties,omitempty"`
}

// AvailabilityStatusListResult the List availabilityStatus operation response.
type AvailabilityStatusListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of availabilityStatuses.
	Value *[]AvailabilityStatus `json:"value,omitempty"`
	// NextLink - The URI to fetch the next page of availabilityStatuses. Call ListNext() with this URI to fetch the next page of availabilityStatuses.
	NextLink *string `json:"nextLink,omitempty"`
}

// AvailabilityStatusListResultIterator provides access to a complete listing of AvailabilityStatus values.
type AvailabilityStatusListResultIterator struct {
	i    int
	page AvailabilityStatusListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AvailabilityStatusListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AvailabilityStatusListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AvailabilityStatusListResultIterator) Response() AvailabilityStatusListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AvailabilityStatusListResultIterator) Value() AvailabilityStatus {
	if !iter.page.NotDone() {
		return AvailabilityStatus{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (aslr AvailabilityStatusListResult) IsEmpty() bool {
	return aslr.Value == nil || len(*aslr.Value) == 0
}

// availabilityStatusListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (aslr AvailabilityStatusListResult) availabilityStatusListResultPreparer() (*http.Request, error) {
	if aslr.NextLink == nil || len(to.String(aslr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(aslr.NextLink)))
}

// AvailabilityStatusListResultPage contains a page of AvailabilityStatus values.
type AvailabilityStatusListResultPage struct {
	fn   func(AvailabilityStatusListResult) (AvailabilityStatusListResult, error)
	aslr AvailabilityStatusListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AvailabilityStatusListResultPage) Next() error {
	next, err := page.fn(page.aslr)
	if err != nil {
		return err
	}
	page.aslr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AvailabilityStatusListResultPage) NotDone() bool {
	return !page.aslr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AvailabilityStatusListResultPage) Response() AvailabilityStatusListResult {
	return page.aslr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AvailabilityStatusListResultPage) Values() []AvailabilityStatus {
	if page.aslr.IsEmpty() {
		return nil
	}
	return *page.aslr.Value
}

// AvailabilityStatusProperties properties of availability state.
type AvailabilityStatusProperties struct {
	// AvailabilityState - Availability status of the resource. Possible values include: 'Available', 'Unavailable', 'Unknown'
	AvailabilityState AvailabilityStateValues `json:"availabilityState,omitempty"`
	// Summary - Summary description of the availability state.
	Summary *string `json:"summary,omitempty"`
	// DetailedStatus - Details of the availability status.
	DetailedStatus *string `json:"detailedStatus,omitempty"`
	// ReasonType - When the resource's availabilityState is Unavailable, it describes where the health impacting event was originated. Examples are planned, unplanned, user initiated or an outage etc.
	ReasonType *string `json:"reasonType,omitempty"`
	// RootCauseAttributionTime - When the resource's availabilityState is Unavailable, it provides the Timestamp for when the health impacting event was received.
	RootCauseAttributionTime *date.Time `json:"rootCauseAttributionTime,omitempty"`
	// ResolutionETA - When the resource's availabilityState is Unavailable and the reasonType is not User Initiated, it provides the date and time for when the issue is expected to be resolved.
	ResolutionETA *date.Time `json:"resolutionETA,omitempty"`
	// OccuredTime - Timestamp for when last change in health status occurred.
	OccuredTime *date.Time `json:"occuredTime,omitempty"`
	// ReasonChronicity - Chronicity of the availability transition. Possible values include: 'Transient', 'Persistent'
	ReasonChronicity ReasonChronicityTypes `json:"reasonChronicity,omitempty"`
	// ReportedTime - Timestamp for when the health was last checked.
	ReportedTime *date.Time `json:"reportedTime,omitempty"`
	// RecentlyResolvedState - An annotation describing a change in the availabilityState to Available from Unavailable with a reasonType of type Unplanned
	RecentlyResolvedState *AvailabilityStatusPropertiesRecentlyResolvedState `json:"recentlyResolvedState,omitempty"`
	// RecommendedActions - Lists actions the user can take based on the current availabilityState of the resource.
	RecommendedActions *[]RecommendedAction `json:"recommendedActions,omitempty"`
	// ServiceImpactingEvents - Lists the service impacting events that may be affecting the health of the resource.
	ServiceImpactingEvents *[]ServiceImpactingEvent `json:"serviceImpactingEvents,omitempty"`
}

// AvailabilityStatusPropertiesRecentlyResolvedState an annotation describing a change in the availabilityState to
// Available from Unavailable with a reasonType of type Unplanned
type AvailabilityStatusPropertiesRecentlyResolvedState struct {
	// UnavailableOccurredTime - Timestamp for when the availabilityState changed to Unavailable
	UnavailableOccurredTime *date.Time `json:"unavailableOccurredTime,omitempty"`
	// ResolvedTime - Timestamp when the availabilityState changes to Available.
	ResolvedTime *date.Time `json:"resolvedTime,omitempty"`
	// UnavailabilitySummary - Brief description of cause of the resource becoming unavailable.
	UnavailabilitySummary *string `json:"unavailabilitySummary,omitempty"`
}

// ErrorResponse error details.
type ErrorResponse struct {
	// Code - The error code.
	Code *string `json:"code,omitempty"`
	// Message - The error message.
	Message *string `json:"message,omitempty"`
	// Details - The error details.
	Details *string `json:"details,omitempty"`
}

// Operation operation available in the resourcehealth resource provider.
type Operation struct {
	// Name - Name of the operation.
	Name *string `json:"name,omitempty"`
	// Display - Properties of the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay properties of the operation.
type OperationDisplay struct {
	// Provider - Provider name.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource name.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation.
	Description *string `json:"description,omitempty"`
}

// OperationListResult lists the operations response.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of operations available in the resourcehealth resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// RecommendedAction lists actions the user can take based on the current availabilityState of the resource.
type RecommendedAction struct {
	// Action - Recommended action.
	Action *string `json:"action,omitempty"`
	// ActionURL - Link to the action
	ActionURL *string `json:"actionUrl,omitempty"`
	// ActionURLText - Substring of action, it describes which text should host the action url.
	ActionURLText *string `json:"actionUrlText,omitempty"`
}

// ServiceImpactingEvent lists the service impacting events that may be affecting the health of the resource.
type ServiceImpactingEvent struct {
	// EventStartTime - Timestamp for when the event started.
	EventStartTime *date.Time `json:"eventStartTime,omitempty"`
	// EventStatusLastModifiedTime - Timestamp for when event was submitted/detected.
	EventStatusLastModifiedTime *date.Time `json:"eventStatusLastModifiedTime,omitempty"`
	// CorrelationID - Correlation id for the event
	CorrelationID *string `json:"correlationId,omitempty"`
	// Status - Status of the service impacting event.
	Status *ServiceImpactingEventStatus `json:"status,omitempty"`
	// IncidentProperties - Properties of the service impacting event.
	IncidentProperties *ServiceImpactingEventIncidentProperties `json:"incidentProperties,omitempty"`
}

// ServiceImpactingEventIncidentProperties properties of the service impacting event.
type ServiceImpactingEventIncidentProperties struct {
	// Title - Title of the incident.
	Title *string `json:"title,omitempty"`
	// Service - Service impacted by the event.
	Service *string `json:"service,omitempty"`
	// Region - Region impacted by the event.
	Region *string `json:"region,omitempty"`
	// IncidentType - Type of Event.
	IncidentType *string `json:"incidentType,omitempty"`
}

// ServiceImpactingEventStatus status of the service impacting event.
type ServiceImpactingEventStatus struct {
	// Value - Current status of the event
	Value *string `json:"value,omitempty"`
}
