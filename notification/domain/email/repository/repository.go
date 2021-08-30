package repository

// EmailRepository provide function that related email repository
type EmailRepository struct {
	partnerRepository PartnerRepository
}

// PartnerRepository method that needed for email repository
type PartnerRepository interface {
	SendEmail(subject, body, to string) error
	SendBulkEmail(emails interface{}) error
}

// NewEmailRepository create new email repository
func NewEmailRepository(partnerRepository PartnerRepository) *EmailRepository {
	return &EmailRepository{partnerRepository}
}
