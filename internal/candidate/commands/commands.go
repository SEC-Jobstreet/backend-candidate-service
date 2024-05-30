package commands

import (
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
)

type CreateProfileCommand struct {
	es.BaseCommand
	Profile models.Profile `json:"profile" bson:"profile" validate:"required"`
}

func NewCreateProfileCommand(aggregateID string, profile models.Profile) *CreateProfileCommand {
	return &CreateProfileCommand{BaseCommand: es.NewBaseCommand(aggregateID), Profile: profile}
}

type UpdateProfileCommand struct {
	es.BaseCommand
	Profile models.Profile `json:"profile" bson:"profile" validate:"required"`
}

func NewUpdateProfileCommand(aggregateID string, profile models.Profile) *UpdateProfileCommand {
	return &UpdateProfileCommand{BaseCommand: es.NewBaseCommand(aggregateID), Profile: profile}
}
