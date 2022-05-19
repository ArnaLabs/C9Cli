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
			MySQL             []string `yaml:"MySQL"`
			Shared_Redis      []string `yaml:"Shared_Redis"`
			OnDemand_Redis    []string `yaml:"Ondemand_Redis"`
			OnDemand_RabbitMQ []string `yaml:"OnDemand_RabbitMQ"`
			Scheduler         []string `yaml:"Scheduler"`
			ConfigServer      []string `yaml:"ConfigServer"`
			ServiceRegistry   []string `yaml:"ServiceRegistry"`
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
	var gitlist GitList

	QuotaYml := cpath + "/" + clustername + "/Quota.yml"
	fileQuotaYml, err := ioutil.ReadFile(QuotaYml)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal([]byte(fileQuotaYml), &Quotas)
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
		err = yaml.Unmarshal([]byte(fileOrgYml), &gitlist)
		if err != nil {
			panic(err)
		}
		LenList = len(gitlist.OrgList)
		//fmt.Println("123")
	}

	//LenGitOrgList := len(GitOrgList.OrgList)
	LenQuota := len(Quotas.Quota)

	for i := 0; i < LenList; i++ {

		var count, totalcount int

		fmt.Println(" ")
		fmt.Println("Quota: ", GitOrgList.OrgList[i].Quota)
		fmt.Println("Org: ", GitOrgList.OrgList[i].Name)
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

			for p := 0; p < LenQuota; p++ {
				ServiceAccessAudit := Quotas.ServiceAccessAudit
				if strings.Trim(Quotas.Quota[p].Name, "") == strings.Trim(GitOrgList.OrgList[i].Quota, "") {

					fmt.Println("MySql Plans to validate: ", Quotas.Quota[p].ServiceAccess.MySQL)
					fmt.Println("Redis On-demand Plans to validate: ", Quotas.Quota[p].ServiceAccess.OnDemand_Redis)
					fmt.Println("RabbitMQ Plans to validate: ", Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ)
					fmt.Println("Redis Shared Plans to validate: ", Quotas.Quota[p].ServiceAccess.Shared_Redis)
					fmt.Println("Scheduler Plans to validate: ", Quotas.Quota[p].ServiceAccess.Scheduler)
					fmt.Println("Config Server Plans to validate: ", Quotas.Quota[p].ServiceAccess.ConfigServer)
					fmt.Println("Service Registry Plans to validate: ", Quotas.Quota[p].ServiceAccess.ServiceRegistry)
					fmt.Println("Validating MySQL Plans")

					// pulling service details

					LenMysql := len(Quotas.Quota[p].ServiceAccess.MySQL)
					for ms := 0; ms < LenMysql; LenMysql++ {

						var getserviceguid *exec.Cmd
						if ostype == "windows" {
							path := "\"" + "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.MySQL[ms] + "service_broker_names=dedicated-mysql-broker" + "\""
							getserviceguid = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "Delete_services.json")
						} else {
							path := "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.MySQL[ms] + "service_broker_names=dedicated-mysql-broker"
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
							var ServiceJson ServiceListJSON
							if err := json.Unmarshal(fileServiceJson, &ServiceJson); err != nil {
								panic(err)
							}
							fmt.Println("Service Name: ", ServiceJson.Resources[0].Name)
							fmt.Println("Broker Name: ", ServiceJson.Resources[0].BrokerCatalog.Metadata.DisplayName)
							fmt.Println("Guid Name: ", ServiceJson.Resources[0].GUID)

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

								for listorgs := 0; listorgs < len(ServiceVisibilityJson.Organizations); listorgs++ {
									fmt.Println("Org Name: ", ServiceVisibilityJson.Organizations[listorgs].Name)
									fmt.Println("Org Guid: ", ServiceVisibilityJson.Organizations[listorgs].GUID)

									var getorgdets *exec.Cmd
									if ostype == "windows" {
										path := "\"" + "/v3/organizations/" + ServiceVisibilityJson.Organizations[listorgs].GUID + "\""
										getorgdets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingvalue.json")
									} else {
										path := "/v3/organizations" + ServiceVisibilityJson.Organizations[listorgs].GUID
										getorgdets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparing.json")
									}

									if _, err := getorgdets.Output(); err != nil {
										fmt.Println("command: ", getorgdets)
										fmt.Println("Err: ", getorgdets.Stdout, err)
									} else {

										filegetquotaquid, err := ioutil.ReadFile("quotadetailsforcomparing.json")
										if err != nil {
											fmt.Println(err)
										}
										var getquotaguid quotaguid
										if err := json.Unmarshal(filegetquotaquid, &getquotaguid); err != nil {
											panic(err)
										}
										fmt.Println("Quota Guid: ", getquotaguid.Relationships.Quota.Data.GUID)

										var getqutadets *exec.Cmd
										if ostype == "windows" {
											path := "\"" + "/v3/organizations_quotas/" + getquotaguid.Relationships.Quota.Data.GUID + "\""
											getqutadets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										} else {
											path := "/v3/organizations_quotas" + getquotaguid.Relationships.Quota.Data.GUID
											getqutadets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										}

										if _, err := getqutadets.Output(); err != nil {
											fmt.Println("command: ", getqutadets)
											fmt.Println("Err: ", getqutadets.Stdout, err)
										} else {

											filegetquotaguidvalue, err := ioutil.ReadFile("quotadetailsforcomparingval.json")
											if err != nil {
												fmt.Println(err)
											}

											var getquotaguidval guidval
											if err := json.Unmarshal(filegetquotaguidvalue, &getquotaguidval); err != nil {
												panic(err)
											}
											fmt.Println("Quota Name: ", getquotaguidval.Name)

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
												fmt.Println("Total MySql Plans accessble: ", totallist)

												for list := 0; list < totallist; list++ {

													if getacclist.Resources[list].Name == Quotas.Quota[p].ServiceAccess.MySQL[ms] {
														count = 1
													} else {
														count = 0
													}
													totalcount = totalcount + count
													if totalcount == 0 {

														if ServiceAccessAudit == "list" {

															fmt.Println("plan access to be revoked manually: ", getacclist.Resources[list].Name)

														} else if ServiceAccessAudit == "delete" {

															var getserviceguid *exec.Cmd
															if ostype == "windows" {
																//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
																getserviceguid = exec.Command("powershell", "-command", "cf", "disable-service-access", strings.TrimSpace("p.mysql"), Quotas.Quota[p].ServiceAccess.MySQL[ms], GitOrgList.OrgList[i].Name)
															} else {
																//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
																getserviceguid = exec.Command("cf", "disable-service-access", strings.TrimSpace("p.mysql"), Quotas.Quota[p].ServiceAccess.MySQL[ms], GitOrgList.OrgList[i].Name)
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

													}
												}
											}
										}
									}
								}
							}
						}
					}

					LenRedisOnDemand := len(Quotas.Quota[p].ServiceAccess.OnDemand_Redis)
					for ms := 0; ms < LenRedisOnDemand; LenRedisOnDemand++ {

						var getserviceguid *exec.Cmd
						if ostype == "windows" {
							path := "\"" + "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.OnDemand_Redis[ms] + "service_broker_names=redis-odb" + "\""
							getserviceguid = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "Delete_services.json")
						} else {
							path := "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.OnDemand_Redis[ms] + "service_broker_names=redis-odb"
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
							var ServiceJson ServiceListJSON
							if err := json.Unmarshal(fileServiceJson, &ServiceJson); err != nil {
								panic(err)
							}
							fmt.Println("Service Name: ", ServiceJson.Resources[0].Name)
							fmt.Println("Broker Name: ", ServiceJson.Resources[0].BrokerCatalog.Metadata.DisplayName)
							fmt.Println("Guid Name: ", ServiceJson.Resources[0].GUID)

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

								for listorgs := 0; listorgs < len(ServiceVisibilityJson.Organizations); listorgs++ {
									fmt.Println("Org Name: ", ServiceVisibilityJson.Organizations[listorgs].Name)
									fmt.Println("Org Guid: ", ServiceVisibilityJson.Organizations[listorgs].GUID)

									var getorgdets *exec.Cmd
									if ostype == "windows" {
										path := "\"" + "/v3/organizations/" + ServiceVisibilityJson.Organizations[listorgs].GUID + "\""
										getorgdets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingvalue.json")
									} else {
										path := "/v3/organizations" + ServiceVisibilityJson.Organizations[listorgs].GUID
										getorgdets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparing.json")
									}

									if _, err := getorgdets.Output(); err != nil {
										fmt.Println("command: ", getorgdets)
										fmt.Println("Err: ", getorgdets.Stdout, err)
									} else {

										filegetquotaquid, err := ioutil.ReadFile("quotadetailsforcomparing.json")
										if err != nil {
											fmt.Println(err)
										}
										var getquotaguid quotaguid
										if err := json.Unmarshal(filegetquotaquid, &getquotaguid); err != nil {
											panic(err)
										}
										fmt.Println("Quota Guid: ", getquotaguid.Relationships.Quota.Data.GUID)

										var getqutadets *exec.Cmd
										if ostype == "windows" {
											path := "\"" + "/v3/organizations_quotas/" + getquotaguid.Relationships.Quota.Data.GUID + "\""
											getqutadets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										} else {
											path := "/v3/organizations_quotas" + getquotaguid.Relationships.Quota.Data.GUID
											getqutadets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										}

										if _, err := getqutadets.Output(); err != nil {
											fmt.Println("command: ", getqutadets)
											fmt.Println("Err: ", getqutadets.Stdout, err)
										} else {

											filegetquotaguidvalue, err := ioutil.ReadFile("quotadetailsforcomparingval.json")
											if err != nil {
												fmt.Println(err)
											}

											var getquotaguidval guidval
											if err := json.Unmarshal(filegetquotaguidvalue, &getquotaguidval); err != nil {
												panic(err)
											}
											fmt.Println("Quota Name: ", getquotaguidval.Name)

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
												fmt.Println("Total Redis Ondemand Plans accessble: ", totallist)

												for list := 0; list < totallist; list++ {

													if getacclist.Resources[list].Name == Quotas.Quota[p].ServiceAccess.OnDemand_Redis[ms] {
														count = 1
													} else {
														count = 0
													}
													totalcount = totalcount + count
													if totalcount == 0 {

														if ServiceAccessAudit == "list" {

															fmt.Println("plan access to be revoked manually: ", getacclist.Resources[list].Name)

														} else if ServiceAccessAudit == "delete" {

															var getserviceguid *exec.Cmd
															if ostype == "windows" {
																//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
																getserviceguid = exec.Command("powershell", "-command", "cf", "disable-service-access", strings.TrimSpace("p.redis"), Quotas.Quota[p].ServiceAccess.OnDemand_Redis[ms], GitOrgList.OrgList[i].Name)
															} else {
																//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
																getserviceguid = exec.Command("cf", "disable-service-access", strings.TrimSpace("p.redis"), Quotas.Quota[p].ServiceAccess.OnDemand_Redis[ms], GitOrgList.OrgList[i].Name)
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

													}
												}
											}
										}
									}
								}
							}
						}
					}

					LenRedisShared := len(Quotas.Quota[p].ServiceAccess.Shared_Redis)
					for ms := 0; ms < LenRedisShared; LenRedisShared++ {

						var getserviceguid *exec.Cmd
						if ostype == "windows" {
							path := "\"" + "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.Shared_Redis[ms] + "service_broker_names=p-redis" + "\""
							getserviceguid = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "Delete_services.json")
						} else {
							path := "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.Shared_Redis[ms] + "service_broker_names=p-redis"
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
							var ServiceJson ServiceListJSON
							if err := json.Unmarshal(fileServiceJson, &ServiceJson); err != nil {
								panic(err)
							}
							fmt.Println("Service Name: ", ServiceJson.Resources[0].Name)
							fmt.Println("Broker Name: ", ServiceJson.Resources[0].BrokerCatalog.Metadata.DisplayName)
							fmt.Println("Guid Name: ", ServiceJson.Resources[0].GUID)

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

								for listorgs := 0; listorgs < len(ServiceVisibilityJson.Organizations); listorgs++ {
									fmt.Println("Org Name: ", ServiceVisibilityJson.Organizations[listorgs].Name)
									fmt.Println("Org Guid: ", ServiceVisibilityJson.Organizations[listorgs].GUID)

									var getorgdets *exec.Cmd
									if ostype == "windows" {
										path := "\"" + "/v3/organizations/" + ServiceVisibilityJson.Organizations[listorgs].GUID + "\""
										getorgdets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingvalue.json")
									} else {
										path := "/v3/organizations" + ServiceVisibilityJson.Organizations[listorgs].GUID
										getorgdets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparing.json")
									}

									if _, err := getorgdets.Output(); err != nil {
										fmt.Println("command: ", getorgdets)
										fmt.Println("Err: ", getorgdets.Stdout, err)
									} else {

										filegetquotaquid, err := ioutil.ReadFile("quotadetailsforcomparing.json")
										if err != nil {
											fmt.Println(err)
										}
										var getquotaguid quotaguid
										if err := json.Unmarshal(filegetquotaquid, &getquotaguid); err != nil {
											panic(err)
										}
										fmt.Println("Quota Guid: ", getquotaguid.Relationships.Quota.Data.GUID)

										var getqutadets *exec.Cmd
										if ostype == "windows" {
											path := "\"" + "/v3/organizations_quotas/" + getquotaguid.Relationships.Quota.Data.GUID + "\""
											getqutadets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										} else {
											path := "/v3/organizations_quotas" + getquotaguid.Relationships.Quota.Data.GUID
											getqutadets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										}

										if _, err := getqutadets.Output(); err != nil {
											fmt.Println("command: ", getqutadets)
											fmt.Println("Err: ", getqutadets.Stdout, err)
										} else {

											filegetquotaguidvalue, err := ioutil.ReadFile("quotadetailsforcomparingval.json")
											if err != nil {
												fmt.Println(err)
											}

											var getquotaguidval guidval
											if err := json.Unmarshal(filegetquotaguidvalue, &getquotaguidval); err != nil {
												panic(err)
											}
											fmt.Println("Quota Name: ", getquotaguidval.Name)

											/////

											var pullseracclist *exec.Cmd
											if ostype == "windows" {
												path := "\"" + "/v3/service_plans/?organization_guids=" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=p-redis" + "\""
												pullseracclist = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "getseracclist.json")
											} else {
												path := "/v3/organizations_quotas" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=p-redis"
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
												fmt.Println("Total Redis Shared Plans accessble: ", totallist)

												for list := 0; list < totallist; list++ {

													if getacclist.Resources[list].Name == Quotas.Quota[p].ServiceAccess.Shared_Redis[ms] {
														count = 1
													} else {
														count = 0
													}
													totalcount = totalcount + count
													if totalcount == 0 {

														if ServiceAccessAudit == "list" {

															fmt.Println("plan access to be revoked manually: ", getacclist.Resources[list].Name)

														} else if ServiceAccessAudit == "delete" {

															var getserviceguid *exec.Cmd
															if ostype == "windows" {
																//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
																getserviceguid = exec.Command("powershell", "-command", "cf", "disable-service-access", strings.TrimSpace("p-redis"), Quotas.Quota[p].ServiceAccess.Shared_Redis[ms], GitOrgList.OrgList[i].Name)
															} else {
																//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
																getserviceguid = exec.Command("cf", "disable-service-access", strings.TrimSpace("p-redis"), Quotas.Quota[p].ServiceAccess.Shared_Redis[ms], GitOrgList.OrgList[i].Name)
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

													}
												}
											}
										}
									}
								}
							}
						}
					}

					LenRabbitMq := len(Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ)
					for ms := 0; ms < LenRabbitMq; LenRabbitMq++ {

						var getserviceguid *exec.Cmd
						if ostype == "windows" {
							path := "\"" + "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ[ms] + "service_broker_names=rabbitmq-odb" + "\""
							getserviceguid = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "Delete_services.json")
						} else {
							path := "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ[ms] + "service_broker_names=rabbitmq-odb"
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
							var ServiceJson ServiceListJSON
							if err := json.Unmarshal(fileServiceJson, &ServiceJson); err != nil {
								panic(err)
							}
							fmt.Println("Service Name: ", ServiceJson.Resources[0].Name)
							fmt.Println("Broker Name: ", ServiceJson.Resources[0].BrokerCatalog.Metadata.DisplayName)
							fmt.Println("Guid Name: ", ServiceJson.Resources[0].GUID)

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

								for listorgs := 0; listorgs < len(ServiceVisibilityJson.Organizations); listorgs++ {
									fmt.Println("Org Name: ", ServiceVisibilityJson.Organizations[listorgs].Name)
									fmt.Println("Org Guid: ", ServiceVisibilityJson.Organizations[listorgs].GUID)

									var getorgdets *exec.Cmd
									if ostype == "windows" {
										path := "\"" + "/v3/organizations/" + ServiceVisibilityJson.Organizations[listorgs].GUID + "\""
										getorgdets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingvalue.json")
									} else {
										path := "/v3/organizations" + ServiceVisibilityJson.Organizations[listorgs].GUID
										getorgdets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparing.json")
									}

									if _, err := getorgdets.Output(); err != nil {
										fmt.Println("command: ", getorgdets)
										fmt.Println("Err: ", getorgdets.Stdout, err)
									} else {

										filegetquotaquid, err := ioutil.ReadFile("quotadetailsforcomparing.json")
										if err != nil {
											fmt.Println(err)
										}
										var getquotaguid quotaguid
										if err := json.Unmarshal(filegetquotaquid, &getquotaguid); err != nil {
											panic(err)
										}
										fmt.Println("Quota Guid: ", getquotaguid.Relationships.Quota.Data.GUID)

										var getqutadets *exec.Cmd
										if ostype == "windows" {
											path := "\"" + "/v3/organizations_quotas/" + getquotaguid.Relationships.Quota.Data.GUID + "\""
											getqutadets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										} else {
											path := "/v3/organizations_quotas" + getquotaguid.Relationships.Quota.Data.GUID
											getqutadets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										}

										if _, err := getqutadets.Output(); err != nil {
											fmt.Println("command: ", getqutadets)
											fmt.Println("Err: ", getqutadets.Stdout, err)
										} else {

											filegetquotaguidvalue, err := ioutil.ReadFile("quotadetailsforcomparingval.json")
											if err != nil {
												fmt.Println(err)
											}

											var getquotaguidval guidval
											if err := json.Unmarshal(filegetquotaguidvalue, &getquotaguidval); err != nil {
												panic(err)
											}
											fmt.Println("Quota Name: ", getquotaguidval.Name)

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
												fmt.Println("Total Redis Shared Plans accessble: ", totallist)

												for list := 0; list < totallist; list++ {

													if getacclist.Resources[list].Name == Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ[ms] {
														count = 1
													} else {
														count = 0
													}
													totalcount = totalcount + count
													if totalcount == 0 {

														if ServiceAccessAudit == "list" {

															fmt.Println("plan access to be revoked manually: ", getacclist.Resources[list].Name)

														} else if ServiceAccessAudit == "delete" {

															var getserviceguid *exec.Cmd
															if ostype == "windows" {
																//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
																getserviceguid = exec.Command("powershell", "-command", "cf", "disable-service-access", strings.TrimSpace("p.rabbitmq"), Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ[ms], GitOrgList.OrgList[i].Name)
															} else {
																//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
																getserviceguid = exec.Command("cf", "disable-service-access", strings.TrimSpace("p.rabbitmq"), Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ[ms], GitOrgList.OrgList[i].Name)
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

													}
												}
											}
										}
									}
								}
							}
						}
					}

					Scheduler := len(Quotas.Quota[p].ServiceAccess.Scheduler)
					for ms := 0; ms < Scheduler; Scheduler++ {

						var getserviceguid *exec.Cmd
						if ostype == "windows" {
							path := "\"" + "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.Scheduler[ms] + "service_broker_names=scheduler-for-pcf" + "\""
							getserviceguid = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "Delete_services.json")
						} else {
							path := "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.Scheduler[ms] + "service_broker_names=scheduler-for-pcf"
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
							var ServiceJson ServiceListJSON
							if err := json.Unmarshal(fileServiceJson, &ServiceJson); err != nil {
								panic(err)
							}
							fmt.Println("Service Name: ", ServiceJson.Resources[0].Name)
							fmt.Println("Broker Name: ", ServiceJson.Resources[0].BrokerCatalog.Metadata.DisplayName)
							fmt.Println("Guid Name: ", ServiceJson.Resources[0].GUID)

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

								for listorgs := 0; listorgs < len(ServiceVisibilityJson.Organizations); listorgs++ {
									fmt.Println("Org Name: ", ServiceVisibilityJson.Organizations[listorgs].Name)
									fmt.Println("Org Guid: ", ServiceVisibilityJson.Organizations[listorgs].GUID)

									var getorgdets *exec.Cmd
									if ostype == "windows" {
										path := "\"" + "/v3/organizations/" + ServiceVisibilityJson.Organizations[listorgs].GUID + "\""
										getorgdets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingvalue.json")
									} else {
										path := "/v3/organizations" + ServiceVisibilityJson.Organizations[listorgs].GUID
										getorgdets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparing.json")
									}

									if _, err := getorgdets.Output(); err != nil {
										fmt.Println("command: ", getorgdets)
										fmt.Println("Err: ", getorgdets.Stdout, err)
									} else {

										filegetquotaquid, err := ioutil.ReadFile("quotadetailsforcomparing.json")
										if err != nil {
											fmt.Println(err)
										}
										var getquotaguid quotaguid
										if err := json.Unmarshal(filegetquotaquid, &getquotaguid); err != nil {
											panic(err)
										}
										fmt.Println("Quota Guid: ", getquotaguid.Relationships.Quota.Data.GUID)

										var getqutadets *exec.Cmd
										if ostype == "windows" {
											path := "\"" + "/v3/organizations_quotas/" + getquotaguid.Relationships.Quota.Data.GUID + "\""
											getqutadets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										} else {
											path := "/v3/organizations_quotas" + getquotaguid.Relationships.Quota.Data.GUID
											getqutadets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										}

										if _, err := getqutadets.Output(); err != nil {
											fmt.Println("command: ", getqutadets)
											fmt.Println("Err: ", getqutadets.Stdout, err)
										} else {

											filegetquotaguidvalue, err := ioutil.ReadFile("quotadetailsforcomparingval.json")
											if err != nil {
												fmt.Println(err)
											}

											var getquotaguidval guidval
											if err := json.Unmarshal(filegetquotaguidvalue, &getquotaguidval); err != nil {
												panic(err)
											}
											fmt.Println("Quota Name: ", getquotaguidval.Name)

											/////

											var pullseracclist *exec.Cmd
											if ostype == "windows" {
												path := "\"" + "/v3/service_plans/?organization_guids=" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=scheduler-for-pcf" + "\""
												pullseracclist = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "getseracclist.json")
											} else {
												path := "/v3/organizations_quotas" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=scheduler-for-pcf"
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
												fmt.Println("Total Redis Shared Plans accessble: ", totallist)

												for list := 0; list < totallist; list++ {

													if getacclist.Resources[list].Name == Quotas.Quota[p].ServiceAccess.Scheduler[ms] {
														count = 1
													} else {
														count = 0
													}
													totalcount = totalcount + count
													if totalcount == 0 {

														if ServiceAccessAudit == "list" {

															fmt.Println("plan access to be revoked manually: ", getacclist.Resources[list].Name)

														} else if ServiceAccessAudit == "delete" {

															var getserviceguid *exec.Cmd
															if ostype == "windows" {
																//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
																getserviceguid = exec.Command("powershell", "-command", "cf", "disable-service-access", strings.TrimSpace("scheduler-for-pcf"), Quotas.Quota[p].ServiceAccess.Scheduler[ms], GitOrgList.OrgList[i].Name)
															} else {
																//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
																getserviceguid = exec.Command("cf", "disable-service-access", strings.TrimSpace("scheduler-for-pcf"), Quotas.Quota[p].ServiceAccess.Scheduler[ms], GitOrgList.OrgList[i].Name)
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

													}
												}
											}
										}
									}
								}
							}
						}
					}

					lenConfigServer := len(Quotas.Quota[p].ServiceAccess.ConfigServer)
					for ms := 0; ms < lenConfigServer; lenConfigServer++ {

						var getserviceguid *exec.Cmd
						if ostype == "windows" {
							path := "\"" + "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.ConfigServer[ms] + "service_broker_names=scs-service-broker" + "\""
							getserviceguid = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "Delete_services.json")
						} else {
							path := "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.ConfigServer[ms] + "service_broker_names=scs-service-broker"
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
							var ServiceJson ServiceListJSON
							if err := json.Unmarshal(fileServiceJson, &ServiceJson); err != nil {
								panic(err)
							}
							fmt.Println("Service Name: ", ServiceJson.Resources[0].Name)
							fmt.Println("Broker Name: ", ServiceJson.Resources[0].BrokerCatalog.Metadata.DisplayName)
							fmt.Println("Guid Name: ", ServiceJson.Resources[0].GUID)

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

								for listorgs := 0; listorgs < len(ServiceVisibilityJson.Organizations); listorgs++ {
									fmt.Println("Org Name: ", ServiceVisibilityJson.Organizations[listorgs].Name)
									fmt.Println("Org Guid: ", ServiceVisibilityJson.Organizations[listorgs].GUID)

									var getorgdets *exec.Cmd
									if ostype == "windows" {
										path := "\"" + "/v3/organizations/" + ServiceVisibilityJson.Organizations[listorgs].GUID + "\""
										getorgdets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingvalue.json")
									} else {
										path := "/v3/organizations" + ServiceVisibilityJson.Organizations[listorgs].GUID
										getorgdets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparing.json")
									}

									if _, err := getorgdets.Output(); err != nil {
										fmt.Println("command: ", getorgdets)
										fmt.Println("Err: ", getorgdets.Stdout, err)
									} else {

										filegetquotaquid, err := ioutil.ReadFile("quotadetailsforcomparing.json")
										if err != nil {
											fmt.Println(err)
										}
										var getquotaguid quotaguid
										if err := json.Unmarshal(filegetquotaquid, &getquotaguid); err != nil {
											panic(err)
										}
										fmt.Println("Quota Guid: ", getquotaguid.Relationships.Quota.Data.GUID)

										var getqutadets *exec.Cmd
										if ostype == "windows" {
											path := "\"" + "/v3/organizations_quotas/" + getquotaguid.Relationships.Quota.Data.GUID + "\""
											getqutadets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										} else {
											path := "/v3/organizations_quotas" + getquotaguid.Relationships.Quota.Data.GUID
											getqutadets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										}

										if _, err := getqutadets.Output(); err != nil {
											fmt.Println("command: ", getqutadets)
											fmt.Println("Err: ", getqutadets.Stdout, err)
										} else {

											filegetquotaguidvalue, err := ioutil.ReadFile("quotadetailsforcomparingval.json")
											if err != nil {
												fmt.Println(err)
											}

											var getquotaguidval guidval
											if err := json.Unmarshal(filegetquotaguidvalue, &getquotaguidval); err != nil {
												panic(err)
											}
											fmt.Println("Quota Name: ", getquotaguidval.Name)

											/////

											var pullseracclist *exec.Cmd
											if ostype == "windows" {
												path := "\"" + "/v3/service_plans/?organization_guids=" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=scs-service-broker" + "\""
												pullseracclist = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "getseracclist.json")
											} else {
												path := "/v3/organizations_quotas" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=scs-service-broker"
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
												fmt.Println("Total Redis Shared Plans accessble: ", totallist)

												for list := 0; list < totallist; list++ {

													if getacclist.Resources[list].Name == Quotas.Quota[p].ServiceAccess.ConfigServer[ms] {
														count = 1
													} else {
														count = 0
													}
													totalcount = totalcount + count
													if totalcount == 0 {

														if ServiceAccessAudit == "list" {

															fmt.Println("plan access to be revoked manually: ", getacclist.Resources[list].Name)

														} else if ServiceAccessAudit == "delete" {

															var getserviceguid *exec.Cmd
															if ostype == "windows" {
																//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
																getserviceguid = exec.Command("powershell", "-command", "cf", "disable-service-access", strings.TrimSpace("p.config-server"), Quotas.Quota[p].ServiceAccess.ConfigServer[ms], GitOrgList.OrgList[i].Name)
															} else {
																//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
																getserviceguid = exec.Command("cf", "disable-service-access", strings.TrimSpace("p.config-server"), Quotas.Quota[p].ServiceAccess.ConfigServer[ms], GitOrgList.OrgList[i].Name)
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

													}
												}
											}
										}
									}
								}
							}
						}
					}

					lenserviceregistry := len(Quotas.Quota[p].ServiceAccess.ServiceRegistry)
					for ms := 0; ms < lenserviceregistry; lenserviceregistry++ {

						var getserviceguid *exec.Cmd
						if ostype == "windows" {
							path := "\"" + "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.ServiceRegistry[ms] + "service_broker_names=scs-service-broker" + "\""
							getserviceguid = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "Delete_services.json")
						} else {
							path := "/v3/service_plans/?names=" + Quotas.Quota[p].ServiceAccess.ServiceRegistry[ms] + "service_broker_names=scs-service-broker"
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
							var ServiceJson ServiceListJSON
							if err := json.Unmarshal(fileServiceJson, &ServiceJson); err != nil {
								panic(err)
							}
							fmt.Println("Service Name: ", ServiceJson.Resources[0].Name)
							fmt.Println("Broker Name: ", ServiceJson.Resources[0].BrokerCatalog.Metadata.DisplayName)
							fmt.Println("Guid Name: ", ServiceJson.Resources[0].GUID)

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

								for listorgs := 0; listorgs < len(ServiceVisibilityJson.Organizations); listorgs++ {
									fmt.Println("Org Name: ", ServiceVisibilityJson.Organizations[listorgs].Name)
									fmt.Println("Org Guid: ", ServiceVisibilityJson.Organizations[listorgs].GUID)

									var getorgdets *exec.Cmd
									if ostype == "windows" {
										path := "\"" + "/v3/organizations/" + ServiceVisibilityJson.Organizations[listorgs].GUID + "\""
										getorgdets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingvalue.json")
									} else {
										path := "/v3/organizations" + ServiceVisibilityJson.Organizations[listorgs].GUID
										getorgdets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparing.json")
									}

									if _, err := getorgdets.Output(); err != nil {
										fmt.Println("command: ", getorgdets)
										fmt.Println("Err: ", getorgdets.Stdout, err)
									} else {

										filegetquotaquid, err := ioutil.ReadFile("quotadetailsforcomparing.json")
										if err != nil {
											fmt.Println(err)
										}
										var getquotaguid quotaguid
										if err := json.Unmarshal(filegetquotaquid, &getquotaguid); err != nil {
											panic(err)
										}
										fmt.Println("Quota Guid: ", getquotaguid.Relationships.Quota.Data.GUID)

										var getqutadets *exec.Cmd
										if ostype == "windows" {
											path := "\"" + "/v3/organizations_quotas/" + getquotaguid.Relationships.Quota.Data.GUID + "\""
											getqutadets = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										} else {
											path := "/v3/organizations_quotas" + getquotaguid.Relationships.Quota.Data.GUID
											getqutadets = exec.Command("cf", "curl", path, "--output", "quotadetailsforcomparingval.json")
										}

										if _, err := getqutadets.Output(); err != nil {
											fmt.Println("command: ", getqutadets)
											fmt.Println("Err: ", getqutadets.Stdout, err)
										} else {

											filegetquotaguidvalue, err := ioutil.ReadFile("quotadetailsforcomparingval.json")
											if err != nil {
												fmt.Println(err)
											}

											var getquotaguidval guidval
											if err := json.Unmarshal(filegetquotaguidvalue, &getquotaguidval); err != nil {
												panic(err)
											}
											fmt.Println("Quota Name: ", getquotaguidval.Name)

											/////

											var pullseracclist *exec.Cmd
											if ostype == "windows" {
												path := "\"" + "/v3/service_plans/?organization_guids=" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=scs-service-broker" + "\""
												pullseracclist = exec.Command("powershell", "-command", "cf", "curl", path, "--output", "getseracclist.json")
											} else {
												path := "/v3/organizations_quotas" + ServiceVisibilityJson.Organizations[listorgs].GUID + "&service_broker_names=scs-service-broker"
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
												fmt.Println("Total Redis Shared Plans accessble: ", totallist)

												for list := 0; list < totallist; list++ {

													if getacclist.Resources[list].Name == Quotas.Quota[p].ServiceAccess.ServiceRegistry[ms] {
														count = 1
													} else {
														count = 0
													}
													totalcount = totalcount + count
													if totalcount == 0 {

														if ServiceAccessAudit == "list" {

															fmt.Println("plan access to be revoked manually: ", getacclist.Resources[list].Name)

														} else if ServiceAccessAudit == "delete" {

															var getserviceguid *exec.Cmd
															if ostype == "windows" {
																//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
																getserviceguid = exec.Command("powershell", "-command", "cf", "disable-service-access", strings.TrimSpace("p.service-registry"), Quotas.Quota[p].ServiceAccess.ServiceRegistry[ms], GitOrgList.OrgList[i].Name)
															} else {
																//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
																getserviceguid = exec.Command("cf", "disable-service-access", strings.TrimSpace("p.service-registry"), Quotas.Quota[p].ServiceAccess.ServiceRegistry[ms], GitOrgList.OrgList[i].Name)
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

													}
												}
											}
										}
									}
								}
							}
						}
					}
					//Pulled SI Guid
				} else {
					fmt.Println("Quota not found in Quota.yml file")
				}

			}
		}
	}
	return err
}
