package messages

const (
	StatusSentEmailAuth          = 251
	StatusInputEmailAuth         = 252
	StatusFailedEmailAuth        = 451
	StatusFailedNotHaveAuthEmail = 452
)

var statusText = map[int]string{
	StatusSentEmailAuth:          "Sent email for login authentication",
	StatusInputEmailAuth:         "Input email address for login authentication",
	StatusFailedEmailAuth:        "Failed email authentication: not match secret key",
	StatusFailedNotHaveAuthEmail: "Failed email authentication: not have auth email",
}

func RestStatusText(code int) string {
	return statusText[code]
}
