package ghwatcher

import "crypto/hmac"
import "crypto/sha1"
import "os"
import log "github.com/Sirupsen/logrus"

// Config ghwatcher config
type Config struct {
	Secret   string
	Debug    bool
	LogPath  string
	RepoPath string
}

// NewConfig returns a default Config instance
func NewConfig() *Config {
	c := &Config{Secret: "", RepoPath: "./"}
	return c
}

// GetSecret returns the config secret
func (c *Config) GetSecret() string {
	return c.Secret
}

func (c *Config) Update(secret, logpath, repopath string, debug bool) {
	if secret != "" {
		c.Secret = secret
	}
	if logpath != "" {
		c.LogPath = logpath
		logFile, err := os.Open(logpath)
		if err == nil {
			log.SetOutput(logFile)
		}
	}
	if repopath != "" {
		c.RepoPath = repopath
	}

	c.Debug = debug
}

var config = NewConfig()

// StringInList check if string in given list
func StringInList(item string, list []string) (result bool) {
	for _, s := range list {
		if item == s {
			result = true
		}
	}
	return
}

// CheckHMAC check if the hmac of secret and message matched the given signature
func CheckHMAC(secret, signature, message []byte) bool {
	mac := hmac.New(sha1.New, secret)
	mac.Write(message)
	calculated := mac.Sum(nil)

	return hmac.Equal(signature, calculated)
}
