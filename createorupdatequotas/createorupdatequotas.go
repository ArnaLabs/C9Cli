package createorupdatequotas

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
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

func CreateOrUpdateQuotas(clustername string, cpath string, ostype string) error {

	var Quotas Quotalist
	var ProtectedQuota ProtectedList
	var cmd *exec.Cmd
	//var quotalistjson QuotaListJson


	QuotaYml := cpath+"/"+clustername+"/Quota.yml"
	fileQuotaYml, err := ioutil.ReadFile(QuotaYml)

	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileQuotaYml), &Quotas)
	if err != nil {
		panic(err)
	}

	ProtectedQuotasYml := cpath+"/"+clustername+"/ProtectedResources.yml"
	fileProtectedQYml, err := ioutil.ReadFile(ProtectedQuotasYml)

	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileProtectedQYml), &ProtectedQuota)
	if err != nil {
		panic(err)
	}

	LenQuota := len(Quotas.Quota)
	LenProtectedQuota := len(ProtectedQuota.Quota)

	for i := 0; i < LenQuota; i++ {

		var count, totalcount int
		fmt.Println(" ")
		fmt.Println("Quota: ", Quotas.Quota[i].Name)


		SerLimit := Quotas.Quota[i].ServiceInstanceLimit
		AppLimt  := Quotas.Quota[i].AppInstanceLimit
		MemLimit := Quotas.Quota[i].MemoryLimit

		if Quotas.Quota[i].ServiceInstanceLimit == ""{
			SerLimit = "0"
		}else {
		}

		if string(Quotas.Quota[i].AppInstanceLimit) == "" {
			AppLimt = "25"
		} else {
		}

		if Quotas.Quota[i].MemoryLimit == "" {
			MemLimit = "1024M"
		} else {
		}

		for p := 0; p < LenProtectedQuota; p++ {
			//fmt.Println("Protected Quota: ", ProtectedQuota.Quota[p],",", Quotas.Quota[i].Name)
			if strings.Trim(ProtectedQuota.Quota[p], "") == strings.Trim(Quotas.Quota[i].Name, "") {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {

			fmt.Println("This is not Protected Quota")
			var getquotas *exec.Cmd
			if ostype == "windows" {
				path := "\""+"/v3/organization_quotas?names="+Quotas.Quota[i].Name+"\""
				getquotas = exec.Command("powershell", "-command","cf", "curl", path, "--output", "CreateOrUpdateQuotas_listquotas.json")
			} else {
				path := "/v3/organization_quotas?names="+Quotas.Quota[i].Name
				getquotas = exec.Command("cf", "curl", path, "--output", "CreateOrUpdateQuotas_listquotas.json")
			}

			if _, err := getquotas.Output(); err != nil {
				fmt.Println("command: ", getquotas)
				fmt.Println("Err: ", getquotas.Stdout, err)
			} else {
				fileQuotaJson, err := ioutil.ReadFile("CreateOrUpdateQuotas_listquotas.json")
				if err != nil {
					fmt.Println(err)
				}
				var quotalistjson QuotaListJson
				if err := json.Unmarshal(fileQuotaJson, &quotalistjson); err != nil {
					panic(err)
				}
				if len(quotalistjson.Resources) == 0 {

					fmt.Println("Creating Quota")

					if Quotas.Quota[i].AllowPaidPlans == true {
						cmd = exec.Command("cf", "create-quota", Quotas.Quota[i].Name, "-m", MemLimit, "-i", "-1", "-r", "-1", "-s", SerLimit, "-a", AppLimt, "--allow-paid-service-plans")
						if _, err := cmd.Output(); err != nil{
							fmt.Println("command: ", cmd)
							fmt.Println("Err: ", cmd.Stdout, err)
						} else {
							//			fmt.Println("command: ", cmd)
							fmt.Println(cmd.Stdout)
						}
						QuotaGet := exec.Command("cf", "quota", Quotas.Quota[i].Name)
						if _, err := QuotaGet.Output(); err != nil{
							fmt.Println("command: ", QuotaGet)
							fmt.Println("Err: ", QuotaGet.Stdout, err)
						} else {
							//			fmt.Println("command: ", QuotaGet)
							fmt.Println(QuotaGet.Stdout)
						}
					} else {
						cmd = exec.Command("cf", "create-quota", Quotas.Quota[i].Name, "-m", MemLimit, "-i", "-1", "-r", "-1", "-s", SerLimit, "-a", AppLimt)
						if _, err := cmd.Output(); err != nil{
							fmt.Println("command: ", cmd)
							fmt.Println("Err: ", cmd.Stdout, err)
						} else {
							//			fmt.Println("command: ", cmd)
							fmt.Println(cmd.Stdout)
						}
						QuotaGet := exec.Command("cf", "quota", Quotas.Quota[i].Name)
						if _, err := QuotaGet.Output(); err != nil{
							fmt.Println("command: ", QuotaGet)
							fmt.Println("Err: ", QuotaGet.Stdout, err)
						} else {
							//			fmt.Println("command: ", QuotaGet)
							fmt.Println(QuotaGet.Stdout)
						}
					}

				} else {

					//fmt.Println("Updating Quota")

					var count, NewMemLimit int

					result1, _ := regexp.MatchString("m", strings.ToLower(MemLimit))
					if result1 == true{
						NewMemLimit, _ = strconv.Atoi(strings.Trim(strings.ToLower(MemLimit),"m"))
					}

					result2, _ := regexp.MatchString("mb", strings.ToLower(MemLimit))
					if result2 == true {
						NewMemLimit, _ = strconv.Atoi(strings.Trim(strings.ToLower(MemLimit), "mb"))
					}

					result3, _ := regexp.MatchString("g", strings.ToLower(MemLimit))
					if result3 == true {
						NewMemLimit, _ = strconv.Atoi(strings.Trim(strings.ToLower(MemLimit), "g"))
						NewMemLimit = 1024*NewMemLimit
					}

					result4, _ := regexp.MatchString("gb", strings.ToLower(MemLimit))
					if result4 == true {
						NewMemLimit, _ = strconv.Atoi(strings.Trim(strings.ToLower(MemLimit), "gb"))
						NewMemLimit = 1024 * NewMemLimit
					}

					if Quotas.Quota[i].MemoryLimit == "" {
						MemLimit = "1024M"
					} else {
					}

					if strings.TrimSpace(string(quotalistjson.Resources[0].Apps.TotalMemoryInMb)) !=  strings.TrimSpace(string(NewMemLimit)) {
						fmt.Println("memory -", quotalistjson.Resources[0].Apps.TotalMemoryInMb)
						if strings.TrimSpace(string(NewMemLimit)) == "" {
							fmt.Println("memory +", MemLimit)
						} else {
							fmt.Println("memory +", strconv.Itoa(NewMemLimit))
						}
						count = 1
					}
					if quotalistjson.Resources[0].Services.PaidServicesAllowed != Quotas.Quota[i].AllowPaidPlans {
						fmt.Println("allow service plans -", quotalistjson.Resources[0].Services.PaidServicesAllowed)
						fmt.Println("allow service plans +", Quotas.Quota[i].AllowPaidPlans)
						count = 1
					}
					if strconv.Itoa(quotalistjson.Resources[0].Apps.TotalInstances) != AppLimt {
						fmt.Println("total app instances -", quotalistjson.Resources[0].Apps.TotalInstances)
						fmt.Println("total app instances +", AppLimt)
						count = 1
					}
					if strconv.Itoa(quotalistjson.Resources[0].Services.TotalServiceInstances) != SerLimit {
						fmt.Println("total service instances -", quotalistjson.Resources[0].Services.TotalServiceInstances)
						fmt.Println("total service instances +", SerLimit)
						count = 1
					}

					if count == 1 {
						fmt.Println("Updating Quota")
						if Quotas.Quota[i].AllowPaidPlans == true {
							cmd = exec.Command("cf", "update-quota", Quotas.Quota[i].Name, "-m", MemLimit, "-i", "-1", "-r", "-1", "-s",  SerLimit, "-a", AppLimt, "--allow-paid-service-plans")
							if _, err := cmd.Output(); err != nil{
								fmt.Println("command: ", cmd)
								fmt.Println("Err: ", cmd.Stdout, err)
							} else {
								//fmt.Println("command: ", cmd)
								fmt.Println(cmd.Stdout)
							}
							QuotaGet := exec.Command("cf", "quota", Quotas.Quota[i].Name)
							if _, err := QuotaGet.Output(); err != nil{
								fmt.Println("command: ", QuotaGet)
								fmt.Println("Err: ", QuotaGet.Stdout, err)
							} else {
								//fmt.Println("command: ", QuotaGet)
								fmt.Println(QuotaGet.Stdout)
							}

						} else {
							cmd = exec.Command("cf", "update-quota", Quotas.Quota[i].Name, "-m", MemLimit, "-i", "-1", "-r", "-1", "-s",  SerLimit, "-a", AppLimt, "--disallow-paid-service-plans")
							if _, err := cmd.Output(); err != nil{
								fmt.Println("command: ", cmd)
								fmt.Println("Err: ", cmd.Stdout, err)
							} else {
								//fmt.Println("command: ", cmd)
								fmt.Println(cmd.Stdout)
							}
							QuotaGet := exec.Command("cf", "quota", Quotas.Quota[i].Name)
							if _, err := QuotaGet.Output(); err != nil{
								fmt.Println("command: ", QuotaGet)
								fmt.Println("Err: ", QuotaGet.Stdout, err)
							} else {
								//fmt.Println("command: ", QuotaGet)
								fmt.Println(QuotaGet.Stdout)
							}

						}
					}
				}
			}
		} else {
			fmt.Println("This is a protected Quota")
		}
	}
	results := exec.Command("cf", "quotas")
	if _, err := results.Output(); err != nil{
		fmt.Println("command: ", results)
		fmt.Println("Err: ", results.Stdout, err)
	} else {
		//fmt.Println("command: ", results)
		fmt.Println(results.Stdout)
	}
	return err
}
