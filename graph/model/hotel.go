package model

import "github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"

func (h *Hotels) ParseModel(record hotelrepo.Hotels) error {
	code := int(record.Code)

	h = &Hotels{
		Code:        &code,
		CountryCode: &record.CountryCode,
	}
	return nil
}
