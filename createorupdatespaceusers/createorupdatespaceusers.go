package createorupdatespaceusers

import (
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

func CreateOrUpdateSpaceUsers(clustername string, cpath string, ostype string) error {

	var ProtectedOrgs ProtectedList
	var list List
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

				var getorg *exec.Cmd
				if ostype == "windows" {
					path := "\""+"/v3/organizations?names="+Orgs.Org.Name+"\""
					getorg = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_orgdetails.json")
				} else {
					path := "/v3/organizations?names="+Orgs.Org.Name
					getorg = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_orgdetails.json")
				}

				//var out bytes.Buffer
				//getorg.Stdout = &out

				err := getorg.Run()
				if err == nil {
					//	fmt.Println(getorg, getorg.Stdout, getorg.Stderr)
				} else {
					fmt.Println("err", getorg, getorg.Stdout, getorg.Stderr)
				}

				fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_orgdetails.json")
				if err != nil {
					fmt.Println(err)
				}
				//var Orgs Orglist
				var orgdetails OrgListJson

				if err := json.Unmarshal(fileSpaceJson, &orgdetails); err != nil {
					panic(err)
				}

				OrgLen := len(orgdetails.Resources)
				if OrgLen == 0 {
					fmt.Println("Org doesn't exist, Please Create Org")
				} else {

					orgguid := orgdetails.Resources[0].GUID

					SpaceLen := len(Orgs.Org.Spaces)
					for j := 0; j < SpaceLen; j++ {

						var getspace *exec.Cmd

						if ostype == "windows" {
							path := "\""+"/v3/spaces?names=" + Orgs.Org.Spaces[j].Name + "&organization_guids=" + orgguid+"\""
							getspace = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpacesUsers_spacedetails.json")
						} else {
							path := "/v3/spaces?names=" + Orgs.Org.Spaces[j].Name + "&organization_guids=" + orgguid
							getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpacesUsers_spacedetails.json")
						}

						err := getspace.Run()
						if err == nil {
							//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
						} else {
							fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
						}

						fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpacesUsers_spacedetails.json")
						if err != nil {
							fmt.Println(err)
						}

						var spacedets SpaceListJson
						if err := json.Unmarshal(fileSpaceJson, &spacedets); err != nil {
							panic(err)
						}
						if len(spacedets.Resources) == 0 {
							fmt.Println("Space does't exist, Create Space")
						} else {

							targetOrg := exec.Command("cf", "t", "-o", Orgs.Org.Name, "-s", Orgs.Org.Spaces[j].Name)
							if _, err := targetOrg.Output(); err == nil {
								//fmt.Println("command: ", targetOrg)
								fmt.Println("Targeted Org: ", targetOrg.Stdout)

								var spaceuserslist *exec.Cmd
								spaceguid := spacedets.Resources[0].GUID
								if ostype == "windows" {
									path := "\""+"/v3/roles/?space_guids="+spaceguid+"\""
									spaceuserslist = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrslist.json")
								} else {
									path := "/v3/roles/?space_guids="+spaceguid
									spaceuserslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrslist.json")
								}
								err := spaceuserslist.Run()
								if err == nil {
									//	fmt.Println(spaceuserslist, spaceuserslist.Stdout, spaceuserslist.Stderr)
								} else {
									fmt.Println("err", spaceuserslist, spaceuserslist.Stdout, spaceuserslist.Stderr)
								}

								fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_spaceusrslist.json")
								if err != nil {
									fmt.Println(err)
								}

								var spaceusrslist SpaceUsersListJson

								if err := json.Unmarshal(fileSpaceJson, &spaceusrslist); err != nil {
									panic(err)
								}

								//SpaceUsrLen := len(spaceusrslist.Resources)
								//fmt.Println("Number of Users currently exist in Org", Orgs.Org.Name, ":", SpaceUsrLen)


								if InitClusterConfigVals.ClusterDetails.SetSpaceManager == true {
									LDAPSpaceManLen := len(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers)
									for k := 0; k < LDAPSpaceManLen; k++ {

										var getspace *exec.Cmd
										if ostype == "windows" {
											path := "\""+"/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers[k] + "&origins=ldap"+"\""
											getspace = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										} else {
											path := "/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers[k] + "&origins=ldap"
											getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										}

										err := getspace.Run()
										if err == nil {
											//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
										} else {
											fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_userguidfind.json")
										if err != nil {
											fmt.Println(err)
										}

										var usedetails UserDetailsJson

										if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
											panic(err)
										}

										if len(usedetails.Resources) != 0 {

											userguid := usedetails.Resources[0].GUID
											var spaceusersdetailslist *exec.Cmd

											if ostype == "windows" {
												path := "\""+"/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_manager"+"\""
												spaceusersdetailslist = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											} else {
												path := "/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_manager"
												spaceusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											}

											err := spaceusersdetailslist.Run()
											if err == nil {
												//	fmt.Println(spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											} else {
												fmt.Println("err", spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											}

											fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											if err != nil {
												fmt.Println(err)
											}
											var spaceusrslist SpaceUsersListJson
											if err := json.Unmarshal(fileSpaceJson, &spaceusrslist); err != nil {
												panic(err)
											}

											SpaceUsrdetailsLen := len(spaceusrslist.Resources)

											if SpaceUsrdetailsLen != 0 {

											} else {

												fmt.Println("+ ", Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers[k], ",", "LDAP SpaceManager")
												cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceManager")
												if _, err := cmd.Output(); err != nil {
													fmt.Println("command: ", cmd)
													fmt.Println("Err: ", cmd.Stdout, err)
												} else {
													//fmt.Println("command: ", cmd)
													fmt.Println(cmd.Stdout)
												}
											}
										} else {
											fmt.Println(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers[k],"LDAP SpaceManagers User does't exist in Space",Orgs.Org.Spaces[j].Name, "please ask user to login to apps manager, and add user to Org as audit user")
										}
									}

									UAASpaceManLen := len(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers)
									for k := 0; k < UAASpaceManLen; k++ {

										var getspace *exec.Cmd
										if ostype == "windows" {
											path := "\""+"/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers[k] + "&origins=uaa"+"\""
											getspace = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										} else {
											path := "/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers[k] + "&origins=uaa"
											getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										}

										err := getspace.Run()
										if err == nil {
											//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
										} else {
											fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_userguidfind.json")
										if err != nil {
											fmt.Println(err)
										}

										var usedetails UserDetailsJson
										if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
											panic(err)
										}

										if len(usedetails.Resources) != 0 {

											var spaceusersdetailslist *exec.Cmd
											userguid := usedetails.Resources[0].GUID
											if ostype == "windows" {
												path := "\""+"/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_manager"+"\""
												spaceusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											} else {
												path := "/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_manager"
												spaceusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											}
											err := spaceusersdetailslist.Run()
											if err == nil {
												//fmt.Println(spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											} else {
												fmt.Println("err", spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											}

											fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											if err != nil {
												fmt.Println(err)
											}

											var spaceusrslist SpaceUsersListJson
											if err := json.Unmarshal(fileSpaceJson, &spaceusrslist); err != nil {
												panic(err)
											}
											SpaceUsrdetailsLen := len(spaceusrslist.Resources)

											if SpaceUsrdetailsLen != 0 {

											} else {

												fmt.Println("+ ", Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers[k], ",", "UAA SpaceManager")
												cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceManager")
												if _, err := cmd.Output(); err != nil {
													fmt.Println("command: ", cmd)
													fmt.Println("Err: ", cmd.Stdout, err)
												} else {
													//fmt.Println("command: ", cmd)
													fmt.Println(cmd.Stdout)
												}
											}
										} else {
											fmt.Println(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers[k],"UAA SpaceManagers User does't exist in Space",Orgs.Org.Spaces[j].Name,"please ask user to login to apps manager, and add user to Org as audit user")
										}
									}

									SSOSpaceManLen := len(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers)
									SSOProvider := InitClusterConfigVals.ClusterDetails.SSOProvider
									for k := 0; k < SSOSpaceManLen; k++ {

										var getspace *exec.Cmd
										if ostype == "windows" {
											path := "\""+"/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers[k] + "&origins="+SSOProvider+"\""
											getspace = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										} else {
											path := "/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers[k] + "&origins="+SSOProvider
											getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										}

										err := getspace.Run()
										if err == nil {
											//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
										} else {
											fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_userguidfind.json")
										if err != nil {
											fmt.Println(err)
										}

										var usedetails UserDetailsJson

										if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
											panic(err)
										}

										if len(usedetails.Resources) != 0 {

											var spaceusersdetailslist *exec.Cmd
											userguid := usedetails.Resources[0].GUID

											if ostype == "windows" {
												path := "\""+"/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_manager"+"\""
												spaceusersdetailslist = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											} else {
												path := "/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_manager"
												spaceusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											}

											err := spaceusersdetailslist.Run()
											if err == nil {
												//fmt.Println(spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											} else {
												fmt.Println("err", spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											}

											fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											if err != nil {
												fmt.Println(err)
											}

											var spaceusrslist SpaceUsersListJson
											if err := json.Unmarshal(fileSpaceJson, &spaceusrslist); err != nil {
												panic(err)
											}

											SpaceUsrdetailsLen := len(spaceusrslist.Resources)

											if SpaceUsrdetailsLen != 0 {

											} else {

												fmt.Println("+ ", Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers[k], ",", "SSO SpaceManager")
												cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceManager")
												if _, err := cmd.Output(); err != nil {
													fmt.Println("command: ", cmd)
													fmt.Println("Err: ", cmd.Stdout, err)
												} else {
													//fmt.Println("command: ", cmd)
													fmt.Println(cmd.Stdout)
												}
											}
										} else {
											fmt.Println(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers[k],"SSO SpaceManagers User does't exist in Space",Orgs.Org.Spaces[j].Name, "please ask user to login to apps manager, and add user to Org as audit user")
										}
									}

								} else {
									fmt.Println("Set SpaceManager:", InitClusterConfigVals.ClusterDetails.SetSpaceManager)
									fmt.Println("set SetSpaceManager field to true to manage SpaceManager for the Org")
								}

								if InitClusterConfigVals.ClusterDetails.SetSpaceAuditor == true {

									LDAPSpaceAuditLen := len(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors)
									for k := 0; k < LDAPSpaceAuditLen; k++ {

										var getspace *exec.Cmd

										if ostype == "windows" {
											path := "\""+"/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors[k] + "&origins=ldap"+"\""
											getspace = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										} else {
											path := "/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors[k] + "&origins=ldap"
											getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										}

										err := getspace.Run()
										if err == nil {
											//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
										} else {
											fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_userguidfind.json")
										if err != nil {
											fmt.Println(err)
										}

										var usedetails UserDetailsJson

										if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
											panic(err)
										}

										if len(usedetails.Resources) != 0 {

											var spaceusersdetailslist *exec.Cmd
											userguid := usedetails.Resources[0].GUID

											if ostype == "windows" {
												path := "\""+"/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_auditor"+"\""
												spaceusersdetailslist = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											} else {
												path := "/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_auditor"
												spaceusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											}
											err := spaceusersdetailslist.Run()
											if err == nil {
												//	fmt.Println(spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											} else {
												fmt.Println("err", spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											}

											fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											if err != nil {
												fmt.Println(err)
											}
											var spaceusrslist SpaceUsersListJson
											if err := json.Unmarshal(fileSpaceJson, &spaceusrslist); err != nil {
												panic(err)
											}

											SpaceUsrdetailsLen := len(spaceusrslist.Resources)

											if SpaceUsrdetailsLen != 0 {

											} else {

												fmt.Println("+ ", Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors[k], ",", "LDAP SpaceAuditor")
												cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceAuditor")
												if _, err := cmd.Output(); err != nil {
													fmt.Println("command: ", cmd)
													fmt.Println("Err: ", cmd.Stdout, err)
												} else {
													//fmt.Println("command: ", cmd)
													fmt.Println(cmd.Stdout)
												}
											}
										} else {
											fmt.Println(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors[k],"LDAP SpaceAuditor User does't exist in Space",Orgs.Org.Spaces[j].Name,"please ask user to login to apps manager, and add user to Org as audit user")
										}
									}

									UAASpaceAuditLen := len(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors)
									for k := 0; k < UAASpaceAuditLen; k++ {

										var getspace *exec.Cmd
										if ostype == "windows" {
											path := "\""+"/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors[k] + "&origins=uaa"+"\""
											getspace = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										} else {
											path := "/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors[k] + "&origins=uaa"
											getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										}

										err := getspace.Run()
										if err == nil {
											//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
										} else {
											fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_userguidfind.json")
										if err != nil {
											fmt.Println(err)
										}

										var usedetails UserDetailsJson

										if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
											panic(err)
										}

										if len(usedetails.Resources) != 0 {

											var spaceusersdetailslist *exec.Cmd
											userguid := usedetails.Resources[0].GUID
											if ostype == "windows" {
												path := "\""+"/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_auditor"+"\""
												spaceusersdetailslist = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											} else {
												path := "/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_auditor"
												spaceusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											}
											err := spaceusersdetailslist.Run()
											if err == nil {
												//fmt.Println(spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											} else {
												fmt.Println("err", spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											}

											fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											if err != nil {
												fmt.Println(err)
											}
											var spaceusrslist SpaceUsersListJson
											if err := json.Unmarshal(fileSpaceJson, &spaceusrslist); err != nil {
												panic(err)
											}

											SpaceUsrdetailsLen := len(spaceusrslist.Resources)

											if SpaceUsrdetailsLen != 0 {

											} else {

												fmt.Println("+ ", Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors[k], ",", "UAA SpaceAuditor")
												cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceAuditor")
												if _, err := cmd.Output(); err != nil {
													fmt.Println("command: ", cmd)
													fmt.Println("Err: ", cmd.Stdout, err)
												} else {
													//fmt.Println("command: ", cmd)
													fmt.Println(cmd.Stdout)
												}
											}
										} else {
											fmt.Println(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors[k],"UAA SpaceAuditor User does't exist in Space",Orgs.Org.Spaces[j].Name, "please ask user to login to apps manager, and add user to Org as audit user")
										}
									}

									SSOSpaceAuditLen := len(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceAuditors)
									SSOProvider := InitClusterConfigVals.ClusterDetails.SSOProvider
									for k := 0; k < SSOSpaceAuditLen; k++ {

										var getspace *exec.Cmd
										if ostype == "windows" {
											path := "\""+"/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceAuditors[k] + "&origins="+SSOProvider+"\""
											getspace = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										} else {
											path := "/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceAuditors[k] + "&origins="+SSOProvider
											getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										}

										err := getspace.Run()
										if err == nil {
											//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
										} else {
											fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_userguidfind.json")
										if err != nil {
											fmt.Println(err)
										}

										var usedetails UserDetailsJson
										if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
											panic(err)
										}

										if len(usedetails.Resources) != 0 {

											var spaceusersdetailslist *exec.Cmd
											userguid := usedetails.Resources[0].GUID
											if ostype == "windows" {
												path := "\""+"/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_auditor"+"\""
												spaceusersdetailslist = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											} else {
												path := "/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_auditor"
												spaceusersdetailslist = exec.Command( "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											}

											err := spaceusersdetailslist.Run()
											if err == nil {
												//fmt.Println(spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											} else {
												fmt.Println("err", spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											}

											fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											if err != nil {
												fmt.Println(err)
											}
											var spaceusrslist SpaceUsersListJson
											if err := json.Unmarshal(fileSpaceJson, &spaceusrslist); err != nil {
												panic(err)
											}

											SpaceUsrdetailsLen := len(spaceusrslist.Resources)

											if SpaceUsrdetailsLen != 0 {

											} else {

												fmt.Println("+ ", Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceAuditors[k], ",", "SSO SpaceAuditor")
												cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceAuditors[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceAuditor")
												if _, err := cmd.Output(); err != nil {
													fmt.Println("command: ", cmd)
													fmt.Println("Err: ", cmd.Stdout, err)
												} else {
													//fmt.Println("command: ", cmd)
													fmt.Println(cmd.Stdout)
												}
											}
										} else {
											fmt.Println(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceAuditors[k],"SSO SpaceAuditor User does't exist in Space",Orgs.Org.Spaces[j].Name,"please ask user to login to apps manager, and add user to Org as audit user")
										}
									}

								} else {

									fmt.Println("Set SpaceAuditor:", InitClusterConfigVals.ClusterDetails.SetSpaceAuditor)
									fmt.Println("set SetSpaceAuditor field to true to manage SpaceAuditor for the Org")
								}

								if InitClusterConfigVals.ClusterDetails.SetSpaceDeveloper == true {

									LDAPSpaceDevLen := len(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers)
									for k := 0; k < LDAPSpaceDevLen; k++ {

										var getspace *exec.Cmd
										if ostype == "windows" {
											path := "\""+"/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers[k] + "&origins=ldap"+"\""
											getspace = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										} else {
											path := "/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers[k] + "&origins=ldap"
											getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										}

										err := getspace.Run()
										if err == nil {
											//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
										} else {
											fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_userguidfind.json")
										if err != nil {
											fmt.Println(err)
										}

										var usedetails UserDetailsJson
										if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
											panic(err)
										}

										if len(usedetails.Resources) != 0 {

											var spaceusersdetailslist *exec.Cmd
											userguid := usedetails.Resources[0].GUID
											if ostype == "windows" {
												path := "\""+"/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_developer"+"\""
												spaceusersdetailslist = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											} else {
												path := "/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_developer"
												spaceusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											}

											err := spaceusersdetailslist.Run()
											if err == nil {
												//fmt.Println(spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											} else {
												fmt.Println("err", spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											}

											fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											if err != nil {
												fmt.Println(err)
											}

											var spaceusrslist SpaceUsersListJson
											if err := json.Unmarshal(fileSpaceJson, &spaceusrslist); err != nil {
												panic(err)
											}
											SpaceUsrdetailsLen := len(spaceusrslist.Resources)

											if SpaceUsrdetailsLen != 0 {

											} else {

												fmt.Println("+ ", Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers[k], ",", "LDAP SpaceDeveloper")
												cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceDeveloper")
												if _, err := cmd.Output(); err != nil {
													fmt.Println("command: ", cmd)
													fmt.Println("Err: ", cmd.Stdout, err)
												} else {
													//fmt.Println("command: ", cmd)
													fmt.Println(cmd.Stdout)
												}
											}
										} else {
											fmt.Println(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers[k],"LDAP SpaceDeveloper User does't exist in Space",Orgs.Org.Spaces[j].Name, "please ask user to login to apps manager, and add user to Org as audit user")
										}
									}

									UAASpaceDevLen := len(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers)
									for k := 0; k < UAASpaceDevLen; k++ {

										var getspace *exec.Cmd
										if ostype == "windows" {
											path := "\""+"/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers[k] + "&origins=uaa"+"\""
											getspace = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										} else {
											path := "/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers[k] + "&origins=uaa"
											getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										}

										err := getspace.Run()
										if err == nil {
											//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
										} else {
											fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_userguidfind.json")
										if err != nil {
											fmt.Println(err)
										}

										var usedetails UserDetailsJson

										if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
											panic(err)
										}

										if len(usedetails.Resources) != 0 {

											var spaceusersdetailslist *exec.Cmd
											userguid := usedetails.Resources[0].GUID

											if ostype == "windows" {
												path := "\""+"/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_developer"+"\""
												spaceusersdetailslist = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											} else {
												path := "/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_developer"
												spaceusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											}
											err := spaceusersdetailslist.Run()
											if err == nil {
												//fmt.Println(spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											} else {
												fmt.Println("err", spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											}

											fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											if err != nil {
												fmt.Println(err)
											}

											var spaceusrslist SpaceUsersListJson
											if err := json.Unmarshal(fileSpaceJson, &spaceusrslist); err != nil {
												panic(err)
											}

											SpaceUsrdetailsLen := len(spaceusrslist.Resources)

											if SpaceUsrdetailsLen != 0 {

											} else {

												fmt.Println("+ ", Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers[k], ",", "UAA SpaceDeveloper")
												cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceDeveloper")
												if _, err := cmd.Output(); err != nil {
													fmt.Println("command: ", cmd)
													fmt.Println("Err: ", cmd.Stdout, err)
												} else {
													//fmt.Println("command: ", cmd)
													fmt.Println(cmd.Stdout)
												}
											}
										} else {
											fmt.Println(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers[k],"UAA SpaceDeveloper User does't exist in Space",Orgs.Org.Spaces[j].Name,"please ask user to login to apps manager, and add user to Org as audit user")
										}
									}

									SSOSpaceDevLen := len(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers)
									SSOProvider := InitClusterConfigVals.ClusterDetails.SSOProvider
									for k := 0; k < SSOSpaceDevLen; k++ {

										var getspace *exec.Cmd

										if ostype == "windows" {
											path := "\""+"/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers[k] + "&origins="+SSOProvider+"\""
											getspace = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										} else {
											path := "/v3/users/?usernames=" + Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers[k] + "&origins="+SSOProvider
											getspace = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_userguidfind.json")
										}

										err := getspace.Run()
										if err == nil {
											//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
										} else {
											fmt.Println("err", getspace, getspace.Stdout, getspace.Stderr)
										}

										fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_userguidfind.json")
										if err != nil {
											fmt.Println(err)
										}

										var usedetails UserDetailsJson
										//var spacedets SpaceListJson
										if err := json.Unmarshal(fileSpaceJson, &usedetails); err != nil {
											panic(err)
										}

										if len(usedetails.Resources) != 0 {
											var spaceusersdetailslist *exec.Cmd
											userguid := usedetails.Resources[0].GUID
											if ostype == "windows" {
												path := "\""+"/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_developer"+"\""
												spaceusersdetailslist = exec.Command("powershell", "-command","cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")

											} else {
												path := "/v3/roles/?space_guids="+spaceguid + "&user_guids=" + userguid + "&types=space_developer"
												spaceusersdetailslist = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaceUsers_spaceusrsroledets.json")

											}
											err := spaceusersdetailslist.Run()
											if err == nil {
												//fmt.Println(spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											} else {
												fmt.Println("err", spaceusersdetailslist, spaceusersdetailslist.Stdout, spaceusersdetailslist.Stderr)
											}

											fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaceUsers_spaceusrsroledets.json")
											if err != nil {
												fmt.Println(err)
											}

											var spaceusrslist SpaceUsersListJson
											if err := json.Unmarshal(fileSpaceJson, &spaceusrslist); err != nil {
												panic(err)
											}

											SpaceUsrdetailsLen := len(spaceusrslist.Resources)

											if SpaceUsrdetailsLen != 0 {

											} else {

												fmt.Println("+ ", Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers[k], ",", "SSO SpaceDeveloper")
												cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceDeveloper")
												if _, err := cmd.Output(); err != nil {
													fmt.Println("command: ", cmd)
													fmt.Println("Err: ", cmd.Stdout, err)
												} else {
													//fmt.Println("command: ", cmd)
													fmt.Println(cmd.Stdout)
												}
											}
										} else {
											fmt.Println(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers[k],"SSO SpaceDeveloper User does't exist in Space",Orgs.Org.Spaces[j].Name,"please ask user to login to apps manager, and add user to Org as audit user")
										}
									}

								} else {

									fmt.Println("Set SpaceDeveloper:", InitClusterConfigVals.ClusterDetails.SetSpaceDeveloper)
									fmt.Println("set SetSpaceDeveloper field to true to manage SpaceDeveloper for the Org")
								}

								results := exec.Command("cf", "space-users", Orgs.Org.Name, Orgs.Org.Spaces[j].Name )
								if _, err := results.Output(); err != nil{
									fmt.Println("command: ", results)
									fmt.Println("Err: ", results.Stdout, err)
								} else {
									//fmt.Println("command: ", results)
									fmt.Println(results.Stdout)
								}

							} else {
								fmt.Println("command: ", targetOrg)
								fmt.Println("Err: ", targetOrg.Stdout,targetOrg.Stderr)
							}
						}
					}
				}

			} else {
				fmt.Println("Org Name does't match with folder name")
			}
		}
	}
	return err
}
