package models

import (
	"time"

	"github.com/google/uuid" // Пакет для работы с UUID
)

type Tender struct {
	ID              uuid.UUID `db:"id" json:"id"`
	Name            string    `db:"name" json:"name"`
	Description     string    `db:"description" json:"description"`
	Status          string    `db:"status" json:"status"`
	ServiceType     string    `db:"service_type" json:"serviceType"`
	CreatorUsername string    `db:"creator_username" json:"creatorUsername"`
	OrganizationID  uuid.UUID `db:"organization_id" json:"organizationId"` // Изменили на UUID
	CreatedAt       time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time `db:"updated_at" json:"updatedAt"`
}

type Bid struct {
	ID              uuid.UUID `db:"id" json:"id"`
	Name            string    `db:"name" json:"name"`
	Description     string    `db:"description" json:"description"`
	Status          string    `db:"status" json:"status"`
	TenderID        uuid.UUID `db:"tender_id" json:"tenderId"`             // Изменили на UUID
	OrganizationID  uuid.UUID `db:"organization_id" json:"organizationId"` // Изменили на UUID
	CreatorUsername string    `db:"creator_username" json:"creatorUsername"`
	CreatedAt       time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time `db:"updated_at" json:"updatedAt"`
}
