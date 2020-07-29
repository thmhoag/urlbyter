package root

import "github.com/thmhoag/urlbyter/pkg/urlbyter"

type Ctx interface {
	BatchRunner() *urlbyter.BatchRunner
}