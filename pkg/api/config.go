package api

import (
	"github.com/joho/godotenv"
)

type Parser interface {
	Parse(path string) error
}

type EnvParser struct {
	path string `default:""`
}

func NewEnvParser(path ...string) *EnvParser {
	if len(path) > 0 {
		return &EnvParser{path: path[0]}
	}
	return &EnvParser{}
}

func (p *EnvParser) Parse() error {
	err := godotenv.Load(p.path)
	if err != nil {
		return err
	}
	return nil
}
