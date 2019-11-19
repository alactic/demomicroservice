package staff

type Staff struct {
	Id        string `json:"id, omitempty"`
	Type      string `json:"type"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
