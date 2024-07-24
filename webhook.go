package auvik

import "time"

type EntityType string
type AlertStatus int
type AlertSeverity int
type AlertSeverityString string
type AlertStatusString string

const (
	EntityTypeDevice    EntityType = "device"
	EntityTypeNetwork   EntityType = "network"
	EntityTypeInterface EntityType = "interface"
	EntityTypeService   EntityType = "service"
)

const (
	AlertTriggered AlertStatus = iota
	AlertCleared
)

const (
	AlertEmergency AlertSeverity = iota + 1
	AlertCritical
	AlertWarning
	AlertInformational
)

const (
	AlertTriggeredString AlertStatusString = "Triggered"
	AlertClearedString   AlertStatusString = "Cleared"
)

const (
	AlertEmergencyString     AlertSeverityString = "Emergency"
	AlertCriticalString      AlertSeverityString = "Cleared"
	AlertWarningString       AlertSeverityString = "Warning"
	AlertInformationalString AlertSeverityString = "Info"
)

func (stat AlertStatusString) String() string {
	return string(stat)
}

func (sev AlertSeverityString) String() string {
	return string(sev)
}

func (stat AlertStatus) String() string {
	switch stat {
	case AlertTriggered:
		return AlertTriggeredString.String()
	case AlertCleared:
		return AlertClearedString.String()
	}
	return "<unknown>"
}

func (sev AlertSeverity) String() string {
	switch sev {
	case AlertEmergency:
		return AlertEmergencyString.String()
	case AlertCritical:
		return AlertCriticalString.String()
	case AlertWarning:
		return AlertWarningString.String()
	case AlertInformational:
		return AlertInformationalString.String()
	}
	return "<unknown>"
}

type AuvikWebhook struct {
	AlertDescription    string        `json:"alertDescription"`
	AlertID             string        `json:"alertId"`
	AlertName           string        `json:"alertName"`
	AlertSeverity       AlertSeverity `json:"alertSeverity"`
	AlertSeverityString string        `json:"alertSeverityString"`
	AlertStatus         AlertStatus   `json:"alertStatus"`
	AlertStatusString   string        `json:"alertStatusString"`
	CompanyID           string        `json:"companyId"`
	CompanyName         string        `json:"companyName"`
	CorrelationID       string        `json:"correlationId"`
	Date                time.Time     `json:"date"`
	EntityID            string        `json:"entityId"`
	EntityName          string        `json:"entityName"`
	EntityType          EntityType    `json:"entityType"`
	Link                string        `json:"link"`
	Subject             string        `json:"subject"`
}
