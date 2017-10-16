package v3

type Server struct {
	Url         string              `json:"url,omitempty"`
	Description string              `json:"description,omitempty"`
	Variables   map[string]Variable `json:"variables,omitempty"`
}

type Variable struct {
	Default     string   `json:"default,omitempty"`
	Description string   `json:"description,omitempty"`
	Enum        []string `json:"enum,omitempty"`
}
