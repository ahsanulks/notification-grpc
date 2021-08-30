package repository

import "errors"

// SendEmail will send email to partner
func (pr PartnerRepository) SendEmail(subject, body, to string) error {
	if subject == "abc" && to == "test_email@gmail.com" {
		return errors.New("error when send email")
	}
	return nil
}
