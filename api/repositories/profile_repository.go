package repositories

type ProfileRepository interface {
}

type profileRepository struct {
}

func NewProfileRepository() ProfileRepository {
	return &profileRepository{}
}
