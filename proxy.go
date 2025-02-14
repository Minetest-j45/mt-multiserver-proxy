/*
Package proxy is a minetest reverse proxy for multiple servers.
It also provides an API for plugins.
*/
package proxy

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

const (
	latestSerializeVer = 28
	latestProtoVer     = 39
	versionString      = "5.5.0-dev-83a7b48bb"
	maxPlayerNameLen   = 20
	bytesPerMediaBunch = 5000
)

var playerNameChars = regexp.MustCompile("^[a-zA-Z0-9-_]+$")

var proxyDir string
var proxyDirOnce sync.Once

// Path prepends the directory the executable is in to the given path.
// It does not follow symlinks.
func Path(path ...string) string {
	proxyDirOnce.Do(func() {
		executable, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}

		proxyDir = filepath.Dir(executable)
	})

	return proxyDir + "/" + strings.Join(path, "")
}
