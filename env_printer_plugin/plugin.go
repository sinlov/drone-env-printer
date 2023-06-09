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

	//log.Printf("=> %s version %s start", p.Name, p.Version)

	if p.Config.Debug {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}

	var err error

	// set default TimeoutSecond
	if p.Config.TimeoutSecond == 0 {
		p.Config.TimeoutSecond = 10
	}
	if p.Config.PaddingLeftMax < 24 {
		p.Config.PaddingLeftMax = 24
	}

	var sb strings.Builder
	_, _ = fmt.Fprint(&sb, "-> just print basic env:\n")
	paddingMax := strconv.Itoa(p.Config.PaddingLeftMax)
	_, _ = fmt.Fprint(&sb, fmt.Sprintf("%-"+paddingMax+"s %s\n", drone_info.EnvDroneStageMachine, p.Drone.Stage.Machine))
	_, _ = fmt.Fprint(&sb, fmt.Sprintf("%-"+paddingMax+"s %s\n", drone_info.EnvDroneStageOs, p.Drone.Stage.Os))
	_, _ = fmt.Fprint(&sb, fmt.Sprintf("%-"+paddingMax+"s %s\n", drone_info.EnvDroneStageArch, p.Drone.Stage.Arch))
	_, _ = fmt.Fprint(&sb, fmt.Sprintf("%-"+paddingMax+"s %s\n", drone_info.EnvDroneRepoName, p.Drone.Repo.ShortName))
	_, _ = fmt.Fprint(&sb, fmt.Sprintf("%-"+paddingMax+"s %s\n", drone_info.EnvDroneRepoOwner, p.Drone.Repo.OwnerName))
	_, _ = fmt.Fprint(&sb, fmt.Sprintf("%-"+paddingMax+"s %s\n", drone_info.EnvDroneRepo, p.Drone.Repo.FullName))
	if p.Drone.Commit.Branch != "" {
		_, _ = fmt.Fprint(&sb, fmt.Sprintf("%-"+paddingMax+"s %s\n", drone_info.EnvDroneCommitBranch, p.Drone.Commit.Branch))
	}
	if p.Drone.Build.Tag != "" {
		_, _ = fmt.Fprint(&sb, fmt.Sprintf("%-"+paddingMax+"s %s\n", drone_info.EnvDroneTag, p.Drone.Build.Tag))
	}

	if len(p.Config.EnvPrintKeys) > 0 {
		_, _ = fmt.Fprint(&sb, "-> start print keys env:\n")
		for _, key := range p.Config.EnvPrintKeys {
			_, _ = fmt.Fprint(&sb, fmt.Sprintf("%-"+paddingMax+"s %s\n", key, os.Getenv(key)))
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

func setEnvFromStr(key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		log.Fatalf("set env key [%v] string err: %v", key, err)
	}
}
