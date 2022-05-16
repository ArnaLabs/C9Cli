package createorupdateorgs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

type List struct {
	OrgList []struct {
		Name string `yaml:"Name"`
		Quota string `yaml:"Quota"`
	} `yaml:"OrgList"`
	Audit string `yaml:"Audit"`
}
type GitList struct {
	OrgList []struct {
		Name string `yaml:"Name"`
		Repo string `yaml:"Repo"`
		Quota string `yaml:"Quota"`
		Branch string `yaml:"Branch"`
	} `yaml:"OrgList"`
	Audit string `yaml:"Audit"`
}
type Orglist struct {
	Org struct {
		Name     string `yaml:"Name"`
		//Quota    string `yaml:"Quota"`
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
			ASG string `yaml:"ASG"`
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
	SpaceAudit string `yaml:"SpaceAudit"`
	UserAudit string `yaml:"UserAudit"`
	ASGAudit string `yaml:"ASGAudit"`
	IsolationAudit string `yaml:"IsolationAudit"`
}
type ProtectedList struct {
	Org   []string `yaml:"Org"`
	Quota []string `yaml:"quota"`
	DefaultRunningSecurityGroup string   `yaml:"DefaultRunningSecurityGroup"`
}
type InitClusterConfigVals struct {
	ClusterDetails struct {
		EndPoint  string `yaml:"EndPoint"`
		User      string `yaml:"User"`
		Org       string `yaml:"Org"`
		Space     string `yaml:"Space"`
		EnableASG bool   `yaml:"EnableASG"`
		EnableGitSubTree bool `yaml:"EnableGitSubTree"`
		GitHost string `yaml:"GitHost"`
		SetOrgAuditor bool	`yaml:"SetOrgAuditor"`
		SetOrgManager bool	`yaml:"SetOrgManager"`
		SetSpaceAuditor bool	`yaml:"SetSpaceAuditor"`
		SetSpaceManager bool	`yaml:"SetSpaceManager"`
		SetSpaceDeveloper bool	`yaml:"SetSpaceDeveloper"`
		EnableSpaceAudit bool `yaml:"EnableSpaceAudit"`
		EnableUserAudit bool `yaml:"EnableUserAudit"`
		EnableASGAudit bool `yaml:"EnableASGAudit"`
		EnableIsolationAudit bool `yaml:"EnableIsolationAudit"`
		SSOProvider string `yaml:"SSOProvider"`
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
		Org     string `yaml:"Org"`
		OrgGuid string `yaml:"OrgGuid"`
		OldSpaceName string `yaml:"OldSpaceName"`
		NewSpaceName string `yaml:"NewSpaceName"`
		SpaceGuid    string `yaml:"SpaceGuid"`
	} `yaml:"SpaceState"`
}
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
type OrgListJson struct {
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
		Suspended     bool      `json:"suspended"`
		Relationships struct {
			Quota struct {
				Data struct {
					GUID string `json:"guid"`
				} `json:"data"`
			} `json:"quota"`
		} `json:"relationships"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Domains struct {
				Href string `json:"href"`
			} `json:"domains"`
			DefaultDomain struct {
				Href string `json:"href"`
			} `json:"default_domain"`
			Quota struct {
				Href string `json:"href"`
			} `json:"quota"`
		} `json:"links"`
		Metadata struct {
			Labels struct {
			} `json:"labels"`
			Annotations struct {
			} `json:"annotations"`
		} `json:"metadata"`
	} `json:"resources"`
}
type QuotaListJson struct {
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
		GUID      string    `json:"guid"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
		Apps      struct {
			TotalMemoryInMb      int         `json:"total_memory_in_mb"`
			PerProcessMemoryInMb interface{} `json:"per_process_memory_in_mb"`
			TotalInstances       int         `json:"total_instances"`
			PerAppTasks          interface{} `json:"per_app_tasks"`
		} `json:"apps"`
		Services struct {
			PaidServicesAllowed   bool        `json:"paid_services_allowed"`
			TotalServiceInstances int         `json:"total_service_instances"`
			TotalServiceKeys      interface{} `json:"total_service_keys"`
		} `json:"services"`
		Routes struct {
			TotalRoutes        interface{} `json:"total_routes"`
			TotalReservedPorts int         `json:"total_reserved_ports"`
		} `json:"routes"`
		Domains struct {
			TotalDomains interface{} `json:"total_domains"`
		} `json:"domains"`
		Relationships struct {
			Organizations struct {
				Data []struct {
					GUID string `json:"guid"`
				} `json:"data"`
			} `json:"organizations"`
		} `json:"relationships"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"resources"`
}
type OrgUsersListJson struct {
	Pagination struct {
		TotalResults int `json:"total_results"`
		TotalPages   int `json:"total_pages"`
		First        struct {
			Href string `json:"href"`
		} `json:"first"`
		Last struct {
			Href string `json:"href"`
		} `json:"last"`
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Previous interface{} `json:"previous"`
	} `json:"pagination"`
	Resources []struct {
		GUID          string    `json:"guid"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		Type          string    `json:"type"`
		Relationships struct {
			User struct {
				Data struct {
					GUID string `json:"guid"`
				} `json:"data"`
			} `json:"user"`
			Organization struct {
				Data struct {
					GUID string `json:"guid"`
				} `json:"data"`
			} `json:"organization"`
			Space struct {
				Data interface{} `json:"data"`
			} `json:"space"`
		} `json:"relationships"`
	} `json:"resources"`
}
type UserDetailsJson struct {
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
		GUID             string    `json:"guid"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		Username         string    `json:"username"`
		PresentationName string    `json:"presentation_name"`
		Origin           string    `json:"origin"`
		Metadata         struct {
			Labels struct {
			} `json:"labels"`
			Annotations struct {
			} `json:"annotations"`
		} `json:"metadata"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"resources"`
}
type SpaceUsersListJson struct {
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
		Type          string    `json:"type"`
		Relationships struct {
			User struct {
				Data struct {
					GUID string `json:"guid"`
				} `json:"data"`
			} `json:"user"`
			Space struct {
				Data struct {
					GUID string `json:"guid"`
				} `json:"data"`
			} `json:"space"`
			Organization struct {
				Data interface{} `json:"data"`
			} `json:"organization"`
		} `json:"relationships"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			User struct {
				Href string `json:"href"`
			} `json:"user"`
			Space struct {
				Href string `json:"href"`
			} `json:"space"`
		} `json:"links"`
	} `json:"resources"`
}
type ASGListJson struct {
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
		GUID      string    `json:"guid"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
		Rules     []struct {
			Destination string `json:"destination"`
			Protocol    string `json:"protocol"`
		} `json:"rules"`
		GloballyEnabled struct {
			Running bool `json:"running"`
			Staging bool `json:"staging"`
		} `json:"globally_enabled"`
		Relationships struct {
			RunningSpaces struct {
				Data []struct {
					GUID string `json:"guid"`
				} `json:"data"`
			} `json:"running_spaces"`
			StagingSpaces struct {
				Data []struct {
					GUID string `json:"guid"`
				} `json:"data"`
			} `json:"staging_spaces"`
		} `json:"relationships"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"resources"`
}
type Quotalist struct {
	Quota []struct {
		Name        string `yaml:"Name"`
		MemoryLimit string `yaml:"memory_limit"`
		AllowPaidPlans       bool `yaml:"allow_paid_plans"`
		AppInstanceLimit     string `yaml:"app_instance_limit"`
		ServiceInstanceLimit string `yaml:"service_instance_limit"`
	} `yaml:"quota"`
	Audit string `yaml:"Audit"`
}

func CreateOrUpdateOrgs(clustername string, cpath string, ostype string) error {

	var list List
	var gitlist GitList
	var ProtectedOrgs ProtectedList
	var InitClusterConfigVals InitClusterConfigVals
	var ListYml string
	spath := cpath+"/"+clustername+"-state/"

	ConfigFile := cpath+"/"+clustername+"/config.yml"
	fileConfigYml, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileConfigYml), &InitClusterConfigVals)
	if err != nil {
		panic(err)
	}
	if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true{
		ListYml = cpath+"/"+clustername+"/OrgsList.yml"
		fileOrgYml, err := ioutil.ReadFile(ListYml)
		if err != nil {
			fmt.Println(err)
		}
		err = yaml.Unmarshal([]byte(fileOrgYml), &list)
		if err != nil {
			panic(err)
		}
	} else {
		ListYml = cpath+"/"+clustername+"/GitOrgsList.yml"
		fileOrgYml, err := ioutil.ReadFile(ListYml)
		if err != nil {
			fmt.Println(err)
		}

		err = yaml.Unmarshal([]byte(fileOrgYml), &gitlist)
		if err != nil {
			panic(err)
		}
	}


	ProtectedOrgsYml := cpath+"/"+clustername+"/ProtectedResources.yml"
	fileProtectedYml, err := ioutil.ReadFile(ProtectedOrgsYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileProtectedYml), &ProtectedOrgs)
	if err != nil {
		panic(err)
	}

	var LenList int
	if 	InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
		LenList = len(list.OrgList)
	} else {
		LenList = len(gitlist.OrgList)
	}

	LenProtectedOrgs := len(ProtectedOrgs.Org)

	for i := 0; i < LenList; i++ {

		var OrgName, QuotaName string
		if 	InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
			OrgName = list.OrgList[i].Name
			QuotaName = list.OrgList[i].Quota
		} else {
			OrgName = gitlist.OrgList[i].Name
			//RepoName = gitlist.OrgList[i].Name
			QuotaName = gitlist.OrgList[i].Quota
		}
		var count, totalcount int
		fmt.Println(" ")
		fmt.Println("Org: ", OrgName)
		for p := 0; p < LenProtectedOrgs; p++ {
			//	fmt.Println("Protected Org: ", ProtectedOrgs.Org[p], ",", list.OrgList[i])
			if ProtectedOrgs.Org[p] == OrgName {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count

		}
		if totalcount == 0 {

			OrgsYml := cpath+"/"+clustername+"/"+OrgName+"/Org.yml"
			fileOrgYml, err := ioutil.ReadFile(OrgsYml)
			if err != nil {
				fmt.Println(err)
			}

			var Orgs Orglist
			err = yaml.Unmarshal([]byte(fileOrgYml), &Orgs)
			if err != nil {
				panic(err)
			}

			if OrgName == Orgs.Org.Name {

				fullpath := spath+OrgName+"_OrgState.yml"
				var OrgGuidPull string

				// Getting Org Guid from State file
				OrgsStateYml := fullpath
				fileOrgStateYml, err := ioutil.ReadFile(OrgsStateYml)
				if err != nil {
					fmt.Println(err)
				}
				var orgstatedetails OrgStateYaml
				err = yaml.Unmarshal([]byte(fileOrgStateYml), &orgstatedetails)
				if err != nil {
					panic(err)
				}
				OrgStateGuid := orgstatedetails.OrgState.OrgGuid

				// Checking if org exist with Guid from State File
				var getorgguid *exec.Cmd
				if ostype == "windows" {
					path := "\""+"/v3/organizations?guids="+OrgStateGuid+"\""
					getorgguid = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgs_orgdetails_guid.json")
				} else {
					path := "/v3/organizations?guids="+OrgStateGuid
					getorgguid = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgs_orgdetails_guid.json")
				}
				err = getorgguid.Run()
				if err == nil {
					//fmt.Println(getorgguid, getorgguid.Stdout, getorgguid.Stderr)
				} else {
					fmt.Println("err", getorgguid, getorgguid.Stdout, getorgguid.Stderr)
				}
				fileSpaceGuidJson, err := ioutil.ReadFile("CreateOrUpdateOrgs_orgdetails_guid.json")
				if err != nil {
					fmt.Println(err)
				}

				var orgdetailsguid OrgListJson
				if err := json.Unmarshal(fileSpaceGuidJson, &orgdetailsguid); err != nil {
					panic(err)
				}

				// Checking if org exist with Org Name
				var getorgname *exec.Cmd
				if ostype == "windows" {
					path := "\""+"/v3/organizations?names="+Orgs.Org.Name+"\""
					getorgname = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgs_orgdetails_name.json")
				} else {
					path := "/v3/organizations?names="+Orgs.Org.Name
					getorgname = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgs_orgdetails_name.json")
				}
				err = getorgname.Run()
				if err == nil {
					//fmt.Println(getorgname, getorgname.Stdout, getorgname.Stderr)
				} else {
					fmt.Println("err", getorgname, getorgname.Stdout, getorgname.Stderr)
				}
				fileSpaceNameJson, err := ioutil.ReadFile("CreateOrUpdateOrgs_orgdetails_name.json")
				if err != nil {
					fmt.Println(err)
				}
				var orgdetailsname OrgListJson
				if err := json.Unmarshal(fileSpaceNameJson, &orgdetailsname); err != nil {
					panic(err)
				}

				OrgStateGuidLen := len(orgdetailsguid.Resources)
				OrgStateNameLen := len(orgdetailsname.Resources)


				if OrgStateGuidLen != 0 {
					//fmt.Println("Org exists in state and platform")
					//Checking if Org name has changed

					orgname := orgdetailsguid.Resources[0].Name

					if Orgs.Org.Name == orgname {
						OrgGuidPull = orgdetailsguid.Resources[0].GUID
					} else {
						fmt.Println("Resetting Org Name")
						fmt.Println("- ", orgname)
						fmt.Println("+ ", Orgs.Org.Name)
						renameorg := exec.Command("cf", "rename-org", orgname,  Orgs.Org.Name)
						err = renameorg.Run()
						if err == nil {
							//	fmt.Println(getorg, getorg.Stdout, getorg.Stderr)
						} else {
							fmt.Println("err", renameorg, renameorg.Stdout, renameorg.Stderr)
						}

						//OrgGuidPull = orgdetailsguid.Resources[0].GUID
						// Update State file
						// Moved Below
						////type OrgState struct {
						////	OldName string `yaml:"OldName"`
						////	NewName string `yaml:"NewName"`
						////	OrgGuid    string `yaml:"OrgGuid"`
						////}

						//spath := cpath+"/"+clustername+"-state/"
						////values := OrgState{OldName: Orgs.Org.Name, NewName: Orgs.Org.Name, OrgGuid: OrgGuidPull}

						////var templates *template.Template
						////var allFiles []string

						////if err != nil {
						////fmt.Println(err)
						////}

						////filename := "OrgGuid.tmpl"
						////fullPath := spath+"OrgGuid.tmpl"
						////if strings.HasSuffix(filename, ".tmpl") {
						////allFiles = append(allFiles, fullPath)
						////}

						//fmt.Println(allFiles)
						////templates, err = template.ParseFiles(allFiles...)
						////if err != nil {
						////fmt.Println(err)
						////}

						////s1 := templates.Lookup("OrgGuid.tmpl")
						////f, err := os.Create(spath+Orgs.Org.Name+"_OrgState.yml")
						////if err != nil {
						////	panic(err)
						////}

						////err = s1.Execute(f, values)
						////defer f.Close() // don't forget to close the file when finished.
						////if err != nil {
						////panic(err)
						////}
					}

					//Checking if Quota has changed
					var getquotaName *exec.Cmd
					quotaguid := orgdetailsguid.Resources[0].Relationships.Quota.Data.GUID
					if ostype == "windows" {
						path := "\""+"/v3/organization_quotas?guids="+quotaguid+"\""
						getquotaName = exec.Command("powershell", "-command","cf", "curl", path, "--output", "CreateOrUpdateOrgs_quotaname.json")
					} else {
						path := "/v3/organization_quotas?guids="+quotaguid
						getquotaName = exec.Command("cf", "curl", path, "--output", "CreateOrUpdateOrgs_quotaname.json")
					}
					err := getquotaName.Run()
					if err == nil {
						//fmt.Println(getquotaName, getquotaName.Stdout, getquotaName.Stderr)
					} else {
						fmt.Println("err", getquotaName, getquotaName.Stdout, getquotaName.Stderr)
					}
					//
					fileNameJson, err := ioutil.ReadFile("CreateOrUpdateOrgs_quotaname.json")
					if err != nil {
						fmt.Println(err)
					}
					var body QuotaListJson
					if err := json.Unmarshal(fileNameJson, &body); err != nil {
						panic(err)
					}
					if body.Resources[0].Name != QuotaName {
						fmt.Println("quota -", body.Resources[0].Name)
						fmt.Println("quota +", QuotaName)

						fmt.Println("Updating Org quota")
						SetQuota := exec.Command("cf", "set-quota", Orgs.Org.Name, QuotaName)
						if _, err := SetQuota.Output(); err != nil{
							fmt.Println("command: ", SetQuota)
							fmt.Println("Err: ", SetQuota.Stdout, err)
						} else {
							// fmt.Println("command: ", SetQuota)
							// fmt.Println(SetQuota.Stdout)
						}
					}
					var out bytes.Buffer
					pullguid := exec.Command("cf", "org", Orgs.Org.Name, "--guid")
					pullguid.Stdout = &out
					err = pullguid.Run()
					OrgGuidPull = out.String()
					out.Reset()

				} else if OrgStateGuidLen == 0 && OrgStateNameLen != 0 {
					fmt.Println("Missing State file, Please use org-init function to create state files")
				} else if OrgStateGuidLen == 0 && OrgStateNameLen == 0 {

					// Org does't exist
					createorg := exec.Command("cf", "create-org", Orgs.Org.Name)
					fmt.Println("Creating Org")
					fmt.Println("+", Orgs.Org.Name)
					if _, err := createorg.Output(); err != nil{
						fmt.Println("command: ", createorg)
						fmt.Println("Err: ", createorg.Stdout, err)
					} else {
						//	fmt.Println("command: ", createorg)
						fmt.Println(createorg.Stdout)
					}
					attachquota := exec.Command("cf", "set-quota", Orgs.Org.Name, QuotaName)
					fmt.Println("Attaching Quota")
					if QuotaName == ""{
						QuotaName = "default"
					}
					fmt.Println("+", QuotaName)
					if _, err := attachquota.Output(); err != nil{
						fmt.Println("command: ", attachquota)
						fmt.Println("Err: ", attachquota.Stdout, err)
					} else {
						//fmt.Println("command: ", attachquota)
						fmt.Println(attachquota.Stdout)
					}
					var out bytes.Buffer
					pullguid := exec.Command("cf", "org", Orgs.Org.Name, "--guid")
					pullguid.Stdout = &out
					err = pullguid.Run()
					OrgGuidPull = out.String()
					out.Reset()
				}

				// Creating state file

				if OrgStateGuidLen != 0 {

					type OrgState struct {
						OldName string `yaml:"OldName"`
						NewName string `yaml:"NewName"`
						OrgGuid    string `yaml:"OrgGuid"`
					}

					//spath := cpath+"/"+clustername+"-state/"
					values := OrgState{OldName: Orgs.Org.Name, NewName: Orgs.Org.Name, OrgGuid: OrgGuidPull}

					var templates *template.Template
					var allFiles []string

					if err != nil {
						fmt.Println(err)
					}

					filename := "OrgGuid.tmpl"
					fullPath := spath+"OrgGuid.tmpl"
					if strings.HasSuffix(filename, ".tmpl") {
						allFiles = append(allFiles, fullPath)
					}

					//fmt.Println(allFiles)
					templates, err = template.ParseFiles(allFiles...)
					if err != nil {
						fmt.Println(err)
					}

					s1 := templates.Lookup("OrgGuid.tmpl")
					f, err := os.Create(spath+Orgs.Org.Name+"_OrgState.yml")
					if err != nil {
						panic(err)
					}

					err = s1.Execute(f, values)
					defer f.Close() // don't forget to close the file when finished.
					if err != nil {
						panic(err)
					}


				} else if OrgStateGuidLen == 0 && OrgStateNameLen != 0 {

				} else if  OrgStateGuidLen == 0 && OrgStateNameLen == 0{

					type OrgState struct {
						OldName string `yaml:"OldName"`
						NewName string `yaml:"NewName"`
						OrgGuid    string `yaml:"OrgGuid"`
					}

					//spath := cpath+"/"+clustername+"-state/"
					values := OrgState{OldName: Orgs.Org.Name, NewName: Orgs.Org.Name, OrgGuid: OrgGuidPull}

					var templates *template.Template
					var allFiles []string

					if err != nil {
						fmt.Println(err)
					}

					filename := "OrgGuid.tmpl"
					fullPath := spath+"OrgGuid.tmpl"
					if strings.HasSuffix(filename, ".tmpl") {
						allFiles = append(allFiles, fullPath)
					}

					//fmt.Println(allFiles)
					templates, err = template.ParseFiles(allFiles...)
					if err != nil {
						fmt.Println(err)
					}

					s1 := templates.Lookup("OrgGuid.tmpl")
					f, err := os.Create(spath+Orgs.Org.Name+"_OrgState.yml")
					if err != nil {
						panic(err)
					}

					err = s1.Execute(f, values)
					defer f.Close() // don't forget to close the file when finished.
					if err != nil {
						panic(err)
					}

				}

			} else {
				fmt.Println("Org Name does't match with folder name")
			}

		} else {
			fmt.Println("This is a protected Org")
		}
	}
	results := exec.Command("cf", "orgs")
	if _, err := results.Output(); err != nil{
		fmt.Println("command: ", results)
		fmt.Println("Err: ", results.Stdout, err)
	} else {
		//fmt.Println("command: ", results)
		fmt.Println(results.Stdout)
	}
	return err
}
