package repository

type PartnerRepository struct{}

// NewPartnerRepository create new partner repository
func NewPartnerRepository() *PartnerRepository {
	return &PartnerRepository{}
}
