package commands

type CandidateCommands struct {
	CreateProfile CreateProfileCommandHandler
	UpdateProfile UpdateProfileCommandHandler
}

func NewCandidateCommands(
	createProfile CreateProfileCommandHandler,
	updateProfile UpdateProfileCommandHandler,
) *CandidateCommands {
	return &CandidateCommands{
		CreateProfile: createProfile,
		UpdateProfile: updateProfile,
	}
}
