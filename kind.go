package erx

//Kind of error
type Kind string

//Prebuilt Kinds
const (
	ValidationError         Kind = "validationError"
	InternalError           Kind = "internalError"
	AuthenticationError     Kind = "authenticationError"
	InvalidCredentialsError Kind = "invalidCredentialsError"
	InvalidArgsError        Kind = "invalidArgsError"
	InitializationError     Kind = "initializationError"
	ResourceNotFoundError   Kind = "resourceNotFound"
	DuplicateRecordError    Kind = "duplicateRecordError"
	ProducerError           Kind = "producerError"
	ConsumerError           Kind = "consumerError"
)
