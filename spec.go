package swaggen

import "encoding/json"

type Spec interface {
	// Scan allows to read from multiple sources
	Scan(sources ...interface{}) error

	// Bytes returns spec's data as bytes
	Bytes() []byte

	// Bytes returns spec's data as string
	String() string
}

func New(version string) Spec {
	var spec Spec
	switch version {
	case "2", "2.0", "2.0.0":
	case "3", "3.0", "3.0.0":
		spec = new(SpecV3)
	}

	if spec == nil {
		panic(ErrInvalidVersion)
	}

	return spec
}

type FactorySpec struct{}

func (s *FactorySpec) Bytes() []byte {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	return b
}

func (s *FactorySpec) String() string {
	return string(s.Bytes())
}
