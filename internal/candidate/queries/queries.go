package queries

type CandidateQueries struct {
	GetProfileByID GetProfileByIDQueryHandler
}

func NewCandidateQueries(getProfileByID GetProfileByIDQueryHandler) *CandidateQueries {
	return &CandidateQueries{GetProfileByID: getProfileByID}
}

type GetProfileByIDQuery struct {
	ID string
}

func NewGetProfileByIDQuery(ID string) *GetProfileByIDQuery {
	return &GetProfileByIDQuery{ID: ID}
}
