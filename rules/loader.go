package rules

import (
	"github.com/prometheus/prometheus/model/rulefmt"
	"github.com/prometheus/prometheus/promql/parser"
)

// GroupLoader is responsible for loading rule groups from arbitrary sources and parsing them.
type GroupLoader interface {
	Load(identifier string) (*rulefmt.RuleGroups, []error)
	Parse(query string) (parser.Expr, error)
}

// FileLoader is the default GroupLoader implementation. It defers to rulefmt.ParseFile
// and parser.ParseExpr
type FileLoader struct{}

func (FileLoader) Load(identifier string) (*rulefmt.RuleGroups, []error) {
	return rulefmt.ParseFile(identifier)
}

func (FileLoader) Parse(query string) (parser.Expr, error) { return parser.ParseExpr(query) }
