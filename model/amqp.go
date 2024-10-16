package model

const (
	AMQPEmailEvent             = "email"
	AMQPProjectEvent           = "project"
	AMQPWorkspaceEvent         = "workspace"
	AMQPDeleteByProjectEvent   = "project.delete"
	AMQPDeleteByWorkspaceEvent = "workspace.delete"

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

type AmqpDeleteByWorkspaceEvent struct {
	WorkspaceID string
}

type AmqpDeleteByProjectEvent struct {
	ID string
}
