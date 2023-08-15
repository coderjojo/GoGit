package repository

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/ini.v1"
)

type GitRepository struct {
	WorkTree string
	GitDir   string
	Conf     *ini.File
}

func Init(path string, option ...bool) (*GitRepository, error) {
	gitRepo := &GitRepository{
		WorkTree: path,
		GitDir:   filepath.Join(path, ".git"),
	}

	force := false

	if len(option) > 0 {
		force = option[0]
	}

	if _, err := os.Stat(gitRepo.GitDir); os.IsNotExist(err) {
		if force {
			err := os.Mkdir(gitRepo.GitDir, os.ModePerm)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("Not a Git repository: %s", path)
		}
	}

	gitConfigPath := filepath.Join(gitRepo.GitDir, "config")

	if _, err := os.Stat(gitConfigPath); err == nil {
		gitRepo.Conf, err = ini.Load(gitConfigPath)
		if err != nil {
			return nil, err
		}
	} else if !force {
		return nil, fmt.Errorf("Configuration file missing")
	}

	if !force {
		vers, err := strconv.Atoi(gitRepo.Conf.Section("core").Key("repositoryformatversion").String())
		if err != nil {
			return nil, err
		}
		if vers != 0 {
			return nil, fmt.Errorf("Unsupported repositoryformatversion: %d", vers)
		}

	}
	return gitRepo, nil
}

func RepoPath(gitRepo *GitRepository, path ...string) string {
	/* Compute path under repo's gitdir */

	paths := append([]string{gitRepo.GitDir}, path...)
	return filepath.Join(paths...)
}

func RepoFile(gitRepo *GitRepository, mkdir bool, path ...string) (string, error) {
	/* Create dirname if absent */
	_, err := RepoDir(gitRepo, mkdir, path[:len(path)-1]...)

	if err != nil {
		return "", err
	}
	return RepoPath(gitRepo, path...), nil
}

func RepoDir(gitRepo *GitRepository, mkdir bool, path ...string) (string, error) {
	/* Similar to RepoFile, but mkdir path if absent if mkdir */
	dirPath := RepoPath(gitRepo, path...)
	if info, err := os.Stat(dirPath); err == nil {
		if !info.IsDir() {
			return "", fmt.Errorf("Not a directory: %s", dirPath)
		}
		return dirPath, nil
	}

	if mkdir {
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			return "", err
		}
		return dirPath, nil
	}
	return "", nil
}

/*
* .git/objects/
* .git/branches/
* .git/refs -> .git/refs/heads/, .git/refs/tags/
* .git/HEAD
* .git/config
* .git/description
 */

func RepoCreate(path string) (*GitRepository, error) {
	repo, err := Init(path, true)
	if err != nil {
		return nil, err
	}

	if info, err := os.Stat(repo.WorkTree); err == nil {
		if !info.IsDir() {
			return nil, fmt.Errorf("%s is not a directory", path)
		}
		numDirs, err := os.ReadDir(repo.GitDir)
		if err == nil && len(numDirs) > 0 {
			return nil, fmt.Errorf("%s is not empty", path)
		}
	} else {
		err := os.MkdirAll(repo.WorkTree, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	if _, err := RepoDir(repo, true, "objects"); err != nil {
		return nil, err
	}
	if _, err := RepoDir(repo, true, "branches"); err != nil {
		return nil, err
	}
	if _, err := RepoDir(repo, true, "refs"); err != nil {
		return nil, err
	}
	if _, err := RepoDir(repo, true, "refs", "heads"); err != nil {
		return nil, err
	}
	if _, err := RepoDir(repo, true, "refs", "tags"); err != nil {
		return nil, err
	}

	descPath, err := RepoFile(repo, true, "description")
	if err != nil {
		return nil, err
	}
	descContent := "Unnamed repository; edit this file 'description', to name the repository.\n"
	if err := os.WriteFile(descPath, []byte(descContent), os.ModePerm); err != nil {
		return nil, err
	}

	headPath, err := RepoFile(repo, true, "HEAD")
	if err != nil {
		return nil, err
	}
	headContent := "ref: refs/heads/main\n"
	if err := os.WriteFile(headPath, []byte(headContent), os.ModePerm); err != nil {
		return nil, err
	}

	configPath, err := RepoFile(repo, true, "config")
	if err != nil {
		return nil, err
	}
	configContent := repo_default_config()
	if err := os.WriteFile(configPath, []byte(configContent), os.ModePerm); err != nil {
		return nil, err
	}

	return repo, nil
}

func repo_default_config() string {

	cfg := ini.Empty()

	coreSection, _ := cfg.NewSection("core")
	coreSection.NewKey("repositoryformatversion", "0")
	coreSection.NewKey("filemode", "false")
	coreSection.NewKey("bare", "false")

	var buffer bytes.Buffer
	cfg.WriteTo(&buffer)

	return buffer.String()
}
