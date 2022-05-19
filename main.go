package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ArnaLabs/C9Cli/CreateOrUpdateProtOrgAsg"
	"github.com/ArnaLabs/C9Cli/createorupdateserviceaccess"
	"github.com/ArnaLabs/C9Cli/delorauditserviceaccess"
	"github.com/ArnaLabs/C9Cli/createorupdateorgs"
	"github.com/ArnaLabs/C9Cli/createorupdateorgusers"
	"github.com/ArnaLabs/C9Cli/createorupdatequotas"
	"github.com/ArnaLabs/C9Cli/createorupdatespaces"
	"github.com/ArnaLabs/C9Cli/createorupdatespacesasgs"
	"github.com/ArnaLabs/C9Cli/createorupdatespaceusers"
	"github.com/ArnaLabs/C9Cli/delorauditorgs"
	"github.com/ArnaLabs/C9Cli/delorauditorgusers"
	"github.com/ArnaLabs/C9Cli/delorauditquotas"
	"github.com/ArnaLabs/C9Cli/delorauditspaceasgs"
	"github.com/ArnaLabs/C9Cli/delorauditspaces"
	"github.com/ArnaLabs/C9Cli/delorauditspaceusers"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
	//"time"
)

type SpaceListJson struct {
	Pagination struct {
		TotalResults int `json:"total_results"`
		TotalPages   int `json:"total_pages"`
		First        struct {
			Href string `json:"href"`
		} `json:"first"`
		Last struct {
			Href string `json:"href"`
		} `json:"last"`
		Next     interface{} `json:"next"`
		Previous interface{} `json:"previous"`
	} `json:"pagination"`
	Resources []struct {
		GUID          string    `json:"guid"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		Name          string    `json:"name"`
		Relationships struct {
			Organization struct {
				Data struct {
					GUID string `json:"guid"`
				} `json:"data"`
			} `json:"organization"`
			Quota struct {
				Data interface{} `json:"data"`
			} `json:"quota"`
		} `json:"relationships"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Organization struct {
				Href string `json:"href"`
			} `json:"organization"`
			Features struct {
				Href string `json:"href"`
			} `json:"features"`
			ApplyManifest struct {
				Href   string `json:"href"`
				Method string `json:"method"`
			} `json:"apply_manifest"`
		} `json:"links"`
		Metadata struct {
			Labels struct {
			} `json:"labels"`
			Annotations struct {
			} `json:"annotations"`
		} `json:"metadata"`
	} `json:"resources"`
}
type List struct {
	OrgList []struct {
		Name  string `yaml:"Name"`
		Quota string `yaml:"Quota"`
	} `yaml:"OrgList"`
	Audit string `yaml:"Audit"`
}
type GitList struct {
	OrgList []struct {
		Name   string `yaml:"Name"`
		Repo   string `yaml:"Repo"`
		Quota  string `yaml:"Quota"`
		Branch string `yaml:"Branch"`
	} `yaml:"OrgList"`
	Audit string `yaml:"Audit"`
}
type Orglist struct {
	Org struct {
		Name string `yaml:"Name"`
		ProtectedUsers []string `yaml:"ProtectedUsers"`
		OrgUsers struct {
			LDAP struct {
				OrgManagers []string `yaml:"OrgManagers"`
				OrgAuditors []string `yaml:"OrgAuditors"`
			} `yaml:"LDAP"`
			SSO struct {
				OrgManagers []string `yaml:"OrgManagers"`
				OrgAuditors []string `yaml:"OrgAuditors"`
			} `yaml:"SSO"`
			UAA struct {
				OrgManagers []string `yaml:"OrgManagers"`
				OrgAuditors []string `yaml:"OrgAuditors"`
			} `yaml:"UAA"`
		} `yaml:"OrgUsers"`
		Spaces []struct {
			Name         string `yaml:"Name"`
			IsolationSeg string `yaml:"IsolationSeg"`
			ASG          string `yaml:"ASG"`
			SpaceUsers   struct {
				LDAP struct {
					SpaceManagers   []string `yaml:"SpaceManagers"`
					SpaceDevelopers []string `yaml:"SpaceDevelopers"`
					SpaceAuditors   []string `yaml:"SpaceAuditors"`
				} `yaml:"LDAP"`
				UAA struct {
					SpaceManagers   []string `yaml:"SpaceManagers"`
					SpaceDevelopers []string `yaml:"SpaceDevelopers"`
					SpaceAuditors   []string `yaml:"SpaceAuditors"`
				} `yaml:"UAA"`
				SSO struct {
					SpaceManagers   []string `yaml:"SpaceManagers"`
					SpaceDevelopers []string `yaml:"SpaceDevelopers"`
					SpaceAuditors   []string `yaml:"SpaceAuditors"`
				} `yaml:"SSO"`
			} `yaml:"SpaceUsers"`
		} `yaml:"Spaces"`
	} `yaml:"Org"`
	SpaceAudit     string `yaml:"SpaceAudit"`
	UserAudit      string `yaml:"UserAudit"`
	ASGAudit       string `yaml:"ASGAudit"`
	IsolationAudit string `yaml:"IsolationAudit"`
}
type ProtectedList struct {
	Org                         []string `yaml:"Org"`
	Quota                       []string `yaml:"quota"`
	DefaultRunningSecurityGroup string   `yaml:"DefaultRunningSecurityGroup"`
}
type InitClusterConfigVals struct {
	ClusterDetails struct {
		EndPoint             string `yaml:"EndPoint"`
		User                 string `yaml:"User"`
		Org                  string `yaml:"Org"`
		Space                string `yaml:"Space"`
		EnableASG            bool   `yaml:"EnableASG"`
		EnableGitSubTree     bool   `yaml:"EnableGitSubTree"`
		GitHost              string `yaml:"GitHost"`
		SetOrgAuditor        bool   `yaml:"SetOrgAuditor"`
		SetOrgManager        bool   `yaml:"SetOrgManager"`
		SetSpaceAuditor      bool   `yaml:"SetSpaceAuditor"`
		SetSpaceManager      bool   `yaml:"SetSpaceManager"`
		SetSpaceDeveloper    bool   `yaml:"SetSpaceDeveloper"`
		EnableSpaceAudit     bool   `yaml:"EnableSpaceAudit"`
		EnableUserAudit      bool   `yaml:"EnableUserAudit"`
		EnableASGAudit       bool   `yaml:"EnableASGAudit"`
		EnableIsolationAudit bool   `yaml:"EnableIsolationAudit"`
		SSOProvider          string `yaml:"SSOProvider"`
	} `yaml:"ClusterDetails"`
}
type OrgStateYaml struct {
	OrgState struct {
		OldName string `yaml:"OldName"`
		NewName string `yaml:"NewName"`
		OrgGuid string `yaml:"OrgGuid"`
	} `yaml:"OrgState"`
}
type SpaceStateYaml struct {
	SpaceState struct {
		Org          string `yaml:"Org"`
		OrgGuid      string `yaml:"OrgGuid"`
		OldSpaceName string `yaml:"OldSpaceName"`
		NewSpaceName string `yaml:"NewSpaceName"`
		SpaceGuid    string `yaml:"SpaceGuid"`
	} `yaml:"SpaceState"`
}

func main() {

	var endpoint, user, pwd, org, space, asg, subtree, githost, sshkey, gitrepo, gitbranch, operation, cpath, orgaudit, orgman, spaceaudit, spaceman, spacedev, ostype string
	var ospath io.Writer

	flag.StringVar(&endpoint, "e", "api.sys-domain", "Use with init operation, Provide PCF Endpoint")
	flag.StringVar(&user, "u", "user", "Use with init operation, Provide UserName")
	flag.StringVar(&pwd, "p", "pwd", "Use with all operation, Provide Password")
	flag.StringVar(&org, "o", "org", "Use with init operation, Provide Org")
	flag.StringVar(&space, "s", "space", "Use with init operation, Provide Space")
	flag.StringVar(&asg, "a", "true", "Use with init operation, Enable ASGs ?.")
	flag.StringVar(&subtree, "st", "false", "Use with init operation, Enable Git Subtree ?.")
	flag.StringVar(&orgaudit, "OrgAuditor", "false", "Use with init operation, Enable setting up OrgAuditors ?.")
	flag.StringVar(&orgman, "OrgManager", "false", "Use with init operation, Enable setting up OrgManagers ?.")
	flag.StringVar(&spaceaudit, "SpaceAuditor", "false", "Use with init operation, Enable setting up SpaceAuditors ?.")
	flag.StringVar(&spaceman, "SpaceManager", "false", "Use with init operation, Enable setting up SpaceManagers ?.")
	flag.StringVar(&spacedev, "SpaceDeveloper", "false", "Use with init operation, Enable setting up SpaceDevelopers ?.")
	flag.StringVar(&operation, "i", "", "Provide Operation to be performed: init, create-{org,space,org-user,space-user,quota, ")
	flag.StringVar(&sshkey, "sshkey", "ssh-key.rsa", "path to SSH Key if Submodule is enabled")
	flag.StringVar(&githost, "githost", "githost.com", "Git Host to be added to knownfile")
	flag.StringVar(&gitrepo, "gr", "", "Provide SSH Git Repo")
	flag.StringVar(&gitbranch, "gb", "master", "Provide Git Branch")
	//flag.StringVar(&statepath, "sp", "ssh-key", "Provide path to create/update state files, use with all operations")
	//flag.StringVar(&cpath, "k", ".", "Provide path to configs, i.e <cluster-name>, i.e, to config folder, use with all operations")

	flag.Parse()
	ClusterName := strings.ReplaceAll(endpoint, ".", "-")
	fmt.Printf("Operation: %v\n", operation)

	// Pull repo name
	// git rev-parse --show-toplevel
	var out bytes.Buffer
	var errDir *exec.Cmd
	if ostype == "windows" {
		//fmt.Println(gitrepo)
		if gitrepo == "" {
			fmt.Println("Please provide SSH Git Repo")
			panic("")
		}
		cmd := gitrepo + ".Split('/')[-1].Replace('.git', '')"
		errDir = exec.Command("powershell", "-command", cmd)
		//errDir.Stderr = &out
		errDir.Stdout = &out
		errDir.Run()
		cpath = strings.TrimSpace(out.String())
		cpath = strings.TrimSpace(cpath)
		fmt.Printf("Repo: %v\n", cpath)
	} else {
		//fmt.Println(gitrepo)
		if gitrepo == "" {
			fmt.Println("Please provide SSH Git Repo")
			panic("")
		}
		cmd := "basename " + gitrepo + " | sed 's/.git//g'"
		errDir = exec.Command("sh", "-c", cmd)
		//errDir.Stderr = &out
		errDir.Stdout = &out
		errDir.Run()
		if _, err := errDir.Output(); err != nil {
			fmt.Println(errDir.Stdout, err, cmd)
		}
		cpath = strings.TrimSpace(out.String())
		cpath = strings.TrimSpace(cpath)
		fmt.Printf("Repo: %v\n", cpath)

	}

	oscmd := exec.Command("cmd", "/C", "echo", "%systemdrive%%homepath%")
	if _, err := oscmd.Output(); err != nil {
		//fmt.Println("Checking OS")
		//fmt.Println("command: ", oscmd)
		fmt.Println(oscmd.Stdout, err)
		oscmd = exec.Command("sh", "-c", "echo", "$HOME")
		if _, err := oscmd.Output(); err != nil {
			fmt.Println("Checking OS failed - Can't find Underlying OS")
			fmt.Println("command: ", oscmd)
			fmt.Println(oscmd.Stdout, err)
			panic(err)
		} else {
			fmt.Println("command: ", oscmd)
			ospath = oscmd.Stdout
			fmt.Println("PATH: ", ospath)
			fmt.Println("Checking OS - Setting up for Mac/Linux/Ubuntu")
			ostype = "non-windows"
		}
	} else {
		//fmt.Println("command: ", oscmd)
		ospath = oscmd.Stdout
		//fmt.Println("PATH: ", ospath)
		//fmt.Println("Checking OS - Setting up for Windows")
		ostype = "windows"
	}

	if operation == "init" {

		fmt.Println("Initializing C9Cli")

		fmt.Printf("ClusterName: %v\n", ClusterName)
		fmt.Printf("EndPoint: %v\n", endpoint)
		fmt.Printf("User: %v\n", user)
		fmt.Printf("Org: %v\n", org)
		fmt.Printf("Space: %v\n", space)
		fmt.Printf("EnableASG: %v\n", asg)
		fmt.Printf("EnableGitSubTree: %v\n", &subtree)
		fmt.Printf("GitHost: %v\n", &githost)
		fmt.Printf("EnableOrgAuditor: %v\n", orgaudit)
		fmt.Printf("EnableOrgManager: %v\n", orgman)
		fmt.Printf("EnableSpaceAuditor: %v\n", spaceaudit)
		fmt.Printf("EnableSpaceManager: %v\n", spaceman)
		fmt.Printf("EnableSpaceDeveloper: %v\n", spacedev)
		fmt.Printf("Path: %v\n", cpath)
		Init(ClusterName, endpoint, user, org, space, asg, subtree, githost, cpath, orgaudit, orgman, spaceaudit, spaceman, spacedev)

	} else if operation == "org-init" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		OrgsInit(ClusterName, cpath, ostype, sshkey)
		GitPush(ClusterName, ostype, cpath, sshkey, gitbranch)

	} else if operation == "audit-quota" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		delorauditquotas.DeleteOrAuditQuotas(ClusterName, cpath)
		GitPush(ClusterName, ostype, cpath, sshkey, gitbranch)

	} else if operation == "audit-org" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		delorauditorgs.DeleteorAuditOrgs(ClusterName, cpath)
		GitPush(ClusterName, ostype, cpath, sshkey, gitbranch)

	} else if operation == "space-init" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		SpaceInit(ClusterName, cpath, ostype, sshkey)
		GitPush(ClusterName, ostype, cpath, sshkey, gitbranch)

	} else if operation == "audit-space" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		delorauditspaces.DeleteorAuditSpaces(ClusterName, cpath, ostype)
		GitPush(ClusterName, ostype, cpath, sshkey, gitbranch)

	} else if operation == "audit-org-user" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		delorauditorgusers.DeleteorAuditOrgUsers(ClusterName, cpath, ostype)
		GitPush(ClusterName, ostype, cpath, sshkey, gitbranch)

	} else if operation == "audit-space-user" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		delorauditspaceusers.DeleteOrAuditSpaceUsers(ClusterName, cpath, ostype)
		GitPush(ClusterName, ostype, cpath, sshkey, gitbranch)

	} else if operation == "audit-space-asg" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		deleteorauditspaceasgs.DeleteOrAuditSpacesASGs(ClusterName, cpath, ostype)
		GitPush(ClusterName, ostype, cpath, sshkey, gitbranch)

	} else if operation == "create-org" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		createorupdateorgs.CreateOrUpdateOrgs(ClusterName, cpath, ostype)
		GitPush(ClusterName, ostype, cpath, sshkey, gitbranch)

	} else if operation == "create-quota" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		createorupdatequotas.CreateOrUpdateQuotas(ClusterName, cpath, ostype)
		//GitPush(ClusterName, ostype, cpath)

	} else if operation == "create-org-user" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		createorupdateorgusers.CreateOrUpdateOrgUsers(ClusterName, cpath, ostype)
		//GitPush(ClusterName, ostype, cpath)

	} else if operation == "create-space" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		createorupdatespaces.CreateOrUpdateSpaces(ClusterName, cpath, ostype)
		GitPush(ClusterName, ostype, cpath, sshkey, gitbranch)

	} else if operation == "create-space-user" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		createorupdatespaceusers.CreateOrUpdateSpaceUsers(ClusterName, cpath, ostype)
		//GitPush(ClusterName, ostype, cpath)

	} else if operation == "create-protected-org-asg" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		createorupdateprotorgasg.CreateOrUpdateProtOrgAsg(ClusterName, cpath, ostype)
		//GitPush(ClusterName, ostype, cpath)

	} else if operation == "create-space-asg" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		createorupdatespacesasgs.CreateOrUpdateSpacesASGs(ClusterName, cpath, ostype)
		//GitPush(ClusterName, ostype, cpath)

	} else if operation == "create-all" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName, pwd, cpath, sshkey, ostype, gitrepo, gitbranch)
		createorupdateprotorgasg.CreateOrUpdateProtOrgAsg(ClusterName, cpath, ostype)
		createorupdatequotas.CreateOrUpdateQuotas(ClusterName, cpath, ostype)
		createorupdateorgs.CreateOrUpdateOrgs(ClusterName, cpath, ostype)
		createorupdateorgusers.CreateOrUpdateOrgUsers(ClusterName, cpath, ostype)
		createorupdatespaces.CreateOrUpdateSpaces(ClusterName, cpath, ostype)
		createorupdatespacesasgs.CreateOrUpdateSpacesASGs(ClusterName, cpath, ostype)
		createorupdatespaceusers.CreateOrUpdateSpaceUsers(ClusterName, cpath, ostype)
		GitPush(ClusterName, ostype, cpath, sshkey, gitbranch)

	} else {
		fmt.Println("Provide Valid input operation")
	}
}
func SetupConnection(clustername string, pwd string, cpath string, sshkey string, ostype string, gitrepo string, gitbranch string) error {

	//Git Clone
	fmt.Printf("Git Repo: %v\n", gitrepo)
	fmt.Printf("Git Branch: %v\n", gitbranch)
	var out bytes.Buffer
	var errDir *exec.Cmd

	//fmt.Println(InitClusterConfigVals.ClusterDetails.EndPoint)
	//fmt.Printf("Pwd: %v\n", pwd)

	if ostype == "windows" {

		cmd := "start-ssh-agent.cmd"
		sshagent := exec.Command("powershell", "-command", cmd)
		if _, err := sshagent.Output(); err != nil {
			fmt.Println("err", sshagent, sshagent.Stdout, sshagent.Stderr)
			log.Fatal(err)
		} else {
			fmt.Println("Setup SSH Agent: ", sshagent, sshagent.Stdout)
		}

	} else {

		cmd := "cat /dev/zero | ssh-keygen -q -N ''"
		sshinit := exec.Command("sh", "-c", cmd)
		if _, err := sshinit.Output(); err != nil {
			fmt.Println("err", sshinit, sshinit.Stdout, sshinit.Stderr)
			log.Fatal(err)
		} else {
			fmt.Println("SSH init: ", sshinit, sshinit.Stdout)
		}

		cmd = "echo 'StrictHostKeyChecking no' >> ~/.ssh/config"
		nofg := exec.Command("sh", "-c", cmd)
		if _, err := nofg.Output(); err != nil {
			fmt.Println("err", nofg, nofg.Stdout, nofg.Stderr)
			log.Fatal(err)
		} else {
			fmt.Println(nofg, nofg.Stdout)
		}
	}

	if ostype == "windows" {
		cmd := "git clone -q -b " + gitbranch + " " + gitrepo
		errDir = exec.Command("powershell", "-command", cmd)
		errDir.Stderr = &out

	} else {
		cmd := "ssh-agent bash -c 'ssh-add " + sshkey + "; git clone -q -b " + gitbranch + " " + gitrepo + "'"
		errDir = exec.Command("sh", "-c", cmd)
		errDir.Stderr = &out

	}
	if _, err := errDir.Output(); err != nil {
		fmt.Println("err", errDir, errDir.Stderr)
		out.Reset()
	} else {
		fmt.Println("Cloning GitRepo: ", errDir, out.String())
		out.Reset()
	}

	ConfigFile := strings.TrimSpace(cpath + "/" + clustername + "/config.yml")
	fileConfigYml, err := ioutil.ReadFile(ConfigFile)
	if err != nil {

		fmt.Println(err)
	}

	var InitClusterConfigVals InitClusterConfigVals
	err = yaml.Unmarshal([]byte(fileConfigYml), &InitClusterConfigVals)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Endpoint: %v\n", InitClusterConfigVals.ClusterDetails.EndPoint)
	fmt.Printf("User: %v\n", InitClusterConfigVals.ClusterDetails.User)
	fmt.Printf("Org: %v\n", InitClusterConfigVals.ClusterDetails.Org)
	fmt.Printf("Space: %v\n", InitClusterConfigVals.ClusterDetails.Space)
	fmt.Printf("EnableASG: %v\n", InitClusterConfigVals.ClusterDetails.EnableASG)
	fmt.Printf("GitHost: %v\n", InitClusterConfigVals.ClusterDetails.GitHost)
	fmt.Printf("SetOrgAuditor: %v\n", InitClusterConfigVals.ClusterDetails.SetOrgAuditor)
	fmt.Printf("SetOrgManager: %v\n", InitClusterConfigVals.ClusterDetails.SetOrgManager)
	fmt.Printf("SetSpaceAuditor: %v\n", InitClusterConfigVals.ClusterDetails.SetSpaceAuditor)
	fmt.Printf("SetSpaceManager: %v\n", InitClusterConfigVals.ClusterDetails.SetSpaceManager)
	fmt.Printf("SetSpaceDeveloper: %v\n", InitClusterConfigVals.ClusterDetails.SetSpaceDeveloper)

	if ostype == "windows" {

	} else {

		cmd := "ssh-keyscan -H " + InitClusterConfigVals.ClusterDetails.GitHost + " > ~/.ssh/known_hosts"
		sshfingerprint := exec.Command("sh", "-c", cmd)
		if _, err = sshfingerprint.Output(); err != nil {
			fmt.Println("err", sshfingerprint, sshfingerprint.Stdout, sshfingerprint.Stderr)
			log.Fatal(err)
		} else {
			fmt.Println("Adding host to knowhosts: ", sshfingerprint, sshfingerprint.Stdout)
		}
	}

	if ostype == "windows" {
		cmd := "git config --global user.email 'you@example.com'"
		gitemail := exec.Command("powershell", "-command", cmd)
		if _, err := gitemail.Output(); err != nil {
			fmt.Println("err", gitemail, gitemail.Stdout, gitemail.Stderr)
			log.Fatal(err)
		} else {
			fmt.Println("Setup Git Email: ", gitemail, gitemail.Stdout)
		}
		cmd = "git config --global user.name C9Cli"
		gituser := exec.Command("powershell", "-command", cmd)
		if _, err := gituser.Output(); err != nil {
			fmt.Println("err", gituser, gituser.Stdout, gituser.Stderr)
			log.Fatal(err)
		} else {
			fmt.Println("Setup Git User: ", gituser, gituser.Stdout)
		}
	} else {
		cmd := "git config --global user.email 'you@example.com'"
		gitemail := exec.Command("sh", "-c", cmd)
		if _, err := gitemail.Output(); err != nil {
			fmt.Println("err", gitemail, gitemail.Stdout, gitemail.Stderr)
			log.Fatal(err)
		} else {
			fmt.Println("Setup Git Email: ", gitemail, gitemail.Stdout)
		}
		cmd = "git config --global user.name C9Cli"
		gituser := exec.Command("sh", "-c", cmd)
		if _, err := gituser.Output(); err != nil {
			fmt.Println("err", gituser, gituser.Stdout, gituser.Stderr)
			log.Fatal(err)
		} else {
			fmt.Println("Setup Git User: ", gituser, gituser.Stdout)
		}
	}

	cmd := exec.Command("cf", "login", "-a", InitClusterConfigVals.ClusterDetails.EndPoint, "-u", InitClusterConfigVals.ClusterDetails.User, "-p", pwd, "-o", InitClusterConfigVals.ClusterDetails.Org, "-s", InitClusterConfigVals.ClusterDetails.Space, "--skip-ssl-validation")
	if _, err := cmd.Output(); err != nil {
		fmt.Println("Connection failed")
		//fmt.Println("command: ", cmd)
		fmt.Println(cmd.Stdout, err)
		panic(err)
	} else {
		//fmt.Println("Connection Passed")
		//fmt.Println("command: ", cmd)
		fmt.Println(cmd.Stdout)
	}

	//if 	InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
	//} else {
	// setup git eval
	//}
	return err
}
func GitPush(clustername string, ostype string, cpath string, sshkey string, gitbranch string) error {

	var InitClusterConfigVals InitClusterConfigVals
	ConfigFile := cpath + "/" + clustername + "/config.yml"

	fileConfigYml, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileConfigYml), &InitClusterConfigVals)
	if err != nil {
		panic(err)
	}

	if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {

	} else {
		var errDir *exec.Cmd
		var out bytes.Buffer
		if ostype == "windows" {
			cmd := "git -C " + cpath + " --git-dir=.git checkout " + gitbranch
			errDir = exec.Command("powershell", "-command", cmd)
			errDir.Stderr = &out
			errDir.Stdout = &out
			err = errDir.Run()
		} else {
			cmd := "ssh-agent bash -c 'ssh-add " + sshkey + "; git -C " + cpath + " --git-dir=.git checkout " + gitbranch + "'"
			errDir = exec.Command("sh", "-c", cmd)
			errDir.Stderr = &out
			errDir.Stdout = &out
			err = errDir.Run()
		}
		if _, err := errDir.Output(); err != nil {
			fmt.Println(errDir, errDir.Stderr)
			out.Reset()
		} else {
			fmt.Println("Checkout Master: ", errDir, out.String(), errDir.Stdout)
			out.Reset()
		}
		if ostype == "windows" {
			cmd := "git -C " + cpath + " --git-dir=.git pull"
			errDir = exec.Command("powershell", "-command", cmd)
			errDir.Stderr = &out
			errDir.Stdout = &out
			err = errDir.Run()
		} else {
			cmd := "ssh-agent bash -c 'ssh-add " + sshkey + "; git -C " + cpath + " --git-dir=.git pull'"
			errDir = exec.Command("sh", "-c", cmd)
			errDir.Stderr = &out
			errDir.Stdout = &out
			err = errDir.Run()
		}
		if _, err := errDir.Output(); err != nil {
			fmt.Println(errDir, errDir.Stderr)
			out.Reset()
		} else {
			fmt.Println("Pulling Repo: ", errDir, out.String(), errDir.Stdout)
			out.Reset()
		}
		if ostype == "windows" {
			cmd := "git -C " + cpath + " --git-dir=.git add ."
			errDir = exec.Command("powershell", "-command", cmd)
			errDir.Stderr = &out
			errDir.Stdout = &out
			err = errDir.Run()
		} else {
			cmd := "ssh-agent bash -c 'ssh-add " + sshkey + "; git -C " + cpath + " --git-dir=.git add .'"
			errDir = exec.Command("sh", "-c", cmd)
			errDir.Stderr = &out
			errDir.Stdout = &out
			err = errDir.Run()
		}
		if _, err := errDir.Output(); err != nil {
			fmt.Println(errDir, errDir.Stderr)
			out.Reset()
			//fmt.Println(errDir, out.String(),errDir.Stdout)
			//log.Fatal(err)
		} else {
			fmt.Println("Adding Changes: ", errDir, out.String(), errDir.Stdout)
			out.Reset()
		}
		if ostype == "windows" {
			cmd := "git -C " + cpath + " --git-dir=.git commit -m 'Adding-Cluster-Updates'"
			errDir = exec.Command("powershell", "-command", cmd)
			errDir.Stderr = &out
			errDir.Stdout = &out
			err = errDir.Run()
		} else {
			cmd := "ssh-agent bash -c 'ssh-add " + sshkey + "; git -C " + cpath + " --git-dir=.git commit -m Adding-Cluster-Updates'"
			errDir = exec.Command("sh", "-c", cmd)
			errDir.Stderr = &out
			errDir.Stdout = &out
			err = errDir.Run()
		}
		if _, err := errDir.Output(); err != nil {
			//fmt.Println("err",errDir, errDir.Stdout, errDir.Stderr)
			fmt.Println(errDir, errDir.Stderr)
			out.Reset()
			//log.Fatal(err)
		} else {
			fmt.Println("Adding Commits: ", errDir, out.String(), errDir.Stdout)
			out.Reset()
		}
		if ostype == "windows" {
			cmd := "git -C " + cpath + " --git-dir=.git push"
			errDir = exec.Command("powershell", "-command", cmd)
			errDir.Stderr = &out
			errDir.Stdout = &out
			err = errDir.Run()
		} else {
			cmd := "ssh-agent bash -c 'ssh-add " + sshkey + "; git -C " + cpath + " --git-dir=.git push'"
			errDir = exec.Command("sh", "-c", cmd)
			errDir.Stderr = &out
			errDir.Stdout = &out
			err = errDir.Run()
		}
		if _, err := errDir.Output(); err != nil {
			fmt.Println(errDir, errDir.Stderr)
			//fmt.Println("err",errDir, errDir.Stdout)
			//log.Fatal(err)
			out.Reset()
		} else {
			fmt.Println("Pushing Changes: ", errDir, out.String(), errDir.Stdout)
			out.Reset()
		}
	}
	return err
}
func OrgsInit(clustername string, cpath string, ostype string, sshkey string) error {

	var list List
	var gitlist GitList
	var ProtectedOrgs ProtectedList
	var InitClusterConfigVals InitClusterConfigVals
	var LenList int

	ConfigFile := cpath + "/" + clustername + "/config.yml"
	fileConfigYml, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileConfigYml), &InitClusterConfigVals)
	if err != nil {
		panic(err)
	}

	var ListYml string

	if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
		ListYml = cpath + "/" + clustername + "/OrgsList.yml"
		fileOrgYml, err := ioutil.ReadFile(ListYml)
		if err != nil {
			fmt.Println(err)
		}
		err = yaml.Unmarshal([]byte(fileOrgYml), &list)
		if err != nil {
			panic(err)
		}
		LenList = len(list.OrgList)
		//fmt.Println("abc")
	} else {
		ListYml = cpath + "/" + clustername + "/GitOrgsList.yml"
		fileOrgYml, err := ioutil.ReadFile(ListYml)
		if err != nil {
			fmt.Println(err)
		}
		err = yaml.Unmarshal([]byte(fileOrgYml), &gitlist)
		if err != nil {
			panic(err)
		}
		LenList = len(gitlist.OrgList)
		//fmt.Println("123")
	}

	spath := cpath + "/" + clustername + "-state/"

	ProtectedOrgsYml := cpath + "/" + clustername + "/ProtectedResources.yml"
	fileProtectedYml, err := ioutil.ReadFile(ProtectedOrgsYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileProtectedYml), &ProtectedOrgs)
	if err != nil {
		panic(err)
	}

	LenProtectedOrgs := len(ProtectedOrgs.Org)

	// Org Renaming/Creating steps

	for i := 0; i < LenList; i++ {
		var OrgName, RepoName, BranchName string
		if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
			OrgName = list.OrgList[i].Name
			fmt.Println("Org: ", OrgName)
		} else {
			OrgName = gitlist.OrgList[i].Name
			RepoName = gitlist.OrgList[i].Repo
			BranchName = gitlist.OrgList[i].Branch
			if BranchName == "" {
				BranchName = "main"
			}
			fmt.Println("Org: ", OrgName)
			fmt.Println("Repo: ", RepoName)
			fmt.Println("Branch: ", BranchName)

		}
		var checkfile *exec.Cmd
		var fullpath string

		fullpath = spath + OrgName + "_OrgState.yml"
		OrgsStateYml := fullpath

		if ostype == "windows" {
			checkfile = exec.Command("powershell", "-command", "Get-Content", fullpath)
		} else {
			checkfile = exec.Command("cat", fullpath)
		}

		if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {

		} else {
			var errDir *exec.Cmd
			var out bytes.Buffer
			RepoPath := RepoName
			if ostype == "windows" {
				cmd := "git -C " + cpath + " --git-dir=.git subtree add --prefix " + "\"" + clustername + "/" + OrgName + "\"" + " " + RepoPath + " " + BranchName + " --squash"
				errDir = exec.Command("powershell", "-command", cmd)
				errDir.Stderr = &out
				errDir.Stdout = &out
				err = errDir.Run()
			} else {
				cmd := "ssh-agent bash -c 'ssh-add " + sshkey + "; git -C " + cpath + " --git-dir=.git subtree add --prefix " + clustername + "/" + OrgName + " " + RepoPath + " " + BranchName + " --squash'"
				errDir = exec.Command("sh", "-c", cmd)
				errDir.Stderr = &out
				errDir.Stdout = &out
				err = errDir.Run()
			}
			if _, err := errDir.Output(); err != nil {
				//fmt.Println("err",errDir, errDir.Stdout)
				fmt.Println(errDir, errDir.Stderr)
				out.Reset()
				//log.Fatal(err)
			} else {
				fmt.Println("Adding Org Repo: ", errDir, out.String(), errDir.Stdout)
				out.Reset()
			}
			RepoPath = RepoName
			if ostype == "windows" {
				cmd := "git -C " + cpath + " --git-dir=.git subtree pull --prefix " + "\"" + clustername + "/" + OrgName + "\"" + " " + RepoPath + " " + BranchName + " --squash -m C9Cli-bot'"
				errDir = exec.Command("powershell", "-command", cmd)
				errDir.Stderr = &out
				errDir.Stdout = &out
				err = errDir.Run()

			} else {
				cmd := "ssh-agent bash -c 'ssh-add " + sshkey + "; git -C " + cpath + " --git-dir=.git subtree pull --prefix " + clustername + "/" + OrgName + " " + RepoPath + " " + BranchName + " --squash -m C9Cli-bot'"
				errDir = exec.Command("sh", "-c", cmd)
				errDir.Stderr = &out
				errDir.Stdout = &out
				err = errDir.Run()
			}
			if _, err := errDir.Output(); err != nil {

				fmt.Println(errDir, errDir.Stderr)
				out.Reset()
				//fmt.Println("err",errDir, errDir.Stdout, errDir.Stderr)
				//log.Fatal(err)
			} else {
				fmt.Println("Pulling Org Repo: ", errDir, out.String(), errDir.Stdout)
				out.Reset()
			}
		}

		if _, err := checkfile.Output(); err == nil {

			if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
				fmt.Println("Org: ", OrgName)

			} else {
				fmt.Println("Org: ", OrgName)
				fmt.Println("Repo: ", RepoName)
			}
			fmt.Println(checkfile.Stdout)
			fileOrgStateYml, err := ioutil.ReadFile(OrgsStateYml)
			if err != nil {
				fmt.Println(err)
			}
			var orgstatedetails OrgStateYaml
			err = yaml.Unmarshal([]byte(fileOrgStateYml), &orgstatedetails)
			if err != nil {
				panic(err)
			}

			OrgNewName := orgstatedetails.OrgState.NewName
			OrgOldName := orgstatedetails.OrgState.OldName

			if OrgNewName == OrgOldName {

				//Org Name is not changing
				//checking if Change in Space Name
				//var count, totalcount int
				//checking if name is listed in protected orgs
				//checking if name aleady exists in orglist.ymll

			} else {
				fmt.Println("Changing Org Name")
				fmt.Println(" ")

				var count, totalcount int

				//checking if name is listed in protected orgs
				for p := 0; p < LenProtectedOrgs; p++ {
					if ProtectedOrgs.Org[p] == OrgNewName {
						count = 1
					} else {
						count = 0
					}
					totalcount = totalcount + count
				}

				//checking if org name aleady exists in orglist.yml
				if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
					for i := 0; i < LenList; i++ {
						if list.OrgList[i].Name == OrgNewName {
							count = 1
						} else {
							count = 0
						}
						totalcount = totalcount + count
					}
				} else {
					for i := 0; i < LenList; i++ {
						if gitlist.OrgList[i].Name == OrgNewName {
							count = 1
						} else {
							count = 0
						}
						totalcount = totalcount + count
					}
				}

				if totalcount == 0 {

					//Changing org name in OrgList.yml or GitOrgList.yml
					//fmt.Println("Changing org name in OrgList.yml or GitOrgList.yml")
					//fmt.Println("Changing Org Name")

					fmt.Println("- OrgList", OrgOldName)
					fmt.Println("+ OrgList", OrgNewName)

					if ostype == "windows" {
						var olpath string
						if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
							olpath = cpath + "/" + clustername + "/OrgsList.yml"
							stng := "((Get-Content -path" + " " + olpath + " -Raw) -replace '" + OrgOldName + "', '" + OrgNewName + "') | Set-Content -path " + olpath
							value := "(Get-Content " + olpath + " -Encoding UTF8) | ForEach-Object {$_ -replace '\"',''}| Out-File " + olpath + " -Encoding UTF8"
							trimquotes := exec.Command("powershell", "-command", value)
							err := trimquotes.Run()
							if err != nil {
								fmt.Println("err :", err, trimquotes, trimquotes.Stdout, trimquotes.Stderr)
								panic(err)
							}
							changestr := exec.Command("powershell", "-command", stng)
							err = changestr.Run()
							if err != nil {
								fmt.Println("err :", err, changestr, changestr.Stdout, changestr.Stderr)
								panic(err)
							} else {
							}
						} else {
							olpath = cpath + "/" + clustername + "/GitOrgsList.yml"
							stng := "((Get-Content -path" + " " + olpath + " -Raw) -replace '  - Name: " + OrgOldName + "', '  - Name: " + OrgNewName + "') | Set-Content -path " + olpath
							value := "(Get-Content " + olpath + " -Encoding UTF8) | ForEach-Object {$_ -replace '\"',''}| Out-File " + olpath + " -Encoding UTF8"
							trimquotes := exec.Command("powershell", "-command", value)
							err := trimquotes.Run()
							if err != nil {
								fmt.Println("err :", err, trimquotes, trimquotes.Stdout, trimquotes.Stderr)
								panic(err)
							}
							changestr := exec.Command("powershell", "-command", stng)
							err = changestr.Run()
							if err != nil {
								fmt.Println("err :", err, changestr, changestr.Stdout, changestr.Stderr)
								panic(err)
							} else {
							}
						}

					} else {
						var olpath string
						if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
							olpath = cpath + "/" + clustername + "/OrgsList.yml"
							stng := "sed -i 's/" + strings.TrimSpace(OrgOldName) + "/" + strings.TrimSpace(OrgNewName) + "/g' " + olpath
							value := "sed -i 's/\\" + "\"//g' " + olpath
							trimquotes := exec.Command("sh", "-c", value)
							err := trimquotes.Run()
							if err != nil {
								fmt.Println("err :", err, trimquotes, trimquotes.Stdout, trimquotes.Stderr)
								panic(err)
							}

							changestr := exec.Command("sh", "-c", stng)
							err = changestr.Run()
							if err != nil {
								fmt.Println("err :", err, changestr, changestr.Stdout, changestr.Stderr)
								//panic(err)
							}
						} else {
							olpath = cpath + "/" + clustername + "/GitOrgsList.yml"
							stng := "sed -i 's/" + "  - Name: " + strings.TrimSpace(OrgOldName) + "/" + "  - Name: " + strings.TrimSpace(OrgNewName) + "/g' " + olpath
							value := "sed -i 's/\\" + "\"//g' " + olpath
							trimquotes := exec.Command("sh", "-c", value)
							err := trimquotes.Run()
							if err != nil {
								fmt.Println("err :", err, trimquotes, trimquotes.Stdout, trimquotes.Stderr)
								panic(err)
							}

							changestr := exec.Command("sh", "-c", stng)
							err = changestr.Run()
							if err != nil {
								fmt.Println("err :", err, changestr, changestr.Stdout, changestr.Stderr)
								//panic(err)
							}
						}

					}

					//changing org state filename
					//spath := cpath+"/"+clustername+"-state/"
					oldfullpath := spath + OrgOldName + "_OrgState.yml"
					newfullpath := spath + OrgNewName + "_OrgState.yml"
					fmt.Println("- ", oldfullpath)
					fmt.Println("+ ", newfullpath)
					if ostype == "windows" {
						changestfile := exec.Command("powershell", "-command", "mv", oldfullpath, newfullpath)
						err := changestfile.Run()
						if err != nil {
							panic(err)
						}
					} else {
						value := "mv" + " " + oldfullpath + " " + newfullpath
						changestfile := exec.Command("sh", "-c", value)
						err := changestfile.Run()
						if err != nil {
							panic(err)
						} else {
							fmt.Println(changestfile, changestfile.Stdout, changestfile.Stderr)
						}
					}

					//Changing Folder name
					oldmgmtpath := cpath + "/" + clustername + "/" + OrgOldName
					//newmgmtpath := cpath + "/" + clustername + "/" + OrgNewName
					newmgmtpath := cpath + "/" + clustername + "/" + OrgOldName + "_old"
					fmt.Println("- ", oldmgmtpath)
					fmt.Println("+ ", newmgmtpath)
					if ostype == "windows" {
						changefolderfile := exec.Command("powershell", "-command", "mv", oldmgmtpath, newmgmtpath)
						err = changefolderfile.Run()
						if err != nil {
							panic(err)
						}
					} else {
						value := "mv" + " " + oldmgmtpath + " " + newmgmtpath
						changefolderfile := exec.Command("sh", "-c", value)
						err = changefolderfile.Run()
						if err != nil {
							panic(err)
						} else {
							fmt.Println(changefolderfile, changefolderfile.Stdout, changefolderfile.Stderr)
						}
					}

					var errDir *exec.Cmd
					var out bytes.Buffer

					if ostype == "windows" {
						cmd := "git -C " + cpath + " --git-dir=.git subtree add --prefix " + "\"" + clustername + "/" + OrgNewName + "\"" + " " + RepoName + " " + BranchName + " --squash"
						errDir = exec.Command("powershell", "-command", cmd)
						errDir.Stderr = &out
						errDir.Stdout = &out
						err = errDir.Run()
					} else {
						cmd := "ssh-agent bash -c 'ssh-add " + sshkey + "; git -C " + cpath + " --git-dir=.git subtree add --prefix " + clustername + "/" + OrgNewName + " " + RepoName + " " + BranchName + " --squash'"
						errDir = exec.Command("sh", "-c", cmd)
						errDir.Stderr = &out
						errDir.Stdout = &out
						err = errDir.Run()
					}
					if _, err := errDir.Output(); err != nil {
						//fmt.Println("err",errDir, errDir.Stdout)
						fmt.Println(errDir, errDir.Stderr)
						out.Reset()
						//log.Fatal(err)
					} else {
						fmt.Println("Adding Org Repo: ", errDir, out.String(), errDir.Stdout)
						out.Reset()
					}
					if ostype == "windows" {
						cmd := "git -C " + cpath + " --git-dir=.git subtree pull --prefix " + "\"" + clustername + "/" + OrgNewName + "\"" + " " + RepoName + " " + BranchName + " --squash -m C9Cli-bot"
						errDir = exec.Command("powershell", "-command", cmd)
						errDir.Stderr = &out
						errDir.Stdout = &out
						err = errDir.Run()

					} else {
						cmd := "ssh-agent bash -c 'ssh-add " + sshkey + "; git -C " + cpath + " --git-dir=.git subtree pull --prefix " + clustername + "/" + OrgNewName + " " + RepoName + " " + BranchName + " --squash -m C9Cli-bot'"
						errDir = exec.Command("sh", "-c", cmd)
						errDir.Stderr = &out
						errDir.Stdout = &out
						err = errDir.Run()
					}
					if _, err := errDir.Output(); err != nil {

						fmt.Println(errDir, errDir.Stderr)
						out.Reset()
						//fmt.Println("err",errDir, errDir.Stdout, errDir.Stderr)
						//log.Fatal(err)
					} else {
						fmt.Println("Pulling Org Repo: ", errDir, out.String(), errDir.Stdout)
						out.Reset()
					}

					// Checking for Space Name Change

				} else {
					fmt.Println("Org Name exist in Org's list, can't be renamed")
					fmt.Println(" ")
				}
			}
		} else {

			//fmt.Println("Missing/New Request")
			// Creating State file for listed Orgs

			var count, totalcount int
			var OrgName string
			//RepoName string
			if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
				fmt.Println("Org: ", list.OrgList[i])
				OrgName = list.OrgList[i].Name
			} else {
				fmt.Println("Org: ", gitlist.OrgList[i].Name)
				fmt.Println("Repo: ", gitlist.OrgList[i].Repo)
				OrgName = gitlist.OrgList[i].Name
				RepoName = gitlist.OrgList[i].Repo
			}
			fmt.Println("Missing State file, Creating State files")
			var out bytes.Buffer

			//pullguid := exec.Command("cf", "org", list.OrgList[i], "--guid")
			pullguid := exec.Command("cf", "org", OrgName, "--guid")
			pullguid.Stdout = &out
			err = pullguid.Run()
			if err == nil {

				// Creating Org state file
				fmt.Println("Org exist, creating state file")
				OrgGuidPull := out.String()
				out.Reset()
				type OrgState struct {
					OldName string `yaml:"OldName"`
					NewName string `yaml:"NewName"`
					OrgGuid string `yaml:"OrgGuid"`
				}

				//spath := cpath+"/"+clustername+"-state/"
				values := OrgState{OldName: OrgName, NewName: OrgName, OrgGuid: OrgGuidPull}

				var templates *template.Template
				var allFiles []string

				if err != nil {
					fmt.Println(err)
				}

				filename := "OrgGuid.tmpl"
				fullPath := spath + "OrgGuid.tmpl"
				if strings.HasSuffix(filename, ".tmpl") {
					allFiles = append(allFiles, fullPath)
				}

				//fmt.Println(allFiles)
				templates, err = template.ParseFiles(allFiles...)
				if err != nil {
					fmt.Println(err)
				}

				s1 := templates.Lookup("OrgGuid.tmpl")
				f, err := os.Create(spath + OrgName + "_OrgState.yml")
				if err != nil {
					panic(err)
				}

				err = s1.Execute(f, values)
				defer f.Close() // don't forget to close the file when finished.
				if err != nil {
					panic(err)
				}

				// Checking Space State file exists

			} else {

				fmt.Println("Org does't exist, Creating config files")
				for p := 0; p < LenProtectedOrgs; p++ {
					//fmt.Println("Protected Org: ", ProtectedOrgs.Org[p])
					if ProtectedOrgs.Org[p] == OrgName {
						count = 1
					} else {
						count = 0
					}
					totalcount = totalcount + count
				}

				if totalcount == 0 {
					//fmt.Println("This is not Protected Org")

					mgmtpath := cpath + "/" + clustername + "/" + OrgName
					ASGPath := cpath + "/" + clustername + "/" + OrgName + "/ASGs/"
					OrgsYml := cpath + "/" + clustername + "/" + OrgName + "/Org.yml"
					JsonPath := cpath + "/" + clustername + "/" + OrgName + "/ASGs/" + "test_test.json"

					_, err = os.Stat(mgmtpath)
					if os.IsNotExist(err) {

						fmt.Println("Creating <cluster>/<Org> folder")

						if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {

							errDir := os.MkdirAll(mgmtpath, 0755)
							if errDir != nil {
								log.Fatal(err)
							}

							var OrgTmp = `---
Org:
  Name:
  Quota:
  OrgUsers:
    LDAP:
      OrgManagers:
        - User1
        - User2
        - User3
      OrgAuditors:
        - User1
        - User2
    SSO:
      OrgManagers:
        - User1
        - User2
        - User3
      OrgAuditors:
        - User1
        - User2
    UAA:
      OrgManagers:
        - User1
        - User2
        - User3
      OrgAuditors:
        - User1
        - User2
  Spaces:
    - Name: Space1
      IsolationSeg: test-segment-1
      ASG: org_space.json
      SpaceUsers:
        LDAP:
          SpaceManagers:
            - User1
            - User2
            - User3
          SpaceDevelopers:
            - User1
            - User2
            - User3
          SpaceAuditors:
            - User1
            - User2
            - User3
    - Name: Space2
      IsolationSeg: test-segment-2
      ASG: org_space.json
      SpaceUsers:
        LDAP:
          SpaceManagers:
            - User1
            - User2
            - User3
          SpaceDevelopers:
            - User1
            - User2
            - User3
          SpaceAuditors:
            - User1
            - User2
            - User3
SpaceAudit: list #delete/rename/list
UserAudit:  list #unset/list
ASGAudit:   list #delete/list
IsolationAudit: list #unbind/list`

							fmt.Println("Creating <cluster>/<Org> sample yaml files")
							err = ioutil.WriteFile(OrgsYml, []byte(OrgTmp), 0644)
							check(err)

							_, err = os.Stat(ASGPath)
							if os.IsNotExist(err) {
								errDir := os.MkdirAll(ASGPath, 0755)
								if errDir != nil {
									log.Fatal(err)
									fmt.Println("<cluster>/<Org>/ASGs exist, please manually edit file to make changes or provide new cluster name")
								} else {
									fmt.Println("Creating <cluster>/<Org>/ASGs")
									var AsgTmp = `---
[
  {
    "protocol": "tcp",
    "destination": "10.x.x.88",
    "ports": "1443",
	"log": true,
	"description": "Allow DNS lookup by default."
  }
]`

									fmt.Println("Creating <cluster>/<Org>/ASGs sample json file")
									err = ioutil.WriteFile(JsonPath, []byte(AsgTmp), 0644)
									check(err)
								}

							}
						} else {
							//fmt.Println("test3")
						}
					} else {
						fmt.Println("<cluster>/<Org> exists, please manually edit file to make changes or provide new cluster name")
					}
				}
			}
		}
	}
	return nil
}
func SpaceInit(clustername string, cpath string, ostype string, sshkey string) error {

	var list List
	var gitlist GitList
	var ProtectedOrgs ProtectedList
	var InitClusterConfigVals InitClusterConfigVals
	var LenList int

	ConfigFile := cpath + "/" + clustername + "/config.yml"
	fileConfigYml, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileConfigYml), &InitClusterConfigVals)
	if err != nil {
		panic(err)
	}

	var ListYml string

	if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
		ListYml = cpath + "/" + clustername + "/OrgsList.yml"
		fileOrgYml, err := ioutil.ReadFile(ListYml)
		if err != nil {
			fmt.Println(err)
		}
		err = yaml.Unmarshal([]byte(fileOrgYml), &list)
		if err != nil {
			panic(err)
		}
		LenList = len(list.OrgList)
		//fmt.Println("abc")
	} else {
		ListYml = cpath + "/" + clustername + "/GitOrgsList.yml"
		fileOrgYml, err := ioutil.ReadFile(ListYml)
		if err != nil {
			fmt.Println(err)
		}
		err = yaml.Unmarshal([]byte(fileOrgYml), &gitlist)
		if err != nil {
			panic(err)
		}
		LenList = len(gitlist.OrgList)
		//fmt.Println("123")
	}

	spath := cpath + "/" + clustername + "-state/"

	ProtectedOrgsYml := cpath + "/" + clustername + "/ProtectedResources.yml"
	fileProtectedYml, err := ioutil.ReadFile(ProtectedOrgsYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileProtectedYml), &ProtectedOrgs)
	if err != nil {
		panic(err)
	}

	// Org Renaming/Creating steps

	for i := 0; i < LenList; i++ {
		var OrgName, RepoName string
		if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
			OrgName = list.OrgList[i].Name
		} else {
			OrgName = gitlist.OrgList[i].Name
			RepoName = gitlist.OrgList[i].Repo
		}
		var checkfile *exec.Cmd
		var fullpath string

		fullpath = spath + OrgName + "_OrgState.yml"
		OrgsStateYml := fullpath

		if ostype == "windows" {
			checkfile = exec.Command("powershell", "-command", "Get-Content", fullpath)
		} else {
			checkfile = exec.Command("cat", fullpath)
		}

		if _, err := checkfile.Output(); err == nil {

			// Org exist
			// checking space
			if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {

				fmt.Println("Org: ", OrgName)

			} else {
				fmt.Println("Org: ", OrgName)
				fmt.Println("Repo: ", RepoName)
			}

			fmt.Println(checkfile.Stdout)
			fileOrgStateYml, err := ioutil.ReadFile(OrgsStateYml)
			if err != nil {
				fmt.Println(err)
			}
			var orgstatedetails OrgStateYaml
			err = yaml.Unmarshal([]byte(fileOrgStateYml), &orgstatedetails)
			if err != nil {
				panic(err)
			}

			// Rename checking finished, checking for missing statefiles
			var OrgsYml string
			if ostype == "windows" {
				OrgsYml = cpath + "\\" + clustername + "\\" + OrgName + "\\Org.yml"
			} else {
				OrgsYml = cpath + "/" + clustername + "/" + OrgName + "/Org.yml"
			}

			fileOrgYml, err := ioutil.ReadFile(OrgsYml)
			if err != nil {
				fmt.Println(err)
			}

			var Orgs Orglist
			err = yaml.Unmarshal([]byte(fileOrgYml), &Orgs)
			if err != nil {
				panic(err)
			}

			var spacelist *exec.Cmd
			if ostype == "windows" {
				path := "\"" + "/v3/spaces/?organization_guids=" + orgstatedetails.OrgState.OrgGuid + "\""
				spacelist = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "DeleteorAuditSpaces_spacelist.json")
			} else {
				path := "/v3/spaces/?organization_guids=" + orgstatedetails.OrgState.OrgGuid
				spacelist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "DeleteorAuditSpaces_spacelist.json")
			}

			err = spacelist.Run()
			if err == nil {
				//	fmt.Println(spacelist, spacelist.Stdout, spacelist.Stderr)
			} else {
				fmt.Println("err", spacelist, spacelist.Stdout, spacelist.Stderr)
			}

			fileSpaceJson, err := ioutil.ReadFile("DeleteorAuditSpaces_spacelist.json")
			if err != nil {
				fmt.Println(err)
			}

			var spacelistjson SpaceListJson

			if err := json.Unmarshal(fileSpaceJson, &spacelistjson); err != nil {
				panic(err)
			}

			SpaceLen1 := len(spacelistjson.Resources)
			fmt.Println("Number of Spaces for Org", OrgName, ":", SpaceLen1)

			SpaceLen2 := len(Orgs.Org.Spaces)

			for j := 0; j < SpaceLen2; j++ {

				fullpath := spath + OrgName + "_" + Orgs.Org.Spaces[j].Name + "_SpaceState.yml"
				SpaceStateYml := fullpath

				if ostype == "windows" {
					checkfile = exec.Command("powershell", "-command", "Get-Content", SpaceStateYml)
				} else {
					checkfile = exec.Command("cat", SpaceStateYml)
				}

				if _, err := checkfile.Output(); err == nil {

				} else {

					// Space state file missing, create state file
					var getspacename *exec.Cmd
					if ostype == "windows" {
						path := "\"" + "/v3/spaces?names=" + Orgs.Org.Spaces[j].Name + "&organization_guids=" + orgstatedetails.OrgState.OrgGuid + "\""
						getspacename = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spacedetails_name.json")
					} else {
						path := "/v3/spaces?names=" + Orgs.Org.Spaces[j].Name + "&organization_guids=" + orgstatedetails.OrgState.OrgGuid
						getspacename = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spacedetails_name.json")
					}
					err = getspacename.Run()
					if err == nil {
						//	fmt.Println(getorg, getorg.Stdout, getorg.Stderr)
					} else {
						fmt.Println("err", getspacename, getspacename.Stdout, getspacename.Stderr)
					}
					fileSpaceNameJson, err := ioutil.ReadFile("CreateOrUpdateSpaces_spacedetails_name.json")
					if err != nil {
						fmt.Println(err)
					}
					var spacedetailsname SpaceListJson
					if err := json.Unmarshal(fileSpaceNameJson, &spacedetailsname); err != nil {
						panic(err)
					}
					OrgGuidPull := orgstatedetails.OrgState.OrgGuid
					SpaceStateNameLen := len(spacedetailsname.Resources)

					if SpaceStateNameLen == 0 {
						// Space not yet created or space name has been changed
						fmt.Println("Space not created yet - Org, Space: ", OrgName, Orgs.Org.Spaces[j].Name)

					} else {

						// creating missing state file
						fmt.Println("Org, Space: ", OrgName, Orgs.Org.Spaces[j].Name)
						fmt.Println("Missing state file, creating state")

						OrgGuidPull = orgstatedetails.OrgState.OrgGuid
						spaceguidpull := spacedetailsname.Resources[0].GUID

						type SpaceState struct {
							Org          string `yaml:"Org"`
							OrgGuid      string `yaml:"OrgGuid"`
							OldSpaceName string `yaml:"OldSpaceName"`
							NewSpaceName string `yaml:"NewSpaceName"`
							SpaceGuid    string `yaml:"SpaceGuid"`
						}

						//spath := cpath+"/"+clustername+"-state/"
						values := SpaceState{Org: Orgs.Org.Name, OrgGuid: OrgGuidPull, OldSpaceName: Orgs.Org.Spaces[j].Name, NewSpaceName: Orgs.Org.Spaces[j].Name, SpaceGuid: spaceguidpull}

						var templates *template.Template
						var allFiles []string

						if err != nil {
							fmt.Println(err)
						}

						filename := "SpaceGuid.tmpl"
						fullPath := spath + "SpaceGuid.tmpl"
						if strings.HasSuffix(filename, ".tmpl") {
							allFiles = append(allFiles, fullPath)
						}

						//fmt.Println(allFiles)
						templates, err = template.ParseFiles(allFiles...)
						if err != nil {
							fmt.Println(err)
						}

						s1 := templates.Lookup("SpaceGuid.tmpl")
						f, err := os.Create(spath + Orgs.Org.Name + "_" + Orgs.Org.Spaces[j].Name + "_SpaceState.yml")
						if err != nil {
							panic(err)
						}

						err = s1.Execute(f, values)
						defer f.Close() // don't forget to close the file when finished.
						if err != nil {
							panic(err)
						}

					}
				}
			}

			// checking rename

			if SpaceLen1 != 0 {

				for i := 0; i < SpaceLen1; i++ {

					ExistingOrgName := spacelistjson.Resources[i].Name
					fullpath := spath + OrgName + "_" + ExistingOrgName + "_SpaceState.yml"
					SpaceStateYml := fullpath

					if ostype == "windows" {
						checkfile = exec.Command("powershell", "-command", "Get-Content", SpaceStateYml)
					} else {
						checkfile = exec.Command("cat", SpaceStateYml)
					}
					if _, err := checkfile.Output(); err == nil {
						fmt.Println(checkfile.Stdout)
						fileOrgStateYml, err := ioutil.ReadFile(SpaceStateYml)
						if err != nil {
							fmt.Println(err)
						}
						var spacestatedetails SpaceStateYaml
						err = yaml.Unmarshal([]byte(fileOrgStateYml), &spacestatedetails)
						if err != nil {
							panic(err)
						}

						SpaceNewName := spacestatedetails.SpaceState.NewSpaceName
						SpaceOldName := spacestatedetails.SpaceState.OldSpaceName

						if SpaceNewName == SpaceOldName {
							// No space rename

						} else {

							fmt.Println("Space has been renamed")
							//Changing File name
							oldstatepath := spath + OrgName + "_" + SpaceOldName + "_SpaceState.yml"
							newstatepath := spath + OrgName + "_" + SpaceNewName + "_SpaceState.yml"
							fmt.Println("- ", oldstatepath)
							fmt.Println("+ ", newstatepath)

							if ostype == "windows" {
								changestfile := exec.Command("powershell", "-command", "mv", oldstatepath, newstatepath)
								err := changestfile.Run()
								if err != nil {
									panic(err)
								}
							} else {
								//value := "\""+"mv "+oldstatepath+" "+newstatepath+"\""
								//changestfile := exec.Command("sh", "-c", value)
								changestfile := exec.Command("mv", oldstatepath, newstatepath)
								err := changestfile.Run()
								if err != nil {
									fmt.Println("err", err, changestfile, changestfile.Stdout, changestfile.Stderr)
									//panic(err)
								} else {
									fmt.Println(changestfile, changestfile.Stdout, changestfile.Stderr)
								}
							}
						}
					} else {
						fmt.Println("Statefile missing for space - Org, Space: ", OrgName, spacelistjson.Resources[i].Name)
					}
				}

			}

		} else {
			fmt.Println("Org Statefile missing, please use Org init and then create-org function to create state file")
		}
	}
	return nil
}
func Init(clustername string, endpoint string, user string, org string, space string, asg string, subtree string, githost string, cpath string, orgaudit string, orgman string, spaceaudit string, spaceman string, spacedev string) (err error) {

	type ClusterDetails struct {
		EndPoint             string `yaml:"EndPoint"`
		User                 string `yaml:"User"`
		Org                  string `yaml:"Org"`
		Space                string `yaml:"Space"`
		EnableASG            string `yaml:"EnableASG"`
		EnableGitSubTree     string `yaml:"EnableGitSubTree"`
		GitHost              string `yaml:"GitHost"`
		SetOrgAuditor        string `yaml:"SetOrgAuditor"`
		SetOrgManager        string `yaml:"SetOrgManager"`
		SetSpaceAuditor      string `yaml:"SetSpaceAuditor"`
		SetSpaceManager      string `yaml:"SetSpaceManager"`
		SetSpaceDeveloper    string `yaml:"SetSpaceDeveloper"`
		EnableSpaceAudit     string `yaml:"EnableSpaceAudit"`
		EnableUserAudit      string `yaml:"EnableUserAudit"`
		EnableASGAudit       string `yaml:"EnableASGAudit"`
		EnableIsolationAudit string `yaml:"EnableIsolationAudit"`
		SSOProvider          string `yaml:"SSOProvider"`
	}

	// Cluster configs
	mgmtpath := cpath + "/" + clustername
	ASGPath := cpath + "/" + clustername + "/ProtectedOrgsASGs/"
	QuotasYml := cpath + "/" + clustername + "/Quota.yml"
	ProtectedResourcesYml := cpath + "/" + clustername + "/ProtectedResources.yml"
	ListOrgsYml := cpath + "/" + clustername + "/OrgsList.yml"
	GitListOrgsYml := cpath + "/" + clustername + "/GitOrgsList.yml"

	_, err = os.Stat(mgmtpath)
	if os.IsNotExist(err) {

		fmt.Println("Creating <cluster> folder")
		errDir := os.MkdirAll(mgmtpath, 0755)

		var data = `---
ClusterDetails:
  EndPoint: {{ .EndPoint }}
  User: {{ .User }}
  Org: {{ .Org }}
  Space: {{ .Space }}
  EnableASG: {{ .EnableASG }}
  EnableGitSubTree: {{ .EnableGitSubTree }}
  GitHost: {{ .GitHost }}
  SetOrgAuditor: {{ .SetOrgAuditor }}
  SetOrgManager: {{ .SetOrgManager }}
  SetSpaceAuditor: {{ .SetSpaceAuditor }}
  SetSpaceManager: {{ .SetSpaceManager }}
  SetSpaceDeveloper: {{ .SetSpaceDeveloper }}
  EnableSpaceAudit: {{ .EnableSpaceAudit }}
  EnableUserAudit: {{ .EnableUserAudit }}
  EnableASGAudit: {{ .EnableASGAudit }}
  EnableIsolationAudit: {{ .EnableIsolationAudit }}
  SSOProvider: {{ .SSOProvider }}`

		// Create the file:
		err = ioutil.WriteFile(mgmtpath+"/config.tmpl", []byte(data), 0644)
		check(err)

		values := ClusterDetails{EndPoint: endpoint, User: user, Org: org, Space: space, EnableASG: asg, EnableGitSubTree: subtree, GitHost: githost, SetOrgAuditor: orgaudit, SetOrgManager: orgman, SetSpaceAuditor: spaceaudit, SetSpaceManager: spaceman, SetSpaceDeveloper: spacedev, SSOProvider: "testsso", EnableSpaceAudit: "false", EnableUserAudit: "false", EnableASGAudit: "false", EnableIsolationAudit: "false"}

		var templates *template.Template
		var allFiles []string

		if err != nil {
			fmt.Println(err)
		}

		filename := "config.tmpl"
		fullPath := mgmtpath + "/config.tmpl"
		if strings.HasSuffix(filename, ".tmpl") {
			allFiles = append(allFiles, fullPath)
		}

		fmt.Println(allFiles)
		templates, err = template.ParseFiles(allFiles...)
		if err != nil {
			fmt.Println(err)
		}

		s1 := templates.Lookup("config.tmpl")
		f, err := os.Create(mgmtpath + "/config.yml")
		if err != nil {
			panic(err)
		}

		fmt.Println("Initializing folder and config files")

		err = s1.Execute(f, values)
		defer f.Close() // don't forget to close the file when finished.
		if err != nil {
			panic(err)
		}

		var QuotasTmp = `---
quota:
  - Name: default
    memory_limit: 1024M
    allow_paid_plans: False
    app_instance_limit: 25
    service_instance_limit: 25
  - Name: small_quota
    memory_limit: 2048M
  - Name: medium_quota
    memory_limit: 2048M
  - Name: large_quota
    memory_limit: 2048M
Audit:  list`

		var ProtectedListTmp = `---
Org:
  - system
  - healthwatch
  - dynatrace
quota:
  - default
DefaultRunningSecurityGroup: default_security_group`

		var ListTmp = `---
OrgList:
  - Name: Org-1
    Quota: 
  - Name: Org-2
    Quota:
  - Name: Org-3
    Quota:
Audit: list`

		var GitSubTreeListTmp = `---
OrgList:
  - Name: Org-1
    Repo: Org-1
    Quota:
    Branch:
  - Name: Org-2
    Repo: Org-2
    Quota:
    Branch:
  - Name: Org-3
    Repo: Org-3
    Quota:
    Branch:
Audit: list`

		fmt.Println("Creating <cluster>/ sample yaml files")
		err = ioutil.WriteFile(QuotasYml, []byte(QuotasTmp), 0644)
		check(err)
		err = ioutil.WriteFile(ProtectedResourcesYml, []byte(ProtectedListTmp), 0644)
		check(err)
		err = ioutil.WriteFile(ListOrgsYml, []byte(ListTmp), 0644)
		check(err)
		err = ioutil.WriteFile(GitListOrgsYml, []byte(GitSubTreeListTmp), 0644)
		check(err)

		if errDir != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("<cluster> exists, please manually edit file to make changes or provide new cluster name")
	}

	_, err = os.Stat(ASGPath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(ASGPath, 0755)
		if errDir != nil {
			log.Fatal(err)
			fmt.Println("<cluster>/ASGs exist, please manually edit file to make changes or provide new cluster name")
		} else {
			fmt.Println("Creating <cluster>/ASGs")
		}
	}

	statepath := cpath + "/" + clustername + "-state/"

	_, err = os.Stat(statepath)
	if os.IsNotExist(err) {

		fmt.Println("Creating <cluster-state> folder")
		errDir := os.MkdirAll(statepath, 0755)

		// State Configs
		// State files

		var orgguid = `---
OrgState:
  OldName: {{ .OldName }}
  NewName: {{ .NewName }}
  OrgGuid: {{ .OrgGuid }}`

		err = ioutil.WriteFile(statepath+"OrgGuid.tmpl", []byte(orgguid), 0644)
		check(err)
		if errDir != nil {
			log.Fatal(err)
		}

		var spaceguid = `---
SpaceState:
  Org: {{ .Org }}
  OrgGuid: {{ .OrgGuid }}
  OldSpaceName: {{ .OldSpaceName }}
  NewSpaceName: {{ .NewSpaceName }}
  SpaceGuid: {{ .SpaceGuid }}`

		err = ioutil.WriteFile(statepath+"SpaceGuid.tmpl", []byte(spaceguid), 0644)
		check(err)
		if errDir != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("<cluster-state> exists, please manually edit file to make changes or provide new cluster name")
	}

	return
}
func check(e error) {
	if e != nil {
		fmt.Println("<cluster>/ yamls exists, please manually edit file to make changes or provide new cluster name")
		panic(e)
	}
}
