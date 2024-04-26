package repositories

type CandidateProfileRepository interface {
}

type candidateProfileRepository struct {
}

func NewCandidateProfileRepository() CandidateProfileRepository {
	return &candidateProfileRepository{}
}
