package registration

type RegistrationStatusID string

const (
	StatusPending   RegistrationStatusID = "PENDING"
	StatusConfirmed RegistrationStatusID = "CONFIRMED"
	StatusCancelled RegistrationStatusID = "CANCELLED"
)
