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
	} `yaml:"resources"`
}
type guidval2 struct {
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
	//LenQuota := len(Quotas.Quota)
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

			// pulling from quota from org, not a protected org
			//getorgs := exec.Command("cf", "curl", "/v3/organizations/?names="+GitOrgList.OrgList[i].Name, "--output", "DeleteorAuditOrgs_listorgs.json")

			var getorgdetails *exec.Cmd
			if ostype == "windows" {
				path := "\"" + "/v3/organizations/?names=" + GitOrgList.OrgList[i].Name + "\""
				getorgdetails = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "getorgdetails.json")
			} else {
				path := "/v3/organizations/?names=" + GitOrgList.OrgList[i].Name
				getorgdetails = exec.Command("cf", "curl", path, "--output", "getorgdetails.json")
			}
			if _, err := getorgdetails.Output(); err != nil {
				fmt.Println("command: ", getorgdetails)
				fmt.Println("Err: ", getorgdetails.Stdout, err)
			} else {
				// get quota from above details
				fileOrgJson, err := ioutil.ReadFile("getorgdetails.json")
				if err != nil {
					fmt.Println(err)
				}
				var orgdetails guidval
				if err := json.Unmarshal(fileOrgJson, &orgdetails); err != nil {
					panic(err)
				}
				// pull quota name from guid
				var getquotadetails *exec.Cmd
				if ostype == "windows" {
					path := "\"" + "/v3/organization_quotas/" + orgdetails.Resources[0].Relationships.Quota.Data.GUID + "\""
					getquotadetails = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "getquotadetails.json")
				} else {
					path := "/v3/organization_quotas/" + orgdetails.Resources[0].Relationships.Quota.Data.GUID
					getquotadetails = exec.Command("cf", "curl", path, "--output", "getquotadetails.json")
				}
				if _, err := getquotadetails.Output(); err != nil {
					fmt.Println("command: ", getquotadetails)
					fmt.Println("Err: ", getquotadetails.Stdout, err)
				} else {
					filequotaJson, err := ioutil.ReadFile("getquotadetails.json")
					if err != nil {
						fmt.Println(err)
					}
					var quotaguid quotaguid
					if err := json.Unmarshal(filequotaJson, &quotaguid); err != nil {
						panic(err)
					}
					// pull quota name from guid
					fmt.Println("Org Name: ", GitOrgList.OrgList[i].Name)
					fmt.Println("Quota Org binded to: ", quotaguid.Name)
					fmt.Println("Quota listed in Gitlist.yml: ", GitOrgList.OrgList[i].Quota)
					if strings.Trim(quotaguid.Name, "") == strings.Trim(GitOrgList.OrgList[i].Quota, "") {

						// Pulling Details for OnDemand MySQL
						fmt.Println("")
						fmt.Println("Validating MySQL OnDemand Plans")
						//validating plans in quota yml
						LenQuota := len(Quotas.Quota)
						for p := 0; p < LenQuota; p++ {
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

									if len(ServiceJson.Resources) == 0 {

										fmt.Println("Ondemand MySql plan '" + Quotas.Quota[p].ServiceAccess.MySQL[ms] + "' doesn't exist")
										fmt.Println("Please correct the plan in Quota, skipping audit")
										break
									} else {
										// auditing plans
										fmt.Println("MySQL Service Plan & Quota: ", ServiceJson.Resources[0].Name, " & ", Quotas.Quota[p].Name)
										var pullseracclist *exec.Cmd
										if ostype == "windows" {
											path := "\"" + "/v3/service_plans/?organization_guids=" + orgdetails.Resources[0].GUID + "&service_broker_names=dedicated-mysql-broker" + "\""
											pullseracclist = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "getseracclist.json")
										} else {
											path := "/v3/service_plans/?organization_guids=" + orgdetails.Resources[0].GUID + "&service_broker_names=dedicated-mysql-broker"
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

											var p int
											for i := 0; i < len(Quotas.Quota); i++ {
												if Quotas.Quota[i].Name == quotaguid.Name {
													p = i
												} else {

												}
											}
											totalacclist := len(Quotas.Quota[p].ServiceAccess.MySQL)

											ServiceAccessAudit := Quotas.ServiceAccessAudit

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
													} else if ServiceAccessAudit == "unbind" {
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
							}
						}

						fmt.Println("")
						fmt.Println("Validating Redis OnDemand Plans")
						//validating plans in quota yml
						for p := 0; p < LenQuota; p++ {
							LenMysql := len(Quotas.Quota[p].ServiceAccess.OnDemand_Redis)
							for ms := 0; ms < LenMysql; ms++ {
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
										var pullseracclist *exec.Cmd
										if ostype == "windows" {

											path := "\"" + "/v3/service_plans/?organization_guids=" + orgdetails.Resources[0].GUID + "&service_broker_names=redis-odb" + "\""
											pullseracclist = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "getseracclist.json")
										} else {
											path := "/v3/service_plans/?organization_guids=" + orgdetails.Resources[0].GUID + "&service_broker_names=redis-odb"
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

											var p int
											for i := 0; i < len(Quotas.Quota); i++ {
												if Quotas.Quota[i].Name == quotaguid.Name {
													p = i
												} else {

												}
											}
											totalacclist := len(Quotas.Quota[p].ServiceAccess.OnDemand_Redis)
											ServiceAccessAudit := Quotas.ServiceAccessAudit
											for list := 0; list < totallist; list++ {
												var totalcount int
												var count int
												for list2 := 0; list2 < totalacclist; list2++ {
													if strings.Trim(getacclist.Resources[list].Name, "") == strings.Trim(Quotas.Quota[p].ServiceAccess.OnDemand_Redis[list2], "") {
														count = 1
													} else {
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
													} else if ServiceAccessAudit == "unbind" {
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
															fmt.Println("err", getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)

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
							}
						}

						fmt.Println("")
						fmt.Println("Validating RabbitMQ OnDemand Plans")
						//validating plans in quota yml
						for p := 0; p < LenQuota; p++ {
							LenMysql := len(Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ)
							for ms := 0; ms < LenMysql; ms++ {
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
										fmt.Println("")
										break
									} else {
										fmt.Println("RabbitMQ Service Plan & Quota: ", ServiceJson.Resources[0].Name, " & ", Quotas.Quota[p].Name)
										var pullseracclist *exec.Cmd
										if ostype == "windows" {
											path := "\"" + "/v3/service_plans/?organization_guids=" + orgdetails.Resources[0].GUID + "&service_broker_names=rabbitmq-odb" + "\""
											pullseracclist = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "getseracclist.json")
										} else {
											path := "/v3/service_plans/?organization_guids=" + orgdetails.Resources[0].GUID + "&service_broker_names=rabbitmq-odb"
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

											var p int
											for i := 0; i < len(Quotas.Quota); i++ {
												if Quotas.Quota[i].Name == quotaguid.Name {
													p = i
												} else {

												}
											}
											totalacclist := len(Quotas.Quota[p].ServiceAccess.OnDemand_Redis)
											ServiceAccessAudit := Quotas.ServiceAccessAudit
											for list := 0; list < totallist; list++ {
												var totalcount int
												var count int
												for list2 := 0; list2 < totalacclist; list2++ {
													if strings.Trim(getacclist.Resources[list].Name, "") == strings.Trim(Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ[list2], "") {
														count = 1
													} else {
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
													} else if ServiceAccessAudit == "unbind" {
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
										fmt.Println(" ")

									}
								}
							}
						}

					} else {
						fmt.Println("Skipping validation for org; Quota hasn't been updated")
					}
				}
			}
		} else {
			fmt.Println(GitOrgList.OrgList[i].Name + " is a protected org")
		}
	}
	return err
}
