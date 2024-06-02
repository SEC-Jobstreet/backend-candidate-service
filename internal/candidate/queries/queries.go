package queries

import "github.com/google/uuid"

type CandidateQueries struct {
	GetProfileByID            GetProfileByIDQueryHandler
	GetSavedJobList           GetSavedJobListQueryHandler
	GetAppliedCandidateList   GetAppliedCandidateListQueryHandler
	GetAppliedCandidateNumber GetAppliedCandidateNumberQueryHandler
}

func NewCandidateQueries(
	getProfileByID GetProfileByIDQueryHandler,
	getSavedJobList GetSavedJobListQueryHandler,
	getAppliedCandidateList GetAppliedCandidateListQueryHandler,
	getAppliedCandidateNumber GetAppliedCandidateNumberQueryHandler,
) *CandidateQueries {
	return &CandidateQueries{
		GetProfileByID:            getProfileByID,
		GetSavedJobList:           getSavedJobList,
		GetAppliedCandidateList:   getAppliedCandidateList,
		GetAppliedCandidateNumber: getAppliedCandidateNumber,
	}
}

type GetProfileByIDQuery struct {
	ID string
}

func NewGetProfileByIDQuery(ID string) *GetProfileByIDQuery {
	return &GetProfileByIDQuery{ID: ID}
}

type GetSavedJobListQuery struct {
	Username string
}

func NewGetSavedJobListQuery(username string) *GetSavedJobListQuery {
	return &GetSavedJobListQuery{Username: username}
}

type GetAppliedCandidateNumberQuery struct {
	JobID uuid.UUID
}

func NewGetAppliedCandidateNumberQuery(jobId uuid.UUID) *GetAppliedCandidateNumberQuery {
	return &GetAppliedCandidateNumberQuery{JobID: jobId}
}

type GetAppliedCandidateListQuery struct {
	JobID    uuid.UUID
	PageID   int
	PageSize int
}

func NewGetAppliedCandidateListQuery(jobId uuid.UUID, pageID int, pageSize int) *GetAppliedCandidateListQuery {
	return &GetAppliedCandidateListQuery{
		JobID:    jobId,
		PageID:   pageID,
		PageSize: pageSize,
	}
}
