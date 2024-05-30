package dto

type TourSearchSagaRequestDTO struct {
	ID            int64                `json:"Id"`
	Name          string               `json:"Name" gorm:"not null;type:string"`
	Description   string               `json:"Description" gorm:"not null;type:string"`
}