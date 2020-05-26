package messages

const (
	StatusEmailAuthConfirm = 251
	StatusInputEmailAuth = 252
	StatusFailedEmailAuth = 451
)

var statusText = map[int]string{
	StatusEmailAuthConfirm: "Sent email for login authentication",
	StatusInputEmailAuth: "Input email address for login authentication",
	StatusFailedEmailAuth: "Failed email authentication: not match secret key",
}

func RestStatusText(code int) string {
	return statusText[code]
}
