package delorauditserviceaccess

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

type pullserviceaccesslistjsin struct {
	Pagination struct {
		TotalResults int `yaml:"total_results"`
		TotalPages   int `yaml:"total_pages"`
		First        struct {
			Href string `yaml:"href"`
		} `yaml:"first"`
		Last struct {
			Href string `yaml:"href"`
		} `yaml:"last"`
		Next     interface{} `yaml:"next"`
		Previous interface{} `yaml:"previous"`
	} `yaml:"pagination"`
	Resources []struct {
		GUID            string        `yaml:"guid"`
		CreatedAt       time.Time     `yaml:"created_at"`
		UpdatedAt       time.Time     `yaml:"updated_at"`
		Name            string        `yaml:"name"`
		VisibilityType  string        `yaml:"visibility_type"`
		Available       bool          `yaml:"available"`
		Free            bool          `yaml:"free"`
		Costs           []interface{} `yaml:"costs"`
		Description     string        `yaml:"description"`
		MaintenanceInfo struct {
			Version     string `yaml:"version"`
			Description string `yaml:"description"`
		} `yaml:"maintenance_info"`
		BrokerCatalog struct {
			ID       string `yaml:"id"`
			Metadata struct {
				Bullets     []string `yaml:"bullets"`
				DisplayName string   `yaml:"displayName"`
			} `yaml:"metadata"`
			MaximumPollingDuration interface{} `yaml:"maximum_polling_duration"`
			Features               struct {
				Bindable       bool `yaml:"bindable"`
				PlanUpdateable bool `yaml:"plan_updateable"`
			} `yaml:"features"`
		} `yaml:"broker_catalog"`
		Schemas struct {
			ServiceInstance struct {
				Create struct {
					Parameters struct {
					} `yaml:"parameters"`
				} `yaml:"create"`
				Update struct {
					Parameters struct {
					} `yaml:"parameters"`
				} `yaml:"update"`
			} `yaml:"service_instance"`
			ServiceBinding struct {
				Create struct {
					Parameters struct {
					} `yaml:"parameters"`
				} `yaml:"create"`
			} `yaml:"service_binding"`
		} `yaml:"schemas"`
		Relationships struct {
			ServiceOffering struct {
				Data struct {
					GUID string `yaml:"guid"`
				} `yaml:"data"`
			} `yaml:"service_offering"`
		} `yaml:"relationships"`
		Metadata struct {
			Labels struct {
			} `yaml:"labels"`
			Annotations struct {
			} `yaml:"annotations"`
		} `yaml:"metadata"`
		Links struct {
			Self struct {
				Href string `yaml:"href"`
			} `yaml:"self"`
			ServiceOffering struct {
				Href string `yaml:"href"`
			} `yaml:"service_offering"`
			Visibility struct {
				Href string `yaml:"href"`
			} `yaml:"visibility"`
		} `yaml:"links"`
	} `yaml:"resources"`
}

type guidval struct {
	GUID      string    `yaml:"guid"`
	CreatedAt time.Time `yaml:"created_at"`
	UpdatedAt time.Time `yaml:"updated_at"`
	Name      string    `yaml:"name"`
	Apps      struct {
		TotalMemoryInMb      int         `yaml:"total_memory_in_mb"`
		PerProcessMemoryInMb interface{} `yaml:"per_process_memory_in_mb"`
		TotalInstances       int         `yaml:"total_instances"`
		PerAppTasks          interface{} `yaml:"per_app_tasks"`
	} `yaml:"apps"`
	Services struct {
		PaidServicesAllowed   bool        `yaml:"paid_services_allowed"`
		TotalServiceInstances int         `yaml:"total_service_instances"`
		TotalServiceKeys      interface{} `yaml:"total_service_keys"`
	} `yaml:"services"`
	Routes struct {
		TotalRoutes        interface{} `yaml:"total_routes"`
		TotalReservedPorts int         `yaml:"total_reserved_ports"`
	} `yaml:"routes"`
	Domains struct {
		TotalDomains interface{} `yaml:"total_domains"`
	} `yaml:"domains"`
	Relationships struct {
		Organizations struct {
			Data []struct {
				GUID string `yaml:"guid"`
			} `yaml:"data"`
		} `yaml:"organizations"`
	} `yaml:"relationships"`
	Links struct {
		Self struct {
			Href string `yaml:"href"`
		} `yaml:"self"`
	} `yaml:"links"`
}
type QuotaDetailsJson struct {
	GUID      string    `yaml:"guid"`
	CreatedAt time.Time `yaml:"created_at"`
	UpdatedAt time.Time `yaml:"updated_at"`
	Name      string    `yaml:"name"`
	Apps      struct {
		TotalMemoryInMb      int         `yaml:"total_memory_in_mb"`
		PerProcessMemoryInMb interface{} `yaml:"per_process_memory_in_mb"`
		TotalInstances       int         `yaml:"total_instances"`
		PerAppTasks          interface{} `yaml:"per_app_tasks"`
	} `yaml:"apps"`
	Services struct {
		PaidServicesAllowed   bool        `yaml:"paid_services_allowed"`
		TotalServiceInstances int         `yaml:"total_service_instances"`
		TotalServiceKeys      interface{} `yaml:"total_service_keys"`
	} `yaml:"services"`
	Routes struct {
		TotalRoutes        interface{} `yaml:"total_routes"`
		TotalReservedPorts int         `yaml:"total_reserved_ports"`
	} `yaml:"routes"`
	Domains struct {
		TotalDomains interface{} `yaml:"total_domains"`
	} `yaml:"domains"`
	Relationships struct {
		Organizations struct {
			Data []struct {
				GUID string `yaml:"guid"`
			} `yaml:"data"`
		} `yaml:"organizations"`
	} `yaml:"relationships"`
	Links struct {
		Self struct {
			Href string `yaml:"href"`
		} `yaml:"self"`
	} `yaml:"links"`
}
type quotaguid struct {
	GUID          string    `yaml:"guid"`
	CreatedAt     time.Time `yaml:"created_at"`
	UpdatedAt     time.Time `yaml:"updated_at"`
	Name          string    `yaml:"name"`
	Suspended     bool      `yaml:"suspended"`
	Relationships struct {
		Quota struct {
			Data struct {
				GUID string `yaml:"guid"`
			} `yaml:"data"`
		} `yaml:"quota"`
	} `yaml:"relationships"`
	Metadata struct {
		Labels struct {
		} `yaml:"labels"`
		Annotations struct {
		} `yaml:"annotations"`
	} `yaml:"metadata"`
	Links struct {
		Self struct {
			Href string `yaml:"href"`
		} `yaml:"self"`
		Domains struct {
			Href string `yaml:"href"`
		} `yaml:"domains"`
		DefaultDomain struct {
			Href string `yaml:"href"`
		} `yaml:"default_domain"`
		Quota struct {
			Href string `yaml:"href"`
		} `yaml:"quota"`
	} `yaml:"links"`
}
type ServiceVisibilityAccess struct {
	Type          string `yaml:"type"`
	Organizations []struct {
		GUID string `yaml:"guid"`
		Name string `yaml:"name"`
	} `yaml:"organizations"`
}
type List struct {
	OrgList []struct {
		Name  string `yaml:"Name"`
		Quota string `yaml:"Quota"`
	} `yaml:"OrgList"`
	Audit string `yaml:"Audit"`
}
type Quotalist struct {
	Quota []struct {
		Name                 string `yaml:"Name"`
		MemoryLimit          string `yaml:"memory_limit"`
		AllowPaidPlans       bool   `yaml:"allow_paid_plans"`
		AppInstanceLimit     string `yaml:"app_instance_limit"`
		ServiceInstanceLimit string `yaml:"service_instance_limit"`
		ServiceAccess        struct {
			MySQL []string `yaml:"MySQL"`
			//Shared_Redis      []string `yaml:"Shared_Redis"`
			OnDemand_Redis    []string `yaml:"OnDemand_Redis"`
			OnDemand_RabbitMQ []string `yaml:"OnDemand_RabbitMQ"`
			//Scheduler         []string `yaml:"Scheduler"`
			//ConfigServer      []string `yaml:"ConfigServer"`
			//ServiceRegistry   []string `yaml:"ServiceRegistry"`
		} `yaml:"ServiceAccess"`
	} `yaml:"quota"`
	Audit              string `yaml:"Audit"`
	ServiceAccessAudit string `yaml:"ServiceAccessAudit"`
}
type ProtectedList struct {
	Org                         []string `yaml:"Org"`
	Quota                       []string `yaml:"quota"`
	DefaultRunningSecurityGroup string   `yaml:"DefaultRunningSecurityGroup"`
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
type ServiceListJSON struct {
	Pagination struct {
		TotalResults int `yaml:"total_results"`
		TotalPages   int `yaml:"total_pages"`
		First        struct {
			Href string `yaml:"href"`
		} `yaml:"first"`
		Last struct {
			Href string `yaml:"href"`
		} `yaml:"last"`
		Next     interface{} `yaml:"next"`
		Previous interface{} `yaml:"previous"`
	} `yaml:"pagination"`
	Resources []struct {
		GUID            string        `yaml:"guid"`
		CreatedAt       time.Time     `yaml:"created_at"`
		UpdatedAt       time.Time     `yaml:"updated_at"`
		Name            string        `yaml:"name"`
		VisibilityType  string        `yaml:"visibility_type"`
		Available       bool          `yaml:"available"`
		Free            bool          `yaml:"free"`
		Costs           []interface{} `yaml:"costs"`
		Description     string        `yaml:"description"`
		MaintenanceInfo struct {
			Version     string `yaml:"version"`
			Description string `yaml:"description"`
		} `yaml:"maintenance_info"`
		BrokerCatalog struct {
			ID       string `yaml:"id"`
			Metadata struct {
				Bullets     []string `yaml:"bullets"`
				DisplayName string   `yaml:"displayName"`
			} `yaml:"metadata"`
			MaximumPollingDuration interface{} `yaml:"maximum_polling_duration"`
			Features               struct {
				Bindable       bool `yaml:"bindable"`
				PlanUpdateable bool `yaml:"plan_updateable"`
			} `yaml:"features"`
		} `yaml:"broker_catalog"`
		Schemas struct {
			ServiceInstance struct {
				Create struct {
					Parameters struct {
					} `yaml:"parameters"`
				} `yaml:"create"`
				Update struct {
					Parameters struct {
					} `yaml:"parameters"`
				} `yaml:"update"`
			} `yaml:"service_instance"`
			ServiceBinding struct {
				Create struct {
					Parameters struct {
					} `yaml:"parameters"`
				} `yaml:"create"`
			} `yaml:"service_binding"`
		} `yaml:"schemas"`
		Relationships struct {
			ServiceOffering struct {
				Data struct {
					GUID string `yaml:"guid"`
				} `yaml:"data"`
			} `yaml:"service_offering"`
		} `yaml:"relationships"`
		Metadata struct {
			Labels struct {
			} `yaml:"labels"`
			Annotations struct {
			} `yaml:"annotations"`
		} `yaml:"metadata"`
		Links struct {
			Self struct {
				Href string `yaml:"href"`
			} `yaml:"self"`
			ServiceOffering struct {
				Href string `yaml:"href"`
			} `yaml:"service_offering"`
			Visibility struct {
				Href string `yaml:"href"`
			} `yaml:"visibility"`
		} `yaml:"links"`
	} `yaml:"resources"`
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

func DeleteorauditServiceAccess(clustername string, cpath string, ostype string) error {

	var Quotas Quotalist
	var GitOrgList GitList
	var InitClusterConfigVals InitClusterConfigVals
	var ListYml string
	var LenList int
	var list List
	//var gitlist GitList

	QuotaYml := cpath + "/" + clustername + "/Quota.yml"
	fileQuotaYml, err := ioutil.ReadFile(QuotaYml)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal([]byte(fileQuotaYml), &Quotas)
	if err != nil {
		panic(err)
	}

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
		//fmt.Println("abc")
	} else {
		ListYml = cpath + "/" + clustername + "/GitOrgsList.yml"
		fileOrgYml, err := ioutil.ReadFile(ListYml)
		if err != nil {
			fmt.Println(err)
		}
		err = yaml.Unmarshal([]byte(fileOrgYml), &GitOrgList)
		if err != nil {
			panic(err)
		}
		LenList = len(GitOrgList.OrgList)
		//fmt.Println("123")
	}

	//LenGitOrgList := len(GitOrgList.OrgList)
	LenQuota := len(Quotas.Quota)

	for i := 0; i < LenList; i++ {

		var count, totalcount int

		var ProtectedOrgs ProtectedList
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

		for p := 0; p < LenProtectedOrgs; p++ {
			if ProtectedOrgs.Org[p] == GitOrgList.OrgList[i].Name {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {

			// pulling from quota yml

			for p := 0; p < LenQuota; p++ {

				ServiceAccessAudit := Quotas.ServiceAccessAudit

				// checking if quota name from quota yml to gitrepo org quota
				// check if passes
				if strings.Trim(Quotas.Quota[p].Name, "") == strings.Trim(GitOrgList.OrgList[i].Quota, "") {

					//fmt.Println("part 1")
					//fmt.Println("Org from GitOrgList: ", GitOrgList.OrgList[i].Name)
					//fmt.Println("Quota from GitOrgList: ", GitOrgList.OrgList[i].Quota)
					//fmt.Println("Quota from QuotaLists: ", 	Quotas.Quota[p].Name)

					fmt.Println(" ")
					fmt.Println("MySql Plans to validate for quota: ", Quotas.Quota[p].ServiceAccess.MySQL, Quotas.Quota[p].Name)
					fmt.Println("Redis On-demand Plans to validate for quota: ", Quotas.Quota[p].ServiceAccess.OnDemand_Redis, Quotas.Quota[p].Name)
					fmt.Println("RabbitMQ Plans to validate for quota: ", Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ, Quotas.Quota[p].Name)

					fmt.Println(" ")
					fmt.Println("Validating MySQL Plans")
					fmt.Println(" ")

					// pulling service details

					LenMysql := len(Quotas.Quota[p].ServiceAccess.MySQL)
					for ms := 0; ms < LenMysql; ms++ {

						var getserviceguid *exec.Cmd
						if ostype == "windows" {
							path := "\"" + "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.MySQL[ms] + "&service_broker_names=dedicated-mysql-broker" + "\""
							getserviceguid = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "Delete_services.json")
						} else {
							path := "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.MySQL[ms] + "&service_broker_names=dedicated-mysql-broker"
							getserviceguid = exec.Command("cf", "curl", path, "--output", "Delete_services.json")
						}

						if _, err := getserviceguid.Output(); err != nil {
							fmt.Println("command: ", getserviceguid)
							fmt.Println("Err: ", getserviceguid.Stdout, err)
						} else {
							fileServiceJson, err := ioutil.ReadFile("Delete_services.json")
							if err != nil {
								fmt.Println(err)
							}
							// pull details from serviceplan - quota/mysql broker
							var ServiceJson ServiceListJSON
							if err := json.Unmarshal(fileServiceJson, &ServiceJson); err != nil {
								panic(err)
							}
							fmt.Println("MySQL Service Plan & Quota: ", ServiceJson.Resources[0].Name, " & ", Quotas.Quota[p].Name)
							//fmt.Println("Service Guid: ", ServiceJson.Resources[0].GUID)

							// using service-plan guid to pull visibility
							var getservicevisibility *exec.Cmd
							if ostype == "windows" {
								path := "\"" + "/v3/service_plans/" + ServiceJson.Resources[0].GUID + "/visibility" + "\""
								getservicevisibility = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "servicesvisibility.json")
							} else {
								path := "/v3/service_plans/?names=" + ServiceJson.Resources[0].GUID + "/visibility"
								getservicevisibility = exec.Command("cf", "curl", path, "--output", "servicesvisibility.json")
							}

							if len(ServiceJson.Resources) == 0 {

								fmt.Println("Ondemand MySql plan '" + Quotas.Quota[p].ServiceAccess.MySQL[ms] + "' doesn't exist")
								fmt.Println("Please correct the plan in Quota, skipping audit")
								break
							} else {
								if _, err := getservicevisibility.Output(); err != nil {
									fmt.Println("command: ", getservicevisibility)
									fmt.Println("Err: ", getservicevisibility.Stdout, err)
								} else {
									fileServiceVisibilityJson, err := ioutil.ReadFile("servicesvisibility.json")
									if err != nil {
										fmt.Println(err)
									}

									var ServiceVisibilityJson ServiceVisibilityAccess
									if err := json.Unmarshal(fileServiceVisibilityJson, &ServiceVisibilityJson); err != nil {
										panic(err)
									}
									// From service-visibility pulling list org binded to this ServicePlans
									//fmt.Println("")

									fmt.Println("Orgs binded to the MySQL Service-plan: ", ServiceJson.Resources[0].Name)
									orglen := len(ServiceVisibilityJson.Organizations)
									for i := 0; i < orglen; i++ {
										fmt.Println(" - " + ServiceVisibilityJson.Organizations[i].Name)
									}

									for listorgs := 0; listorgs < len(ServiceVisibilityJson.Organizations); listorgs++ {

										fmt.Println("")
										fmt.Println("Org: ", ServiceVisibilityJson.Organizations[listorgs].Name)
										//	fmt.Println("Org Guid binded to service: ", ServiceVisibilityJson.Organizations[listorgs].GUID)

										var count, totalcount int
										//fmt.Println(" ")
										//fmt.Println("Org: ", ServiceVisibilityJson.Organizations[listorgs].Name)
										for p := 0; p < LenProtectedOrgs; p++ {
											//	fmt.Println("Protected Org: ", ProtectedOrgs.Org[p], ",", list.OrgList[i])
											if ProtectedOrgs.Org[p] == ServiceVisibilityJson.Organizations[listorgs].Name {
												count = 1
											} else {
												count = 0
											}
											totalcount = totalcount + count
										}
										if totalcount == 0 {
											var getorgdets *exec.Cmd
											if ostype == "windows" {
												path := "\"" + "/v3/organizations/" + ServiceVisibilityJson.Organizations[listorgs].GUID + "\""
												getorgdets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingvalue.json")
											} else {
												path := "/v3/organizations/" + ServiceVisibilityJson.Organizations[listorgs].GUID
												getorgdets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparingvalue.json")
											}
											if _, err := getorgdets.Output(); err != nil {
												//fmt.Println("command: ", getorgdets)
												fmt.Println("Err: ", getorgdets.Stdout, err)
											} else {
												//fmt.Println("command: ", getorgdets)
												filegetquotaquid, err := ioutil.ReadFile("quotadetailsforcomparingvalue.json")
												if err != nil {
													fmt.Println(err)
												}
												var getquotaguid quotaguid
												if err := json.Unmarshal(filegetquotaquid, &getquotaguid); err != nil {
													panic(err)
												}
												// Pulling Quota from Quota GUID
												//fmt.Println("Quota Guid: ", getquotaguid.Relationships.Quota.Data.GUID)

												var getqutadets *exec.Cmd
												if ostype == "windows" {
													path := "\"" + "/v3/organization_quotas/" + getquotaguid.Relationships.Quota.Data.GUID + "\""
													getqutadets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
												} else {
													path := "/v3/organization_quotas/" + getquotaguid.Relationships.Quota.Data.GUID
													getqutadets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
												}

												if _, err := getqutadets.Output(); err != nil {
													//fmt.Println("command: ", getqutadets)
													fmt.Println("Err: ", getqutadets.Stdout, err)
												} else {
													//fmt.Println("command: ", getqutadets)
													filegetquotaguidvalue, err := ioutil.ReadFile("quotadetailsforcomparingval.json")
													if err != nil {
														fmt.Println(err)
													}

													var getquotaguidval guidval
													if err := json.Unmarshal(filegetquotaguidvalue, &getquotaguidval); err != nil {
														panic(err)
													}
													//fmt.Println("Quota Name pulled from org guid: ", getquotaguidval.Name)
													/////

													var pullseracclist *exec.Cmd

													if ostype == "windows" {
														path := "\"" + "/v3/service_plans/?organization_guids=" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=dedicated-mysql-broker" + "\""
														pullseracclist = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "getseracclist.json")
													} else {
														path := "/v3/organizations_quotas" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=dedicated-mysql-broker"
														pullseracclist = exec.Command("cf", "curl", path, "--output", "getseracclist.json")
													}

													if _, err := pullseracclist.Output(); err != nil {
														fmt.Println("command: ", pullseracclist)
														fmt.Println("Err: ", pullseracclist.Stdout, err)
													} else {
														filegetseraccdets, err := ioutil.ReadFile("getseracclist.json")
														if err != nil {
															fmt.Println(err)
														}

														var getacclist pullserviceaccesslistjsin
														if err := json.Unmarshal(filegetseraccdets, &getacclist); err != nil {
															panic(err)
														}
														totallist := len(getacclist.Resources)
														totalacclist := len(Quotas.Quota[p].ServiceAccess.MySQL)

														//fmt.Println("Total MySql Plans Org has access: ", totallist)
														//fmt.Println("Total MySql Plans Org should have access: ", totalacclist)

														for list := 0; list < totallist; list++ {
															var totalcount int
															var count int
															for list2 := 0; list2 < totalacclist; list2++ {
																if strings.Trim(getacclist.Resources[list].Name, "") == strings.Trim(Quotas.Quota[p].ServiceAccess.MySQL[list2], "") {
																	//fmt.Println("Hi1")
																	//fmt.Println("part 1", getacclist.Resources[list].Name)
																	//fmt.Println("part 1", Quotas.Quota[p].ServiceAccess.MySQL[list2])
																	count = 1
																} else {
																	//fmt.Println("Hi1")
																	//fmt.Println("part 2", getacclist.Resources[list].Name)
																	//fmt.Println("part 2", Quotas.Quota[p].ServiceAccess.MySQL[list2])
																	count = 0
																}
																totalcount = totalcount + count
															}
															//	fmt.Println(totalcount)
															if totalcount == 0 {
																//fmt.Println(ServiceAccessAudit)
																if ServiceAccessAudit == "list" {
																	fmt.Println("plan access to be revoked manually for org: ", getacclist.Resources[list].Name, GitOrgList.OrgList[i].Name)
																	//fmt.Println(" ")
																} else if ServiceAccessAudit == "delete" {
																	fmt.Println("plan access getting revoked for org: ", getacclist.Resources[list].Name, GitOrgList.OrgList[i].Name)

																	var getserviceguid *exec.Cmd
																	if ostype == "windows" {
																		//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
																		getserviceguid = exec.Command("powershell", "-command", "cf", "disable-service-access", strings.TrimSpace("p.mysql"), "-p", getacclist.Resources[list].Name, "-o", GitOrgList.OrgList[i].Name)
																	} else {
																		//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
																		getserviceguid = exec.Command("cf", "disable-service-access", strings.TrimSpace("p.mysql"), "-p", getacclist.Resources[list].Name, "-o", GitOrgList.OrgList[i].Name)
																	}
																	err = getserviceguid.Run()
																	if err == nil {
																	} else {
																		fmt.Println("err", getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
																	}
																} else {
																	fmt.Println("please provide valid input for ServiceAccessAudit flag")
																}
															} else {

																totalcount = 0
															}

														}

													}
												}
											}
										} else {
											fmt.Println(ServiceVisibilityJson.Organizations[listorgs].Name + " is a protected Org")
										}
									}
								}
							}
						}
					}

					fmt.Println(" ")
					fmt.Println("Validating Redis OnDemand Plans")
					fmt.Println(" ")

					LenRedis := len(Quotas.Quota[p].ServiceAccess.OnDemand_Redis)
					for ms := 0; ms < LenRedis; ms++ {

						var getserviceguid *exec.Cmd
						if ostype == "windows" {
							path := "\"" + "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.OnDemand_Redis[ms] + "&service_broker_names=redis-odb" + "\""
							getserviceguid = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "Delete_services.json")
						} else {
							path := "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.OnDemand_Redis[ms] + "&service_broker_names=redis-odb"
							getserviceguid = exec.Command("cf", "curl", path, "--output", "Delete_services.json")
						}

						if _, err := getserviceguid.Output(); err != nil {
							fmt.Println("command: ", getserviceguid)
							fmt.Println("Err: ", getserviceguid.Stdout, err)
						} else {
							fileServiceJson, err := ioutil.ReadFile("Delete_services.json")
							if err != nil {
								fmt.Println(err)
							}
							// pull details from serviceplan - quota/mysql broker
							var ServiceJson ServiceListJSON
							if err := json.Unmarshal(fileServiceJson, &ServiceJson); err != nil {
								panic(err)
							}
							if len(ServiceJson.Resources) == 0 {

								fmt.Println("Ondemand Redis plan '" + Quotas.Quota[p].ServiceAccess.OnDemand_Redis[ms] + "' doesn't exist")
								fmt.Println("Please correct the plan in Quota, skipping audit")
								break
							} else {

								fmt.Println("Redis Service Plan & Quota: ", ServiceJson.Resources[0].Name, " & ", Quotas.Quota[p].Name)
								var getservicevisibility *exec.Cmd
								if ostype == "windows" {
									path := "\"" + "/v3/service_plans/" + ServiceJson.Resources[0].GUID + "/visibility" + "\""
									getservicevisibility = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "servicesvisibility.json")
								} else {
									path := "/v3/service_plans/?names=" + ServiceJson.Resources[0].GUID + "/visibility"
									getservicevisibility = exec.Command("cf", "curl", path, "--output", "servicesvisibility.json")
								}
								if _, err := getservicevisibility.Output(); err != nil {
									fmt.Println("command: ", getservicevisibility)
									fmt.Println("Err: ", getservicevisibility.Stdout, err)
								} else {
									fileServiceVisibilityJson, err := ioutil.ReadFile("servicesvisibility.json")
									if err != nil {
										fmt.Println(err)
									}

									var ServiceVisibilityJson ServiceVisibilityAccess
									if err := json.Unmarshal(fileServiceVisibilityJson, &ServiceVisibilityJson); err != nil {
										panic(err)
									}
									// From service-visibility pulling list org binded to this ServicePlans
									//fmt.Println("")

									fmt.Println("Orgs binded to the Redis Ondemand Service-plan: ", ServiceJson.Resources[0].Name)
									orglen := len(ServiceVisibilityJson.Organizations)
									for i := 0; i < orglen; i++ {
										fmt.Println(" - " + ServiceVisibilityJson.Organizations[i].Name)
									}

									for listorgs := 0; listorgs < len(ServiceVisibilityJson.Organizations); listorgs++ {

										fmt.Println("")
										fmt.Println("Org: ", ServiceVisibilityJson.Organizations[listorgs].Name)
										//	fmt.Println("Org Guid binded to service: ", ServiceVisibilityJson.Organizations[listorgs].GUID)

										var count, totalcount int
										//fmt.Println(" ")
										//fmt.Println("Org: ", ServiceVisibilityJson.Organizations[listorgs].Name)
										for p := 0; p < LenProtectedOrgs; p++ {
											//	fmt.Println("Protected Org: ", ProtectedOrgs.Org[p], ",", list.OrgList[i])
											if ProtectedOrgs.Org[p] == ServiceVisibilityJson.Organizations[listorgs].Name {
												count = 1
											} else {
												count = 0
											}
											totalcount = totalcount + count
										}
										if totalcount == 0 {
											var getorgdets *exec.Cmd
											if ostype == "windows" {
												path := "\"" + "/v3/organizations/" + ServiceVisibilityJson.Organizations[listorgs].GUID + "\""
												getorgdets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingvalue.json")
											} else {
												path := "/v3/organizations/" + ServiceVisibilityJson.Organizations[listorgs].GUID
												getorgdets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparingvalue.json")
											}
											if _, err := getorgdets.Output(); err != nil {
												//fmt.Println("command: ", getorgdets)
												fmt.Println("Err: ", getorgdets.Stdout, err)
											} else {
												//fmt.Println("command: ", getorgdets)
												filegetquotaquid, err := ioutil.ReadFile("quotadetailsforcomparingvalue.json")
												if err != nil {
													fmt.Println(err)
												}
												var getquotaguid quotaguid
												if err := json.Unmarshal(filegetquotaquid, &getquotaguid); err != nil {
													panic(err)
												}
												// Pulling Quota from Quota GUID
												//fmt.Println("Quota Guid: ", getquotaguid.Relationships.Quota.Data.GUID)

												var getqutadets *exec.Cmd
												if ostype == "windows" {
													path := "\"" + "/v3/organization_quotas/" + getquotaguid.Relationships.Quota.Data.GUID + "\""
													getqutadets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
												} else {
													path := "/v3/organization_quotas/" + getquotaguid.Relationships.Quota.Data.GUID
													getqutadets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
												}

												if _, err := getqutadets.Output(); err != nil {
													//fmt.Println("command: ", getqutadets)
													fmt.Println("Err: ", getqutadets.Stdout, err)
												} else {
													//fmt.Println("command: ", getqutadets)
													filegetquotaguidvalue, err := ioutil.ReadFile("quotadetailsforcomparingval.json")
													if err != nil {
														fmt.Println(err)
													}

													var getquotaguidval guidval
													if err := json.Unmarshal(filegetquotaguidvalue, &getquotaguidval); err != nil {
														panic(err)
													}
													//fmt.Println("Quota Name pulled from org guid: ", getquotaguidval.Name)
													/////

													var pullseracclist *exec.Cmd

													if ostype == "windows" {
														path := "\"" + "/v3/service_plans/?organization_guids=" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=redis-odb" + "\""
														pullseracclist = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "getseracclist.json")
													} else {
														path := "/v3/organizations_quotas" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=redis-odb"
														pullseracclist = exec.Command("cf", "curl", path, "--output", "getseracclist.json")
													}

													if _, err := pullseracclist.Output(); err != nil {
														fmt.Println("command: ", pullseracclist)
														fmt.Println("Err: ", pullseracclist.Stdout, err)
													} else {
														filegetseraccdets, err := ioutil.ReadFile("getseracclist.json")
														if err != nil {
															fmt.Println(err)
														}

														var getacclist pullserviceaccesslistjsin
														if err := json.Unmarshal(filegetseraccdets, &getacclist); err != nil {
															panic(err)
														}
														totallist := len(getacclist.Resources)
														totalacclist := len(Quotas.Quota[p].ServiceAccess.OnDemand_Redis)

														//fmt.Println("Total MySql Plans Org has access: ", totallist)
														//fmt.Println("Total MySql Plans Org should have access: ", totalacclist)

														for list := 0; list < totallist; list++ {
															var totalcount int
															var count int
															for list2 := 0; list2 < totalacclist; list2++ {
																if strings.Trim(getacclist.Resources[list].Name, "") == strings.Trim(Quotas.Quota[p].ServiceAccess.OnDemand_Redis[list2], "") {
																	//fmt.Println("Hi1")
																	//fmt.Println("part 1", getacclist.Resources[list].Name)
																	//fmt.Println("part 1", Quotas.Quota[p].ServiceAccess.MySQL[list2])
																	count = 1
																} else {
																	//fmt.Println("Hi1")
																	//fmt.Println("part 2", getacclist.Resources[list].Name)
																	//fmt.Println("part 2", Quotas.Quota[p].ServiceAccess.MySQL[list2])
																	count = 0
																}
																totalcount = totalcount + count
															}
															//	fmt.Println(totalcount)
															if totalcount == 0 {
																//fmt.Println(ServiceAccessAudit)
																if ServiceAccessAudit == "list" {
																	fmt.Println("plan access to be revoked manually for org: ", getacclist.Resources[list].Name, GitOrgList.OrgList[i].Name)
																	//fmt.Println(" ")
																} else if ServiceAccessAudit == "delete" {
																	fmt.Println("plan access getting revoked for org: ", getacclist.Resources[list].Name, GitOrgList.OrgList[i].Name)

																	var getserviceguid *exec.Cmd
																	if ostype == "windows" {
																		//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
																		getserviceguid = exec.Command("powershell", "-command", "cf", "disable-service-access", strings.TrimSpace("p.redis"), "-p", getacclist.Resources[list].Name, "-o", GitOrgList.OrgList[i].Name)
																	} else {
																		//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
																		getserviceguid = exec.Command("cf", "disable-service-access", strings.TrimSpace("p.redis"), "-p", getacclist.Resources[list].Name, "-o", GitOrgList.OrgList[i].Name)
																	}
																	err = getserviceguid.Run()
																	if err == nil {
																	} else {
																		fmt.Println("err", getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
																	}
																} else {
																	fmt.Println("please provide valid input for ServiceAccessAudit flag")
																}
															} else {

																totalcount = 0
															}

														}

													}
												}
											}
										} else {
											fmt.Println(ServiceVisibilityJson.Organizations[listorgs].Name + " is a protected Org")
										}
									}
								}
							}
						}
					}

					fmt.Println(" ")
					fmt.Println("Validating RabbitMQ OnDemand Plans")
					fmt.Println(" ")

					LenRabbitMQ := len(Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ)
					for ms := 0; ms < LenRabbitMQ; ms++ {

						var getserviceguid *exec.Cmd
						if ostype == "windows" {
							path := "\"" + "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ[ms] + "&service_broker_names=rabbitmq-odb" + "\""
							getserviceguid = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "Delete_services.json")
						} else {
							path := "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ[ms] + "&service_broker_names=rabbitmq-odb"
							getserviceguid = exec.Command("cf", "curl", path, "--output", "Delete_services.json")
						}

						if _, err := getserviceguid.Output(); err != nil {
							fmt.Println("command: ", getserviceguid)
							fmt.Println("Err: ", getserviceguid.Stdout, err)
						} else {
							fileServiceJson, err := ioutil.ReadFile("Delete_services.json")
							if err != nil {
								fmt.Println(err)
							}
							// pull details from serviceplan - quota/mysql broker
							var ServiceJson ServiceListJSON
							if err := json.Unmarshal(fileServiceJson, &ServiceJson); err != nil {
								panic(err)
							}
							if len(ServiceJson.Resources) == 0 {

								fmt.Println("Ondemand RabbitMQ plan '" + Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ[ms] + "' doesn't exist")
								fmt.Println("Please correct the plan in Quota, skipping audit")
								break
							} else {

								fmt.Println("Redis Service Plan & Quota: ", ServiceJson.Resources[0].Name, " & ", Quotas.Quota[p].Name)
								var getservicevisibility *exec.Cmd
								if ostype == "windows" {
									path := "\"" + "/v3/service_plans/" + ServiceJson.Resources[0].GUID + "/visibility" + "\""
									getservicevisibility = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "servicesvisibility.json")
								} else {
									path := "/v3/service_plans/?names=" + ServiceJson.Resources[0].GUID + "/visibility"
									getservicevisibility = exec.Command("cf", "curl", path, "--output", "servicesvisibility.json")
								}
								if _, err := getservicevisibility.Output(); err != nil {
									fmt.Println("command: ", getservicevisibility)
									fmt.Println("Err: ", getservicevisibility.Stdout, err)
								} else {
									fileServiceVisibilityJson, err := ioutil.ReadFile("servicesvisibility.json")
									if err != nil {
										fmt.Println(err)
									}

									var ServiceVisibilityJson ServiceVisibilityAccess
									if err := json.Unmarshal(fileServiceVisibilityJson, &ServiceVisibilityJson); err != nil {
										panic(err)
									}
									// From service-visibility pulling list org binded to this ServicePlans
									//fmt.Println("")

									fmt.Println("Orgs binded to the RabbitMQ Ondemand Service-plan: ", ServiceJson.Resources[0].Name)
									orglen := len(ServiceVisibilityJson.Organizations)
									for i := 0; i < orglen; i++ {
										fmt.Println(" - " + ServiceVisibilityJson.Organizations[i].Name)
									}

									for listorgs := 0; listorgs < len(ServiceVisibilityJson.Organizations); listorgs++ {

										fmt.Println("")
										fmt.Println("Org: ", ServiceVisibilityJson.Organizations[listorgs].Name)
										//	fmt.Println("Org Guid binded to service: ", ServiceVisibilityJson.Organizations[listorgs].GUID)

										var count, totalcount int
										//fmt.Println(" ")
										//fmt.Println("Org: ", ServiceVisibilityJson.Organizations[listorgs].Name)
										for p := 0; p < LenProtectedOrgs; p++ {
											//	fmt.Println("Protected Org: ", ProtectedOrgs.Org[p], ",", list.OrgList[i])
											if ProtectedOrgs.Org[p] == ServiceVisibilityJson.Organizations[listorgs].Name {
												count = 1
											} else {
												count = 0
											}
											totalcount = totalcount + count
										}
										if totalcount == 0 {
											var getorgdets *exec.Cmd
											if ostype == "windows" {
												path := "\"" + "/v3/organizations/" + ServiceVisibilityJson.Organizations[listorgs].GUID + "\""
												getorgdets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingvalue.json")
											} else {
												path := "/v3/organizations/" + ServiceVisibilityJson.Organizations[listorgs].GUID
												getorgdets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparingvalue.json")
											}
											if _, err := getorgdets.Output(); err != nil {
												//fmt.Println("command: ", getorgdets)
												fmt.Println("Err: ", getorgdets.Stdout, err)
											} else {
												//fmt.Println("command: ", getorgdets)
												filegetquotaquid, err := ioutil.ReadFile("quotadetailsforcomparingvalue.json")
												if err != nil {
													fmt.Println(err)
												}
												var getquotaguid quotaguid
												if err := json.Unmarshal(filegetquotaquid, &getquotaguid); err != nil {
													panic(err)
												}
												// Pulling Quota from Quota GUID
												//fmt.Println("Quota Guid: ", getquotaguid.Relationships.Quota.Data.GUID)

												var getqutadets *exec.Cmd
												if ostype == "windows" {
													path := "\"" + "/v3/organization_quotas/" + getquotaguid.Relationships.Quota.Data.GUID + "\""
													getqutadets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
												} else {
													path := "/v3/organization_quotas/" + getquotaguid.Relationships.Quota.Data.GUID
													getqutadets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
												}

												if _, err := getqutadets.Output(); err != nil {
													//fmt.Println("command: ", getqutadets)
													fmt.Println("Err: ", getqutadets.Stdout, err)
												} else {
													//fmt.Println("command: ", getqutadets)
													filegetquotaguidvalue, err := ioutil.ReadFile("quotadetailsforcomparingval.json")
													if err != nil {
														fmt.Println(err)
													}

													var getquotaguidval guidval
													if err := json.Unmarshal(filegetquotaguidvalue, &getquotaguidval); err != nil {
														panic(err)
													}
													//fmt.Println("Quota Name pulled from org guid: ", getquotaguidval.Name)
													/////

													var pullseracclist *exec.Cmd

													if ostype == "windows" {
														path := "\"" + "/v3/service_plans/?organization_guids=" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=rabbitmq-odb" + "\""
														pullseracclist = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "getseracclist.json")
													} else {
														path := "/v3/organizations_quotas" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=rabbitmq-odb"
														pullseracclist = exec.Command("cf", "curl", path, "--output", "getseracclist.json")
													}

													if _, err := pullseracclist.Output(); err != nil {
														fmt.Println("command: ", pullseracclist)
														fmt.Println("Err: ", pullseracclist.Stdout, err)
													} else {
														filegetseraccdets, err := ioutil.ReadFile("getseracclist.json")
														if err != nil {
															fmt.Println(err)
														}

														var getacclist pullserviceaccesslistjsin
														if err := json.Unmarshal(filegetseraccdets, &getacclist); err != nil {
															panic(err)
														}
														totallist := len(getacclist.Resources)
														totalacclist := len(Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ)

														//fmt.Println("Total MySql Plans Org has access: ", totallist)
														//fmt.Println("Total MySql Plans Org should have access: ", totalacclist)

														for list := 0; list < totallist; list++ {
															var totalcount int
															var count int
															for list2 := 0; list2 < totalacclist; list2++ {
																if strings.Trim(getacclist.Resources[list].Name, "") == strings.Trim(Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ[list2], "") {
																	//fmt.Println("Hi1")
																	//fmt.Println("part 1", getacclist.Resources[list].Name)
																	//fmt.Println("part 1", Quotas.Quota[p].ServiceAccess.MySQL[list2])
																	count = 1
																} else {
																	//fmt.Println("Hi1")
																	//fmt.Println("part 2", getacclist.Resources[list].Name)
																	//fmt.Println("part 2", Quotas.Quota[p].ServiceAccess.MySQL[list2])
																	count = 0
																}
																totalcount = totalcount + count
															}
															//	fmt.Println(totalcount)
															if totalcount == 0 {
																//fmt.Println(ServiceAccessAudit)
																if ServiceAccessAudit == "list" {
																	fmt.Println("plan access to be revoked manually for org: ", getacclist.Resources[list].Name, GitOrgList.OrgList[i].Name)
																	//fmt.Println(" ")
																} else if ServiceAccessAudit == "delete" {
																	fmt.Println("plan access getting revoked for org: ", getacclist.Resources[list].Name, GitOrgList.OrgList[i].Name)

																	var getserviceguid *exec.Cmd
																	if ostype == "windows" {
																		//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
																		getserviceguid = exec.Command("powershell", "-command", "cf", "disable-service-access", strings.TrimSpace("p.rabbitmq"), "-p", getacclist.Resources[list].Name, "-o", GitOrgList.OrgList[i].Name)
																	} else {
																		//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
																		getserviceguid = exec.Command("cf", "disable-service-access", strings.TrimSpace("p.rabbitmq"), "-p", getacclist.Resources[list].Name, "-o", GitOrgList.OrgList[i].Name)
																	}
																	err = getserviceguid.Run()
																	if err == nil {
																	} else {
																		fmt.Println("err", getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
																	}
																} else {
																	fmt.Println("please provide valid input for ServiceAccessAudit flag")
																}
															} else {

																totalcount = 0
															}

														}

													}
												}
											}
										} else {
											fmt.Println(ServiceVisibilityJson.Organizations[listorgs].Name + " is a protected Org")
										}
									}
								}
							}
						}
					}

				} else {

					//fmt.Println("part 2")
					//fmt.Println("Org from GitOrgList: ", GitOrgList.OrgList[i].Name)
					//fmt.Println("Quota from GitOrgList: ", GitOrgList.OrgList[i].Quota)
					//fmt.Println("Quota from QuotaLists: ", 	Quotas.Quota[p].Name)

					// checking if quota name from quota yml to gitrepo org quota
					// if failed

					//fmt.Println("")
					//fmt.Println("Org : ", GitOrgList.OrgList[i].Name)
					//fmt.Println("Quota GIT : ", GitOrgList.OrgList[i].Quota)
					//fmt.Println("Quota Quota yml: ", Quotas.Quota[p].Name)
					//fmt.Println("Quota is not binded to org in Quota.yml file")
					//fmt.Println("")
				}
			}
		}
	}
	return err
}
