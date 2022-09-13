package model

import "github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"

func (h *Hotel) ParseModel(record hotelrepo.Hotels) error {
	code := int(record.Code)

	h = &Hotel{
		Code:        &code,
		CountryCode: &record.CountryCode,
	}
	return nil
}
