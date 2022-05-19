package createorupdatespaces

import (
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
type SpaceIsoJson struct {
	Data struct {
		GUID string `json:"guid"`
	} `json:"data"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Related struct {
			Href string `json:"href"`
		} `json:"related"`
	} `json:"links"`
}
type IsoJson struct {
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
		GUID      string    `json:"guid"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Links     struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Organizations struct {
				Href string `json:"href"`
			} `json:"organizations"`
		} `json:"links"`
		Metadata struct {
			Annotations struct {
			} `json:"annotations"`
			Labels struct {
			} `json:"labels"`
		} `json:"metadata"`
	} `json:"resources"`
}
type ProtectedList struct {
	Org                         []string `yaml:"Org"`
	Quota                       []string `yaml:"quota"`
	DefaultRunningSecurityGroup string   `yaml:"DefaultRunningSecurityGroup"`
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
type Orglist struct {
	Org struct {
		Name string `yaml:"Name"`
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
type SpaceStateYaml struct {
	SpaceState struct {
		Org          string `yaml:"Org"`
		OrgGuid      string `yaml:"OrgGuid"`
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

func CreateOrUpdateSpaces(clustername string, cpath string, ostype string) error {

	var ProtectedOrgs ProtectedList
	var list List
	var gitlist GitList
	var InitClusterConfigVals InitClusterConfigVals
	var ListYml string
	var LenList int

	spath := cpath + "/" + clustername + "-state/"
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

	}

	var OrgsYml string

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
	EnableIsolationAudit := InitClusterConfigVals.ClusterDetails.EnableIsolationAudit

	for i := 0; i < LenList; i++ {

		var count, totalcount int

		var OrgName string
		if InitClusterConfigVals.ClusterDetails.EnableGitSubTree != true {
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

			if OrgName == Orgs.Org.Name {
				var getorg *exec.Cmd
				if ostype == "windows" {
					path := "\"" + "/v3/organizations?names=" + Orgs.Org.Name + "\""
					getorg = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_orgdetails.json")

				} else {
					path := "/v3/organizations?names=" + Orgs.Org.Name
					getorg = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_orgdetails.json")

				}

				err := getorg.Run()
				if err == nil {
					//	fmt.Println(getorg, getorg.Stdout, getorg.Stderr)
				} else {
					fmt.Println("err", getorg, getorg.Stdout, getorg.Stderr)
				}

				fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaces_orgdetails.json")
				if err != nil {
					fmt.Println(err)
				}

				//var Orgs Orglist
				var orgdetails OrgListJson

				if err := json.Unmarshal(fileSpaceJson, &orgdetails); err != nil {
					panic(err)
				}
				OrgLen := len(orgdetails.Resources)

				if OrgLen != 0 {

					target := exec.Command("cf", "t", "-o", Orgs.Org.Name)
					if _, err := target.Output(); err == nil {
						orgguid := orgdetails.Resources[0].GUID
						SpaceLen := len(Orgs.Org.Spaces)

						for j := 0; j < SpaceLen; j++ {

							//Getting Org Guid from state file
							fullpath := spath + Orgs.Org.Name + "_" + Orgs.Org.Spaces[j].Name + "_SpaceState.yml"
							var SpaceGuidPull string

							SpaceStateYml := fullpath
							fileSpaceStateYml, err := ioutil.ReadFile(SpaceStateYml)
							if err != nil {
								fmt.Println(err)
							}
							var spacestatedetails SpaceStateYaml
							err = yaml.Unmarshal([]byte(fileSpaceStateYml), &spacestatedetails)
							if err != nil {
								panic(err)
							}
							SpaceStateGuid := spacestatedetails.SpaceState.SpaceGuid

							//Checking if space exist with Guid from State File
							var getspaceguid *exec.Cmd
							if ostype == "windows" {
								path := "\"" + "/v3/spaces?guids=" + SpaceStateGuid + "\""
								getspaceguid = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spacedetails_guid.json")

							} else {
								path := "/v3/spaces?guids=" + SpaceStateGuid
								getspaceguid = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spacedetails_guid.json")

							}
							err = getspaceguid.Run()
							if err == nil {
								//	fmt.Println(getorg, getorg.Stdout, getorg.Stderr)
							} else {
								fmt.Println("err", getspaceguid, getspaceguid.Stdout, getspaceguid.Stderr)
							}
							fileSpaceGuidJson, err := ioutil.ReadFile("CreateOrUpdateSpaces_spacedetails_guid.json")
							if err != nil {
								fmt.Println(err)
							}
							var spacedetailsguid SpaceListJson
							if err := json.Unmarshal(fileSpaceGuidJson, &spacedetailsguid); err != nil {
								panic(err)
							}

							// Checking if space exist with space Name in Org
							var getspacename *exec.Cmd
							if ostype == "windows" {
								path := "\"" + "/v3/spaces?names=" + Orgs.Org.Spaces[j].Name + "&organization_guids=" + orgguid + "\""
								getspacename = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spacedetails_name.json")
							} else {
								path := "/v3/spaces?names=" + Orgs.Org.Spaces[j].Name + "&organization_guids=" + orgguid
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

							SpaceStateGuidLen := len(spacedetailsguid.Resources)
							SpaceStateNameLen := len(spacedetailsname.Resources)

							if SpaceStateGuidLen != 0 {

								//fmt.Println("Space exists in state and platform")
								spacename := spacedetailsguid.Resources[0].Name

								if Orgs.Org.Spaces[j].Name == spacename {

									SpaceGuidPull = spacedetailsguid.Resources[0].GUID
									var spacedets SpaceListJson

									if err := json.Unmarshal(fileSpaceNameJson, &spacedets); err != nil {
										panic(err)
									}
									spaceguid := spacedets.Resources[0].GUID
									var getisoguid *exec.Cmd
									if ostype == "windows" {
										path := "\"" + "/v3/spaces/" + spaceguid + "/relationships/isolation_segment" + "\""
										getisoguid = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spaceisodetails.json")

									} else {
										path := "/v3/spaces/" + spaceguid + "/relationships/isolation_segment"
										getisoguid = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spaceisodetails.json")

									}

									err := getisoguid.Run() // it can be nill
									if err == nil {
										//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
									} else {
										fmt.Println("err", getisoguid, getisoguid.Stdout, getisoguid.Stderr)
									}
									fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaces_spaceisodetails.json")
									if err != nil {
										fmt.Println(err)
									}
									var spaceiso SpaceIsoJson
									if err := json.Unmarshal(fileSpaceJson, &spaceiso); err != nil {
										panic(err)
									}
									isoguid := spaceiso.Data.GUID // will be null
									// iso segment guid if noting specific it will be default
									////
									// Pulling isolation segment Name from Guid
									var existingisoguid *exec.Cmd
									if ostype == "windows" {
										path := "\"" + "/v3/isolation_segments?guids=" + isoguid + "\""
										existingisoguid = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_existingisodetails.json")
									} else {
										path := "/v3/isolation_segments?guids=" + isoguid
										existingisoguid = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_existingisodetails.json")
									}
									err = existingisoguid.Run()
									if err == nil {
										//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
									} else {
										fmt.Println("err", existingisoguid, existingisoguid.Stdout, existingisoguid.Stderr)
									}
									fileSpaceJson, err = ioutil.ReadFile("CreateOrUpdateSpaces_existingisodetails.json")
									if err != nil {
										fmt.Println(err)
									}
									var isoexistingdetails IsoJson
									if err := json.Unmarshal(fileSpaceJson, &isoexistingdetails); err != nil {
										panic(err)
									}
									//// 	pulled info of iso

									///////////From YAML
									// 	Pulling Guid of Iso Specified in YAML file
									var detailsisoguid *exec.Cmd
									segname := Orgs.Org.Spaces[j].IsolationSeg
									if ostype == "windows" {
										path := "\"" + "/v3/isolation_segments?names=" + segname + "\""
										detailsisoguid = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_isodetails.json")
									} else {
										path := "/v3/isolation_segments?names=" + segname
										detailsisoguid = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_isodetails.json")
									}

									err = detailsisoguid.Run()
									if err == nil {
										//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
									} else {
										fmt.Println("err", detailsisoguid, detailsisoguid.Stdout, detailsisoguid.Stderr)
									}
									fileSpaceJson, err = ioutil.ReadFile("CreateOrUpdateSpaces_isodetails.json")
									if err != nil {
										fmt.Println(err)
									}
									var isodetails IsoJson
									if err := json.Unmarshal(fileSpaceJson, &isodetails); err != nil {
										panic(err)
									}
									// check if iso defined in yaml exist in platform

									if len(isodetails.Resources) == 0 {
										// If some name is defined in YAML but that does/t exits
										if segname == "" && isoguid == "" {

											// No Iso defined in yaml and No Iso required
											// No action needs to be done

										} else if segname != "" && isoguid == "" {

											// This is just for stdout
											// Iso is asked in YAML but that is installed in platform
											fmt.Println("+ isolation segment", Orgs.Org.Spaces[j].IsolationSeg)
											fmt.Println("+ Org, Space, Isolation Segment: ", Orgs.Org.Name, ",", Orgs.Org.Spaces[j].Name, ",", segname)
											fmt.Println("No Isolation segment exists with name: ", segname)
										} else if segname == "" && isoguid != "" {

											if strings.ToLower(Orgs.IsolationAudit) == "unbind" {
												fmt.Println("- isolation segment", isoexistingdetails.Resources[0].Name)
												fmt.Println("Removing Isolation Segment")
												if EnableIsolationAudit == true {
													resetspace := exec.Command("cf", "reset-space-isolation-segment", Orgs.Org.Spaces[j].Name)
													if _, err := resetspace.Output(); err != nil {
														fmt.Println("command: ", resetspace)
														fmt.Println("Err: ", resetspace.Stdout, err)
													} else {
														fmt.Println("command: ", resetspace)
														fmt.Println(resetspace.Stdout)
													}
												} else {
													fmt.Println("IsolationSegment Audit flag is not enabled, please work with pcf operator to unbind")
												}
											} else if strings.ToLower(Orgs.IsolationAudit) == "list" {
												fmt.Println("Unbind!Unbind!")
												fmt.Println("isolation segment:", isoexistingdetails.Resources[0].Name)

											} else {
												fmt.Println("Provide valid input")
											}
										}
									} else {

										if isodetails.Resources[0].GUID == isoguid {
											//fmt.Println(isodetails.Resources[0].GUID, isoguid)
											// iso guid defined yaml == iso attached to space
											// Iso defined in YAML exist in platform
											// Check if that is same as currectly binded
											// No action needed

										} else {
											if isoguid == "" {

												// Iso defined in YAML exist
												// but Currently space is not binded to any Iso
												// This is a new request, binding to isolation segment

												fmt.Println("+ isolation segment", Orgs.Org.Spaces[j].IsolationSeg)
												//fmt.Println("- isolation segment", isoexistingdetails.Resources[0].Name)
												fmt.Println("Enabling Space Isolation Segment ")
												fmt.Println("Org, Space, Isolation Segment: ", Orgs.Org.Name, ",", Orgs.Org.Spaces[j].Name, ",", Orgs.Org.Spaces[j].IsolationSeg)
												iso := exec.Command("cf", "enable-org-isolation", Orgs.Org.Name, Orgs.Org.Spaces[j].IsolationSeg)
												if _, err := iso.Output(); err != nil {
													fmt.Println("command: ", iso)
													fmt.Println("Err: ", iso.Stdout, err)
												} else {
													//	fmt.Println("command: ", iso)
													fmt.Println(iso.Stdout)
												}
												isospace := exec.Command("cf", "set-space-isolation-segment", Orgs.Org.Spaces[j].Name, Orgs.Org.Spaces[j].IsolationSeg)
												if _, err := isospace.Output(); err != nil {
													fmt.Println("command: ", isospace)
													fmt.Println("Err: ", isospace.Stdout, err)
												} else {
													//	fmt.Println("command: ", isospace)
													fmt.Println(isospace.Stdout)
												}

											} else {
												// Iso defined in YAML exist
												// but Currently space is not binded to any Iso
												// This is a change in request to remove isolation segment
												fmt.Println("+ isolation segment", Orgs.Org.Spaces[j].IsolationSeg)
												fmt.Println("- isolation segment", isoexistingdetails.Resources[0].Name)
												fmt.Println("Currently doesn't support changing Isolation Segments, Please change isolation segment assigned to space manually")
											}
										}
									}
								} else {

									// Checking space name
									fmt.Println("Resetting Space Name")
									fmt.Println("- ", spacename)
									fmt.Println("+ ", Orgs.Org.Spaces[j].Name)
									var path string
									//path := "v3/spaces/"+spacedetailsguid.Resources[0].GUID+"/?name="+Orgs.Org.Spaces[j].Name
									//path := "v3/spaces/"+spacedetailsguid.Resources[0].GUID+"/?name="+Orgs.Org.Spaces[j].Name
									//renamespace := exec.Command("cf", "curl", path)
									renamespace := exec.Command("cf", "rename-space", spacename, Orgs.Org.Spaces[j].Name)

									err = renamespace.Run()

									if err == nil {
										//	fmt.Println(err, renamespace, renamespace.Stdout, renamespace.Stderr)

									} else {
										fmt.Println("err", renamespace, renamespace.Stdout, renamespace.Stderr)
									}

									// checking isolation segments
									////////////// From CF
									// pulling if any isolation segment assigned to org

									//	var getspacename *exec.Cmd
									if ostype == "windows" {
										path := "\"" + "/v3/spaces?names=" + Orgs.Org.Spaces[j].Name + "&organization_guids=" + orgguid + "\""
										getspacename = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spacedetails_name.json")
									} else {
										path := "/v3/spaces?names=" + Orgs.Org.Spaces[j].Name + "&organization_guids=" + orgguid
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

									var spacedets SpaceListJson
									if err := json.Unmarshal(fileSpaceNameJson, &spacedets); err != nil {
										panic(err)
									}

									//	fmt.Println(fileSpaceNameJson)
									//	fmt.Println(spacedets)

									var getisoguid *exec.Cmd
									spaceguid := spacedets.Resources[0].GUID
									if ostype == "windows" {
										path = "\"" + "/v3/spaces/" + spaceguid + "/relationships/isolation_segment" + "\""
										getisoguid = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spaceisodetails.json")

									} else {
										path = "/v3/spaces/" + spaceguid + "/relationships/isolation_segment"
										getisoguid = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spaceisodetails.json")

									}
									err = getisoguid.Run() // it can be nill
									if err == nil {
										//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
									} else {
										fmt.Println("err", getisoguid, getisoguid.Stdout, getisoguid.Stderr)
									}
									fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaces_spaceisodetails.json")
									if err != nil {
										fmt.Println(err)
									}
									var spaceiso SpaceIsoJson
									if err := json.Unmarshal(fileSpaceJson, &spaceiso); err != nil {
										panic(err)
									}
									isoguid := spaceiso.Data.GUID // will be null
									// iso segment guid if noting specific it will be default
									////
									// Pulling isolation segment Name from Guid
									var existingisoguid *exec.Cmd
									if ostype == "windows" {
										path = "\"" + "/v3/isolation_segments?guids=" + isoguid + "\""
										existingisoguid = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_existingisodetails.json")
									} else {
										path = "/v3/isolation_segments?guids=" + isoguid
										existingisoguid = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_existingisodetails.json")
									}
									err = existingisoguid.Run()
									if err == nil {
										//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
									} else {
										fmt.Println("err", existingisoguid, existingisoguid.Stdout, existingisoguid.Stderr)
									}
									fileSpaceJson, err = ioutil.ReadFile("CreateOrUpdateSpaces_existingisodetails.json")
									if err != nil {
										fmt.Println(err)
									}
									var isoexistingdetails IsoJson
									if err := json.Unmarshal(fileSpaceJson, &isoexistingdetails); err != nil {
										panic(err)
									}
									//// 	pulled info of iso

									///////////From YAML
									// 	Pulling Guid of Iso Specified in YAML file
									segname := Orgs.Org.Spaces[j].IsolationSeg
									var detailsisoguid *exec.Cmd
									if ostype == "windows" {
										path = "\"" + "/v3/isolation_segments?names=" + segname + "\""
										detailsisoguid = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_isodetails.json")

									} else {
										path = "/v3/isolation_segments?names=" + segname
										detailsisoguid = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_isodetails.json")

									}
									err = detailsisoguid.Run()
									if err == nil {
										//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
									} else {
										fmt.Println("err", detailsisoguid, detailsisoguid.Stdout, detailsisoguid.Stderr)
									}
									fileSpaceJson, err = ioutil.ReadFile("CreateOrUpdateSpaces_isodetails.json")
									if err != nil {
										fmt.Println(err)
									}
									var isodetails IsoJson
									if err := json.Unmarshal(fileSpaceJson, &isodetails); err != nil {
										panic(err)
									}

									// check if iso defined in yaml exist in platform
									if len(isodetails.Resources) == 0 {
										// If some name is defined in YAML but that does/t exits
										if segname == "" && isoguid == "" {

											// No Iso defined in yaml and No Iso required
											// No action needs to be done

										} else if segname != "" && isoguid == "" {

											// This is just for stdout
											// Iso is asked in YAML but that is installed in platform
											fmt.Println("+ isolation segment", Orgs.Org.Spaces[j].IsolationSeg)
											fmt.Println("+ Org, Space, Isolation Segment: ", Orgs.Org.Name, ",", Orgs.Org.Spaces[j].Name, ",", segname)
											fmt.Println("No Isolation segment exists with name: ", segname)
										} else if segname == "" && isoguid != "" {
											fmt.Println("- isolation segment", isoexistingdetails.Resources[0].Name)
											fmt.Println("Removing Isolation Segment")
											resetspace := exec.Command("cf", "reset-space-isolation-segment", Orgs.Org.Spaces[j].Name)
											if _, err := resetspace.Output(); err != nil {
												fmt.Println("command: ", resetspace)
												fmt.Println("Err: ", resetspace.Stdout, err)
											} else {
												fmt.Println("command: ", resetspace)
												fmt.Println(resetspace.Stdout)
											}
										}
									} else {
										if isodetails.Resources[0].GUID == isoguid {
											//fmt.Println(isodetails.Resources[0].GUID, isoguid)

											// iso guid defined yaml == iso attached to space

											// Iso defined in YAML exist in platform
											// Check if that is same as currectly binded
											// No action needed

										} else {
											if isoguid == "" {

												// Iso defined in YAML exist
												// but Currently space is not binded to any Iso
												// This is a new request, binding to isolation segment

												fmt.Println("+ isolation segment", Orgs.Org.Spaces[j].IsolationSeg)
												//fmt.Println("- isolation segment", isoexistingdetails.Resources[0].Name)
												fmt.Println("Enabling Space Isolation Segment")
												fmt.Println("Org, Space, Isolation Segment: ", Orgs.Org.Name, ",", Orgs.Org.Spaces[j].Name, ",", Orgs.Org.Spaces[j].IsolationSeg)
												iso := exec.Command("cf", "enable-org-isolation", Orgs.Org.Name, Orgs.Org.Spaces[j].IsolationSeg)
												if _, err := iso.Output(); err != nil {
													fmt.Println("command: ", iso)
													fmt.Println("Err: ", iso.Stdout, err)
												} else {
													//	fmt.Println("command: ", iso)
													fmt.Println(iso.Stdout)
												}
												isospace := exec.Command("cf", "set-space-isolation-segment", Orgs.Org.Spaces[j].Name, Orgs.Org.Spaces[j].IsolationSeg)
												if _, err := isospace.Output(); err != nil {
													fmt.Println("command: ", isospace)
													fmt.Println("Err: ", isospace.Stdout, err)
												} else {
													//	fmt.Println("command: ", isospace)
													fmt.Println(isospace.Stdout)
												}

											} else {

												// Iso defined in YAML exist
												// but Currently space is not binded to any Iso
												// This is a change in request to remove isolation segment

												fmt.Println("+ isolation segment", Orgs.Org.Spaces[j].IsolationSeg)
												fmt.Println("- isolation segment", isoexistingdetails.Resources[0].Name)
												fmt.Println("Currently doesn't support changing Isolation Segments, Please change isolation segment assigned to space manually")
											}
										}
									}

									// Pulling Space GUID
									var getspacename *exec.Cmd
									if ostype == "windows" {
										path := "\"" + "/v3/spaces?names=" + Orgs.Org.Spaces[j].Name + "&organization_guids=" + orgguid + "\""
										getspacename = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spacedetails_name.json")

									} else {
										path := "/v3/spaces?names=" + Orgs.Org.Spaces[j].Name + "&organization_guids=" + orgguid
										getspacename = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spacedetails_name.json")

									}
									err = getspacename.Run()
									if err == nil {
										//	fmt.Println(getorg, getorg.Stdout, getorg.Stderr)
									} else {
										fmt.Println("err", getspacename, getspacename.Stdout, getspacename.Stderr)
									}
									fileSpaceNameJson, err = ioutil.ReadFile("CreateOrUpdateSpaces_spacedetails_name.json")
									if err != nil {
										fmt.Println(err)
									}
									var spacedetailsname SpaceListJson
									if err := json.Unmarshal(fileSpaceNameJson, &spacedetailsname); err != nil {
										panic(err)
									}
									SpaceGuidPull = spacedetailsname.Resources[0].GUID
								}
							} else if SpaceStateGuidLen == 0 && SpaceStateNameLen != 0 {
								fmt.Println("Missing State file, Please use org-init then space-init function to create state files")
							} else if SpaceStateGuidLen == 0 && SpaceStateNameLen == 0 {

								//fmt.Println("Creating Spaces")
								fmt.Println("+ ", Orgs.Org.Spaces[j].Name)
								CreateSpace := exec.Command("cf", "create-space", Orgs.Org.Spaces[j].Name, "-o", Orgs.Org.Name)
								if _, err := CreateSpace.Output(); err != nil {
									fmt.Println("command: ", CreateSpace)
									fmt.Println("Err: ", CreateSpace.Stdout, err)
								} else {
									fmt.Println(CreateSpace.Stdout)
									if Orgs.Org.Spaces[j].IsolationSeg != "" {
										fmt.Println("Enabling Space Isolation Segment")
										fmt.Println("SegName: ", Orgs.Org.Spaces[j].IsolationSeg)
										fmt.Println("+ isolation segment", Orgs.Org.Spaces[j].IsolationSeg)
										iso := exec.Command("cf", "enable-org-isolation", Orgs.Org.Name, Orgs.Org.Spaces[j].IsolationSeg)
										if _, err := iso.Output(); err != nil {
											fmt.Println("command: ", iso)
											fmt.Println("Err: ", iso.Stdout, err)
										} else {
											//fmt.Println("command: ", iso)
											fmt.Println(iso.Stdout)
										}
										fmt.Println("Attaching to space")
										isospace := exec.Command("cf", "set-space-isolation-segment", Orgs.Org.Spaces[j].Name, Orgs.Org.Spaces[j].IsolationSeg)
										if _, err := isospace.Output(); err != nil {
											fmt.Println("command: ", isospace)
											fmt.Println("Err: ", isospace.Stdout, err)
										} else {
											//fmt.Println("command: ", isospace)
											fmt.Println(isospace.Stdout)
										}
									} else {
										fmt.Println("No Isolation Segment Provided, Will be attached to Org default, if it exist")
									}

									// Pulling Space GUID
									var getspacename *exec.Cmd
									if ostype == "windows" {
										path := "\"" + "/v3/spaces?names=" + Orgs.Org.Spaces[j].Name + "&organization_guids=" + orgguid + "\""
										getspacename = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spacedetails_name.json")

									} else {
										path := "/v3/spaces?names=" + Orgs.Org.Spaces[j].Name + "&organization_guids=" + orgguid
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
									SpaceGuidPull = spacedetailsname.Resources[0].GUID
								}
							}

							// Creating state file
							if SpaceStateGuidLen != 0 {

								orgguid := orgdetails.Resources[0].GUID

								type SpaceState struct {
									Org          string
									OrgGuid      string
									OldSpaceName string
									NewSpaceName string
									SpaceGuid    string
								}

								values := SpaceState{Org: Orgs.Org.Name, OrgGuid: orgguid, OldSpaceName: Orgs.Org.Spaces[j].Name, NewSpaceName: Orgs.Org.Spaces[j].Name, SpaceGuid: SpaceGuidPull}

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

							} else if SpaceStateGuidLen == 0 && SpaceStateNameLen != 0 {

							} else if SpaceStateGuidLen == 0 && SpaceStateNameLen == 0 {

								orgguid := orgdetails.Resources[0].GUID

								type SpaceState struct {
									Org          string
									OrgGuid      string
									OldSpaceName string
									NewSpaceName string
									SpaceGuid    string
								}

								values := SpaceState{Org: Orgs.Org.Name, OrgGuid: orgguid, OldSpaceName: Orgs.Org.Spaces[j].Name, NewSpaceName: Orgs.Org.Spaces[j].Name, SpaceGuid: SpaceGuidPull}

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

						var getspacelit *exec.Cmd
						if ostype == "windows" {
							path := "\"" + "/v3/spaces?organization_guids=" + orgguid + "\""
							getspacelit = exec.Command("powershell", "-command", "cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spacelist.json")

						} else {
							path := "/v3/spaces?organization_guids=" + orgguid
							getspacelit = exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "CreateOrUpdateSpaces_spacelist.json")

						}
						err := getspacelit.Run()
						if err == nil {
							//fmt.Println(getspace, getspace.Stdout, getspace.Stderr)
						} else {
							fmt.Println("err", getspacelit, getspacelit.Stdout, getspacelit.Stderr)
						}

						fileSpaceJson, err := ioutil.ReadFile("CreateOrUpdateSpaces_spacelist.json")
						if err != nil {
							fmt.Println(err)
						}
						var spacedets SpaceListJson
						if err := json.Unmarshal(fileSpaceJson, &spacedets); err != nil {
							panic(err)
						}

						noofspace := len(spacedets.Resources)
						fmt.Println("Spaces: ")
						for s := 0; s < noofspace; s++ {
							fmt.Println(" ", spacedets.Resources[s].Name)
						}
					} else {
						fmt.Println("err: ", target.Stdout, target.Stderr)
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
	return err
}
