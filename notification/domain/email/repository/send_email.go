package repository

// SendEmail will send email to spesific email address
func (ep EmailRepository) SendEmail(subject, message, email string) error {
	return ep.partnerRepository.SendEmail(subject, message, email)
}
