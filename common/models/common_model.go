package models


/**
 *  For ReCAPTCHA
 */
const SecretKey = "6LdrobkZAAAAAHmavqXT7wztFLOBgZXCMKdiy79Z"
const GoogleVerifyUrl = "https://www.google.com/recaptcha/api/siteverify"

type GoogleKey struct {
	HumanKey	string	`json:"humanKey"`
}
