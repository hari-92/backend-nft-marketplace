package user_converters

import (
	userFilters "gitlab.com/hari-92/nft-market-server/internal/modules/user/filters"
	userRequests "gitlab.com/hari-92/nft-market-server/internal/modules/user/requests"
)

type RequestToFilterConverter struct{}

func NewRequestToFilterConverter() *RequestToFilterConverter {
	return &RequestToFilterConverter{}
}

func (c *RequestToFilterConverter) ToUserFilter(req *userRequests.GetUserRequest) *userFilters.UserFilter {
	if req == nil {
		return nil
	}

	return &userFilters.UserFilter{
		ID:    uint64(req.ID), // Convert uint32 to uint64
		Email: req.Email,
	}
}
