package swaggen

import (
	. "github.com/goline/swaggen/spec/v3"
)

type SpecV3 struct {
	FactorySpec `json:"-"`

	Openapi string   `json:"openapi"`
	Info    Info     `json:"info:omitempty"`
	Servers []Server `json:"servers:omitempty"`
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
	}

	return ErrInvalidSource
}

func (s *SpecV3) scanApp(r AppSource) error {
	return nil
}
