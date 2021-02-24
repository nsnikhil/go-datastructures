package erx

//Severity of the error
type Severity string

//Prebuilt Severities
const (
	SeverityError Severity = "error"
	SeverityWarn  Severity = "warn"
	SeverityInfo  Severity = "info"
	SeverityDebug Severity = "debug"
)
