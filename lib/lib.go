package lib

import "encoding/json"

// Schema of the JSON message sent to Pub/Sub.
//
// https://cloud.google.com/monitoring/support/notification-options#creating_channels
type AlertingNotification struct {
	Version string `json:"version"`

	// AlertingNotificationIncident
	Incident json.RawMessage `json:"incident"`
}

type AlertingNotificationIncident struct {
	IncidentID           string `json:"incident_id"`
	ScopingProjectID     string `json:"scoping_project_id"`
	ScopingProjectNumber int64  `json:"scoping_project_number"`
	// URL                     string
	StartedAt int64 `json:"started_at"`
	EndedAt   int64 `json:"ended_at"`
	State     string
	Summary   string
	// ApigeeURL               string `json:"apigee_url"`
	// ObservedValue           string `json:"observed_value"`
	// Resource                AlertingNotificationResource
	// ResourceTypeDisplayName string `json:"resource_type_display_name"`
	// Severity                string
	Condition AlertingNotificationCondition
}

type AlertingNotificationResource struct {
	Type   string
	Labels map[string]string
}

type AlertingNotificationCondition struct {
}
