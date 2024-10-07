package generator

import "cypher/pkg/schema"

type Generator interface {
	Generate(s *schema.Schema, outputFile string) error
}
