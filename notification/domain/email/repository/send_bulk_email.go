package repository

import (
	"notification/model"
)

// SendBulkEmail will send multiple email at same time
func (ep EmailRepository) SendBulkEmail(emails []model.Email) error {
	return ep.partnerRepository.SendBulkEmail(emails)
}
