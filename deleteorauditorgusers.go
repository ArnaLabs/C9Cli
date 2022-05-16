package delorauditorgusers

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

func DeleteorAuditOrgUsers(clustername string, cpath string, ostype string) error {

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
	// Org List
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


	//Config File

	var OrgsYml string

	//Protected Orgs
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
	MasterUserAudit := InitClusterConfigVals.ClusterDetails.EnableUserAudit
	for i := 0; i < LenList; i++ {
		var OrgName string
		if 	InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
			OrgName = list.OrgList[i].Name
		} else {
			OrgName = gitlist.OrgList[i].Name
			//RepoName = gitlist.OrgList[i].Name
		}
		var count, totalcount int
		fmt.Println(" ")
		fmt.Println("Org: ", OrgName)
		for p := 0; p < LenProtectedOrgs; p++ {
			//fmt.Println("Protected Org: ", ProtectedOrgs.Org[p], ",",list.OrgList[i])
			if ProtectedOrgs.Org[p] == OrgName {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {

			//fmt.Println("This is not Protected Org")

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

			if OrgName == Orgs.Org.Name {

				var out bytes.Buffer

				guid := exec.Command("cf", "org", Orgs.Org.Name, "--guid")
				guid.Stdout = &out
				err := guid.Run()

				if err == nil {

					//fmt.Println("Org exists: ", Orgs.Org.Name,",", out.String())
					target := exec.Command("cf", "t", "-o", Orgs.Org.Name)

					if _, err := target.Output(); err == nil {

						//	fmt.Println("command: ", target)
						fmt.Println(target.Stdout)

						//var out bytes.Buffer
						//orguserslist.Stdout = &out
						// OrgMan Users

						if InitClusterConfigVals.ClusterDetails.SetOrgManager == true {
							var orgmanuserslist *exec.Cmd
							if ostype == "windows" {
								path := "\""+"/v3/roles/?types=organization_manager"+"&organization_guids="+out.String()+"\""
								orgmanuserslist = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "DeleteorAuditOrgUsers_orgmanusrslist.json")

							} else {
								path := "/v3/roles/?types=organization_manager"+"&organization_guids="+out.String()
								orgmanuserslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "DeleteorAuditOrgUsers_orgmanusrslist.json")

							}
							err := orgmanuserslist.Run()

							if err == nil {
								//	fmt.Println(orguserslist, orguserslist.Stdout, orguserslist.Stderr)
							} else {
								fmt.Println("err", orgmanuserslist, orgmanuserslist.Stdout, orgmanuserslist.Stderr)
							}

							fileSpaceJson, err := ioutil.ReadFile("DeleteorAuditOrgUsers_orgmanusrslist.json")
							if err != nil {
								fmt.Println(err)
							}

							var orgmanusrslist OrgUsersListJson

							if err := json.Unmarshal(fileSpaceJson, &orgmanusrslist); err != nil {
								panic(err)
							}

							OrgManUsrLen := len(orgmanusrslist.Resources)

							//fmt.Println("Number of OrgManager Users currently exist in Org", Orgs.Org.Name, ":", OrgManUsrLen)

							if OrgManUsrLen != 0 {

								for i := 0; i < OrgManUsrLen; i++ {

									OrgUsLenSSOManagers := len(Orgs.Org.OrgUsers.SSO.OrgManagers)
									OrgUsLenUAAManagers := len(Orgs.Org.OrgUsers.UAA.OrgManagers)
									OrgUsLenLDAPManagers := len(Orgs.Org.OrgUsers.LDAP.OrgManagers)

									var orgusrssomangscount, aorusrssomangtotalcount int
									var orgusruaamangscount, aorusruaamangtotalcount int
									var orgusrldapmangscount, aorusrldapmangtotalcount int

									var userdetails *exec.Cmd
									userguid := orgmanusrslist.Resources[i].Relationships.User.Data.GUID
									//var out bytes.Buffer
									if ostype == "windows" {
										path := "\""+"/v3/users/?guids=" + userguid+"\""
										userdetails = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "DeleteorAuditOrgUsers_usrdetails.json")
									} else {
										path := "/v3/users/?guids=" + userguid
										userdetails = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "DeleteorAuditOrgUsers_usrdetails.json")
									}
									//userdetails.Stdout = &out
									err := userdetails.Run()

									fileusrdetlsjson, err := ioutil.ReadFile("DeleteorAuditOrgUsers_usrdetails.json")
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
												for q := 0; q < OrgUsLenSSOManagers; q++ {

													//fmt.Println("SSO Org Managers Usr: ", strings.ToLower(Orgs.Org.OrgUsers.SSO.OrgManagers[q]), ",", username)
													if strings.TrimSpace(Orgs.Org.OrgUsers.SSO.OrgManagers[q]) == username {
														orgusrssomangscount = 1
													} else {
														orgusrssomangscount = 0
													}
													aorusrssomangtotalcount = aorusrssomangtotalcount + orgusrssomangscount
												}

												if aorusrssomangtotalcount == 0 {

													//fmt.Println("SSO OrgManager User", username,"has not be listed in Org.yml for Org", Orgs.Org.Name)

													if Audit == "unset" {
														if strings.TrimSpace(username) == "admin" {
															//fmt.Println("Skipping unset for admin user")
														} else {
															fmt.Println("SSO OrgManager User", username, "has not be listed in Org.yml for Org", Orgs.Org.Name)
															fmt.Println("UNSET!UNSET!")
															fmt.Println("Unsetting user")
															if MasterUserAudit == true {
																unset := exec.Command("cf", "unset-org-role", username, Orgs.Org.Name, "OrgManager")
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
															fmt.Println("User to be deleted: ", username, "SSO OrgManager")
														}
													} else {
														fmt.Println("Provide Valid Input")
													}
												} else {
													//fmt.Println("User is listed in Org.yml as SSO Org Manager User: ", username)
												}
											}
										} else if origin == "uaa" {
											if err == nil {
												for q := 0; q < OrgUsLenUAAManagers; q++ {

													//		fmt.Println("UAA Org Managers Usr: ", strings.ToLower(Orgs.Org.OrgUsers.UAA.OrgManagers[q]), ",", username)
													if strings.TrimSpace(Orgs.Org.OrgUsers.UAA.OrgManagers[q]) == username {
														orgusruaamangscount = 1
													} else {
														orgusruaamangscount = 0
													}
													aorusruaamangtotalcount = aorusruaamangtotalcount + orgusruaamangscount
												}
												if aorusruaamangtotalcount == 0 {

													//fmt.Println("UAA OrgManager User",username,"has not be listed in Org.yml for Org", Orgs.Org.Name)

													if Audit == "unset" {
														if strings.TrimSpace(username) == "admin" {
															//fmt.Println("Skipping unset for admin user")
														} else {
															fmt.Println("UAA OrgManager User", username, "has not be listed in Org.yml for Org", Orgs.Org.Name)
															fmt.Println("UNSET!UNSET!")
															fmt.Println("Unsetting user")
															if MasterUserAudit ==  true {
																unset := exec.Command("cf", "unset-org-role", username, Orgs.Org.Name, "OrgManager")
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
															fmt.Println("User to be deleted: ", username, "UAA OrgManager")
														}
													} else {
														fmt.Println("Provide Valid Input")
													}
												} else {
													//fmt.Println("User is listed in Org.yml as UAA Org Manager User: ", username)
												}
											}
										} else if origin == "ldap" {
											if err == nil {
												for q := 0; q < OrgUsLenLDAPManagers; q++ {

													//	fmt.Println("LDAP Org Managers Usr: ", strings.ToLower(Orgs.Org.OrgUsers.LDAP.OrgManagers[q]), ",", username)
													if strings.TrimSpace(Orgs.Org.OrgUsers.LDAP.OrgManagers[q]) == username {
														orgusrldapmangscount = 1
													} else {
														orgusrldapmangscount = 0
													}
													aorusrldapmangtotalcount = aorusrldapmangtotalcount + orgusrldapmangscount
												}
												if aorusrldapmangtotalcount == 0 {
													//fmt.Println("LDAP OrgManager User", username,"has not be listed in Org.yml for Org", Orgs.Org.Name)

													if Audit == "unset" {
														//fmt.Println("UNSET!UNSET!")
														//fmt.Println("Unsetting user")
														if strings.TrimSpace(username) == "admin" {
															//fmt.Println("Skipping unset for admin user")
														} else {
															fmt.Println("LDAP OrgManager User", username, "has not be listed in Org.yml for Org", Orgs.Org.Name)
															fmt.Println("UNSET!UNSET!")
															fmt.Println("Unsetting user")
															if MasterUserAudit ==  true {
																unset := exec.Command("cf", "unset-org-role", username, Orgs.Org.Name, "OrgManager")
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
															fmt.Println("User to be deleted: ", username, "LDAP OrgManager")
														}
													} else {
														fmt.Println("Provide Valid Input")
													}
												} else {
													//fmt.Println("User is listed in Org.yml as LDAP Org Manager User: ", username)
												}
											}
										} else {
											fmt.Println("Authentication of type", origin, "is not monitored")
										}
									}
								}
							} else {
								fmt.Println("No Org Manager users exist")
							}
						} else {
							fmt.Println("Set OrgManager:", InitClusterConfigVals.ClusterDetails.SetOrgManager)
							fmt.Println("set SetOrgManager field to true to manage OrgManagers for the Org")
						}
						if InitClusterConfigVals.ClusterDetails.SetOrgAuditor == true {
							//OrgAudit Users
							var orgaudituserslist *exec.Cmd
							if ostype == "windows" {
								path := "\""+"/v3/roles/?types=organization_auditor"+"&organization_guids="+out.String()+"\""
								orgaudituserslist = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "DeleteorAuditOrgUsers_orgauditusrslist.json")
							} else {
								path := "/v3/roles/?types=organization_auditor"+"&organization_guids="+out.String()
								orgaudituserslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "DeleteorAuditOrgUsers_orgauditusrslist.json")
							}

							err := orgaudituserslist.Run()

							if err == nil {
								//	fmt.Println(orguserslist, orguserslist.Stdout, orguserslist.Stderr)
							} else {
								fmt.Println("err", orgaudituserslist, orgaudituserslist.Stdout, orgaudituserslist.Stderr)
							}

							fileSpaceJson, err := ioutil.ReadFile("DeleteorAuditOrgUsers_orgauditusrslist.json")
							if err != nil {
								fmt.Println(err)
							}

							var orgauditusrslist OrgUsersListJson

							if err := json.Unmarshal(fileSpaceJson, &orgauditusrslist); err != nil {
								panic(err)
							}

							OrgAuditUsrLen := len(orgauditusrslist.Resources)

							//fmt.Println("Number of OrgAuditor Users currently exist in Org", Orgs.Org.Name, ":", OrgAuditUsrLen)

							if OrgAuditUsrLen != 0 {
								for i := 0; i < OrgAuditUsrLen; i++ {

									OrgUsLenSSOAuditor := len(Orgs.Org.OrgUsers.SSO.OrgAuditors)
									OrgUsLenUAAAuditor := len(Orgs.Org.OrgUsers.UAA.OrgAuditors)
									OrgUsLenLDAPAuditor := len(Orgs.Org.OrgUsers.LDAP.OrgAuditors)

									var orgusrssoauditscount, aorusrssoaudittotalcount int
									var orgusruaaauditscount, aorusruaaaudittotalcount int
									var orgusrldapauditscount, aorusrldapaudittotalcount int

									var userdetails *exec.Cmd
									userguid := orgauditusrslist.Resources[i].Relationships.User.Data.GUID
									if ostype == "windows" {
										path := "\""+"/v3/users/?guids=" + userguid+"\""
										userdetails = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "DeleteorAuditOrgUsers_usrdetails.json")
									} else {
										path := "/v3/users/?guids=" + userguid
										userdetails = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "DeleteorAuditOrgUsers_usrdetails.json")
									}
									err := userdetails.Run()
									fileusrdetlsjson, err := ioutil.ReadFile("DeleteorAuditOrgUsers_usrdetails.json")
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
												for q := 0; q < OrgUsLenSSOAuditor; q++ {
													//fmt.Println("SSO Audit Usr: ", strings.ToLower(Orgs.Org.OrgUsers.SSO.OrgAuditors[q]), ",", username)
													if strings.TrimSpace(Orgs.Org.OrgUsers.SSO.OrgAuditors[q]) == username {
														orgusrssoauditscount = 1
													} else {
														orgusrssoauditscount = 0
													}
													aorusrssoaudittotalcount = aorusrssoaudittotalcount + orgusrssoauditscount
												}

												if aorusrssoaudittotalcount == 0 {
													if Audit == "unset" {
														if strings.TrimSpace(username) == "admin" {
															//fmt.Println("Skipping unset for admin user")
														} else {
															fmt.Println("SSO OrgAuditor", username, "has not be listed in Org.yml for Org", Orgs.Org.Name)
															fmt.Println("UNSET!UNSET!")
															fmt.Println("Unsetting user")
															if MasterUserAudit ==  true {
																unset := exec.Command("cf", "unset-org-role", username, Orgs.Org.Name, "OrgAuditor")
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
															fmt.Println("User to be deleted: ", username, "SSO OrgAuditor")
														}
													} else {
														fmt.Println("Provide Valid Input")
													}
												} else {
													//fmt.Println("User is listed in Org.yml as SSO Org Audit User: ", username)
												}
											}
										} else if origin == "uaa" {
											if err == nil {
												for q := 0; q < OrgUsLenUAAAuditor; q++ {

													//fmt.Println("UAA Audit Usr: ", strings.ToLower(Orgs.Org.OrgUsers.UAA.OrgAuditors[q]), ",", username)
													if strings.TrimSpace(Orgs.Org.OrgUsers.UAA.OrgAuditors[q]) == username {
														orgusruaaauditscount = 1
													} else {
														orgusruaaauditscount = 0
													}
													aorusruaaaudittotalcount = aorusruaaaudittotalcount + orgusruaaauditscount
												}

												if aorusruaaaudittotalcount == 0 {
													//fmt.Println("UAA OrgAuditor User", username,"has not be listed in Org.yml for Org", Orgs.Org.Name)

													if Audit == "unset" {
														//fmt.Println("UNSET!UNSET!")
														//fmt.Println("Unsetting user")
														if strings.TrimSpace(username) == "admin" {
															//fmt.Println("Skipping unset for admin user")
														} else {
															fmt.Println("UAA OrgAuditor User", username, "has not be listed in Org.yml for Org", Orgs.Org.Name)
															fmt.Println("UNSET!UNSET!")
															fmt.Println("Unsetting user")
															if MasterUserAudit == true {
																unset := exec.Command("cf", "unset-org-role", username, Orgs.Org.Name, "OrgAuditor")
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
															fmt.Println("User to be deleted: ", username, "UAA OrgAuditor")
														}
													} else {
														fmt.Println("Provide Valid Input")
													}
												} else {
													//fmt.Println("User is listed in Org.yml as UAA Org Audit User: ", username)
												}
											}
										} else if origin == "ldap" {
											if err == nil {
												for q := 0; q < OrgUsLenLDAPAuditor; q++ {
													//fmt.Println("LDAP Audit Usr: ", strings.ToLower(Orgs.Org.OrgUsers.LDAP.OrgAuditors[q]), ",", username)
													if strings.TrimSpace(Orgs.Org.OrgUsers.LDAP.OrgAuditors[q]) == username {
														orgusrldapauditscount = 1
													} else {
														orgusrldapauditscount = 0
													}
													aorusrldapaudittotalcount = aorusrldapaudittotalcount + orgusrldapauditscount

												}

												if aorusrldapaudittotalcount == 0 {
													//fmt.Println("LDAP OrgAuditor User",username,"has not be listed in Org.yml for Org", Orgs.Org.Name)
													if Audit == "unset" {
														//	fmt.Println("UNSET!UNSET!")
														//	fmt.Println("Unsetting user")
														if strings.TrimSpace(username) == "admin" {
															//fmt.Println("Skipping unset for admin user")
														} else {
															fmt.Println("LDAP OrgAuditor User", username, "has not be listed in Org.yml for Org", Orgs.Org.Name)
															fmt.Println("UNSET!UNSET!")
															fmt.Println("Unsetting user")
															if MasterUserAudit == true {
																unset := exec.Command("cf", "unset-org-role", username, Orgs.Org.Name, "OrgAuditor")
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
															fmt.Println("User to be deleted: ", username, "LDAP OrgAuditor")
														}
													} else {
														fmt.Println("Provide Valid Input")
													}
												} else {
													//	fmt.Println("User is listed in Org.yml as LDAP Org Audit User: ", username)
												}
											}
										} else {
											fmt.Println("Authentication of type", origin, "is not monitored")
										}
									}
								}
							} else {
								fmt.Println("No Org Auditor users exist")
							}
						} else {
							fmt.Println("Set OrgAuditor:", InitClusterConfigVals.ClusterDetails.SetOrgAuditor)
							fmt.Println("set SetOrgAuditor field to true to manage OrgAuditors for the Org")
						}

					} else {
						fmt.Println("command: ", target)
						fmt.Println("Err: ", target.Stdout, target.Stderr)
					}

					results := exec.Command("cf", "org-users", Orgs.Org.Name)
					if _, err := results.Output(); err != nil{
						fmt.Println("command: ", results)
						fmt.Println("Err: ", results.Stdout, err)
					} else {
						//fmt.Println("command: ", results)
						fmt.Println(results.Stdout)
					}

				} else {
					fmt.Println("command: ", guid )
					fmt.Println("Err: ", guid.Stdout, err)
					fmt.Println("Org doesn't exist")
				}
			} else {
				fmt.Println("Org Name doesn't match with folder name")
			}
		} else {
			fmt.Println("This is a protected Org")
		}
	}
	return err
}
