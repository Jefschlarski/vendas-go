package services

import (
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"api/src/interface/api/dtos"
	"net/http"
)

type getAllAddresses struct {
	addressRepository interfaces.GetAllAddresses
}

func NewGetAllAddresses(repo interfaces.GetAllAddresses) *getAllAddresses {
	return &getAllAddresses{addressRepository: repo}
}

func (s *getAllAddresses) GetAll() (addressesDto []dtos.AddressDto, err *errors.Error) {

	addressesDto, error := s.addressRepository.GetAll()
	if error != nil {
		err = errors.NewError(error.Error(), http.StatusInternalServerError)
		return
	}

	return
}
