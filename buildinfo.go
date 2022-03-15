package buildinfo

import "runtime/debug"

var Info map[string]string // mapped version of BuildInfo.Settings
var VCSRevision string     // Key "vcs.revision"
var VCSRev string          // Key "vcs.revision" truncated to 7 characters

func init() {
	Info = make(map[string]string)
	bi, _ := debug.ReadBuildInfo()
	for _, setting := range bi.Settings {
		Info[setting.Key] = setting.Value
	}
	VCSRevision = Info["vcs.revision"]
	VCSRev = Info["vcs.revision"][:7]
}
