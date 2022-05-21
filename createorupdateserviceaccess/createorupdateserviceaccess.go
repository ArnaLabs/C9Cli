package createorupdateserviceaccess

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

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
			//Shared_Redis     []string `yaml:"Shared_Redis"`
			OnDemand_Redis    []string `yaml:"OnDemand_Redis"`
			OnDemand_RabbitMQ []string `yaml:"OnDemand_RabbitMQ"`
			//Scheduler []string `yaml:"Scheduler"`
			///ConfigServer []string `yaml:"ConfigServer"`
			//ServiceRegistry []string `yaml:"ServiceRegistry"`
		} `yaml:"ServiceAccess"`
	} `yaml:"quota"`
	Audit string `yaml:"Audit"`
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

func CreateOrUpdateServiceAccess(clustername string, cpath string, ostype string) error {

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

	//spath := cpath+"/"+clustername+"-state/"

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
		err = yaml.Unmarshal([]byte(fileOrgYml), &GitOrgList)
		if err != nil {
			panic(err)
		}
		LenList = len(GitOrgList.OrgList)
	}

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
			for p := 0; p < LenQuota; p++ {

				if strings.ToLower(strings.Trim(Quotas.Quota[p].Name, "")) == strings.ToLower(strings.Trim(GitOrgList.OrgList[i].Quota, "")) {

					fmt.Println(" ")
					fmt.Println("Org: ", GitOrgList.OrgList[i].Name)
					fmt.Println("Quota: ", GitOrgList.OrgList[i].Quota)
					fmt.Println("MySql Plans to enable: ", Quotas.Quota[p].ServiceAccess.MySQL)
					fmt.Println("Redis On-demand Plans to enable: ", Quotas.Quota[p].ServiceAccess.OnDemand_Redis)
					fmt.Println("RabbitMQ Plans to enable: ", Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ)
					//	fmt.Println("Redis Shared Plans to enable: ", Quotas.Quota[p].ServiceAccess.Shared_Redis)
					//	fmt.Println("Scheduler Plans to enable: ", Quotas.Quota[p].ServiceAccess.Scheduler)
					//	fmt.Println("Config Server Plans to enable: ", Quotas.Quota[p].ServiceAccess.ConfigServer)
					//	fmt.Println("Service Registry Plans to enable: ", Quotas.Quota[p].ServiceAccess.ServiceRegistry)

					fmt.Println("Enabling MySQL Plans")
					// pulling service details
					LenMysql := len(Quotas.Quota[p].ServiceAccess.MySQL)
					for ms := 0; ms < LenMysql; ms++ {
						fmt.Println("++", Quotas.Quota[p].ServiceAccess.MySQL[ms])
						var getserviceguid *exec.Cmd
						if ostype == "windows" {
							//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
							getserviceguid = exec.Command("powershell", "-command", "cf", "enable-service-access", strings.TrimSpace("p.mysql"), "-p", Quotas.Quota[p].ServiceAccess.MySQL[ms], "-o", GitOrgList.OrgList[i].Name)
						} else {
							//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
							getserviceguid = exec.Command("cf", "enable-service-access", strings.TrimSpace("p.mysql"), "-p", Quotas.Quota[p].ServiceAccess.MySQL[ms], "-o", GitOrgList.OrgList[i].Name)
						}
						err = getserviceguid.Run()
						if err == nil {
							//fmt.Println(err, getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
							//fmt.Println(Quotas.Quota[p].ServiceAccess.MySQL[ms], "Plan exist")
						} else {
							fmt.Println("err", err, getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
						}
					}

					fmt.Println("Enabling Redis OnDemand Plans")
					LenRedisOnDemand := len(Quotas.Quota[p].ServiceAccess.OnDemand_Redis)
					for ro := 0; ro < LenRedisOnDemand; ro++ {
						fmt.Println("++", Quotas.Quota[p].ServiceAccess.OnDemand_Redis[ro])
						var getserviceguid *exec.Cmd
						if ostype == "windows" {
							//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
							getserviceguid = exec.Command("powershell", "-command", "cf", "enable-service-access", strings.TrimSpace("p.redis"), "-p", Quotas.Quota[p].ServiceAccess.OnDemand_Redis[ro], "-o", GitOrgList.OrgList[i].Name)
						} else {
							//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
							getserviceguid = exec.Command("cf", "enable-service-access", strings.TrimSpace("p.redis"), "-p", Quotas.Quota[p].ServiceAccess.OnDemand_Redis[ro], "-o", GitOrgList.OrgList[i].Name)
						}
						err = getserviceguid.Run()
						if err == nil {
							//fmt.Println(err,getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
							//fmt.Println("Plan exist")
						} else {
							fmt.Println("err", err, getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
						}
					}

					//fmt.Println("Enabling Redis Shared Plans")
					//LenRedisShared := len(Quotas.Quota[p].ServiceAccess.Shared_Redis)
					//for rs := 0; rs < LenRedisShared; rs++ {
					//	var getserviceguid *exec.Cmd
					//	if ostype == "windows" {
					//		//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
					//		getserviceguid = exec.Command("powershell", "-command", "cf", "enable-service-access", strings.TrimSpace("p-redis"), "-p",Quotas.Quota[p].ServiceAccess.Shared_Redis[rs], "-o",GitOrgList.OrgList[i].Name)
					//	} else {
					//		//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
					//		getserviceguid = exec.Command("cf", "enable-service-access", strings.TrimSpace("p-redis"), "-p",Quotas.Quota[p].ServiceAccess.Shared_Redis[rs], "-o" ,GitOrgList.OrgList[i].Name)
					//	}
					//	err = getserviceguid.Run()
					//	if err == nil {
					//		fmt.Println(err, getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
					//	} else {
					//		fmt.Println("err", getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
					//	}
					//}

					fmt.Println("Enabling RabbitMQ Ondemand Plans")
					LenRabbitMq := len(Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ)
					for rb := 0; rb < LenRabbitMq; rb++ {
						fmt.Println("++", Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ[rb])
						var getserviceguid *exec.Cmd
						if ostype == "windows" {
							//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
							getserviceguid = exec.Command("powershell", "-command", "cf", "enable-service-access", strings.TrimSpace("p.rabbitmq"), "-p", Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ[rb], "-o", GitOrgList.OrgList[i].Name)
						} else {
							//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
							getserviceguid = exec.Command("cf", "enable-service-access", strings.TrimSpace("p.rabbitmq"), "-p", Quotas.Quota[p].ServiceAccess.OnDemand_RabbitMQ[rb], "-o", GitOrgList.OrgList[i].Name)
						}
						err = getserviceguid.Run()
						if err == nil {
							//fmt.Println("Plan exist")
							//fmt.Println(err,getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
						} else {
							fmt.Println("err", err, getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
						}
					}

					//fmt.Println("Enabling Scheduler Plans")
					//Scheduler := len(Quotas.Quota[p].ServiceAccess.Scheduler)
					//for s := 0; s < Scheduler; s++ {

					//	var getserviceguid *exec.Cmd
					//	if ostype == "windows" {
					//		//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
					//		getserviceguid = exec.Command("powershell", "-command", "cf", "enable-service-access", strings.TrimSpace("scheduler-for-pcf"), "-p",Quotas.Quota[p].ServiceAccess.Scheduler[s],"-o" ,GitOrgList.OrgList[i].Name)
					//	} else {
					//		//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
					//		getserviceguid = exec.Command("cf", "enable-service-access", strings.TrimSpace("scheduler-for-pcf"), "-p",Quotas.Quota[p].ServiceAccess.Scheduler[s], "-o",GitOrgList.OrgList[i].Name)
					//	}
					//	err = getserviceguid.Run()
					//	if err == nil {
					//		fmt.Println(err, getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
					//	} else {
					//		fmt.Println("err", getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr, err)
					//	}
					//}

					//fmt.Println("Enabling ConfigServer Plans")
					//lenConfigServer := len(Quotas.Quota[p].ServiceAccess.ConfigServer)
					//for cs := 0; cs < lenConfigServer; cs++ {

					//	var getserviceguid *exec.Cmd
					//	if ostype == "windows" {
					//		//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
					//		getserviceguid = exec.Command("powershell", "-command", "cf", "enable-service-access", strings.TrimSpace("p.config-server"), "-p",Quotas.Quota[p].ServiceAccess.ConfigServer[cs], "-o",GitOrgList.OrgList[i].Name)
					//	} else {
					//		//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
					//		getserviceguid = exec.Command("cf", "enable-service-access", strings.TrimSpace("p.config-server"), "-p",Quotas.Quota[p].ServiceAccess.ConfigServer[cs], "-o",GitOrgList.OrgList[i].Name)
					//	}
					//	err = getserviceguid.Run()
					//	if err == nil {
					//		fmt.Println(err, getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
					//	} else {
					//		fmt.Println("err", err, getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
					//	}
					//}

					//fmt.Println("Enabling ServiceRegistry Plans")
					//lenserviceregistry := len(Quotas.Quota[p].ServiceAccess.ServiceRegistry)
					//for sr := 0; sr < lenserviceregistry; sr++ {

					//	var getserviceguid *exec.Cmd
					//	if ostype == "windows" {
					//		//path := "\""+"/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[ms]+"service_broker_names=dedicated-mysql-broker"+"\""
					//		getserviceguid = exec.Command("powershell", "-command", "cf", "enable-service-access", strings.TrimSpace("p.service-registry"), "-p",Quotas.Quota[p].ServiceAccess.ServiceRegistry[sr], "-o",GitOrgList.OrgList[i].Name)
					//	} else {
					//		//path := "/v3/service_plans/?names="+Quotas.Quota[p].ServiceAccess.MySQL[0]+"service_broker_names=dedicated-mysql-broker"
					//		getserviceguid = exec.Command("cf", "enable-service-access", strings.TrimSpace("p.service-registry"), "-p",Quotas.Quota[p].ServiceAccess.ServiceRegistry[sr], "-o",GitOrgList.OrgList[i].Name)
					//	}
					//	err = getserviceguid.Run()
					//	if err == nil {
					//		fmt.Println(err, getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
					//	} else {
					//		fmt.Println("err", err,getserviceguid, getserviceguid.Stdout, getserviceguid.Stderr)
					//	}
					//}

				} else {
					//fmt.Println("Org: ", GitOrgList.OrgList[i].Name)
					//fmt.Println("Quota: ", Quotas.Quota[p].Name)
					//fmt.Println("Quota is not assiged in Quota.yml file")
				}
			}
		}
	}
	return err
}
