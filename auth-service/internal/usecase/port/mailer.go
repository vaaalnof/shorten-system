package port

type Mailer interface {
	Send(
		to string,
		subject string,
		body string,
	) error
}
