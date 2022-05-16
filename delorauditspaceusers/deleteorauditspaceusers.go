package delorauditspaceusers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
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

func DeleteOrAuditSpaceUsers(clustername string, cpath string, ostype string) error {

	var list List
	var ProtectedOrgs ProtectedList
	var gitlist GitList
	var InitClusterConfigVals InitClusterConfigVals
	var ListYml string
	var LenList int
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
	MasterUserAuditor := InitClusterConfigVals.ClusterDetails.EnableUserAudit
	for z := 0; z < LenList; z++ {

		var count, totalcount int
		var OrgName string
		if 	InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
			OrgName = list.OrgList[z].Name
		} else {
			OrgName = gitlist.OrgList[z].Name
			//RepoName = gitlist.OrgList[i].Name
		}
		fmt.Println(" ")
		fmt.Println("Org: ", OrgName)
		for p := 0; p < LenProtectedOrgs; p++ {
			if ProtectedOrgs.Org[p] == OrgName {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {

			var OrgsYml string
			if ostype == "windows" {
				OrgsYml = cpath+"\\"+clustername+"\\"+OrgName+"\\Org.yml"
			} else {
				OrgsYml = cpath+"/"+clustername+"/"+OrgName+"/Org.yml"
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

			Audit := strings.ToLower(Orgs.UserAudit)
			if Audit == "" {
				Audit = "list"
			}

			fmt.Println("Audit: ", Audit)
			if OrgName == Orgs.Org.Name {

				guid := exec.Command("cf", "org", Orgs.Org.Name, "--guid")

				if _, err := guid.Output(); err == nil {

					fmt.Println("command: ", guid)
					//fmt.Println("Org exists: ", guid.Stdout)
					targetOrg := exec.Command("cf", "t", "-o", Orgs.Org.Name)
					if _, err := targetOrg.Output(); err == nil {
						fmt.Println("command: ", targetOrg)
						fmt.Println("Targeted Org: ", targetOrg.Stdout)
					} else {
						fmt.Println("command: ", targetOrg)
						fmt.Println("Err: ", targetOrg.Stdout, targetOrg.Stderr)
					}

					SpaceLen := len(Orgs.Org.Spaces)

					for j := 0; j < SpaceLen; j++ {

						var outguid bytes.Buffer
						guid = exec.Command("cf", "space", Orgs.Org.Spaces[j].Name, "--guid")
						guid.Stdout = &outguid
						err := guid.Run()

						if err == nil {

							fmt.Println("Space exists: ", Orgs.Org.Spaces[j].Name,",", outguid.String())

							target := exec.Command("cf", "t", "-o", Orgs.Org.Name,  "-s", Orgs.Org.Spaces[j].Name)
							if _, err := target.Output(); err == nil {

								fmt.Println("command: ", target)
								fmt.Println(target.Stdout)
								//var out bytes.Buffer

								if InitClusterConfigVals.ClusterDetails.SetSpaceAuditor == true {
									//fmt.Println(Audit)
									var spaceaudituserslist *exec.Cmd

									if ostype == "windows" {
										path := "\""+"/v3/roles/?types=space_auditor"+"&space_guids="+outguid.String()+"\""
										spaceaudituserslist = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "DeleteOrAuditSpaceUsers_spaceauditusrslist.json")
									} else {
										path := "/v3/roles/?types=space_auditor"+"&space_guids="+outguid.String()
										spaceaudituserslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "DeleteOrAuditSpaceUsers_spaceauditusrslist.json")
									}

									err := spaceaudituserslist.Run()
									if err == nil {
										//	fmt.Println(spaceuserslist, spaceuserslist.Stdout)
									} else {
										fmt.Println("err", spaceaudituserslist, spaceaudituserslist.Stdout, spaceaudituserslist.Stderr)
									}

									var spaceauditusrslist SpaceUsersListJson
									fileSpaceJson, err := ioutil.ReadFile("DeleteOrAuditSpaceUsers_spaceauditusrslist.json")
									if err != nil {
										fmt.Println(err)
									}
									if err := json.Unmarshal(fileSpaceJson, &spaceauditusrslist); err != nil {
										panic(err)
									}

									SpaceAuditUsrLen := len(spaceauditusrslist.Resources)

									//fmt.Println("Number of Space Audit Users currently exist in Space",Orgs.Org.Spaces[j].Name,":", SpaceAuditUsrLen)

									if SpaceAuditUsrLen != 0 {

										for i := 0; i < SpaceAuditUsrLen; i++ {
											//
											SpaceUsLenSSOAuditor := len(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceAuditors)
											SpaceUsLenUAAAuditor := len(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors)
											SpaceUsLenLDAPAuditor := len(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors)
											var spaceusrssoauditscount, spaceusrssoaudittotalcount int
											var spaceusruaaauditscount, spaceusruaaaudittotalcount int
											var spaceusrldapauditscount, spaceusrldapaudittotalcount int
											userguid := spaceauditusrslist.Resources[i].Relationships.User.Data.GUID
											var userdetails *exec.Cmd
											//var out1 bytes.Buffer
											if ostype == "windows" {
												path := "\""+"/v3/users/?guids=" + userguid+"\""
												userdetails = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "DeleteOrAuditSpaceUsers_usrdetails.json")
											} else {
												path := "/v3/users/?guids=" + userguid
												userdetails = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "DeleteOrAuditSpaceUsers_usrdetails.json")

											}
											//userdetails.Stdout = &out1
											err := userdetails.Run()

											fileusrdetlsjson, err := ioutil.ReadFile("DeleteOrAuditSpaceUsers_usrdetails.json")
											if err != nil {
												fmt.Println(err)
											}

											var usedetails UserDetailsJson
											err = json.Unmarshal([]byte(fileusrdetlsjson), &usedetails)
											if err != nil {
												panic(err)
											} else {
												username := usedetails.Resources[0].Username
												origin := usedetails.Resources[0].Origin
												SSOProvider := InitClusterConfigVals.ClusterDetails.SSOProvider
												if origin == SSOProvider {
													if err == nil {
														for q := 0; q < SpaceUsLenSSOAuditor; q++ {
															if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceAuditors[q]) == username {
																spaceusrssoauditscount = 1
															} else {
																spaceusrssoauditscount = 0
															}
															spaceusrssoaudittotalcount = spaceusrssoaudittotalcount + spaceusrssoauditscount
														}
														if spaceusrssoaudittotalcount == 0 {
															//fmt.Println("SSO SpaceAuditor", username, "has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
															if Audit == "unset" {
																//	fmt.Println("UNSET!UNSET!")
																//	fmt.Println("Unsetting user")
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("SSO SpaceAuditor", username, "has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("Unsetting user")
																	if MasterUserAuditor == true {
																		unset := exec.Command("cf", "unset-space-role", username, Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceAuditor")
																		if _, err := unset.Output(); err == nil {
																			fmt.Println("command: ", unset)
																			fmt.Println(unset.Stdout)
																		} else {
																			fmt.Println("command: ", unset)
																			fmt.Println("Err: ", unset.Stdout, unset.Stderr)
																		}
																	} else {
																		fmt.Println("UserAudit is not enabled")
																	}

																}
															} else if Audit == "list" {
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("User to be deleted: ", username, "SSO SpaceAuditor")
																}
															} else {
																fmt.Println("Provide Valid Input")
															}
														} else {
															//fmt.Println("User is listed in Org.yml as SSO Space Audit User: ", username)
														}
													}
												}
												if origin == "uaa" {
													if err == nil {
														for q := 0; q < SpaceUsLenUAAAuditor; q++ {

															//fmt.Println("UAA Space Audit Usr: ", strings.ToLower(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors[q]),",", username)
															if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors[q]) == username {
																spaceusruaaauditscount = 1
															} else {
																spaceusruaaauditscount = 0
															}
															spaceusruaaaudittotalcount = spaceusruaaaudittotalcount + spaceusruaaauditscount
														}
														if spaceusruaaaudittotalcount == 0 {
															//fmt.Println("UAA SpaceAuditor", username,"has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)

															if Audit == "unset" {
																//fmt.Println("UNSET!UNSET!")
																//fmt.Println("Unsetting user")
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("UAA SpaceAuditor", username, "has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("Unsetting user")
																	if MasterUserAuditor == true {
																		unset := exec.Command("cf", "unset-space-role", username, Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceAuditor")
																		if _, err := unset.Output(); err == nil {
																			fmt.Println("command: ", unset)
																			fmt.Println(unset.Stdout)
																		} else {
																			fmt.Println("command: ", unset)
																			fmt.Println("Err: ", unset.Stdout, unset.Stderr)
																		}
																	} else {
																		fmt.Println("UserAudit is not enabled")
																	}

																}
															} else if Audit == "list" {
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("User to be deleted: ", username, "UAA SpaceAuditor")
																}
															} else {
																fmt.Println("Provide Valid Input")
															}
														} else {
															//fmt.Println("User is listed in Org.yml as UAA Space Audit User: ", username)
														}
													}
												}
												if origin == "ldap" {
													if err == nil {
														for q := 0; q < SpaceUsLenLDAPAuditor; q++ {
															//fmt.Println("LDAP Space Audit Usr: ", strings.ToLower(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors[q]),",", username)
															if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors[q]) == username {
																spaceusrldapauditscount = 1
															} else {
																spaceusrldapauditscount = 0
															}
															spaceusrldapaudittotalcount = spaceusrldapaudittotalcount + spaceusrldapauditscount
														}

														if spaceusrldapaudittotalcount == 0 {
															//	fmt.Println("LDAP SpaceAuditor", username, "has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
															if Audit == "unset" {
																//		fmt.Println("UNSET!UNSET!")
																//		fmt.Println("Unsetting user")
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("LDAP SpaceAuditor", username, "has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("Unsetting user")
																	if MasterUserAuditor == true {
																		unset := exec.Command("cf", "unset-space-role", username, Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceAuditor")
																		if _, err := unset.Output(); err == nil {
																			fmt.Println("command: ", unset)
																			fmt.Println(unset.Stdout)
																		} else {
																			fmt.Println("command: ", unset)
																			fmt.Println("Err: ", unset.Stdout, unset.Stderr)
																		}
																	} else {
																		fmt.Println("UserAudit is not enabled")
																	}

																}
															} else if Audit == "list" {
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("User to be deleted: ", username, "LDAP SpaceAuditor")
																}
															} else {
																fmt.Println("Provide Valid Input")
															}
														} else {
															//fmt.Println("User is listed in Org.yml as LDAP Space Audit User: ", username)
														}
													}
												}
											}
										}
									} else {
										fmt.Println("No space Audit users exist")
									}
								} else {
									fmt.Println("Set SpaceAuditor:", InitClusterConfigVals.ClusterDetails.SetSpaceAuditor)
									fmt.Println("set SetSpaceAuditor field to true to manage SpaceAuditor for the Org")
								}

								if InitClusterConfigVals.ClusterDetails.SetSpaceDeveloper == true {
									var spacedevuserslist *exec.Cmd
									if ostype == "windows" {
										path := "\""+"/v3/roles/?types=space_developer"+"&space_guids="+outguid.String()+"\""
										spacedevuserslist = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "DeleteOrAuditSpaceUsers_spacedevusrslist.json")

									} else {
										path := "/v3/roles/?types=space_developer"+"&space_guids="+outguid.String()
										spacedevuserslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "DeleteOrAuditSpaceUsers_spacedevusrslist.json")

									}
									err := spacedevuserslist.Run()
									if err == nil {
										//	fmt.Println(spaceuserslist, spaceuserslist.Stdout)
									} else {
										fmt.Println("err", spacedevuserslist, spacedevuserslist.Stdout, spacedevuserslist.Stderr)
									}

									var spacedevusrslist SpaceUsersListJson
									fileSpaceJson, err := ioutil.ReadFile("DeleteOrAuditSpaceUsers_spacedevusrslist.json")
									if err != nil {
										fmt.Println(err)
									}
									if err := json.Unmarshal(fileSpaceJson, &spacedevusrslist); err != nil {
										panic(err)
									}

									SpaceDevUsrLen := len(spacedevusrslist.Resources)

									//fmt.Println("Number of Space Developer Users currently exist in Space",Orgs.Org.Spaces[j].Name,":", SpaceDevUsrLen)

									if SpaceDevUsrLen != 0 {

										for i := 0; i < SpaceDevUsrLen; i++ {
											//
											SpaceUsLenSSODeveloper := len(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers)
											SpaceUsLenUAADeveloper := len(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers)
											SpaceUsLenLDAPDeveloper := len(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers)
											var spaceusrldapdevscount, spaceusrldapdevtotalcount int
											var spaceusrssodevscount, spaceusrssodevtotalcount int
											var spaceusruaadevscount, spaceusruaadevtotalcount int

											userguid := spacedevusrslist.Resources[i].Relationships.User.Data.GUID
											//var out bytes.Buffer

											var userdetails *exec.Cmd
											if ostype == "windows" {
												path := "\""+"/v3/users/?guids=" + userguid+"\""
												userdetails = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "DeleteOrAuditSpaceUsers_usrdetails.json")
											} else {
												path := "/v3/users/?guids=" + userguid
												userdetails = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "DeleteOrAuditSpaceUsers_usrdetails.json")
											}
											//userdetails.Stdout = &out
											err := userdetails.Run()

											fileusrdetlsjson, err := ioutil.ReadFile("DeleteOrAuditSpaceUsers_usrdetails.json")
											if err != nil {
												fmt.Println(err)
											}
											//var Orgs Orglist
											//var spaceusrslist SpaceUsersListJson
											var usedetails UserDetailsJson
											err = json.Unmarshal([]byte(fileusrdetlsjson), &usedetails)
											if err != nil {
												panic(err)
											} else {
												username := usedetails.Resources[0].Username
												origin := usedetails.Resources[0].Origin
												SSOProvider := InitClusterConfigVals.ClusterDetails.SSOProvider
												if origin == SSOProvider {
													if err == nil {
														for q := 0; q < SpaceUsLenSSODeveloper; q++ {

															//fmt.Println("SSO Space Dev Usr: ", strings.ToLower(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers[q]), ",", username)
															if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers[q]) == username {
																spaceusrssodevscount = 1
															} else {
																spaceusrssodevscount = 0
															}
															spaceusrssodevtotalcount = spaceusrssodevtotalcount + spaceusrssodevscount
														}

														if spaceusrssodevtotalcount == 0 {
															//fmt.Println("SSO Space Dev", username,"has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
															if Audit == "unset" {
																//fmt.Println("UNSET!UNSET!")
																//fmt.Println("Unsetting user")
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("SSO Space Dev", username, "has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("Unsetting user")
																	if MasterUserAuditor == true {
																		unset := exec.Command("cf", "unset-space-role", username, Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceDeveloper")
																		if _, err := unset.Output(); err == nil {
																			fmt.Println("command: ", unset)
																			fmt.Println(unset.Stdout)
																		} else {
																			fmt.Println("command: ", unset)
																			fmt.Println("Err: ", unset.Stdout, unset.Stderr)
																		}
																	} else {
																		fmt.Println("UserAudit is not enabled")
																	}
																}
															} else if Audit == "list" {
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("User to be deleted: ", username, "SSO SpaceDeveloper")
																}
															} else {
																fmt.Println("Provide Valid Input")
															}
														} else {
															//fmt.Println("User is listed in Org.yml as SSO Space Dev User: ", username)
														}
													}
												}
												if origin == "uaa" {
													if err == nil {
														for q := 0; q < SpaceUsLenUAADeveloper; q++ {

															//fmt.Println("UAA Space Dev Usr: ", strings.ToLower(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers[q]),",", username)
															if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers[q]) == username {
																spaceusruaadevscount = 1
															} else {
																spaceusruaadevscount = 0
															}
															spaceusruaadevtotalcount = spaceusruaadevtotalcount + spaceusruaadevscount
														}

														if spaceusruaadevtotalcount == 0 {
															//fmt.Println("UAA SpaceDeveloper User", username, "has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
															if Audit == "unset" {
																//	fmt.Println("UNSET!UNSET!")
																//	fmt.Println("Unsetting user")
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("UAA SpaceDeveloper User", username, "has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("Unsetting user")
																	if MasterUserAuditor == true {
																		unset := exec.Command("cf", "unset-space-role", username, Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceDeveloper")
																		if _, err := unset.Output(); err == nil {
																			fmt.Println("command: ", unset)
																			fmt.Println(unset.Stdout)
																		} else {
																			fmt.Println("command: ", unset)
																			fmt.Println("Err: ", unset.Stdout, unset.Stderr)
																		}
																	} else {
																		fmt.Println("UserAudit is not enabled")
																	}

																}
															} else if Audit == "list" {
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("User to be deleted: ", username, "UAA SpaceDeveloper")
																}
															} else {
																fmt.Println("Provide Valid Input")
															}
														} else {
															//fmt.Println("User is listed in Org.yml as UAA Space Dev User: ", username)
														}
													}
												}
												if origin == "ldap" {
													if err == nil {
														for q := 0; q < SpaceUsLenLDAPDeveloper; q++ {

															//fmt.Println("LDAP Space Dev Usr: ", strings.ToLower(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers[q]), ",", username)
															if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers[q]) == username {
																spaceusrldapdevscount = 1
															} else {
																spaceusrldapdevscount = 0
															}
															spaceusrldapdevtotalcount = spaceusrldapdevtotalcount + spaceusrldapdevscount
														}

														if spaceusrldapdevtotalcount == 0 {
															//fmt.Println("LDAP SpaceDeveloper User", username, "has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
															if Audit == "unset" {
																//	fmt.Println("UNSET!UNSET!")
																//	fmt.Println("Unsetting user")
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("LDAP SpaceDeveloper User", username, "has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("Unsetting user")
																	if MasterUserAuditor == true {
																		unset := exec.Command("cf", "unset-space-role", username, Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceDeveloper")
																		if _, err := unset.Output(); err == nil {
																			fmt.Println("command: ", unset)
																			fmt.Println(unset.Stdout)
																		} else {
																			fmt.Println("command: ", unset)
																			fmt.Println("Err: ", unset.Stdout, unset.Stderr)
																		}
																	} else {
																		fmt.Println("UserAudit is not enabled")
																	}

																}

															} else if Audit == "list" {
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("User to be deleted: ", username, "LDAP SpaceDeveloper")
																}
															} else {
																fmt.Println("Provide Valid Input")
															}
														} else {
															//fmt.Println("User is listed in Org.yml as LDAP Space Dev User: ", username)
														}
													}
												}
											}
											//
										}
									} else {
										fmt.Println("No space Developer users exist")
									}
								} else {
									fmt.Println("Set SpaceDeveloper:", InitClusterConfigVals.ClusterDetails.SetSpaceDeveloper)
									fmt.Println("set SetSpaceDeveloper field to true to manage SpaceDeveloper for the Org")
								}

								if InitClusterConfigVals.ClusterDetails.SetSpaceManager == true{
									var spacemanuserslist *exec.Cmd
									if ostype == "windows" {
										path := "\""+"/v3/roles/?types=space_manager"+"&space_guids="+outguid.String()+"\""
										spacemanuserslist = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "DeleteOrAuditSpaceUsers_spacemanusrslist.json")
									} else {
										path := "/v3/roles/?types=space_manager"+"&space_guids="+outguid.String()
										spacemanuserslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "DeleteOrAuditSpaceUsers_spacemanusrslist.json")
									}
									err := spacemanuserslist.Run()
									if err == nil {
										//	fmt.Println(spaceuserslist, spaceuserslist.Stdout)
									} else {
										fmt.Println("err", spacemanuserslist, spacemanuserslist.Stdout, spacemanuserslist.Stderr)
									}

									var spacemanusrslist SpaceUsersListJson
									fileSpaceJson, err := ioutil.ReadFile("DeleteOrAuditSpaceUsers_spacemanusrslist.json")
									if err != nil {
										fmt.Println(err)
									}
									if err := json.Unmarshal(fileSpaceJson, &spacemanusrslist); err != nil {
										panic(err)
									}

									SpaceManUsrLen := len(spacemanusrslist.Resources)

									//fmt.Println("Number of Space Developer Users currently exist in Space",Orgs.Org.Spaces[j].Name,":", SpaceManUsrLen)

									if SpaceManUsrLen != 0 {

										for i := 0; i < SpaceManUsrLen; i++ {

											SpaceUsLenSSOMan := len(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers)
											SpaceUsLenUAAMan := len(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers)
											SpaceUsLenLDAPMan := len(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers)
											var spaceusruaamanscount, spaceusruaamantotalcount int
											var spaceusrssomanscount, spaceusrssomantotalcount int
											var spaceusrldapmanscount, spaceusrldapmantotalcount int
											userguid := spacemanusrslist.Resources[i].Relationships.User.Data.GUID
											//var out bytes.Buffer
											var userdetails *exec.Cmd
											if ostype == "windows" {
												path := "\""+"/v3/users/?guids=" + userguid+"\""
												userdetails = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "DeleteOrAuditSpaceUsers_usrdetails.json")

											} else {
												path := "/v3/users/?guids=" + userguid
												userdetails = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "DeleteOrAuditSpaceUsers_usrdetails.json")

											}
											//userdetails.Stdout = &out
											err := userdetails.Run()

											fileusrdetlsjson, err := ioutil.ReadFile("DeleteOrAuditSpaceUsers_usrdetails.json")
											if err != nil {
												fmt.Println(err)
											}
											//var Orgs Orglist
											//var spaceusrslist SpaceUsersListJson
											var usedetails UserDetailsJson
											err = json.Unmarshal([]byte(fileusrdetlsjson), &usedetails)
											if err != nil {
												panic(err)
											} else {
												username := usedetails.Resources[0].Username
												origin := usedetails.Resources[0].Origin
												SSOProvider := InitClusterConfigVals.ClusterDetails.SSOProvider
												if origin == SSOProvider {
													if err == nil {
														for q := 0; q < SpaceUsLenSSOMan; q++ {

															//fmt.Println("SSO Space Manager Usr: ", strings.ToLower(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers[q]), ",", username)
															if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers[q]) == username {
																spaceusrssomanscount = 1
															} else {
																spaceusrssomanscount = 0
															}
															spaceusrssomantotalcount = spaceusrssomantotalcount + spaceusrssomanscount
														}

														if spaceusrssomantotalcount == 0 {
															//fmt.Println("SSO Space Manager", username, "has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)

															if Audit == "unset" {
																//	fmt.Println("UNSET!UNSET!")
																//	fmt.Println("Unsetting user")
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("SSO Space Manager", username, "has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("Unsetting user")
																	if MasterUserAuditor == true {
																		unset := exec.Command("cf", "unset-space-role", username, Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceManager")
																		if _, err := unset.Output(); err == nil {
																			fmt.Println("command: ", unset)
																			fmt.Println(unset.Stdout)
																		} else {
																			fmt.Println("command: ", unset)
																			fmt.Println("Err: ", unset.Stdout, unset.Stderr)
																		}
																	} else {
																		fmt.Println("UserAudit is not enabled")
																	}

																}
															} else if Audit == "list" {
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("User to be deleted: ", username, "SSO SpaceManager")
																}
															} else {
																fmt.Println("Provide Valid Input")
															}
														} else {
															//fmt.Println("User is listed in Org.yml as SSO Space Manager User: ", username)
														}
													}
												}
												if origin == "uaa" {
													if err == nil {
														for q := 0; q < SpaceUsLenUAAMan; q++ {

															//fmt.Println("UAA Space Dev Usr: ", strings.ToLower(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers[q]),",", username)
															if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers[q]) == username {
																spaceusruaamanscount = 1
															} else {
																spaceusruaamanscount = 0
															}
															spaceusruaamantotalcount = spaceusruaamantotalcount + spaceusruaamanscount
														}

														if spaceusruaamantotalcount == 0 {
															//fmt.Println("UAA SpaceManager", username,"has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
															if Audit == "unset" {
																//	fmt.Println("UNSET!UNSET!")
																//	fmt.Println("Unsetting user")
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("UAA SpaceManager", username, "has not be listed in Org.yml for Org/Space", Orgs.Org.Name, Orgs.Org.Spaces[j].Name)
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("Unsetting user")
																	if MasterUserAuditor == true {
																		unset := exec.Command("cf", "unset-space-role", username, Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceManager")
																		if _, err := unset.Output(); err == nil {
																			fmt.Println("command: ", unset)
																			fmt.Println(unset.Stdout)
																		} else {
																			fmt.Println("command: ", unset)
																			fmt.Println("Err: ", unset.Stdout, unset.Stderr)
																		}
																	} else {
																		fmt.Println("UserAudit is not enabled")
																	}

																}
															} else if Audit == "list" {
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("User to be deleted: ", username, "UAA SpaceManager")
																}
															} else {
																fmt.Println("Provide Valid Input")
															}
														} else {
															//fmt.Println("User is listed in Org.yml as UAA Space Manager User: ", username)
														}
													}
												}
												if origin == "ldap" {
													if err == nil {
														for q := 0; q < SpaceUsLenLDAPMan; q++ {

															//fmt.Println("LDAP Space Manager Usr: ", strings.ToLower(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers[q]), ",", username)
															if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers[q]) == username {
																spaceusrldapmanscount = 1
															} else {
																spaceusrldapmanscount = 0
															}
															spaceusrldapmantotalcount = spaceusrldapmantotalcount + spaceusrldapmanscount
														}

														if spaceusrldapmantotalcount == 0 {
															//fmt.Println("LDAP SpaceManager", username,"has not be listed in Org.yml for Org/Space\", Orgs.Org.Name, Orgs.Org.Spaces[j].Name")
															if Audit == "unset" {
																//	fmt.Println("UNSET!UNSET!")
																//	fmt.Println("Unsetting user")
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("LDAP SpaceManager", username, "has not be listed in Org.yml for Org/Space\", Orgs.Org.Name, Orgs.Org.Spaces[j].Name")
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("Unsetting user")
																	if MasterUserAuditor == true {
																		unset := exec.Command("cf", "unset-space-role", username, Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceManager")
																		if _, err := unset.Output(); err == nil {
																			fmt.Println("command: ", unset)
																			fmt.Println(unset.Stdout)
																		} else {
																			fmt.Println("command: ", unset)
																			fmt.Println("Err: ", unset.Stdout, unset.Stderr)
																		}
																	} else {
																		fmt.Println("UserAudit is not enabled")
																	}
																}
															} else if Audit == "list" {
																if strings.TrimSpace(username) == "admin" {
																	//fmt.Println("Skipping unset for admin user")
																} else {
																	fmt.Println("UNSET!UNSET!")
																	fmt.Println("User to be deleted: ", username, "LDAP SpaceManager")
																}
															} else {
																fmt.Println("Provide Valid Input")
															}
														} else {
															//fmt.Println("User is listed in Org.yml as LDAP Space Manager User: ", username)
														}
													}
												}
											}
										}

									} else {
										fmt.Println("No space Manager users exist")
									}

								}else {
									fmt.Println("Set SpaceManager:", InitClusterConfigVals.ClusterDetails.SetSpaceManager)
									fmt.Println("set SetSpaceManager field to true to manage SpaceManager for the Org")
								}


							} else {
								fmt.Println("command: ", target)
								fmt.Println("Err: ", target.Stdout, target.Stderr)
							}
						} else {
							fmt.Println("command: ",guid)
							fmt.Println("Err: ", guid.Stdout, err)
							fmt.Println("Space doesn't exists")
						}
					}
				} else {
					fmt.Println("command: ", guid)
					fmt.Println("Err: ", guid.Stdout, err)
					fmt.Println("Org doesn't exists")
				}
			} else {
				fmt.Println("Org Name does't match with folder name")
			}
		}  else {
			fmt.Println("This is a protected Org")
		}
	}
	return err
}
