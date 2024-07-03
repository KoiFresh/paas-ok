package profile

type Profile struct {
	Filepath string
}

func NewProfileFromFile(filepath string) *Profile {
	return &Profile{
		Filepath: filepath,
	}
}
