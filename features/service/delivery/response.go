package delivery

import (
	"capstone-alta1/features/service"
	"time"
)

type ServiceResponse struct {
	ID               uint    `json:"id"`
	ServiceName      string  `json:"service_name"`
	ServiceCategory  string  `json:"service_category"`
	ServicePrice     uint    `json:"service_price"`
	AverageRating    float64 `json:"average_rating"`
	ServiceImageFile string  `json:"service_image_file"`
	City             string  `json:"city"`
	PartnerID        uint    `json:"partner_id"`
}

type ServiceAdditionalResponse struct {
	ID              uint   `json:"id"`
	AdditionalName  string `json:"additional_name"`
	AdditionalPrice uint   `json:"additional_price"`
	PartnerID       uint   `json:"partner_id"`
	ServiceID       uint   `json:"service_id"`
}

type ServiceReviewResponse struct {
	ID        uint    `json:"id"`
	Review    string  `json:"review"`
	Rating    float64 `json:"rating"`
	OrderID   uint    `json:"order_id"`
	ClientID  uint    `json:"client_id"`
	ServiceID uint    `json:"service_id"`
}

type ServiceDiscussionResponse struct {
	ID        uint      `json:"id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	PartnerID uint      `json:"partner_id"`
	ClientID  uint      `json:"client_id"`
	ServiceID uint      `json:"service_id"`
}

type ServiceAvailabilityResponse struct {
	ServiceName        string
	StartDate          time.Time
	EndDate            time.Time
	AvailabilityStatus string
}

func fromCore(dataCore service.Core) ServiceResponse {
	return ServiceResponse{
		ID:               dataCore.ID,
		ServiceName:      dataCore.ServiceName,
		ServiceCategory:  dataCore.ServiceCategory,
		ServicePrice:     dataCore.ServicePrice,
		AverageRating:    dataCore.AverageRating,
		ServiceImageFile: dataCore.ServiceImageFile,
		City:             dataCore.City,
		PartnerID:        dataCore.PartnerID,
	}
}

func fromCoreAdditional(dataCore service.Additional) ServiceAdditionalResponse {
	return ServiceAdditionalResponse{
		ID:              dataCore.ID,
		AdditionalName:  dataCore.AdditionalName,
		AdditionalPrice: dataCore.AdditionalPrice,
		PartnerID:       dataCore.PartnerID,
		ServiceID:       dataCore.ServiceID,
	}
}

func fromCoreReview(dataCore service.Review) ServiceReviewResponse {
	return ServiceReviewResponse{
		ID:        dataCore.ID,
		Review:    dataCore.Review,
		Rating:    dataCore.Rating,
		OrderID:   dataCore.OrderID,
		ClientID:  dataCore.ClientID,
		ServiceID: dataCore.ServiceID,
	}
}

func fromCoreDiscussion(dataCore service.Discussion) ServiceDiscussionResponse {
	return ServiceDiscussionResponse{
		ID:        dataCore.ID,
		Comment:   dataCore.Comment,
		CreatedAt: dataCore.CreatedAt,
		ClientID:  dataCore.ClientID,
		PartnerID: dataCore.PartnerID,
		ServiceID: dataCore.ServiceID,
	}
}

func fromCoreAvailability(dataCore service.Order) ServiceAvailabilityResponse {
	return ServiceAvailabilityResponse{
		ServiceName:        dataCore.ServiceName,
		StartDate:          dataCore.StartDate,
		EndDate:            dataCore.EndDate,
		AvailabilityStatus: dataCore.AvailabilityStatus,
	}
}

func fromCoreList(dataCore []service.Core) []ServiceResponse {
	var dataResponse []ServiceResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

func fromCoreListAdditional(dataCore []service.Additional) []ServiceAdditionalResponse {
	var dataResponse []ServiceAdditionalResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreAdditional(v))
	}
	return dataResponse
}

func fromCoreListReview(dataCore []service.Review) []ServiceReviewResponse {
	var dataResponse []ServiceReviewResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreReview(v))
	}
	return dataResponse
}

func fromCoreListDiscussion(dataCore []service.Discussion) []ServiceDiscussionResponse {
	var dataResponse []ServiceDiscussionResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreDiscussion(v))
	}
	return dataResponse
}
