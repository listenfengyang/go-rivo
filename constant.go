package rivo

const (
	SignTypeSHA512 = "SHA512"
	Version10      = "1.0"
	Version11      = "1.1"

	// Backward compatibility aliases.
	SIGN_TYPE_SHA512 = SignTypeSHA512
	VERSION_1_0      = Version10
	VERSION_1_1      = Version11

	// Pay-In Status
	PayInStatusPending = 0
	PayInStatusSuccess = 1
	PayInStatusFailed  = 2

	// Backward compatibility aliases.
	PAY_STATUS_PENDING = PayInStatusPending
	PAY_STATUS_SUCCESS = PayInStatusSuccess
	PAY_STATUS_FAILED  = PayInStatusFailed

	// Pay-Out Status
	PayOutStatusPending    = 0
	PayOutStatusSuccess    = 1
	PayOutStatusFailed     = 2
	PayOutStatusProcessing = 3

	// Backward compatibility aliases.
	STATUS_PENDING    = PayOutStatusPending
	STATUS_SUCCESS    = PayOutStatusSuccess
	STATUS_FAILED     = PayOutStatusFailed
	STATUS_PROCESSING = PayOutStatusProcessing
)
