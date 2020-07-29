package cli

import (
	"github.com/thmhoag/urlbyter/pkg/cli/version"
	"github.com/thmhoag/urlbyter/pkg/urlbyter"
)

type globalCtx struct {
	version     version.Properties
	batchRunner *urlbyter.BatchRunner
}

func (c *globalCtx) CurrentVersion() *version.Properties {
	return &c.version
}

func (c *globalCtx) BatchRunner() *urlbyter.BatchRunner {
	return c.batchRunner
}