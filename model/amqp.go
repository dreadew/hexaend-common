package model

const (
	AMQPEmailEvent = "email"

	TwoStepActivatePath = "/account/two-step"
	ConfirmEmailPath    = "/account/email"
)

type AmqpEvent struct {
	EventType string
	Payload   []byte
}

type AmqpEmailEvent struct {
	EventType CodeType
	Payload   []byte
}

type AmqpEmailBody struct {
	Email    string
	Username string
	Code     string
	Path     string
	Type     string
}
