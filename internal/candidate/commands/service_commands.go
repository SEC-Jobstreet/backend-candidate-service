package commands

type CandidateCommands struct {
	CreateProfile CreateProfileCommandHandler
	UpdateProfile UpdateProfileCommandHandler
	ApllyJob      ApplyJobCommandHandler
	SaveJob       SaveJobCommandHandler
	UnsaveJob     UnsaveJobCommandHandler
}

func NewCandidateCommands(
	createProfile CreateProfileCommandHandler,
	updateProfile UpdateProfileCommandHandler,
	apllyJob ApplyJobCommandHandler,
	saveJob SaveJobCommandHandler,
	unsaveJob UnsaveJobCommandHandler,
) *CandidateCommands {
	return &CandidateCommands{
		CreateProfile: createProfile,
		UpdateProfile: updateProfile,
		ApllyJob:      apllyJob,
		SaveJob:       saveJob,
		UnsaveJob:     unsaveJob,
	}
}
