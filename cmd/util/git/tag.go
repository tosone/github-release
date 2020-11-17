package git

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/unknwon/com"

	"github.com/tosone/release/common"
)

// ChangeLog get current tag between last tag chang log
func ChangeLog(dir string) (changeLog []byte, tag string, err error) {
	var tagList []string
	if tagList, err = tags(dir); err != nil {
		return
	}
	if len(tagList) == 0 {
		log.Fatal("Cannot find any tags. Please add a tag and try again.")
	} else if len(tagList) == 1 {
		tag = tagList[0]
		if changeLog, err = run(dir, fmt.Sprintf(
			"git log --pretty=format:'* %%s [%%h](%s/commit/%%H)'", common.HostRepoURL())); err != nil {
			return
		}
	} else {
		tag = tagList[len(tagList)-1]
		if changeLog, err = run(dir,
			fmt.Sprintf("git log --pretty=format:'* %%s [%%h](%s/commit/%%H)' ", common.HostRepoURL())+
				fmt.Sprintf("%s..%s", tagList[len(tagList)-2], tag)); err != nil {
			return
		}
	}
	return
}

func tags(dir string) (tags []string, err error) {
	var isGitRepo bool
	if isGitRepo, err = isRepo(dir); err != nil {
		return
	} else if !isGitRepo {
		err = fmt.Errorf("dir is not a working git directory: %s", dir)
		return
	}
	var stdout []byte
	if stdout, err = run(dir, "git tag"); err != nil {
		return
	}
	tags = strings.Split(strings.TrimSpace(string(stdout)), "\n")
	return
}

func run(dir, script string) (stdout []byte, err error) {
	cmd := exec.Command("sh", "-c", script)
	cmd.Dir = dir
	if stdout, err = cmd.Output(); err != nil {
		return
	}
	return
}

func isRepo(dir string) (isRepo bool, err error) {
	var stdout []byte

	if !com.IsDir(dir) {
		return
	}

	cmd := exec.Command("sh", "-c", "git rev-parse --is-inside-work-tree")
	cmd.Dir = dir

	if stdout, err = cmd.Output(); err != nil {
		return
	}

	gitWorkDir := strings.TrimSpace(string(stdout))
	isRepo = gitWorkDir == "true"
	return
}
