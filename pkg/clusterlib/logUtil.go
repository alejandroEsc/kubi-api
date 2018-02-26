package clusterlib

import (
	"github.com/juju/loggo"
)

var logger = GetModuleLogger("pkg.clusterlib", loggo.INFO)

// GetModuleLogger get a logger given a module name and level
func GetModuleLogger(module string, level loggo.Level) loggo.Logger {
	log := loggo.GetLogger(module)
	log.SetLogLevel(level)
	return log

}
