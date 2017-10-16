package v3

type Info struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Contact     Contact `json:"contact,omitempty"`
	License     License `json:"license,omitempty"`
	Version     string  `json:"version,omitempty"`
}

type Contact struct {
	Name  string `json:"name,omitempty"`
	Url   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type License struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}
