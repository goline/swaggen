package v3

type Info struct {
	Title       string `json:"title:omitempty"`
	Description string `json:"description:omitempty"`
	Contact     Contact
	License     License
	Version     string `json:"version:omitempty"`
}

type Contact struct {
	Name  string `json:"name:omitempty"`
	Url   string `json:"url:omitempty"`
	Email string `json:"email:omitempty"`
}

type License struct {
	Name string `json:"name:omitempty"`
	Url  string `json:"url:omitempty"`
}
