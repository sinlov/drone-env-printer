package env_printer_plugin

import (
	"fmt"
	"github.com/sinlov/drone-info-tools/drone_info"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// Plugin env_printer_plugin all config
	Plugin struct {
		Name    string
		Version string
		Drone   drone_info.Drone
		Config  Config
	}
)

func (p *Plugin) CleanResultEnv() error {
	for _, envItem := range cleanResultEnvList {
		err := os.Unsetenv(envItem)
		if err != nil {
			return fmt.Errorf("at FileBrowserPlugin.CleanResultEnv [ %s ], err: %v", envItem, err)
		}
	}
	return nil
}

func (p *Plugin) Exec() error {

	var err error

	if p.Config.PaddingLeftMax < 24 {
		p.Config.PaddingLeftMax = 24
	}

	var sb strings.Builder
	_, _ = fmt.Fprint(&sb, "-> just print basic env:\n")
	paddingMax := strconv.Itoa(p.Config.PaddingLeftMax)

	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneBuildTrigger, p.Drone.Build.Trigger)
	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneStageName, p.Drone.Stage.Name)
	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneStageMachine, p.Drone.Stage.Machine)
	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneStageOs, p.Drone.Stage.Os)
	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneStageArch, p.Drone.Stage.Arch)
	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneStageType, p.Drone.Stage.Type)
	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneStageKind, p.Drone.Stage.Kind)
	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneRepoName, p.Drone.Repo.ShortName)
	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneRepoOwner, p.Drone.Repo.OwnerName)
	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneRepo, p.Drone.Repo.FullName)
	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneRepoBranch, p.Drone.Build.RepoBranch)
	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneCommitRef, p.Drone.Commit.Ref)

	_, _ = fmt.Fprintf(&sb, "\n")
	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneBuildEvent, p.Drone.Build.Event)
	switch p.Drone.Build.Event {
	default:
		_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneCommitBranch, p.Drone.Commit.Branch)
	case "tag":
		_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneTag, p.Drone.Build.Tag)
	case "pull_request":
		_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneSourceBranch, p.Drone.Build.SourceBranch)
		_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneTargetBranch, p.Drone.Build.TargetBranch)
	case "push":
		_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneCommitBranch, p.Drone.Commit.Branch)
	}
	_, _ = fmt.Fprintf(&sb, "\n")

	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneBuildLink, p.Drone.Build.Link)
	_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", drone_info.EnvDroneCommitLink, p.Drone.Commit.Link)

	if len(p.Config.EnvPrintKeys) > 0 {
		_, _ = fmt.Fprint(&sb, "-> start print keys env:\n")
		for _, key := range p.Config.EnvPrintKeys {
			_, _ = fmt.Fprintf(&sb, "%-"+paddingMax+"s %s\n", key, os.Getenv(key))
		}
		_, _ = fmt.Fprint(&sb, "-> end print keys env\n")
	}
	log.Printf("%s", sb.String())

	//log.Printf("=> %s version %s end", p.Name, p.Version)

	if p.Config.Debug {
		log.Printf("=> debug: %s version %s", p.Name, p.Version)
	}
	return err
}

// randomStr
// new random string by cnt
//
//nolint:golint,unused
func randomStr(cnt uint) string {
	var letters = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	result := make([]byte, cnt)
	keyL := len(letters)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(keyL)]
	}
	return string(result)
}

// randomStr
// new random string by cnt
//
//nolint:golint,unused
func randomStrBySed(cnt uint, sed string) string {
	var letters = []byte(sed)
	result := make([]byte, cnt)
	keyL := len(letters)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(keyL)]
	}
	return string(result)
}

//nolint:golint,unused
func setEnvFromStr(key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		log.Fatalf("set env key [%v] string err: %v", key, err)
	}
}
