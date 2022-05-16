package createorupdateorgusers

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
	"strings"
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

func CreateOrUpdateOrgUsers(clustername string, cpath string, ostype string) error {

	var list List
	var ProtectedOrgs ProtectedList
	var gitlist GitList
	var InitClusterConfigVals InitClusterConfigVals
	var ListYml string
	var  LenList int
	ConfigFile := cpath+"/"+clustername+"/config.yml"
	fileConfigYml, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileConfigYml), &InitClusterConfigVals)
	if err != nil {
		panic(err)
	}

	if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {

		ListYml = cpath+"/"+clustername+"/OrgsList.yml"
		fileOrgYml, err := ioutil.ReadFile(ListYml)
		if err != nil {
			fmt.Println(err)
		}
		err = yaml.Unmarshal([]byte(fileOrgYml), &list)
		if err != nil {
			panic(err)
		}
		LenList = len(list.OrgList)

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
		LenList = len(gitlist.OrgList)

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


	LenProtectedOrgs := len(ProtectedOrgs.Org)
	for i := 0; i < LenList; i++ {

		var count, totalcount int
		var OrgName string
		if 	InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
			OrgName = list.OrgList[i].Name
		} else {
			OrgName = gitlist.OrgList[i].Name
			//RepoName = gitlist.OrgList[i].Name
		}
		fmt.Println(" ")
		fmt.Println("Org: ", OrgName)
		for p := 0; p < LenProtectedOrgs; p++ {
			//fmt.Println("Protected Org: ", ProtectedOrgs.Org[p], ",", list.OrgList[i])
			if ProtectedOrgs.Org[p] == OrgName {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {

			//fmt.Println("This is not Protected Org")
			//fmt.Println("test empty",Orgs)

			var Orgs Orglist
			OrgsYml := cpath+"/"+clustername+"/"+OrgName+"/Org.yml"
			//fmt.Println(OrgsYml)

			fileOrgdetsYml, err := ioutil.ReadFile(OrgsYml)

			if err != nil {
				fmt.Println(err)
			}

			err = yaml.Unmarshal([]byte(fileOrgdetsYml), &Orgs)
			if err != nil {
				panic(err)
			}

			if OrgName == Orgs.Org.Name {

				var getorg *exec.Cmd
				//fmt.Println(Orgs)
				//fmt.Println(list.OrgList[i], Orgs.Org.Name)

				if ostype == "windows" {
					path := strings.TrimSpace("\""+"/v3/organizations?names="+Orgs.Org.Name+"\"")
					getorg = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgdetails.json")
				} else {
					path := strings.TrimSpace("/v3/organizations?names="+Orgs.Org.Name)
					getorg = exec.Command("cf", "curl", path, "--output", "CreateOrUpdateOrgsUsers_orgdetails.json")
				}

				//path := strings.TrimSpace("/v3/organizations?names="+Orgs.Org.Name)
				//getorg := exec.Command("cf", "curl", "\""+path+"\"", "--output", "CreateOrUpdateOrgsUsers_orgdetails.json")
				//var out bytes.Buffer
				//getorg.Stdout = &out

				err := getorg.Run()
				if err == nil {
					//fmt.Println(getorg, getorg.Stdout, getorg.Stderr)
				} else {
					fmt.Println("err", getorg, getorg.Stdout, getorg.Stderr)
				}

				fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_orgdetails.json")
				if err != nil {
					fmt.Println(err)
				}

				var orgdetails OrgListJson
				if err := json.Unmarshal(fileSpaceJson, &orgdetails); err != nil {
					panic(err)
				}

				OrgLen := len(orgdetails.Resources)
				if OrgLen == 0 {
					fmt.Println("Org doesn't exist, Please Create Org")
				} else {

					target := exec.Command("cf", "t", "-o", Orgs.Org.Name)
					if _, err := target.Output(); err == nil {

						//		fmt.Println("command: ", target)
						//		fmt.Println(target.Stdout)

						var orguserslist *exec.Cmd
						orgguid := orgdetails.Resources[0].GUID
						if ostype == "windows" {
							path := "\""+"/v3/roles/?organization_guids="+orgguid+"\""
							orguserslist = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrslist.json")
						} else {
							path := "/v3/roles/?organization_guids="+orgguid
							orguserslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrslist.json")
						}
						err := orguserslist.Run()
						if err == nil {
							//fmt.Println(orguserslist, orguserslist.Stdout, orguserslist.Stderr)
						} else {
							fmt.Println("err", orguserslist, orguserslist.Stdout, orguserslist.Stderr)
						}

						fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_orgusrslist.json")
						if err != nil {
							fmt.Println(err)
						}

						var orgusrslist OrgUsersListJson

						if err := json.Unmarshal(fileSpaceJson, &orgusrslist); err != nil {
							panic(err)
						}

						OrgUsrLen := len(orgusrslist.Resources)

						//fmt.Println("Number of Users currently exist in Org", Orgs.Org.Name, ":", OrgUsrLen)
						if OrgUsrLen != 0 {

							if InitClusterConfigVals.ClusterDetails.SetOrgManager == true {
								LDAPOrgManLen := len(Orgs.Org.OrgUsers.LDAP.OrgManagers)
								for j := 0; j < LDAPOrgManLen; j++ {

									var getspace *exec.Cmd
									if ostype == "windows" {
										path := "\""+"/v3/users/?usernames=" + Orgs.Org.OrgUsers.LDAP.OrgManagers[j] + "&origins=ldap"+"\""
										getspace = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_userguidfind.json")

									} else {
										path := "/v3/users/?usernames=" + Orgs.Org.OrgUsers.LDAP.OrgManagers[j] + "&origins=ldap"
										getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_userguidfind.json")

									}
									err := getspace.Run()

									if err == nil {
										//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
									} else {
										fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
									}

									var usedetails UserDetailsJson
									fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_userguidfind.json")
									if err != nil {
										fmt.Println(err)
									}
									if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
										panic(err)
									}

									if len(usedetails.Resources) != 0 {

										userguid := usedetails.Resources[0].GUID
										var orgusersdetailslist *exec.Cmd

										if ostype == "windows" {
											path := "\""+"/v3/roles/?organization_guids=" + orgguid + "&user_guids=" + userguid + "&types=organization_manager"+"\""
											orgusersdetailslist = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										} else {
											path := "/v3/roles/?organization_guids=" + orgguid + "&user_guids=" + userguid + "&types=organization_manager"
											orgusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										}

										err := orgusersdetailslist.Run()
										if err == nil {
											//fmt.Println(orgusersdetailslist, orgusersdetailslist.Stdout, orgusersdetailslist.Stderr)
										} else {
											fmt.Println("err", orgusersdetailslist, orgusersdetailslist.Stdout, orgusersdetailslist.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										if err != nil {
											fmt.Println(err)
										}

										var orgusrslist OrgUsersListJson
										if err := json.Unmarshal(fileSpaceJson, &orgusrslist); err != nil {
											panic(err)
										}
										OrgUsrdetailsLen := len(orgusrslist.Resources)

										if OrgUsrdetailsLen != 0 {

										} else {

											fmt.Println("+ ", Orgs.Org.OrgUsers.LDAP.OrgManagers[j], ",", "LDAP OrgManager")
											cmd := exec.Command("cf", "set-org-role", Orgs.Org.OrgUsers.LDAP.OrgManagers[j], Orgs.Org.Name, "OrgManager")
											if _, err := cmd.Output(); err != nil {
												fmt.Println("command: ", cmd)
												fmt.Println("Err: ", cmd.Stdout, err)
											} else {
												//fmt.Println("command: ", cmd)
												fmt.Println(cmd.Stdout)
											}
										}
									} else {
										fmt.Println(Orgs.Org.OrgUsers.LDAP.OrgManagers[j], "LDAP OrgManager User does't exist in foundation, please ask user to login to apps manager")
									}
								}

								UAAOrgManLen := len(Orgs.Org.OrgUsers.UAA.OrgManagers)
								for j := 0; j < UAAOrgManLen; j++ {

									//fmt.Println(Orgs.Org.Name)
									//fmt.Println(Orgs.Org.OrgUsers.UAA.OrgManagers[j])
									var getspace *exec.Cmd

									if ostype == "windows" {
										path := "\""+"/v3/users/?usernames=" + Orgs.Org.OrgUsers.UAA.OrgManagers[j] + "&origins=uaa"+"\""
										getspace = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_userguidfind.json")
									} else {
										path := "/v3/users/?usernames=" + Orgs.Org.OrgUsers.UAA.OrgManagers[j] + "&origins=uaa"
										getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_userguidfind.json")
									}
									err := getspace.Run()
									if err == nil {
										//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
									} else {
										fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
									}

									fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_userguidfind.json")
									if err != nil {
										fmt.Println(err)
									}

									var usedetails UserDetailsJson
									if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
										panic(err)
									}

									if len(usedetails.Resources) != 0 {

										var orgusersdetailslist *exec.Cmd
										userguid := usedetails.Resources[0].GUID

										if ostype == "windows" {
											path := "\""+"/v3/roles/?organization_guids=" + orgguid + "&user_guids=" + userguid + "&types=organization_manager"+"\""
											orgusersdetailslist = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										} else {
											path := "/v3/roles/?organization_guids=" + orgguid + "&user_guids=" + userguid + "&types=organization_manager"
											orgusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										}
										err := orgusersdetailslist.Run()
										if err == nil {
											//fmt.Println(orgusersdetailslist, orgusersdetailslist.Stdout, orgusersdetailslist.Stderr)
										} else {
											fmt.Println("err", orgusersdetailslist, orgusersdetailslist.Stdout, orgusersdetailslist.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										if err != nil {
											fmt.Println(err)
										}

										var orgusrslist OrgUsersListJson
										if err := json.Unmarshal(fileSpaceJson, &orgusrslist); err != nil {
											panic(err)
										}

										OrgUsrdetailsLen := len(orgusrslist.Resources)

										if OrgUsrdetailsLen != 0 {

										} else {

											fmt.Println("+ ", Orgs.Org.OrgUsers.UAA.OrgManagers[j], ",", "UAA OrgManager")
											cmd := exec.Command("cf", "set-org-role", Orgs.Org.OrgUsers.UAA.OrgManagers[j], Orgs.Org.Name, "OrgManager")
											if _, err := cmd.Output(); err != nil {
												fmt.Println("command: ", cmd)
												fmt.Println("Err: ", cmd.Stdout, err)
											} else {
												//fmt.Println("command: ", cmd)
												fmt.Println(cmd.Stdout)
											}
										}
									} else {
										fmt.Println(Orgs.Org.OrgUsers.UAA.OrgManagers[j], "UAA OrgManager User does't exist in foundation, please ask user to login to apps manager")
									}

								}

								SSOOrgManLen := len(Orgs.Org.OrgUsers.SSO.OrgManagers)
								SSOProvider := InitClusterConfigVals.ClusterDetails.SSOProvider
								for j := 0; j < SSOOrgManLen; j++ {

									var getspace *exec.Cmd

									if ostype == "windows" {
										path := "\""+"/v3/users/?usernames=" + Orgs.Org.OrgUsers.SSO.OrgManagers[j] + "&origins="+SSOProvider+"\""
										getspace = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_userguidfind.json")
									} else {
										path := "/v3/users/?usernames=" + Orgs.Org.OrgUsers.SSO.OrgManagers[j] + "&origins="+SSOProvider
										getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_userguidfind.json")
									}

									err := getspace.Run()
									if err == nil {
										//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
									} else {
										fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
									}

									fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_userguidfind.json")
									if err != nil {
										fmt.Println(err)
									}
									var usedetails UserDetailsJson
									if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
										panic(err)
									}

									if len(usedetails.Resources) != 0 {

										var orgusersdetailslist *exec.Cmd
										userguid := usedetails.Resources[0].GUID
										if ostype == "windows" {
											path := "\""+"/v3/roles/?organization_guids=" + orgguid + "&user_guids=" + userguid + "&types=organization_manager"+"\""
											orgusersdetailslist = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										} else {
											path := "/v3/roles/?organization_guids=" + orgguid + "&user_guids=" + userguid + "&types=organization_manager"
											orgusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										}
										err := orgusersdetailslist.Run()
										if err == nil {
											//fmt.Println(orgusersdetailslist, orgusersdetailslist.Stdout, orgusersdetailslist.Stderr)
										} else {
											fmt.Println("err", orgusersdetailslist, orgusersdetailslist.Stdout, orgusersdetailslist.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										if err != nil {
											fmt.Println(err)
										}
										var orgusrslist OrgUsersListJson
										if err := json.Unmarshal(fileSpaceJson, &orgusrslist); err != nil {
											panic(err)
										}

										OrgUsrdetailsLen := len(orgusrslist.Resources)

										if OrgUsrdetailsLen != 0 {

										} else {

											fmt.Println("+ ", Orgs.Org.OrgUsers.SSO.OrgManagers[j], ",", "SSO OrgManager")
											cmd := exec.Command("cf", "set-org-role", Orgs.Org.OrgUsers.SSO.OrgManagers[j], Orgs.Org.Name, "OrgManager")
											if _, err := cmd.Output(); err != nil {
												fmt.Println("command: ", cmd)
												fmt.Println("Err: ", cmd.Stdout, err)
											} else {
												//fmt.Println("command: ", cmd)
												fmt.Println(cmd.Stdout)
											}
										}
									} else {
										fmt.Println(Orgs.Org.OrgUsers.SSO.OrgManagers[j], "SSO OrgManager User does't exist in foundation, please ask user to login to apps manager")
									}
								}
							} else {
								fmt.Println("Set OrgManager:", InitClusterConfigVals.ClusterDetails.SetOrgManager)
								fmt.Println("set SetOrgManager field to true to manage OrgManagers for the Org")
							}

							if InitClusterConfigVals.ClusterDetails.SetOrgAuditor == true {

								LDAPOrgAudLen := len(Orgs.Org.OrgUsers.LDAP.OrgAuditors)
								for j := 0; j < LDAPOrgAudLen; j++ {

									//fmt.Println(Orgs.Org.Name)
									//fmt.Println(Orgs.Org.OrgUsers.LDAP.OrgAuditors[j])

									var getspace *exec.Cmd
									if ostype == "windows" {
										path := "\""+"/v3/users/?usernames=" + Orgs.Org.OrgUsers.LDAP.OrgAuditors[j] + "&origins=ldap"+"\""
										getspace = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_userguidfind.json")

									} else {
										path := "/v3/users/?usernames=" + Orgs.Org.OrgUsers.LDAP.OrgAuditors[j] + "&origins=ldap"
										getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_userguidfind.json")
									}

									err := getspace.Run()
									if err == nil {
										//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
									} else {
										fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
									}

									fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_userguidfind.json")
									if err != nil {
										fmt.Println(err)
									}

									var usedetails UserDetailsJson
									if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
										panic(err)
									}

									if len(usedetails.Resources) != 0 {

										//fmt.Println("Hi")
										var orgusersdetailslist *exec.Cmd
										userguid := usedetails.Resources[0].GUID
										if ostype == "windows" {
											path := "\""+"/v3/roles/?organization_guids=" + orgguid + "&user_guids=" + userguid + "&types=organization_auditor"+"\""
											orgusersdetailslist = exec.Command( "powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										} else {
											path := "/v3/roles/?organization_guids=" + orgguid + "&user_guids=" + userguid + "&types=organization_auditor"
											orgusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										}

										err := orgusersdetailslist.Run()
										if err == nil {
											//fmt.Println(orgusersdetailslist, orgusersdetailslist.Stdout, orgusersdetailslist.Stderr)
										} else {
											fmt.Println("err", orgusersdetailslist, orgusersdetailslist.Stdout, orgusersdetailslist.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										if err != nil {
											fmt.Println(err)
										}
										var orgusrslist OrgUsersListJson
										if err := json.Unmarshal(fileSpaceJson, &orgusrslist); err != nil {
											panic(err)
										}

										OrgUsrdetailsLen := len(orgusrslist.Resources)

										if OrgUsrdetailsLen != 0 {

										} else {

											fmt.Println("+ ", Orgs.Org.OrgUsers.LDAP.OrgAuditors[j], ",", "LDAP OrgAuditor")
											cmd := exec.Command("cf", "set-org-role", Orgs.Org.OrgUsers.LDAP.OrgAuditors[j], Orgs.Org.Name, "OrgAuditor")

											if _, err := cmd.Output(); err != nil {
												fmt.Println("command: ", cmd)
												fmt.Println("Err: ", cmd.Stdout, err)
											} else {
												//fmt.Println("command: ", cmd)
												fmt.Println(cmd.Stdout)
											}
										}

									} else {
										fmt.Println(Orgs.Org.OrgUsers.LDAP.OrgAuditors[j], "LDAP OrgAuditor User does't exist in foundation, please ask user to login to apps manager")
									}
								}

								UAAOrgAudLen := len(Orgs.Org.OrgUsers.UAA.OrgAuditors)
								for j := 0; j < UAAOrgAudLen; j++ {

									//fmt.Println(Orgs.Org.Name)
									//fmt.Println(Orgs.Org.OrgUsers.UAA.OrgAuditors[j])
									var getspace *exec.Cmd
									if ostype == "windows" {
										path := "\""+"/v3/users/?usernames=" + Orgs.Org.OrgUsers.UAA.OrgAuditors[j] + "&origins=uaa"+"\""
										getspace = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_userguidfind.json")

									} else {
										path := "/v3/users/?usernames=" + Orgs.Org.OrgUsers.UAA.OrgAuditors[j] + "&origins=uaa"
										getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_userguidfind.json")
									}

									err := getspace.Run()
									if err == nil {
										//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
									} else {
										fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
									}

									fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_userguidfind.json")
									if err != nil {
										fmt.Println(err)
									}
									var usedetails UserDetailsJson
									if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
										panic(err)
									}

									if len(usedetails.Resources) != 0 {

										var orgusersdetailslist *exec.Cmd
										userguid := usedetails.Resources[0].GUID
										if ostype == "windows" {
											path := "\""+"/v3/roles/?organization_guids=" + orgguid + "&user_guids=" + userguid + "&types=organization_auditor"+"\""
											orgusersdetailslist = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										} else {
											path := "/v3/roles/?organization_guids=" + orgguid + "&user_guids=" + userguid + "&types=organization_auditor"
											orgusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrsroledets.json")

										}

										err := orgusersdetailslist.Run()
										if err == nil {
											//fmt.Println(orgusersdetailslist, orgusersdetailslist.Stdout, orgusersdetailslist.Stderr)
										} else {
											fmt.Println("err", orgusersdetailslist, orgusersdetailslist.Stdout, orgusersdetailslist.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										if err != nil {
											fmt.Println(err)
										}
										var orgusrslist OrgUsersListJson
										if err := json.Unmarshal(fileSpaceJson, &orgusrslist); err != nil {
											panic(err)
										}

										OrgUsrdetailsLen := len(orgusrslist.Resources)

										if OrgUsrdetailsLen != 0 {

										} else {

											fmt.Println("+ ", Orgs.Org.OrgUsers.UAA.OrgAuditors[j], ",", "UAA OrgAuditor")
											cmd := exec.Command("cf", "set-org-role", Orgs.Org.OrgUsers.UAA.OrgAuditors[j], Orgs.Org.Name, "OrgAuditor")
											if _, err := cmd.Output(); err != nil {
												fmt.Println("command: ", cmd)
												fmt.Println("Err: ", cmd.Stdout, err)
											} else {
												//fmt.Println("command: ", cmd)
												fmt.Println(cmd.Stdout)
											}
										}

									} else {
										fmt.Println(Orgs.Org.OrgUsers.UAA.OrgAuditors[j], "UAA OrgAuditor User does't exist in foundation, please ask user to login to apps manager")
									}
								}

								SSOOrgAudLen := len(Orgs.Org.OrgUsers.SSO.OrgAuditors)
								SSOProvider := InitClusterConfigVals.ClusterDetails.SSOProvider
								for j := 0; j < SSOOrgAudLen; j++ {

									var getspace *exec.Cmd

									if ostype == "windows" {
										path := "\""+"/v3/users/?usernames=" + Orgs.Org.OrgUsers.SSO.OrgAuditors[j] + "&origins="+SSOProvider+"\""
										getspace = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_userguidfind.json")
									} else {
										path := "/v3/users/?usernames=" + Orgs.Org.OrgUsers.SSO.OrgAuditors[j] + "&origins="+SSOProvider
										getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_userguidfind.json")
									}

									err := getspace.Run()
									if err == nil {
										//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
									} else {
										fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
									}

									fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_userguidfind.json")
									if err != nil {
										fmt.Println(err)
									}
									var usedetails UserDetailsJson
									if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
										panic(err)
									}

									if len(usedetails.Resources) != 0 {

										var orgusersdetailslist *exec.Cmd
										userguid := usedetails.Resources[0].GUID
										if ostype == "windows" {
											path := "\""+"/v3/roles/?organization_guids=" + orgguid + "&user_guids=" + userguid + "&types=organization_auditor"+"\""
											orgusersdetailslist = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										} else {
											path := "/v3/roles/?organization_guids=" + orgguid + "&user_guids=" + userguid + "&types=organization_auditor"
											orgusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										}
										err := orgusersdetailslist.Run()
										if err == nil {
											//fmt.Println(orgusersdetailslist, orgusersdetailslist.Stdout, orgusersdetailslist.Stderr)
										} else {
											fmt.Println("err", orgusersdetailslist, orgusersdetailslist.Stdout, orgusersdetailslist.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateOrgsUsers_orgusrsroledets.json")
										if err != nil {
											fmt.Println(err)
										}

										var orgusrslist OrgUsersListJson
										if err := json.Unmarshal(fileSpaceJson, &orgusrslist); err != nil {
											panic(err)
										}

										OrgUsrdetailsLen := len(orgusrslist.Resources)

										if OrgUsrdetailsLen != 0 {

										} else {

											fmt.Println("+ ", Orgs.Org.OrgUsers.SSO.OrgAuditors[j], ",", "SSO OrgAuditor")
											cmd := exec.Command("cf", "set-org-role", Orgs.Org.OrgUsers.SSO.OrgAuditors[j], Orgs.Org.Name, "OrgAuditor")
											if _, err := cmd.Output(); err != nil {
												fmt.Println("command: ", cmd)
												fmt.Println("Err: ", cmd.Stdout, err)
											} else {
												//fmt.Println("command: ", cmd)
												fmt.Println(cmd.Stdout)
											}
										}
									} else {
										fmt.Println(Orgs.Org.OrgUsers.SSO.OrgAuditors[j],"SSO OrgAuditor User does't exist in foundation, please ask user to login to apps manager")
									}
								}

							} else {
								fmt.Println("Set OrgAuditor:", InitClusterConfigVals.ClusterDetails.SetOrgAuditor)
								fmt.Println("set SetOrgAuditor field to true to manage OrgAuditors for the Org")
							}
						}
					}
					results := exec.Command("cf", "org-users", Orgs.Org.Name)
					if _, err := results.Output(); err != nil{
						fmt.Println("command: ", results)
						fmt.Println("Err: ", results.Stdout, err)
					} else {
						//fmt.Println("command: ", results)
						fmt.Println(results.Stdout)
					}
				}
			} else {
				fmt.Println("Org Name does't match with folder name")
			}
		} else {
			fmt.Println("This is a protected Org")
		}
	}
	return err
}
