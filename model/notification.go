package model

type NotificationType int

const (
	Info NotificationType = iota
	Warning
	Error
)
