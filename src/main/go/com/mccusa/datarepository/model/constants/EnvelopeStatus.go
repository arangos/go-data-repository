package constants

type EnvelopeStatus string

const (
	// EnvelopeStatusSent means the envelope has been sent but not yet completed
	EnvelopeStatusSent EnvelopeStatus = "sent"

	// EnvelopeStatusCompleted means the envelope process has finished successfully
	EnvelopeStatusCompleted EnvelopeStatus = "completed"

	// EnvelopeStatusVoid means the envelope was voided/cancelled
	EnvelopeStatusVoid EnvelopeStatus = "void"
)
