package model

type CodeType int
type CodeStatus int

const (
	EmailConfirmation CodeType = iota
	VerifyTwoStep
	ActivateTwoStep
	RecoverPassword
)

const (
	Created CodeStatus = iota
	Approved
	Expired
)
