package ghwatcher

import "io/ioutil"
import "net/http"
import log "github.com/Sirupsen/logrus"
import "os"
import "os/exec"

var (
	HeaderEvent = "X-GitHub-Event"

	HeaderGUID = "X-GitHub-Delivery"

	HeaderSignature = "X-Hub-Signature"

	HeaderUserAgent = "User-Agent"
)

// HandleGithubPost accepts post request from github webhook
func HandleGithubPost(w http.ResponseWriter, r *http.Request) {
	err := ValidateGithubRequest(r)

	if err != nil {
		log.WithError(err).Errorln("Request validating failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event := r.Header.Get(HeaderEvent)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.WithError(err).Errorln("Read request body error")
		return
	}

	err = Dispatch(event, body)

	if err != nil {
		log.WithError(err).Errorln("Can not dispath event: ", event)
	}
}

func Dispatch(gitEvent string, message []byte) (err error) {
	switch gitEvent {
	case "push":
		// run git pull
		go GitPull()
		return
	}
	return
}

func GitPull() {
	os.Chdir(config.RepoPath)
	err := exec.Command("git", "pull").Run()
	if err != nil {
		log.WithError(err).Errorln("Git pull failed")
		return
	}
	log.Infoln("successfully pull from github")
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
