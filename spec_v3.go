package swaggen

import (
	. "github.com/goline/swaggen/spec/v3"

	"gopkg.in/yaml.v2"

	"encoding/json"
	"go/parser"
	"go/token"
	"regexp"
)

func newV3() *SpecV3 {
	return &SpecV3{
		Openapi: "3.0.0",
	}
}

type SpecV3 struct {
	FactorySpec `json:"-"`

	Openapi string   `json:"openapi"`
	Info    Info     `json:"info,omitempty"`
	Servers []Server `json:"servers,omitempty"`
}

func (s *SpecV3) Scan(sources ...interface{}) error {
	for _, source := range sources {
		if err := s.scanSource(source); err != nil {
			return err
		}
	}

	return nil
}

func (s *SpecV3) scanSource(source interface{}) error {
	switch r := source.(type) {
	case AppSource:
		return s.scanApp(r)
	case FileSource:
		return s.scanFile(r)
	}

	return ErrInvalidSource
}

func (s *SpecV3) scanApp(r AppSource) error {
	return nil
}

func (s *SpecV3) scanFile(r FileSource) error {
	f, err := parser.ParseFile(token.NewFileSet(), r.Path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	rg, _ := regexp.Compile(`^\{`)
	for _, comment := range f.Comments {
		t := comment.Text()
		if rg.MatchString(t) {
			// parse with json
			if err := json.Unmarshal([]byte(t), s); err != nil {
				return err
			}
		} else {
			// parse with yaml
			if err := yaml.Unmarshal([]byte(t), s); err != nil {
				return err
			}
		}
	}

	return nil
}
