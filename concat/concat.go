package concat

type ConcatRequest struct {
	Name   string `json:"name"`
	Friend string `json:"friend"`
}

func ConcatString(name string, friend string) string {
	return name + friend
}
