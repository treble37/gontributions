package gontrib

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/jubalh/gontributions/util"
	"github.com/jubalh/gontributions/vcs/git"
	"github.com/jubalh/gontributions/vcs/hg"
	"github.com/jubalh/gontributions/vcs/mediawiki"
	"github.com/jubalh/gontributions/vcs/obs"
)

// From contributions.go
var PullSources bool
var logwriter *bufio.Writer

// Project hold all important information
// about a project.
type Project struct {
	Name        string
	Description string
	URL         string
	Gitrepos    []string
	Hgrepos     []string
	MediaWikis  []mediawiki.MediaWiki
	Obs         []obs.OpenBuildService
}

// Configuration holds the users E-Mail adresses
// and Projects he contributed to.
type Configuration struct {
	Emails   []string
	Projects []Project
}

// Contribution holds the Projects the user
// contributed to and a the ammounts of contributions
type Contribution struct {
	Project Project
	Count   int
}

// scanGit is a helper function for ScanContributions which takes care of the git part
func scanGit(project Project, emails []string, contributions []Contribution) (int, error) {
	var sum int
	for _, repo := range project.Gitrepos {
		util.PrintInfo(nil, "Working on "+repo, util.PI_TASK)
		if PullSources {
			err := git.GetLatestRepo(repo)
			if err != nil {
				util.PrintInfoF(logwriter, "Problem loading repo %s: %s", util.PI_MILD_ERROR, repo, err.Error())
				return 0, err
			}
		}
		for _, email := range emails {
			path := filepath.Join("repos-git", util.LocalRepoName(repo))
			gitCount, err := git.CountCommits(path, email)
			if err != nil {
				return 0, err
			}

			if gitCount != 0 {
				util.PrintInfoF(nil, "%s: %d commits", util.PI_RESULT, email, gitCount)
				sum += gitCount
			}
		}
	}
	return sum, nil
}

// scanHg is a helper function for ScanContributions which takes care of the git part
func scanHg(project Project, emails []string, contributions []Contribution) (int, error) {
	var sum int
	for _, repo := range project.Hgrepos {
		util.PrintInfo(nil, "Working on "+repo, util.PI_TASK)
		if PullSources {
			err := hg.GetLatestRepo(repo)
			if err != nil {
				util.PrintInfo(logwriter, "Problem loading repo "+repo+": "+err.Error(), util.PI_MILD_ERROR)
				return 0, err
			}
		}
		for _, email := range emails {
			path := filepath.Join("repos-hg", util.LocalRepoName(repo))
			hgCount, err := hg.CountCommits(path, email)
			if err != nil {
				return 0, err
			}

			if hgCount != 0 {
				util.PrintInfoF(nil, "%s: %d commits", util.PI_RESULT, email, hgCount)
				sum += hgCount
			}
		}
	}
	return sum, nil
}

// scanWiki is a helper function for ScanContributions which takes care of the MediaWiki part
func scanWiki(project Project, emails []string, contributions []Contribution) int {
	var sum int
	for _, wiki := range project.MediaWikis {
		util.PrintInfoF(nil, "Working on MediaWiki %s as %s", util.PI_TASK, wiki.BaseUrl, wiki.User)

		wikiCount, err := mediawiki.GetUserEdits(wiki.BaseUrl, wiki.User)
		if err != nil {
			switch err.Error() {
			case "Not a valid URL":
				util.PrintInfo(logwriter, err.Error(), util.PI_MILD_ERROR)
				break
			case "Not able to HTTP Get",
				"Not able to decode JSON",
				"Did not get a 'user' returned":
				util.PrintInfo(logwriter, err.Error(), util.PI_ERROR)
				break
			}
		}

		if wikiCount == 0 {
			util.PrintInfoF(logwriter, "No edits for user %s", util.PI_MILD_ERROR, wiki.User)
		} else {
			util.PrintInfoF(nil, "%d edits", util.PI_RESULT, wikiCount)
			sum += wikiCount
		}
	}
	return sum
}

// scanOBS is a helper function for ScanContributions which takes care of the OBS part
func scanOBS(project Project, emails []string, contributions []Contribution) (int, error) {
	var sum int
Loop_obs:
	for _, obsEntry := range project.Obs {
		util.PrintInfo(nil, "Working on "+obsEntry.Repo, util.PI_TASK)

		if PullSources {
			err := obs.GetLatestRepo(obsEntry)
			if err != nil {
				util.PrintInfoF(logwriter, "Problem loading repo %s: %s", util.PI_MILD_ERROR, obsEntry.Repo, err.Error())
				return 0, err
			}
		}
		for _, email := range emails {
			obsCount, err := obs.CountCommits("repos-obs"+"/"+obsEntry.Repo, email)
			if err != nil {
				if err == obs.ErrNoChangesFileFound {
					util.PrintInfo(logwriter, "No .changes file found", util.PI_MILD_ERROR)
					break Loop_obs
				}
				util.PrintInfo(logwriter, err.Error(), util.PI_ERROR) // TODO: return?
			}

			if obsCount != 0 {
				util.PrintInfoF(nil, "%s: %d changes", util.PI_RESULT, email, obsCount)
				sum += obsCount
			}
		}
	}
	return sum, nil
}

func checkNeededBinaries() map[string]bool {
	m := make(map[string]bool)
	if util.BinaryInstalled("git") {
		m["git"] = true
		os.Mkdir("repos-git", 0755)
	}
	if util.BinaryInstalled("hg") {
		m["hg"] = true
		os.Mkdir("repos-hg", 0755)
	}
	if util.BinaryInstalled("osc") {
		m["osc"] = true
		os.Mkdir("repos-obs", 0755)
	}
	return m
}

func printBinaryInfos(binary map[string]bool) {
	if binary["git"] == false {
		util.PrintInfo(logwriter, "git is not installed. git repositories will be skipped", util.PI_MILD_ERROR)
	}
	if binary["hg"] == false {
		util.PrintInfo(logwriter, "hg is not installed. Mercurial repositories will be skipped", util.PI_MILD_ERROR)
	}
	if binary["osc"] == false {
		util.PrintInfo(logwriter, "osc is not installed. osc repositories will be skipped", util.PI_MILD_ERROR)
	}
}

// ScanContributions takes a Configuration containing a list of emails
// and a list of projects and returns a list of Contributions
// which contain of a project, how often a user contributed to it and
// description of the project.
func ScanContributions(configuration Configuration) ([]Contribution, error) {
	contributions := []Contribution{}

	logfile, err := os.Create("errors.log")
	if err != nil {
		return nil, err
	}
	defer logfile.Close()
	logwriter = bufio.NewWriter(logfile)

	binary := checkNeededBinaries()
	printBinaryInfos(binary)

	for _, project := range configuration.Projects {
		var sumCount int

		if binary["git"] {
			sum, err := scanGit(project, configuration.Emails, contributions)
			if err != nil {
				//logwriter.WriteString(err.Error())
				//return nil, err
			}
			sumCount += sum
		}

		if binary["hg"] {
			sum, err := scanHg(project, configuration.Emails, contributions)
			if err != nil {
				//logwriter.WriteString(err.Error())
				//return nil, err
			}
			sumCount += sum
		}

		sum := scanWiki(project, configuration.Emails, contributions)
		sumCount += sum

		if binary["osc"] {
			sum, err := scanOBS(project, configuration.Emails, contributions)
			if err != nil {
				//logwriter.WriteString(err.Error())
				//return nil, err
			}
			sumCount += sum
		}

		if sumCount > 0 {
			c := Contribution{project, sumCount}
			contributions = append(contributions, c)
		}
	}

	logwriter.Flush()
	return contributions, nil
}
