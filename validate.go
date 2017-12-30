package ghwatcher

import "errors"
import "io/ioutil"
import "net/http"
import "strings"

var errUnSupportedMethod = errors.New("Unsupported Method")

var errMissingHeader = errors.New("Missing Header")

var errInvalidUserAgent = errors.New("Error User-Agent")

var errInvalidSignature = errors.New("Invalid Signature")

var allowedMethods = []string{"GET", "POST"}

var requiredHeaders = [3]string{HeaderEvent, HeaderGUID, HeaderSignature}

//ValidateGithubRequest haha
func ValidateGithubRequest(r *http.Request) (err error) {
	// check method
	if !StringInList(r.Method, allowedMethods) {
		err = errUnSupportedMethod
		return
	}

	// check Headers
	for _, header := range requiredHeaders {
		if r.Header.Get(header) == "" {
			return errMissingHeader
		}
	}

	userAgent := r.Header.Get(HeaderUserAgent)
	if !strings.HasPrefix(userAgent, "GitHub-Hookshot/") {
		return errInvalidUserAgent
	}

	// check signature
	signature := r.Header.Get(HeaderSignature)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	if !CheckHMAC([]byte(config.GetSecret()), []byte(signature), body) {
		return errInvalidSignature
	}
	return nil
}
