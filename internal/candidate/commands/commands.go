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

type ApplyJobCommand struct {
	es.BaseCommand
	Application models.Application `json:"application" bson:"application" validate:"required"`
}

func NewApplyJobCommand(aggregateID string, application models.Application) *ApplyJobCommand {
	return &ApplyJobCommand{BaseCommand: es.NewBaseCommand(aggregateID), Application: application}
}

type SaveJobCommand struct {
	es.BaseCommand
	SavedJob models.SavedJob `json:"savedjob" bson:"savedjob" validate:"required"`
}

func NewSaveJobCommand(aggregateID string, savedjob models.SavedJob) *SaveJobCommand {
	return &SaveJobCommand{BaseCommand: es.NewBaseCommand(aggregateID), SavedJob: savedjob}
}

type UnSaveJobCommand struct {
	es.BaseCommand
	SavedJob models.SavedJob `json:"application" bson:"application" validate:"required"`
}

func NewUnsaveJobCommand(aggregateID string, unsavedjob models.SavedJob) *UnSaveJobCommand {
	return &UnSaveJobCommand{BaseCommand: es.NewBaseCommand(aggregateID), SavedJob: unsavedjob}
}
