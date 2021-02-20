package main

//import "C"
import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
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
	OrgList []string `yaml:"OrgList"`
	Audit string `yaml:"Audit"`
}
type Orglist struct {
	Org struct {
		Name     string `yaml:"Name"`
		Quota    string `yaml:"Quota"`
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
	} `yaml:"ClusterDetails"`
}
func main()  {

	var endpoint, user, pwd, org, space, asg, operation, cpath, ostype string
	var ospath io.Writer

	flag.StringVar(&endpoint, "e", "api.sys-domain", "Use with init operation, Provide PCF Endpoint")
	flag.StringVar(&user, "u", "user", "Use with init operation, Provide UserName")
	flag.StringVar(&pwd, "p", "pwd", "Use with all operation, Provide Password")
	flag.StringVar(&org, "o", "org", "Use with init operation, Provide Org")
	flag.StringVar(&space, "s", "space", "Use with init operation, Provide Space")
	flag.StringVar(&asg, "a", "true", "Use with init operation, Enable ASGs ?.")
	flag.StringVar(&operation, "i", "", "Provide Operation to be performed: init, create-{org,space,org-user,space-user,quota, ")
	flag.StringVar(&cpath, "k", ".", "Provide path to configs, i.e, to config folder, use with all operations")
	flag.Parse()

	ClusterName := strings.ReplaceAll(endpoint, ".", "-")

	fmt.Printf("Operation: %v\n", operation)

	oscmd := exec.Command("cmd", "/C","echo","%systemdrive%%homepath%")
	if _, err := oscmd.Output(); err != nil{
		fmt.Println("Checking OS")
		fmt.Println("command: ", oscmd)
		fmt.Println(oscmd.Stdout, err)
		oscmd = exec.Command("sh", "-c", "echo","$HOME")
		if _, err := oscmd.Output(); err != nil{
			fmt.Println("Checking OS failed - Can't find Underlying OS")
			fmt.Println("command: ", oscmd)
			fmt.Println(oscmd.Stdout, err)
			panic(err)
		} else {
			fmt.Println("command: ", oscmd)
			ospath = oscmd.Stdout
			fmt.Println("PATH: ", ospath)
			fmt.Println("Checking OS - Setting up for Mac/Linux/Ubuntu")
			ostype = "non-windows"
		}
	} else {
		fmt.Println("command: ", oscmd)
		ospath = oscmd.Stdout
		fmt.Println("PATH: ", ospath)
		fmt.Println("Checking OS - Setting up for Windows")
		ostype = "windows"
	}

	if operation == "init" {

		fmt.Println("Initializing C9Cli")

		fmt.Printf("ClusterName: %v\n", ClusterName)
		fmt.Printf("EndPoint: %v\n", endpoint)
		fmt.Printf("User: %v\n", user)
		fmt.Printf("Org: %v\n", org)
		fmt.Printf("Space: %v\n", space)
		fmt.Printf("EnableASG: %v\n", asg)
		fmt.Printf("Path: %v\n", cpath)
		Init(ClusterName, endpoint, user, org, space, asg, cpath)
	} else if operation == "org-init" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		OrgsInit(ClusterName, cpath)

	} else if operation == "audit-quota"{

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection (ClusterName, pwd, cpath)
		DeleteOrAuditQuotas (ClusterName, cpath)

	} else if operation == "audit-org"{

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection (ClusterName, pwd, cpath)
		DeleteorAuditOrgs (ClusterName, cpath)

	} else if operation == "audit-space"{

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection (ClusterName, pwd, cpath)
		DeleteorAuditSpaces (ClusterName, cpath, ostype)

	}  else if operation == "audit-org-user"{

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection (ClusterName, pwd, cpath)
		DeleteorAuditOrgUsers (ClusterName, cpath, ostype)

	} else if operation == "audit-space-user"{

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection (ClusterName, pwd, cpath)
		DeleteOrAuditSpaceUsers (ClusterName, cpath, ostype)

	} else if operation == "audit-asg"{

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection (ClusterName, pwd, cpath)
		DeleteOrAuditSpacesASGs (ClusterName, cpath, ostype)

	} else if operation == "create-org"{

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection (ClusterName, pwd, cpath)
		CreateOrUpdateOrgs (ClusterName, cpath)

	} else if operation == "create-quota" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection (ClusterName,  pwd, cpath)
		CreateOrUpdateQuotas(ClusterName, cpath)

	} else if operation == "create-org-user" {

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection(ClusterName,  pwd, cpath)
		CreateOrUpdateOrgUsers(ClusterName, cpath)
	} else if operation == "create-space"{

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection (ClusterName,  pwd, cpath)
		CreateOrUpdateSpaces (ClusterName, cpath, ostype)

	} else if operation == "create-space-user"{

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection (ClusterName,  pwd, cpath)
		CreateOrUpdateSpaceUsers (ClusterName, cpath)

	} else if operation == "create-protected-org-asg"{

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection (ClusterName,  pwd, cpath)
		CreateOrUpdateProtOrgAsg (ClusterName, cpath, ostype)

	} else if operation == "create-space-asg"{

		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection (ClusterName,  pwd, cpath)
		CreateOrUpdateSpacesASGs (ClusterName, cpath, ostype)

	}else if operation == "create-all" {
		fmt.Printf("ClusterName: %v\n", ClusterName)
		SetupConnection (ClusterName,  pwd, cpath)
		CreateOrUpdateProtOrgAsg (ClusterName, cpath, ostype)
		CreateOrUpdateQuotas(ClusterName, cpath)
		CreateOrUpdateOrgs (ClusterName, cpath)
		CreateOrUpdateOrgUsers(ClusterName, cpath)
		CreateOrUpdateSpaces (ClusterName, cpath, ostype)
		CreateOrUpdateSpacesASGs (ClusterName, cpath, ostype)
		CreateOrUpdateSpaceUsers (ClusterName, cpath)
	} else {
		fmt.Println("Provide Valid input operation")
	}
}
func SetupConnection(clustername string, pwd string, cpath string) error {

	var InitClusterConfigVals InitClusterConfigVals
	ConfigFile := cpath+"/"+clustername+"/config.yml"

	fileConfigYml, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileConfigYml), &InitClusterConfigVals)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Endpoint: %v\n", InitClusterConfigVals.ClusterDetails.EndPoint)
	fmt.Printf("User: %v\n", InitClusterConfigVals.ClusterDetails.User)
	fmt.Printf("Pwd: %v\n", pwd)
	fmt.Printf("Org: %v\n", InitClusterConfigVals.ClusterDetails.Org)
	fmt.Printf("Space: %v\n", InitClusterConfigVals.ClusterDetails.Space)
	//fmt.Println(InitClusterConfigVals.ClusterDetails.EndPoint)

	cmd := exec.Command("cf", "login", "-a", InitClusterConfigVals.ClusterDetails.EndPoint, "-u", InitClusterConfigVals.ClusterDetails.User, "-p", pwd, "-o", InitClusterConfigVals.ClusterDetails.Org, "-s", InitClusterConfigVals.ClusterDetails.Space, "--skip-ssl-validation")
	if _, err := cmd.Output(); err != nil{
		fmt.Println("Connection failed")
		fmt.Println("command: ", cmd)
		fmt.Println(cmd.Stdout, err)
		panic(err)
	} else {
		fmt.Println("Connection Passed")
		fmt.Println("command: ", cmd)
		fmt.Println(cmd.Stdout)
	}
	return err
}
func DeleteOrAuditQuotas(clustername string, cpath string) error {

	var Quotas Quotalist
	var ProtectedQuota ProtectedList

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
	Audit := Quotas.Audit

	getquotas := exec.Command("cf", "curl", "/v3/organization_quotas", "--output", "listquotas.json")

	if _, err := getquotas.Output(); err == nil {
		fmt.Println("command: ", getquotas)

		var body QuotaListJson

		fileOrgJson, err := ioutil.ReadFile("listquotas.json")
		if err != nil {
			fmt.Println(err)
		}

		if err := json.Unmarshal(fileOrgJson, &body); err != nil {
			panic(err)
		}

		QuotaLen := len(body.Resources)

		fmt.Println("No of Quotas: ", QuotaLen)

		if QuotaLen != 0 {

			for i := 0; i < QuotaLen; i++ {

				var count, totalcount int
				fmt.Println(" ")
				fmt.Println("Quota: ", body.Resources[i].Name)

				for p := 0; p < LenProtectedQuota; p++ {
					fmt.Println("Protected Quota: ", ProtectedQuota.Quota[p], ",", body.Resources[i].Name)
					if strings.Trim(ProtectedQuota.Quota[p], "") == strings.Trim(body.Resources[i].Name, "") {
						count = 1
					} else {
						count = 0
					}
					totalcount = totalcount + count
				}

				if totalcount == 0 {

					fmt.Println("This is not Protected Quota")

					var quotascount, quotastotalcount int

					for q := 0; q < LenQuota; q++ {

						fmt.Println("Quota: ", Quotas.Quota[q].Name,",", body.Resources[i].Name)
						if Quotas.Quota[q].Name == body.Resources[i].Name {
							quotascount = 1
						} else {
							quotastotalcount = 0
						}
						quotastotalcount = quotastotalcount + quotascount
					}

					if quotastotalcount == 0 {
						fmt.Println("Quota has not be listed in Quota.yml: ")
						fmt.Println("Auditing Quota: ", body.Resources[i].Name)
						if Audit == "Delete" {
							fmt.Println("DELETE!DELETE!")
							fmt.Println("Deleting Quota: ", body.Resources[i].Name)
							delete := exec.Command("cf", "delete-quota", body.Resources[i].Name, "-f")
							if _, err := delete.Output(); err == nil {
								fmt.Println("command: ", delete)
								fmt.Println(delete.Stdout)
							} else {
								fmt.Println("command: ", delete)
								fmt.Println("Err: ", delete.Stdout, delete.Stderr)
							}
						} else if Audit == "Rename" {
							fmt.Println("DELETE!DELETE!")
							fmt.Println("Renaming Quota: ", body.Resources[i].Name)
							result, _ := regexp.MatchString("_tobedeleted", body.Resources[i].Name)
							if result == true{
								fmt.Println("Quota already renamed")
							} else {
								rename := exec.Command("cf", "update-quota", body.Resources[i].Name, "-n", body.Resources[i].Name+"_tobedeleted")
								if _, err := rename.Output(); err == nil {
									fmt.Println("command: ", rename)
									fmt.Println(rename.Stdout)
								} else {
									fmt.Println("command: ", rename)
									fmt.Println("Err: ", rename.Stdout, rename.Stderr)
								}
							}

						} else if Audit == "List" {
							fmt.Println("DELETE!DELETE!")
							fmt.Println("Quota to be deleted: ", body.Resources[i].Name)
						} else {
							fmt.Println("Provide Valid Input")
						}
					} else {
						fmt.Println("Quota exists in Quota.yml: ", body.Resources[i].Name)
					}
				} else {
					fmt.Println("This is a protected Quota:", body.Resources[i].Name)
				}
			}
		} else {
			fmt.Println("No Quota exist")
		}

	} else {
		fmt.Println("command: ", getquotas)
		fmt.Println("Err: ", getquotas.Stderr)
	}
	return err
}
func DeleteorAuditOrgs(clustername string, cpath string) error {

	var list List
	var ProtectedOrgs ProtectedList

	ListYml := cpath+"/"+clustername+"/OrgsList.yml"
	fileOrgYml, err := ioutil.ReadFile(ListYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileOrgYml), &list)
	if err != nil {
		panic(err)
	}
	LenList := len(list.OrgList)
	Audit := list.Audit

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

	getorgs := exec.Command("cf", "curl", "/v3/organizations", "--output", "listorgs.json")

	if _, err := getorgs.Output(); err == nil {
		fmt.Println("command: ", getorgs)

		var body OrgListJson

		fileOrgJson, err := ioutil.ReadFile("listorgs.json")
		if err != nil {
			fmt.Println(err)
		}

		if err := json.Unmarshal(fileOrgJson, &body); err != nil {
			panic(err)
		}

		OrgsLen := len(body.Resources)

		fmt.Println("No of Orgs: ", OrgsLen)

		if OrgsLen != 0 {

			for i := 0; i < OrgsLen; i++ {

				var count, totalcount int
				fmt.Println("Org: ", body.Resources[i].Name)

				for p := 0; p < LenProtectedOrgs; p++ {
					fmt.Println("Protected Org: ", ProtectedOrgs.Org[p], ",", body.Resources[i].Name)
					if strings.Trim(ProtectedOrgs.Org[p], "") == strings.Trim(body.Resources[i].Name, "") {
						count = 1
					} else {
						count = 0
					}
					totalcount = totalcount + count
				}

				if totalcount == 0 {

					fmt.Println("This is not Protected Org")

					var orgscount, orgstotalcount int

					for q := 0; q < LenList; q++ {

						fmt.Println("Org: ", list.OrgList[q], ",", body.Resources[i].Name)
						if list.OrgList[q] == body.Resources[i].Name {
							orgscount = 1
						} else {
							orgscount = 0
						}
						orgstotalcount = orgstotalcount + orgscount
					}

					if orgstotalcount == 0 {
						fmt.Println("Org has not be listed in Orglist.yml: ")
						fmt.Println("Auditing Org: ", body.Resources[i].Name)
						if Audit == "Delete" {
							delete := exec.Command("cf", "delete-org", body.Resources[i].Name)
							if _, err := delete.Output(); err == nil {
								fmt.Println("command: ", delete)
								fmt.Println(delete.Stdout)
							} else {
								fmt.Println("command: ", delete)
								fmt.Println("Err: ", delete.Stdout, delete.Stderr)
							}
						} else if Audit == "Rename" {
							rename := exec.Command("cf", "rename-org", body.Resources[i].Name, body.Resources[i].Name+"tobedeleted")
							if _, err := rename.Output(); err == nil {
								fmt.Println("command: ", rename)
								fmt.Println(rename.Stdout)
							} else {
								fmt.Println("command: ", rename)
								fmt.Println("Err: ", rename.Stdout, rename.Stderr)
							}
						} else if Audit == "List" {
							fmt.Println("Org to be deleted: ", body.Resources[i].Name)
						} else {
							fmt.Println("Provide Valid Input")
						}
					} else {
						fmt.Println("Org exists in Orgslist.yml: ", body.Resources[i].Name)
					}
				} else {
					fmt.Println("This is a protected Org:", body.Resources[i].Name)
				}
			}
		} else {
			fmt.Println("No Orgs exist")
		}

	} else {
		fmt.Println("command: ", getorgs)
		fmt.Println("Err: ", getorgs.Stderr)
	}
	return err
}
func DeleteorAuditSpaces(clustername string, cpath string, ostype string) error {

	var Orgs Orglist
	var ProtectedOrgs ProtectedList
	var list List
	var spacelistjson SpaceListJson

	// Org List

	ListYml := cpath+"/"+clustername+"/OrgsList.yml"
	fileOrgYml, err := ioutil.ReadFile(ListYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileOrgYml), &list)
	if err != nil {
		panic(err)
	}

	var InitClusterConfigVals InitClusterConfigVals

	//Config File
	ConfigFile := cpath+"/"+clustername+"/config.yml"
	fileConfigYml, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileConfigYml), &InitClusterConfigVals)
	if err != nil {
		panic(err)
	}
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

	LenList := len(list.OrgList)
	LenProtectedOrgs := len(ProtectedOrgs.Org)

	for i := 0; i < LenList; i++ {

		var count, totalcount int
		fmt.Println("Org: ", list.OrgList[i])
		for p := 0; p < LenProtectedOrgs; p++ {
			fmt.Println("Protected Org: ", ProtectedOrgs.Org[p])
			if ProtectedOrgs.Org[p] == list.OrgList[i] {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {

			fmt.Println("This is not Protected Org")

			if ostype == "windows" {
				OrgsYml = cpath+"\\"+clustername+"\\"+list.OrgList[i]+"\\Org.yml"
			} else {
				OrgsYml = cpath+"/"+clustername+"/"+list.OrgList[i]+"/Org.yml"
			}

			fileOrgYml, err := ioutil.ReadFile(OrgsYml)
			if err != nil {
				fmt.Println(err)
			}
			err = yaml.Unmarshal([]byte(fileOrgYml), &Orgs)
			if err != nil {
				panic(err)
			}

			Audit := Orgs.SpaceAudit

			if list.OrgList[i] == Orgs.Org.Name {

				var out bytes.Buffer
				guid := exec.Command("cf", "org", Orgs.Org.Name, "--guid")
				guid.Stdout = &out

				err := guid.Run()

				if err == nil {

					fmt.Println("Org exists: ", Orgs.Org.Name,",", out.String())
					path := "/v3/spaces/?organization_guids="+out.String()
					spacelist := exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "spacelist.json")

					var out bytes.Buffer
					spacelist.Stdout = &out
					err := spacelist.Run()
					if err == nil {
						fmt.Println(spacelist, spacelist.Stdout, spacelist.Stderr)
					} else {
						fmt.Println("err", spacelist, spacelist.Stdout, spacelist.Stderr)

					}

					fileSpaceJson, err := ioutil.ReadFile("spacelist.json")
					if err != nil {
						fmt.Println(err)
					}
					if err := json.Unmarshal(fileSpaceJson, &spacelistjson); err != nil {
						panic(err)
					}

					SpaceLen := len(spacelistjson.Resources)
					fmt.Println("No of Spaces for Org",Orgs.Org.Name,":", SpaceLen)

					if SpaceLen != 0 {

						for i := 0; i < SpaceLen; i++ {

							SpaceLen := len(Orgs.Org.Spaces)
							var Spacescount, Spacestotalcount int

							for q := 0; q < SpaceLen; q++ {

								fmt.Println("Space: ", Orgs.Org.Spaces[q].Name, ",", spacelistjson.Resources[i].Name)
								if Orgs.Org.Spaces[q].Name == spacelistjson.Resources[i].Name {
									Spacescount = 1
								} else {
									Spacescount = 0
								}
								Spacestotalcount = Spacestotalcount + Spacescount
							}

							if Spacestotalcount == 0 {
//								fmt.Println("Space has not be listed in Orgs.yml")
//								fmt.Println("Delete Space: ", spacelistjson.Resources[i].Name)
								fmt.Println("Space has not be listed in Org.yml: ")
								fmt.Println("Auditing Space: ", spacelistjson.Resources[i].Name)
								if Audit == "Delete" {
									delete := exec.Command("cf", "delete-space", spacelistjson.Resources[i].Name, "-o", Orgs.Org.Name)
									if _, err := delete.Output(); err == nil {
										fmt.Println("command: ", delete)
										fmt.Println(delete.Stdout)
									} else {
										fmt.Println("command: ", delete)
										fmt.Println("Err: ", delete.Stdout, delete.Stderr)
									}
								} else if Audit == "Rename" {
									rename := exec.Command("cf", "rename", spacelistjson.Resources[i].Name, spacelistjson.Resources[i].Name+"tobedeleted")
									if _, err := rename.Output(); err == nil {
										fmt.Println("command: ", rename)
										fmt.Println(rename.Stdout)
									} else {
										fmt.Println("command: ", rename)
										fmt.Println("Err: ", rename.Stdout, rename.Stderr)
									}
								} else if Audit == "List" {
									fmt.Println("Space to be deleted: ", spacelistjson.Resources[i].Name)
								} else {
									fmt.Println("Provide Valid Input")
								}
							} else {
								fmt.Println("Space exists in Orgs.yml: ", spacelistjson.Resources[i].Name)
							}
						}
					} else {
						fmt.Println("No Spaces Exist")
					}
				} else {
					fmt.Println("command: ", guid )
					fmt.Println("Err: ", guid.Stdout)
					fmt.Println("Err Code: ", err)
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
func DeleteorAuditOrgUsers(clustername string, cpath string, ostype string) error {

	var Orgs Orglist
	var ProtectedOrgs ProtectedList
	var list List
	var orgusrslist OrgUsersListJson

	// Org List
	ListYml := cpath+"/"+clustername+"/OrgsList.yml"
	fileOrgYml, err := ioutil.ReadFile(ListYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileOrgYml), &list)
	if err != nil {
		panic(err)
	}

	var InitClusterConfigVals InitClusterConfigVals

	//Config File
	ConfigFile := cpath+"/"+clustername+"/config.yml"
	fileConfigYml, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileConfigYml), &InitClusterConfigVals)
	if err != nil {
		panic(err)
	}
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

	LenList := len(list.OrgList)
	LenProtectedOrgs := len(ProtectedOrgs.Org)

	for i := 0; i < LenList; i++ {

		var count, totalcount int
		fmt.Println("Org: ", list.OrgList[i])
		for p := 0; p < LenProtectedOrgs; p++ {
			fmt.Println("Protected Org: ", ProtectedOrgs.Org[p])
			if ProtectedOrgs.Org[p] == list.OrgList[i] {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {

			fmt.Println("This is not Protected Org")

			if ostype == "windows" {
				OrgsYml = cpath+"\\"+clustername+"\\"+list.OrgList[i]+"\\Org.yml"
			} else {
				OrgsYml = cpath+"/"+clustername+"/"+list.OrgList[i]+"/Org.yml"
			}

			fileOrgYml, err := ioutil.ReadFile(OrgsYml)
			if err != nil {
				fmt.Println(err)
			}
			err = yaml.Unmarshal([]byte(fileOrgYml), &Orgs)
			if err != nil {
				panic(err)
			}

			Audit := Orgs.UserAudit

			if list.OrgList[i] == Orgs.Org.Name {

				var out bytes.Buffer

				guid := exec.Command("cf", "org", Orgs.Org.Name, "--guid")
				guid.Stdout = &out

				err := guid.Run()

				if err == nil {

					fmt.Println("Org exists: ", Orgs.Org.Name,",", out.String())

					path := "/v3/roles/?organization_guids="+out.String()
					orguserslist := exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "orgusrslist.json")

					var out bytes.Buffer
					orguserslist.Stdout = &out
					err := orguserslist.Run()
					if err == nil {
						fmt.Println(orguserslist, orguserslist.Stdout, orguserslist.Stderr)
					} else {
						fmt.Println("err", orguserslist, orguserslist.Stdout, orguserslist.Stderr)
					}

					fileSpaceJson, err := ioutil.ReadFile("orgusrslist.json")
					if err != nil {
						fmt.Println(err)
					}
					if err := json.Unmarshal(fileSpaceJson, &orgusrslist); err != nil {
						panic(err)
					}

					OrgUsrLen := len(orgusrslist.Resources)
					fmt.Println("No of Users currently exist in Org",Orgs.Org.Name,":", OrgUsrLen)

					if OrgUsrLen != 0 {

						for i := 0; i < OrgUsrLen; i++ {

							//SSO
							OrgUsLenSSOAuditor := len(Orgs.Org.OrgUsers.SSO.OrgAuditors)

							var orgusrssoauditscount, aorusrssoaudittotalcount int

							if orgusrslist.Resources[i].Type == "organization_auditor"{

								userguid := orgusrslist.Resources[i].GUID
								path := "/v3/users/?guids="+userguid
								var out bytes.Buffer
								username := exec.Command("cf", "curl", strings.TrimSpace(path))
								username.Stdout = &out
								err := username.Run()

								if err == nil {
									for q := 0; q < OrgUsLenSSOAuditor; q++ {

										if out.String() == "admin" {
											fmt.Println("Admin user skipping Validation")
											aorusrssoaudittotalcount = 2
										} else {
											fmt.Println("SSO Audit Usr: ", Orgs.Org.OrgUsers.SSO.OrgAuditors[q], ",", username)
											if strings.TrimSpace(Orgs.Org.OrgUsers.SSO.OrgAuditors[q]) == strings.TrimSpace(out.String()) {
												orgusrssoauditscount = 1
											} else {
												orgusrssoauditscount = 0
											}
											aorusrssoaudittotalcount = aorusrssoaudittotalcount + orgusrssoauditscount
										}
									}

									if aorusrssoaudittotalcount == 0 {

										fmt.Println("User has not be listed in Org.yml: ")
										fmt.Println("SSO OrgAuditor User: ", strings.TrimSpace(out.String()))

										if Audit == "unset" {
											unset := exec.Command("cf", "unset-org-role", strings.TrimSpace(out.String()), Orgs.Org.Name, "OrgAuditor")
											if _, err := unset.Output(); err == nil {
												fmt.Println("command: ", unset)
												fmt.Println(unset.Stdout)
											} else {
												fmt.Println("command: ", unset)
												fmt.Println("Err: ", unset.Stdout, unset.Stderr)
											}
										} else if Audit == "List" {
											fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "SSO OrgAuditor")
										} else {
											fmt.Println("Provide Valid Input")
										}
									} else if aorusrssoaudittotalcount ==2 {
										fmt.Println("Admin set as SSO Org Audit User, skipping unset: ", strings.TrimSpace(out.String()))
									} else {
										fmt.Println("User is listed in Org.yml as SSO Org Audit User: ", strings.TrimSpace(out.String()))
									}
								} else {
									fmt.Println("err", username, username.Stdout, username.Stderr)
								}
							}

							OrgUsLenSSOManagers := len(Orgs.Org.OrgUsers.SSO.OrgManagers)

							var orgusrssomangscount, aorusrssomangtotalcount int

							if orgusrslist.Resources[i].Type == "organization_managers"{

								userguid := orgusrslist.Resources[i].GUID
								path := "/v3/users/?guids="+userguid
								var out bytes.Buffer
								username := exec.Command("cf", "curl", strings.TrimSpace(path))
								username.Stdout = &out
								err := username.Run()

								if err == nil {
									for q := 0; q < OrgUsLenSSOManagers; q++ {

										if out.String() == "admin" {
											fmt.Println("Admin user skipping Validation")
											aorusrssomangtotalcount = 2
										} else {

											fmt.Println("SSO Org Managers Usr: ", Orgs.Org.OrgUsers.SSO.OrgManagers[q], ",", username)
											if strings.TrimSpace(Orgs.Org.OrgUsers.SSO.OrgAuditors[q]) == strings.TrimSpace(out.String()) {
												orgusrssomangscount = 1
											} else {
												orgusrssomangscount = 0
											}
											aorusrssomangtotalcount = aorusrssomangtotalcount + orgusrssomangscount
										}
									}
									if aorusrssomangtotalcount == 0 {

										fmt.Println("User has not be listed in Org.yml: ")
										fmt.Println("SSO OrgManager User: ", strings.TrimSpace(out.String()))

										if Audit == "unset" {
											unset := exec.Command("cf", "unset-org-role", strings.TrimSpace(out.String()), Orgs.Org.Name, "OrgManager")
											if _, err := unset.Output(); err == nil {
												fmt.Println("command: ", unset)
												fmt.Println(unset.Stdout)
											} else {
												fmt.Println("command: ", unset)
												fmt.Println("Err: ", unset.Stdout, unset.Stderr)
											}
										} else if Audit == "List" {
											fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "SSO OrgManager")
										} else {
											fmt.Println("Provide Valid Input")
										}
									} else if aorusrssomangtotalcount ==2 {
										fmt.Println("Admin set as SSO Org Manager User, skipping unset: ", strings.TrimSpace(out.String()))
									} else {
										fmt.Println("User is listed in Org.yml as SSO Org Manager User: ", strings.TrimSpace(out.String()))
									}
								} else {
									fmt.Println("err", username, username.Stdout, username.Stderr)
								}
							}

							//UAA
							OrgUsLenUAAAuditor := len(Orgs.Org.OrgUsers.UAA.OrgAuditors)

							var orgusruaaauditscount, aorusruaaaudittotalcount int

							if orgusrslist.Resources[i].Type == "organization_auditor"{

								userguid := orgusrslist.Resources[i].GUID
								path := "/v3/users/?guids="+userguid
								var out bytes.Buffer
								username := exec.Command("cf", "curl", strings.TrimSpace(path))
								username.Stdout = &out
								err := username.Run()

								if err == nil {
									for q := 0; q < OrgUsLenUAAAuditor; q++ {

										if out.String() == "admin" {
											fmt.Println("Admin user skipping Validation")
											aorusruaaaudittotalcount = 2
										} else {
											fmt.Println("UAA Audit Usr: ", Orgs.Org.OrgUsers.UAA.OrgAuditors[q], ",", username)
											if strings.TrimSpace(Orgs.Org.OrgUsers.UAA.OrgAuditors[q]) == strings.TrimSpace(out.String()) {
												orgusruaaauditscount = 1
											} else {
												orgusruaaauditscount = 0
											}
											aorusruaaaudittotalcount = aorusruaaaudittotalcount + orgusruaaauditscount
										}
									}

									if aorusruaaaudittotalcount == 0 {
										fmt.Println("User has not be listed in Org.yml: ")
										fmt.Println("UAA OrgAuditor User: ", strings.TrimSpace(out.String()))

										if Audit == "unset" {
											unset := exec.Command("cf", "unset-org-role", strings.TrimSpace(out.String()), Orgs.Org.Name, "OrgAuditor")
											if _, err := unset.Output(); err == nil {
												fmt.Println("command: ", unset)
												fmt.Println(unset.Stdout)
											} else {
												fmt.Println("command: ", unset)
												fmt.Println("Err: ", unset.Stdout, unset.Stderr)
											}
										} else if Audit == "List" {
											fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "UAA OrgAuditor")
										} else {
											fmt.Println("Provide Valid Input")
										}
									} else if aorusruaaaudittotalcount ==2 {
										fmt.Println("Admin set as UAA Org Audit User, skipping unset: ", strings.TrimSpace(out.String()))
									} else {
										fmt.Println("User is listed in Org.yml as UAA Org Audit User: ", strings.TrimSpace(out.String()))
									}
								} else {
									fmt.Println("err", username, username.Stdout, username.Stderr)
								}
							}


							OrgUsLenUAAManagers := len(Orgs.Org.OrgUsers.UAA.OrgManagers)

							var orgusruaamangscount, aorusruaamangtotalcount int

							if orgusrslist.Resources[i].Type == "organization_managers"{

								userguid := orgusrslist.Resources[i].GUID
								path := "/v3/users/?guids="+userguid
								var out bytes.Buffer
								username := exec.Command("cf", "curl", strings.TrimSpace(path))
								username.Stdout = &out
								err := username.Run()

								if err == nil {
									for q := 0; q < OrgUsLenUAAManagers; q++ {

										if out.String() == "admin" {
											fmt.Println("Admin user skipping Validation")
											aorusruaamangtotalcount = 2
										} else {

											fmt.Println("UAA Org Managers Usr: ", Orgs.Org.OrgUsers.UAA.OrgManagers[q], ",", username)
											if strings.TrimSpace(Orgs.Org.OrgUsers.UAA.OrgAuditors[q]) == strings.TrimSpace(out.String()) {
												orgusruaamangscount = 1
											} else {
												orgusruaamangscount = 0
											}
											aorusruaamangtotalcount = aorusruaamangtotalcount + orgusruaamangscount
										}
									}
									if aorusruaamangtotalcount == 0 {

										fmt.Println("User has not be listed in Org.yml: ")
										fmt.Println("UAA OrgManager User: ", strings.TrimSpace(out.String()))

										if Audit == "unset" {
											unset := exec.Command("cf", "unset-org-role", strings.TrimSpace(out.String()), Orgs.Org.Name, "OrgManager")
											if _, err := unset.Output(); err == nil {
												fmt.Println("command: ", unset)
												fmt.Println(unset.Stdout)
											} else {
												fmt.Println("command: ", unset)
												fmt.Println("Err: ", unset.Stdout, unset.Stderr)
											}
										} else if Audit == "List" {
											fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "UAA OrgManager")
										} else {
											fmt.Println("Provide Valid Input")
										}

									} else if aorusruaamangtotalcount ==2 {
										fmt.Println("Admin set as UAA Org Manager User, skipping unset: ", strings.TrimSpace(out.String()))
									} else {
										fmt.Println("User is listed in Org.yml as UAA Org Manager User: ", strings.TrimSpace(out.String()))
									}
								} else {
									fmt.Println("err", username, username.Stdout, username.Stderr)
								}
							}

							//LDAP

							OrgUsLenLDAPAuditor := len(Orgs.Org.OrgUsers.LDAP.OrgAuditors)
							var orgusrldapauditscount, aorusrldapaudittotalcount int

							if orgusrslist.Resources[i].Type == "organization_auditor"{

								userguid := orgusrslist.Resources[i].GUID
								path := "/v3/users/?guids="+userguid
								var out bytes.Buffer
								username := exec.Command("cf", "curl", strings.TrimSpace(path))
								username.Stdout = &out
								err := username.Run()

								if err == nil {
									for q := 0; q < OrgUsLenLDAPAuditor; q++ {

										if out.String() == "admin" {
											fmt.Println("Admin user skipping Validation")
											aorusrldapaudittotalcount = 2
										} else {
											fmt.Println("LDAP Audit Usr: ", Orgs.Org.OrgUsers.LDAP.OrgAuditors[q], ",", username)
											if strings.TrimSpace(Orgs.Org.OrgUsers.LDAP.OrgAuditors[q]) == strings.TrimSpace(out.String()) {
												orgusrldapauditscount = 1
											} else {
												orgusrldapauditscount = 0
											}
											aorusrldapaudittotalcount = aorusrldapaudittotalcount + orgusrldapauditscount

										}
									}

									if aorusrldapaudittotalcount == 0 {
										fmt.Println("User has not be listed in Org.yml: ")
										fmt.Println("LDAP OrgAuditor User: ", strings.TrimSpace(out.String()))

										if Audit == "unset" {
											unset := exec.Command("cf", "unset-org-role", strings.TrimSpace(out.String()), Orgs.Org.Name, "OrgAuditor")
											if _, err := unset.Output(); err == nil {
												fmt.Println("command: ", unset)
												fmt.Println(unset.Stdout)
											} else {
												fmt.Println("command: ", unset)
												fmt.Println("Err: ", unset.Stdout, unset.Stderr)
											}
										} else if Audit == "List" {
											fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "LDAP OrgAuditor")
										} else {
											fmt.Println("Provide Valid Input")
										}
									} else if aorusrldapaudittotalcount ==2 {
										fmt.Println("Admin set as LDAP Org Audit User, skipping unset: ", strings.TrimSpace(out.String()))
									} else {
										fmt.Println("User is listed in Org.yml as LDAP Org Audit User: ", strings.TrimSpace(out.String()))
 									}
								} else {
									fmt.Println("err", username, username.Stdout, username.Stderr)
								}
							}


							OrgUsLenLDAPManagers := len(Orgs.Org.OrgUsers.LDAP.OrgManagers)

							var orgusrldapmangscount, aorusrldapmangtotalcount int

							if orgusrslist.Resources[i].Type == "organization_managers"{

								userguid := orgusrslist.Resources[i].GUID
								path := "/v3/users/?guids="+userguid
								var out bytes.Buffer
								username := exec.Command("cf", "curl", strings.TrimSpace(path))
								username.Stdout = &out
								err := username.Run()

								if err == nil {
									for q := 0; q < OrgUsLenLDAPManagers; q++ {

										if out.String() == "admin" {
											fmt.Println("Admin user skipping Validation")
											aorusrldapmangtotalcount = 2
										} else {

											fmt.Println("LDAP Org Managers Usr: ", Orgs.Org.OrgUsers.LDAP.OrgManagers[q], ",", username)
											if strings.TrimSpace(Orgs.Org.OrgUsers.LDAP.OrgAuditors[q]) == strings.TrimSpace(out.String()) {
												orgusrldapmangscount = 1
											} else {
												orgusrldapmangscount = 0
											}
											aorusrldapmangtotalcount = aorusrldapmangtotalcount + orgusrldapmangscount
										}
									}
									if aorusrldapmangtotalcount == 0 {
										fmt.Println("User has not be listed in Org.yml: ")
										fmt.Println("LDAP OrgManager User: ", strings.TrimSpace(out.String()))

										if Audit == "unset" {
											unset := exec.Command("cf", "unset-org-role", strings.TrimSpace(out.String()), Orgs.Org.Name, "OrgManager")
											if _, err := unset.Output(); err == nil {
												fmt.Println("command: ", unset)
												fmt.Println(unset.Stdout)
											} else {
												fmt.Println("command: ", unset)
												fmt.Println("Err: ", unset.Stdout, unset.Stderr)
											}
										} else if Audit == "List" {
											fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "LDAP OrgManager")
										} else {
											fmt.Println("Provide Valid Input")
										}
									} else if aorusrldapmangtotalcount ==2 {
										fmt.Println("Admin set as LDAP Org Manager User, skipping unset: ", strings.TrimSpace(out.String()))
									} else {
										fmt.Println("User is listed in Org.yml as LDAP Org Manager User: ", strings.TrimSpace(out.String()))
									}
								} else {
									fmt.Println("err", username, username.Stdout, username.Stderr)
								}
							}
						}
					} else {
						fmt.Println("No users exist")
					}
				} else {
					fmt.Println("command: ", guid )
					fmt.Println("Err: ", guid.Stdout)
					fmt.Println("Err Code: ", err)
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
func DeleteOrAuditSpaceUsers(clustername string, cpath string, ostype string) error {

	var Orgs Orglist
	var ProtectedOrgs ProtectedList
	var list List
	var spaceusrslist SpaceUsersListJson

	ListYml := cpath+"/"+clustername+"/OrgsList.yml"
	fileOrgYml, err := ioutil.ReadFile(ListYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileOrgYml), &list)
	if err != nil {
		panic(err)
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

	LenList := len(list.OrgList)

	for i := 0; i < LenList; i++ {

		var count, totalcount int
		fmt.Println("Org: ", list.OrgList[i])
		for p := 0; p < LenProtectedOrgs; p++ {
			fmt.Println("Protected Org: ", ProtectedOrgs.Org[p])
			if ProtectedOrgs.Org[p] == list.OrgList[i] {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {


			var OrgsYml string
			if ostype == "windows" {
				OrgsYml = cpath+"\\"+clustername+"\\"+list.OrgList[i]+"\\Org.yml"
			} else {
				OrgsYml = cpath+"/"+clustername+"/"+list.OrgList[i]+"/Org.yml"
			}

			fileOrgYml, err := ioutil.ReadFile(OrgsYml)

			if err != nil {
				fmt.Println(err)
			}

			err = yaml.Unmarshal([]byte(fileOrgYml), &Orgs)
			if err != nil {
				panic(err)
			}

			Audit := Orgs.SpaceAudit

			if list.OrgList[i] == Orgs.Org.Name {
				guid := exec.Command("cf", "org", Orgs.Org.Name, "--guid")

				if _, err := guid.Output(); err == nil {

					fmt.Println("command: ", guid)
					fmt.Println("Org exists: ", guid.Stdout)
					targetOrg := exec.Command("cf", "t", "-o", Orgs.Org.Name)
					if _, err := targetOrg.Output(); err == nil {
						fmt.Println("command: ", targetOrg)
						fmt.Println("Targeted Org: ", targetOrg.Stdout)
					} else {
						fmt.Println("command: ", targetOrg)
						fmt.Println("Err: ", targetOrg.Stdout)
						fmt.Println("Err Code: ", targetOrg.Stderr)
					}

					SpaceLen := len(Orgs.Org.Spaces)

					for j := 0; j < SpaceLen; j++ {

						var out bytes.Buffer
						guid.Stdout = &out
						guid = exec.Command("cf", "space", Orgs.Org.Spaces[j].Name, "--guid")
						err := guid.Run()

						if err == nil {

							fmt.Println("Space exists: ", Orgs.Org.Spaces[j].Name,",", out.String())

							path := "/v3/roles/?space_guids="+out.String()
							spaceuserslist := exec.Command("cf", "curl", strings.TrimSpace(path), "--output", "spaceusrslist.json")

							var out bytes.Buffer
							spaceuserslist.Stdout = &out
							err := spaceuserslist.Run()
							if err == nil {
								fmt.Println("no err",spaceuserslist, spaceuserslist.Stdout)
							} else {
								fmt.Println("err", spaceuserslist, spaceuserslist.Stdout, spaceuserslist.Stderr)
							}

							fileSpaceJson, err := ioutil.ReadFile("spaceuserslist.json")
							if err != nil {
								fmt.Println(err)
							}
							if err := json.Unmarshal(fileSpaceJson, &spaceusrslist); err != nil {
								panic(err)
							}

							SpaceUsrLen := len(spaceusrslist.Resources)
							fmt.Println("No of Users currently exist in Space",Orgs.Org.Spaces[j].Name,":", SpaceUsrLen)

							if SpaceUsrLen != 0 {

								for i := 0; i < SpaceUsrLen; i++ {

									//SSO Audit

									SpaceUsLenSSOAuditor := len(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceAuditors)

									var spaceusrssoauditscount, spaceusrssoaudittotalcount int

									if spaceusrslist.Resources[i].Type == "space_auditor" {

										userguid := spaceusrslist.Resources[i].GUID
										path := "/v3/users/?guids=" + userguid
										var out bytes.Buffer
										username := exec.Command("cf", "curl", strings.TrimSpace(path))
										username.Stdout = &out
										err := username.Run()

										if err == nil {
											for q := 0; q < SpaceUsLenSSOAuditor; q++ {

												if out.String() == "admin" {
													fmt.Println("Admin user skipping Validation")
													spaceusrssoaudittotalcount = 2
												} else {
													fmt.Println("SSO Space Audit Usr: ", Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceAuditors[q], ",", username)
													if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceAuditors[q]) == strings.TrimSpace(out.String()) {
														spaceusrssoauditscount = 1
													} else {
														spaceusrssoauditscount = 0
													}
													spaceusrssoaudittotalcount = spaceusrssoaudittotalcount + spaceusrssoauditscount
												}
											}

											if spaceusrssoaudittotalcount == 0 {

												fmt.Println("User has not be listed in Org.yml: ")
												fmt.Println("SSO SpaceAuditor User: ", strings.TrimSpace(out.String()))

												if Audit == "unset" {
													unset := exec.Command("cf", "unset-space-role", strings.TrimSpace(out.String()), Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceAuditor")
													if _, err := unset.Output(); err == nil {
														fmt.Println("command: ", unset)
														fmt.Println(unset.Stdout)
													} else {
														fmt.Println("command: ", unset)
														fmt.Println("Err: ", unset.Stdout, unset.Stderr)
													}
												} else if Audit == "List" {
													fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "SSO SpaceAuditor")
												} else {
													fmt.Println("Provide Valid Input")
												}

											} else if spaceusrssoaudittotalcount == 2 {
												fmt.Println("Admin set as SSO Space Audit User, skipping unset: ", strings.TrimSpace(out.String()))
											} else {
												fmt.Println("User is listed in Org.yml as SSO Space Audit User: ", strings.TrimSpace(out.String()))
											}
										} else {
											fmt.Println("err", username, username.Stdout, username.Stderr)
										}
									}

									SpaceUsLenSSODeveloper := len(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers)

									var spaceusrssodevscount, spaceusrssodevtotalcount int

									if spaceusrslist.Resources[i].Type == "space_developer" {

										userguid := spaceusrslist.Resources[i].GUID
										path := "/v3/users/?guids=" + userguid
										var out bytes.Buffer
										username := exec.Command("cf", "curl", strings.TrimSpace(path))
										username.Stdout = &out
										err := username.Run()

										if err == nil {
											for q := 0; q < SpaceUsLenSSODeveloper; q++ {

												if out.String() == "admin" {
													fmt.Println("Admin user skipping Validation")
													spaceusrssodevtotalcount = 2
												} else {
													fmt.Println("SSO Space Dev Usr: ", Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers[q], ",", username)
													if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers[q]) == strings.TrimSpace(out.String()) {
														spaceusrssodevscount = 1
													} else {
														spaceusrssodevscount = 0
													}
													spaceusrssodevtotalcount = spaceusrssodevtotalcount + spaceusrssodevscount
												}
											}

											//SSO Developer

											if spaceusrssodevtotalcount == 0 {
												fmt.Println("User has not be listed in Org.yml: ")
												fmt.Println("SSO Space Dev User: ", strings.TrimSpace(out.String()))

												if Audit == "unset" {
													unset := exec.Command("cf", "unset-space-role", strings.TrimSpace(out.String()), Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceDeveloper")
													if _, err := unset.Output(); err == nil {
														fmt.Println("command: ", unset)
														fmt.Println(unset.Stdout)
													} else {
														fmt.Println("command: ", unset)
														fmt.Println("Err: ", unset.Stdout, unset.Stderr)
													}
												} else if Audit == "List" {
													fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "SSO SpaceDeveloper")
												} else {
													fmt.Println("Provide Valid Input")
												}

											} else if spaceusrssodevtotalcount == 2 {
												fmt.Println("Admin set as SSO Space Dev User, skipping unset: ", strings.TrimSpace(out.String()))
											} else {
												fmt.Println("User is listed in Org.yml as SSO Space Dev User: ", strings.TrimSpace(out.String()))
											}
										} else {
											fmt.Println("err", username, username.Stdout, username.Stderr)
										}
									}

									//SSO Manager

									SpaceUsLenSSOMan := len(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers)

									var spaceusrssomanscount, spaceusrssomantotalcount int

									if spaceusrslist.Resources[i].Type == "space_manager" {

										userguid := spaceusrslist.Resources[i].GUID
										path := "/v3/users/?guids=" + userguid
										var out bytes.Buffer
										username := exec.Command("cf", "curl", strings.TrimSpace(path))
										username.Stdout = &out
										err := username.Run()

										if err == nil {
											for q := 0; q < SpaceUsLenSSOMan; q++ {

												if out.String() == "admin" {
													fmt.Println("Admin user skipping Validation")
													spaceusrssomantotalcount = 2
												} else {
													fmt.Println("SSO Space Manager Usr: ", Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers[q],",", username)
													if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers[q]) == strings.TrimSpace(out.String()) {
														spaceusrssomanscount = 1
													} else {
														spaceusrssomanscount = 0
													}
													spaceusrssomantotalcount = spaceusrssomantotalcount + spaceusrssomanscount
												}
											}

											if spaceusrssomantotalcount == 0 {
												fmt.Println("User has not be listed in Org.yml: ")
												fmt.Println("SSO Space Manager User: ", strings.TrimSpace(out.String()))

												if Audit == "unset" {
													unset := exec.Command("cf", "unset-space-role", strings.TrimSpace(out.String()), Orgs.Org.Name,Orgs.Org.Spaces[j].Name, "SpaceManager")
													if _, err := unset.Output(); err == nil {
														fmt.Println("command: ", unset)
														fmt.Println(unset.Stdout)
													} else {
														fmt.Println("command: ", unset)
														fmt.Println("Err: ", unset.Stdout, unset.Stderr)
													}
												} else if Audit == "List" {
													fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "SSO SpaceManager")
												} else {
													fmt.Println("Provide Valid Input")
												}
											} else if spaceusrssomantotalcount == 2 {
												fmt.Println("Admin set as SSO Space Manager User, skipping unset: ", strings.TrimSpace(out.String()))
											} else {
												fmt.Println("User is listed in Org.yml as SSO Space Manager User: ", strings.TrimSpace(out.String()))
											}
										} else {
											fmt.Println("err", username, username.Stdout, username.Stderr)
										}
									}

									//UAA Audit
									SpaceUsLenUAAAuditor := len(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors)

									var spaceusruaaauditscount, spaceusruaaaudittotalcount int

									if spaceusrslist.Resources[i].Type == "space_auditor" {

										userguid := spaceusrslist.Resources[i].GUID
										path := "/v3/users/?guids=" + userguid
										var out bytes.Buffer
										username := exec.Command("cf", "curl", strings.TrimSpace(path))
										username.Stdout = &out
										err := username.Run()

										if err == nil {
											for q := 0; q < SpaceUsLenUAAAuditor; q++ {

												if out.String() == "admin" {
													fmt.Println("Admin user skipping Validation")
													spaceusruaaaudittotalcount = 2
												} else {
													fmt.Println("UAA Space Audit Usr: ", Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors[q],",", username)
													if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors[q]) == strings.TrimSpace(out.String()) {
														spaceusruaaauditscount = 1
													} else {
														spaceusruaaauditscount = 0
													}
													spaceusruaaaudittotalcount = spaceusruaaaudittotalcount + spaceusruaaauditscount
												}
											}

											if spaceusruaaaudittotalcount == 0 {
												fmt.Println("User has not be listed in Org.yml: ")
												fmt.Println("UAA SpaceAuditor User: ", strings.TrimSpace(out.String()))

												if Audit == "unset" {
													unset := exec.Command("cf", "unset-space-role", strings.TrimSpace(out.String()), Orgs.Org.Name, Orgs.Org.Spaces[j].Name,"SpaceAuditor")
													if _, err := unset.Output(); err == nil {
														fmt.Println("command: ", unset)
														fmt.Println(unset.Stdout)
													} else {
														fmt.Println("command: ", unset)
														fmt.Println("Err: ", unset.Stdout, unset.Stderr)
													}
												} else if Audit == "List" {
													fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "UAA SpaceAuditor")
												} else {
													fmt.Println("Provide Valid Input")
												}
											} else if spaceusruaaaudittotalcount == 2 {
												fmt.Println("Admin set as UAA Space Audit User, skipping unset: ", strings.TrimSpace(out.String()))
											} else {
												fmt.Println("User is listed in Org.yml as UAA Space Audit User: ", strings.TrimSpace(out.String()))
											}
										} else {
											fmt.Println("err", username, username.Stdout, username.Stderr)
										}
									}

									//UAA Developer

									SpaceUsLenUAADeveloper := len(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers)

									var spaceusruaadevscount, spaceusruaadevtotalcount int

									if spaceusrslist.Resources[i].Type == "space_developer" {

										userguid := spaceusrslist.Resources[i].GUID
										path := "/v3/users/?guids=" + userguid
										var out bytes.Buffer
										username := exec.Command("cf", "curl", strings.TrimSpace(path))
										username.Stdout = &out
										err := username.Run()

										if err == nil {
											for q := 0; q < SpaceUsLenUAADeveloper; q++ {

												if out.String() == "admin" {
													fmt.Println("Admin user skipping Validation")
													spaceusruaadevtotalcount = 2
												} else {
													fmt.Println("UAA Space Dev Usr: ", Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers[q],",", username)
													if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers[q]) == strings.TrimSpace(out.String()) {
														spaceusruaadevscount = 1
													} else {
														spaceusruaadevscount = 0
													}
													spaceusruaadevtotalcount = spaceusruaadevtotalcount + spaceusruaadevscount
												}
											}

											if spaceusruaadevtotalcount == 0 {
												fmt.Println("User has not be listed in Org.yml: ")
												fmt.Println("UAA SpaceDeveloper User: ", strings.TrimSpace(out.String()))
												if Audit == "unset" {
													unset := exec.Command("cf", "unset-space-role", strings.TrimSpace(out.String()), Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceDeveloper")
													if _, err := unset.Output(); err == nil {
														fmt.Println("command: ", unset)
														fmt.Println(unset.Stdout)
													} else {
														fmt.Println("command: ", unset)
														fmt.Println("Err: ", unset.Stdout, unset.Stderr)
													}
												} else if Audit == "List" {
													fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "UAA SpaceDeveloper")
												} else {
													fmt.Println("Provide Valid Input")
												}
											} else if spaceusruaadevtotalcount == 2 {
												fmt.Println("Admin set as UAA Space Dev User, skipping unset: ", strings.TrimSpace(out.String()))
											} else {
												fmt.Println("User is listed in Org.yml as UAA Space Dev User: ", strings.TrimSpace(out.String()))
											}
										} else {
											fmt.Println("err", username, username.Stdout, username.Stderr)
										}
									}

									//UAA Manager

									SpaceUsLenUAAMan := len(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers)

									var spaceusruaamanscount, spaceusruaamantotalcount int

									if spaceusrslist.Resources[i].Type == "space_manager" {

										userguid := spaceusrslist.Resources[i].GUID
										path := "/v3/users/?guids="+userguid
										var out bytes.Buffer
										username := exec.Command("cf", "curl", strings.TrimSpace(path))
										username.Stdout = &out
										err := username.Run()

										if err == nil {
											for q := 0; q < SpaceUsLenUAAMan; q++ {

												if out.String() == "admin" {
													fmt.Println("Admin user skipping Validation")
													spaceusruaamantotalcount = 2
												} else {
													fmt.Println("UAA Space Dev Usr: ", Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers[q],",", username)
													if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers[q]) == strings.TrimSpace(out.String()) {
														spaceusruaamanscount = 1
													} else {
														spaceusruaamanscount = 0
													}
													spaceusruaamantotalcount = spaceusruaamantotalcount + spaceusruaamanscount
												}
											}

											if spaceusruaamantotalcount == 0 {
												fmt.Println("User has not be listed in Org.yml: ")
												fmt.Println("UAA SpaceManager User: ", strings.TrimSpace(out.String()))
												if Audit == "unset" {
													unset := exec.Command("cf", "unset-space-role", strings.TrimSpace(out.String()), Orgs.Org.Name, Orgs.Org.Spaces[j].Name,"SpaceManager")
													if _, err := unset.Output(); err == nil {
														fmt.Println("command: ", unset)
														fmt.Println(unset.Stdout)
													} else {
														fmt.Println("command: ", unset)
														fmt.Println("Err: ", unset.Stdout, unset.Stderr)
													}
												} else if Audit == "List" {
													fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "UAA SpaceManager")
												} else {
													fmt.Println("Provide Valid Input")
												}
											} else if spaceusruaamantotalcount == 2 {
												fmt.Println("Admin set as UAA Space Manager User, skipping unset: ", strings.TrimSpace(out.String()))
											} else {
												fmt.Println("User is listed in Org.yml as UAA Space Manager User: ", strings.TrimSpace(out.String()))
											}
										} else {
											fmt.Println("err", username, username.Stdout, username.Stderr)
										}
									}

									//LDAP Audit
									SpaceUsLenLDAPAuditor := len(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors)

									var spaceusrldapauditscount, spaceusrldapaudittotalcount int

									if spaceusrslist.Resources[i].Type == "space_auditor" {

										userguid := spaceusrslist.Resources[i].GUID
										path := "/v3/users/?guids=" + userguid
										var out bytes.Buffer
										username := exec.Command("cf", "curl", strings.TrimSpace(path))
										username.Stdout = &out
										err := username.Run()

										if err == nil {

											for q := 0; q < SpaceUsLenLDAPAuditor; q++ {

												if out.String() == "admin" {
													fmt.Println("Admin user skipping Validation")
													spaceusrldapaudittotalcount = 2
												} else {
													fmt.Println("LDAP Space Audit Usr: ", Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors[q],",", username)
													if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors[q]) == strings.TrimSpace(out.String()) {
														spaceusrldapauditscount = 1
													} else {
														spaceusrldapauditscount = 0
													}
													spaceusrldapaudittotalcount = spaceusrldapaudittotalcount + spaceusrldapauditscount
												}
											}

											if spaceusrldapaudittotalcount == 0 {
												fmt.Println("User has not be listed in Org.yml: ")
												fmt.Println("LDAP SpaceAuditor User: ", strings.TrimSpace(out.String()))
												if Audit == "unset" {
													unset := exec.Command("cf", "unset-space-role", strings.TrimSpace(out.String()), Orgs.Org.Name, Orgs.Org.Spaces[j].Name,  "SpaceAuditor")
													if _, err := unset.Output(); err == nil {
														fmt.Println("command: ", unset)
														fmt.Println(unset.Stdout)
													} else {
														fmt.Println("command: ", unset)
														fmt.Println("Err: ", unset.Stdout, unset.Stderr)
													}
												} else if Audit == "List" {
													fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "LDAP SpaceAuditor")
												} else {
													fmt.Println("Provide Valid Input")
												}
											} else if spaceusrldapaudittotalcount == 2 {
												fmt.Println("Admin set as LDAP Space Audit User, skipping unset: ", strings.TrimSpace(out.String()))
											} else {
												fmt.Println("User is listed in Org.yml as LDAP Space Audit User: ", strings.TrimSpace(out.String()))
											}
										} else {
											fmt.Println("err", username, username.Stdout, username.Stderr)
										}
									}

									//LDAP Developer

									SpaceUsLenLDAPDeveloper := len(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers)

									var spaceusrldapdevscount, spaceusrldapdevtotalcount int

									if spaceusrslist.Resources[i].Type == "space_developer" {

										userguid := spaceusrslist.Resources[i].GUID
										path := "/v3/users/?guids=" + userguid
										var out bytes.Buffer
										username := exec.Command("cf", "curl", strings.TrimSpace(path))
										username.Stdout = &out
										err := username.Run()

										if err == nil {
											for q := 0; q < SpaceUsLenLDAPDeveloper; q++ {

												if out.String() == "admin" {
													fmt.Println("Admin user skipping Validation")
													spaceusrldapdevtotalcount = 2
												} else {
													fmt.Println("LDAP Space Dev Usr: ", Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers[q], ",", username)
													if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers[q]) == strings.TrimSpace(out.String()) {
														spaceusrldapdevscount = 1
													} else {
														spaceusrldapdevscount = 0
													}
													spaceusrldapdevtotalcount = spaceusrldapdevtotalcount + spaceusrldapdevscount
												}
											}

											if spaceusrldapdevtotalcount == 0 {
												if Audit == "unset" {
													unset := exec.Command("cf", "unset-space-role", strings.TrimSpace(out.String()), Orgs.Org.Name, Orgs.Org.Spaces[j].Name,  "SpaceDeveloper")
													if _, err := unset.Output(); err == nil {
														fmt.Println("command: ", unset)
														fmt.Println(unset.Stdout)
													} else {
														fmt.Println("command: ", unset)
														fmt.Println("Err: ", unset.Stdout, unset.Stderr)
													}
												} else if Audit == "List" {
													fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "LDAP SpaceDeveloper")
												} else {
													fmt.Println("Provide Valid Input")
												}
											} else if spaceusrldapdevtotalcount == 2 {
												fmt.Println("Admin set as LDAP Space Dev User, skipping unset: ", strings.TrimSpace(out.String()))
											} else {
												fmt.Println("User is listed in Org.yml as LDAP Space Dev User: ", strings.TrimSpace(out.String()))
											}
										} else {
											fmt.Println("err", username, username.Stdout, username.Stderr)
										}
									}

									//LDAP Manager

									SpaceUsLenLDAPMan := len(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers)

									var spaceusrldapmanscount, spaceusrldapmantotalcount int

									if spaceusrslist.Resources[i].Type == "space_manager" {

										userguid := spaceusrslist.Resources[i].GUID
										path := "/v3/users/?guids=" + userguid
										var out bytes.Buffer
										username := exec.Command("cf", "curl", strings.TrimSpace(path))
										username.Stdout = &out
										err := username.Run()

										if err == nil {
											for q := 0; q < SpaceUsLenLDAPMan; q++ {

												if out.String() == "admin" {
													fmt.Println("Admin user skipping Validation")
													spaceusrldapmantotalcount = 2
												} else {
													fmt.Println("LDAP Space Manager Usr: ", Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers[q], ",", username)
													if strings.TrimSpace(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers[q]) == strings.TrimSpace(out.String()) {
														spaceusrldapmanscount = 1
													} else {
														spaceusrldapmanscount = 0
													}
													spaceusrldapmantotalcount = spaceusrldapmantotalcount + spaceusrldapmanscount
												}
											}

											if spaceusrldapmantotalcount == 0 {
												if Audit == "unset" {
													unset := exec.Command("cf", "unset-space-role", strings.TrimSpace(out.String()), Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceManager")
													if _, err := unset.Output(); err == nil {
														fmt.Println("command: ", unset)
														fmt.Println(unset.Stdout)
													} else {
														fmt.Println("command: ", unset)
														fmt.Println("Err: ", unset.Stdout, unset.Stderr)
													}
												} else if Audit == "List" {
													fmt.Println("User to be deleted: ", strings.TrimSpace(out.String()), "LDAP SpaceManager")
												} else {
													fmt.Println("Provide Valid Input")
												}
											} else if spaceusrldapmantotalcount == 2 {
												fmt.Println("Admin set as LDAP Space Manager User, skipping unset: ", strings.TrimSpace(out.String()))
											} else {
												fmt.Println("User is listed in Org.yml as LDAP Space Manager User: ", strings.TrimSpace(out.String()))
											}
										} else {
											fmt.Println("err", username, username.Stdout, username.Stderr)
										}
									}
								}
							} else {
								fmt.Println("No space users exist")
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
func DeleteOrAuditSpacesASGs(clustername string, cpath string, ostype string) error {

	var Orgs Orglist
	var ProtectedOrgs ProtectedList
	var list List

	ListYml := cpath+"/"+clustername+"/OrgsList.yml"
	fileOrgYml, err := ioutil.ReadFile(ListYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileOrgYml), &list)
	if err != nil {
		panic(err)
	}

	var InitClusterConfigVals InitClusterConfigVals
	ConfigFile := cpath+"/"+clustername+"/config.yml"

	fileConfigYml, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileConfigYml), &InitClusterConfigVals)
	if err != nil {
		panic(err)
	}
	var ASGPath, OrgsYml string

	ProtectedOrgsYml := cpath+"/"+clustername+"/ProtectedResources.yml"
	fileProtectedYml, err := ioutil.ReadFile(ProtectedOrgsYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileProtectedYml), &ProtectedOrgs)
	if err != nil {
		panic(err)
	}

	LenList := len(list.OrgList)
	LenProtectedOrgs := len(ProtectedOrgs.Org)


	for i := 0; i < LenList; i++ {

		var count, totalcount int
		fmt.Println("Org: ", list.OrgList[i])
		for p := 0; p < LenProtectedOrgs; p++ {
			fmt.Println("Protected Org: ", ProtectedOrgs.Org[p])
			if ProtectedOrgs.Org[p] == list.OrgList[i] {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {

			fmt.Println("This is not Protected Org")

			if ostype == "windows" {
				ASGPath = cpath+"\\"+clustername+"\\"+list.OrgList[i]+"\\ASGs\\"
				OrgsYml = cpath+"\\"+clustername+"\\"+list.OrgList[i]+"\\Org.yml"
			} else {
				ASGPath = cpath+"/"+clustername+"/"+list.OrgList[i]+"/ASGs/"
				OrgsYml = cpath+"/"+clustername+"/"+list.OrgList[i]+"/Org.yml"
			}


			fileOrgYml, err := ioutil.ReadFile(OrgsYml)

			if err != nil {
				fmt.Println(err)
			}

			err = yaml.Unmarshal([]byte(fileOrgYml), &Orgs)
			if err != nil {
				panic(err)
			}

			Audit := Orgs.ASGAudit

			if list.OrgList[i] == Orgs.Org.Name {
				guid := exec.Command("cf", "org", Orgs.Org.Name, "--guid")

				if _, err := guid.Output(); err == nil {

					fmt.Println("command: ", guid)
					fmt.Println("Org exists: ", guid.Stdout)
					SpaceLen := len(Orgs.Org.Spaces)

					TargetOrg := exec.Command("cf", "t", "-o", Orgs.Org.Name)
					if _, err := TargetOrg.Output(); err == nil {
						fmt.Println("command: ", TargetOrg)
						fmt.Println("Targeting: ", TargetOrg.Stdout)
					} else {
						fmt.Println("command: ", TargetOrg)
						fmt.Println("Err: ", TargetOrg.Stdout, err)
					}

					for j := 0; j < SpaceLen; j++ {

						fmt.Println("Auditing Spaces ASGs")
						guid = exec.Command("cf", "space", Orgs.Org.Spaces[j].Name, "--guid")

						if _, err := guid.Output(); err == nil{

							fmt.Println("command: ", guid)
							fmt.Println("Space exists: ", guid.Stdout)


							fmt.Println("Deleting or Auditing ASGs")
							if InitClusterConfigVals.ClusterDetails.EnableASG == true {
								fmt.Println("Enable ASGs: ", InitClusterConfigVals.ClusterDetails.EnableASG)
								DeleteOrAuditASGs(Orgs.Org.Name, Orgs.Org.Spaces[j].Name, ASGPath, ostype, Audit)
							} else {
								fmt.Println("Enable ASGs: ", InitClusterConfigVals.ClusterDetails.EnableASG)
								fmt.Println("ASGs not enabled")
							}
						} else {
							fmt.Println("command: ", guid)
							fmt.Println("Pulling Space Guid ID: ", guid.Stdout )
							fmt.Println("Space doesn't exist, please create space")
						}
					}
				} else {
					fmt.Println("command: ", guid )
					fmt.Println("Err: ", guid.Stdout, err)
					fmt.Println("Org doesn't exists, Please create Org")
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
func DeleteOrAuditASGs(Org string, Space string, asgpath string, ostype string, audit string) {

	ASGPath := asgpath
	ASGName := Org+"_"+Space+".json"
	path := ASGPath+ASGName

	var asglist ASGListJson

	var check *exec.Cmd

	if ostype == "windows" {
		check = exec.Command("powershell", "-command","Get-Content", path)
	} else {
		check = exec.Command("cat", path)
	}

	if _, err := check.Output(); err != nil {
		fmt.Println("command: ", check)
		fmt.Println("Err: ", check.Stdout, err)
		fmt.Println("No ASG defined for Org and Space combination")

		path := "/v3/security_groups?="+ASGName
		checkasg := exec.Command("cf", "curl", path, "--output", "asg.json")
		if _, err := checkasg.Output(); err != nil {
			fmt.Println("command: ", checkasg)
			fmt.Println("Err: ", checkasg.Stdout, err)
		} else {
			fileAsgJson, err := ioutil.ReadFile("asg.json")
			if err != nil {
				fmt.Println(err)
			}
			if err := json.Unmarshal(fileAsgJson, &asglist); err != nil {
				panic(err)
			}

			if len(asglist.Resources) == 0 {
				fmt.Println("No running ASG exist with name",ASGName,"for deleting")
			} else {

				if audit == "Delete" {
					fmt.Println("Unbinding running ASG: ", ASGName)
					unbind := exec.Command("cf", "unbind-running-security-group", ASGName, Org, Space, "--lifecycle", "running")
					if _, err := unbind.Output(); err != nil {
						fmt.Println("command: ", unbind)
						fmt.Println("Err: ", unbind.Stdout, err)
					} else {
						fmt.Println("Deleting running ASG: ", ASGName)
						delete := exec.Command("cf", "delete-security-group", ASGName)
						if _, err := delete.Output(); err != nil {
							fmt.Println("command: ", delete)
							fmt.Println("Err: ", delete.Stdout, err)
						} else {
							fmt.Println("command: ", delete)
							fmt.Println(delete.Stdout)
						}
					}
				} else if audit == "List" {
					fmt.Println("ASG to be deleted, Org, Space: ",ASGName, Org, Space)
				} else {
					fmt.Println("Provide Valid Input")
				}
			}
		}
	} else {
		fmt.Println("command: ", check)
		fmt.Println(check.Stdout)
		fmt.Println("Running ASG defined for Org, Space combination", ASGName)
	}
	return
}
func CreateOrUpdateOrgs(clustername string, cpath string) error {

	var Orgs Orglist
	var list List
	var ProtectedOrgs ProtectedList

	ListYml := cpath+"/"+clustername+"/OrgsList.yml"
	fileOrgYml, err := ioutil.ReadFile(ListYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileOrgYml), &list)
	if err != nil {
		panic(err)
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

	LenList := len(list.OrgList)
	LenProtectedOrgs := len(ProtectedOrgs.Org)

	for i := 0; i < LenList; i++ {
		var count, totalcount int
		fmt.Println("Org: ", list.OrgList[i])
		for p := 0; p < LenProtectedOrgs; p++ {
			fmt.Println("Protected Org: ", ProtectedOrgs.Org[p])
			if ProtectedOrgs.Org[p] == list.OrgList[i] {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count

		}

		if totalcount == 0 {
			fmt.Println("This is not Protected Org")

			OrgsYml := cpath+"/"+clustername+"/"+list.OrgList[i]+"/Org.yml"
			fileOrgYml, err := ioutil.ReadFile(OrgsYml)

			if err != nil {
				fmt.Println(err)
			}

			err = yaml.Unmarshal([]byte(fileOrgYml), &Orgs)
			if err != nil {
				panic(err)
			}
			if list.OrgList[i] == Orgs.Org.Name {
				guid := exec.Command("cf", "org", Orgs.Org.Name, "--guid")
				if _, err := guid.Output(); err == nil{

					fmt.Println("command: ", guid)
					fmt.Println("Org exists: ", guid.Stdout)
					fmt.Println("Updating Org quota")
					SetQuota := exec.Command("cf", "set-quota", Orgs.Org.Name, Orgs.Org.Quota)
					if _, err := SetQuota.Output(); err != nil{
						fmt.Println("command: ", SetQuota)
						fmt.Println("Err: ", SetQuota.Stdout)
						fmt.Println("Err Code: ", err)
					} else {
						fmt.Println("command: ", SetQuota)
						fmt.Println(SetQuota.Stdout)
					}
				} else {
					fmt.Println("command: ", guid)
					fmt.Println("Err: ", guid.Stdout)
					fmt.Println("Err Code: ", err)
					fmt.Println("Pulling Guid Id: ", guid.Stdout)
					fmt.Println("Org doesn't exists, Creating Org")
					createorg := exec.Command("cf", "create-org", Orgs.Org.Name)
					if _, err := createorg.Output(); err != nil{
						fmt.Println("command: ", createorg)
						fmt.Println("Err: ", createorg.Stdout)
						fmt.Println("Err Code: ", err)
					} else {
						fmt.Println("command: ", createorg)
						fmt.Println(createorg.Stdout)
					}
					attachquota := exec.Command("cf", "set-quota", Orgs.Org.Name, Orgs.Org.Quota)
					if _, err := attachquota.Output(); err != nil{
						fmt.Println("command: ", attachquota)
						fmt.Println("Err: ", attachquota.Stdout)
						fmt.Println("Err Code: ", err)
					} else {
						fmt.Println("command: ", attachquota)
						fmt.Println(attachquota.Stdout)
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
func CreateOrUpdateSpaces(clustername string, cpath string, ostype string) error {

	var Orgs Orglist
	var ProtectedOrgs ProtectedList
	var list List

	ListYml := cpath+"/"+clustername+"/OrgsList.yml"
	fileOrgYml, err := ioutil.ReadFile(ListYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileOrgYml), &list)
	if err != nil {
		panic(err)
	}

	var InitClusterConfigVals InitClusterConfigVals
	ConfigFile := cpath+"/"+clustername+"/config.yml"

	fileConfigYml, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileConfigYml), &InitClusterConfigVals)
	if err != nil {
		panic(err)
	}
	var OrgsYml string
	//var ASGPath, OrgsYml string

	ProtectedOrgsYml := cpath+"/"+clustername+"/ProtectedResources.yml"
	fileProtectedYml, err := ioutil.ReadFile(ProtectedOrgsYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileProtectedYml), &ProtectedOrgs)
	if err != nil {
		panic(err)
	}

	LenList := len(list.OrgList)
	LenProtectedOrgs := len(ProtectedOrgs.Org)


	for i := 0; i < LenList; i++ {

		var count, totalcount int
		fmt.Println("Org: ", list.OrgList[i])
		for p := 0; p < LenProtectedOrgs; p++ {
			fmt.Println("Protected Org: ", ProtectedOrgs.Org[p])
			if ProtectedOrgs.Org[p] == list.OrgList[i] {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {
			fmt.Println("This is not Protected Org")

			if ostype == "windows" {
				OrgsYml = cpath+"\\"+clustername+"\\"+list.OrgList[i]+"\\Org.yml"
			} else {
				OrgsYml = cpath+"/"+clustername+"/"+list.OrgList[i]+"/Org.yml"
			}

			fileOrgYml, err := ioutil.ReadFile(OrgsYml)
			if err != nil {
				fmt.Println(err)
			}

			err = yaml.Unmarshal([]byte(fileOrgYml), &Orgs)
			if err != nil {
				panic(err)
			}

			if list.OrgList[i] == Orgs.Org.Name {

				guid := exec.Command("cf", "org", Orgs.Org.Name, "--guid")

				if _, err := guid.Output(); err == nil {

					fmt.Println("command: ", guid)
					fmt.Println("Org exists: ", guid.Stdout)

					SpaceLen := len(Orgs.Org.Spaces)
					TargetOrg := exec.Command("cf", "t", "-o", Orgs.Org.Name)

					if _, err := TargetOrg.Output(); err == nil {
						fmt.Println("command: ", TargetOrg)
						fmt.Println("Targeting: ", TargetOrg.Stdout)
					} else {
						fmt.Println("command: ", TargetOrg)
						fmt.Println("Err: ", TargetOrg.Stdout)
						fmt.Println("Err Code: ", err)
					}

					for j := 0; j < SpaceLen; j++ {

						fmt.Println("Creating Spaces")

						guid = exec.Command("cf", "space", Orgs.Org.Spaces[j].Name, "--guid")

						if _, err := guid.Output(); err == nil{

							fmt.Println("command: ", guid)
							fmt.Println("Space exists: ", guid.Stdout)

							fmt.Println("Enabling Space Isolation Segment")
							fmt.Println("Org, Space: ", Orgs.Org.Name,",",Orgs.Org.Spaces[j].Name)
							fmt.Println("SegName: ", Orgs.Org.Spaces[j].IsolationSeg)
							if Orgs.Org.Spaces[j].IsolationSeg != "" {
								iso := exec.Command("cf", "enable-org-isolation", Orgs.Org.Name, Orgs.Org.Spaces[j].IsolationSeg)
								if _, err := iso.Output(); err != nil {
									fmt.Println("command: ", iso)
									fmt.Println("Err: ", iso.Stdout)
									fmt.Println("Err Code: ", err)
								} else {
									fmt.Println("command: ", iso)
									fmt.Println(iso.Stdout)
								}
								isospace := exec.Command("cf", "set-space-isolation-segment", Orgs.Org.Spaces[j].Name, Orgs.Org.Spaces[j].IsolationSeg)

								if _, err := isospace.Output(); err != nil {
									fmt.Println("command: ", isospace)
									fmt.Println("Err: ", isospace.Stdout)
									fmt.Println("Err Code: ", err)
								} else {
									fmt.Println("command: ", isospace)
									fmt.Println(isospace.Stdout)
								}
							} else {
								fmt.Println("No Isolation Segment Provided, Will be attached to Default")
							}


							//fmt.Println("Creating or updating ASGs")
							//if InitClusterConfigVals.ClusterDetails.EnableASG == true {
							//	fmt.Println("Enable ASGs: ", InitClusterConfigVals.ClusterDetails.EnableASG)
							//	CreateOrUpdateASGs(Orgs.Org.Name, Orgs.Org.Spaces[j].Name, ASGPath, ostype)
							//} else {
							//	fmt.Println("Enable ASGs: ", InitClusterConfigVals.ClusterDetails.EnableASG)
							//	fmt.Println("ASGs not enabled")
							//}
						} else {
							fmt.Println("command: ", guid)
							fmt.Println("Pulling Space Guid ID: ", guid.Stdout )
							fmt.Println("Creating Space")

							CreateSpace := exec.Command("cf", "create-space", Orgs.Org.Spaces[j].Name, "-o", Orgs.Org.Name)

							if _, err := CreateSpace.Output(); err != nil {
								fmt.Println("command: ", CreateSpace)
								fmt.Println("Err: ", CreateSpace.Stdout)
								fmt.Println("Err Code: ", err)
							} else {
								fmt.Println("command: ", CreateSpace)
								fmt.Println(CreateSpace.Stdout)

								fmt.Println("Enabling Space Isolation Segment")
								fmt.Println("SegName: ", Orgs.Org.Spaces[j].IsolationSeg)
								if Orgs.Org.Spaces[j].IsolationSeg != "" {
									iso := exec.Command("cf", "enable-org-isolation", Orgs.Org.Name, Orgs.Org.Spaces[j].IsolationSeg)

									if _, err := iso.Output(); err != nil {
										fmt.Println("command: ", iso)
										fmt.Println("Err: ", iso.Stdout)
										fmt.Println("Err Code: ", err)
									} else {
										fmt.Println("command: ", iso)
										fmt.Println(iso.Stdout)
									}

									isospace := exec.Command("cf", "set-space-isolation-segment", Orgs.Org.Spaces[j].Name, Orgs.Org.Spaces[j].IsolationSeg)
									if _, err := isospace.Output(); err != nil {
										fmt.Println("command: ", isospace)
										fmt.Println("Err: ", isospace.Stdout)
										fmt.Println("Err Code: ", err)
									} else {
										fmt.Println("command: ", isospace)
										fmt.Println(isospace.Stdout)
									}

								} else {
									fmt.Println("No Isolation Segment Provided, Will be attached to Default")
								}

				//				fmt.Println("Creating ASGs")
				//				if InitClusterConfigVals.ClusterDetails.EnableASG == true {
				//					fmt.Println("Enable ASGs: ", InitClusterConfigVals.ClusterDetails.EnableASG)
				//					CreateOrUpdateASGs(Orgs.Org.Name, Orgs.Org.Spaces[j].Name, ASGPath, ostype)
				//				} else {
				//					fmt.Println("Enable ASGs: ", InitClusterConfigVals.ClusterDetails.EnableASG)
				//					fmt.Println("ASGs not enabled")
				//				}
							}
						}
					}
				} else {
					fmt.Println("command: ", guid )
					fmt.Println("Err: ", guid.Stdout)
					fmt.Println("Err Code: ", err)
					fmt.Println("Org doesn't exists, Please create Org")
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
func CreateOrUpdateQuotas(clustername string, cpath string) error {

	var Quotas Quotalist
	var ProtectedQuota ProtectedList
	var cmd *exec.Cmd


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
			fmt.Println("Protected Quota: ", ProtectedQuota.Quota[p])
			if strings.Trim(ProtectedQuota.Quota[p], "") == strings.Trim(Quotas.Quota[i].Name, "") {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {

			fmt.Println("This is not Protected Quota")
			Quotadetails := exec.Command("cf", "quota", Quotas.Quota[i].Name)

			if _, err := Quotadetails.Output(); err != nil{
				fmt.Println("command: ", Quotadetails)
				fmt.Println("Err: ", Quotadetails.Stdout)
				fmt.Println("Err Code: ", err)
				//fmt.Println("Quota Doesn't exits: ", Quotadetails.Stdout)
				fmt.Println("Creating Quota")

				if Quotas.Quota[i].AllowPaidPlans == true {
					cmd = exec.Command("cf", "create-quota", Quotas.Quota[i].Name, "-m", MemLimit, "-i", "-1", "-r", "-1", "-s", SerLimit, "-a", AppLimt, "--allow-paid-service-plans")
				} else {
					cmd = exec.Command("cf", "create-quota", Quotas.Quota[i].Name, "-m", MemLimit, "-i", "-1", "-r", "-1", "-s", SerLimit, "-a", AppLimt, "--disallow-paid-service-plans")
				}

				if _, err := cmd.Output(); err != nil{
					fmt.Println("command: ", cmd)
					fmt.Println("Err: ", cmd.Stdout)
					fmt.Println("Err Code: ", err)
				} else {
					fmt.Println("command: ", cmd)
					fmt.Println(cmd.Stdout)
				}
				QuotaGet := exec.Command("cf", "quota", Quotas.Quota[i].Name)
				if _, err := QuotaGet.Output(); err != nil{
					fmt.Println("command: ", QuotaGet)
					fmt.Println("Err: ", QuotaGet.Stdout)
					fmt.Println("Err Code: ", err)
				} else {
					fmt.Println("command: ", QuotaGet)
					fmt.Println(QuotaGet.Stdout)
				}
			} else {
				fmt.Println("command: ", Quotadetails)
				fmt.Println("Quota exists: ", Quotadetails.Stdout)
				fmt.Println("Updating Quota")

				if Quotas.Quota[i].AllowPaidPlans == true {
					cmd = exec.Command("cf", "update-quota", Quotas.Quota[i].Name, "-m", MemLimit, "-i", "-1", "-r", "-1", "-s",  SerLimit, "-a", AppLimt, "--allow-paid-service-plans")
				} else {
					cmd = exec.Command("cf", "update-quota", Quotas.Quota[i].Name, "-m", MemLimit, "-i", "-1", "-r", "-1", "-s",  SerLimit, "-a", AppLimt, "--disallow-paid-service-plans")
				}

				if _, err := cmd.Output(); err != nil{
					fmt.Println("command: ", cmd)
					fmt.Println("Err: ", cmd.Stdout)
					fmt.Println("Err Code: ", err)
				} else {
					fmt.Println("command: ", cmd)
					fmt.Println(cmd.Stdout)
				}
				QuotaGet := exec.Command("cf", "quota", Quotas.Quota[i].Name)
				if _, err := QuotaGet.Output(); err != nil{
					fmt.Println("command: ", QuotaGet)
					fmt.Println("Err: ", QuotaGet.Stdout)
					fmt.Println("Err Code: ", err)
				} else {
					fmt.Println("command: ", QuotaGet)
					fmt.Println(QuotaGet.Stdout)
				}
			}
		} else {
			fmt.Println("This is a protected Org")
		}
	}
	return err
}
func CreateOrUpdateOrgUsers(clustername string, cpath string) error {

	var list List
	var Orgs Orglist
	var ProtectedOrgs ProtectedList

	ListYml := cpath+"/"+clustername+"/OrgsList.yml"
	fileOrgYml, err := ioutil.ReadFile(ListYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileOrgYml), &list)
	if err != nil {
		panic(err)
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
	LenList := len(list.OrgList)
	for i := 0; i < LenList; i++ {

		var count, totalcount int
		fmt.Println("Org: ", list.OrgList[i])
		for p := 0; p < LenProtectedOrgs; p++ {
			fmt.Println("Protected Org: ", ProtectedOrgs.Org[p])
			if ProtectedOrgs.Org[p] == list.OrgList[i] {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {

			fmt.Println("This is not Protected Org")

			OrgsYml := cpath+"/"+clustername+"/"+list.OrgList[i]+"/Org.yml"
			fileOrgYml, err := ioutil.ReadFile(OrgsYml)

			if err != nil {
				fmt.Println(err)
			}

			err = yaml.Unmarshal([]byte(fileOrgYml), &Orgs)
			if err != nil {
				panic(err)
			}

			if list.OrgList[i] == Orgs.Org.Name {
				guid := exec.Command("cf", "org", Orgs.Org.Name, "--guid")
				if _, err := guid.Output(); err == nil{

					fmt.Println("command: ", guid)
					fmt.Println("Org exists: ", guid.Stdout)
					fmt.Println("Updating Org Users")
					fmt.Println("Updating LDAP Users")

					LDAPOrgManLen := len(Orgs.Org.OrgUsers.LDAP.OrgManagers)

					for j := 0; j < LDAPOrgManLen; j++ {

						cmd := exec.Command("cf", "set-org-role", Orgs.Org.OrgUsers.LDAP.OrgManagers[j], Orgs.Org.Name, "OrgManager")

						if _, err := cmd.Output(); err != nil{
							fmt.Println("command: ", cmd)
							fmt.Println("Err: ", cmd.Stdout)
							fmt.Println("Err Code: ", err)
						} else {
							fmt.Println("command: ", cmd)
							fmt.Println(cmd.Stdout)
						}
					}

					LDAPOrgAudLen := len(Orgs.Org.OrgUsers.LDAP.OrgAuditors)

					for j := 0; j < LDAPOrgAudLen; j++ {

						cmd := exec.Command("cf", "set-org-role", Orgs.Org.OrgUsers.LDAP.OrgAuditors[j], Orgs.Org.Name, "OrgAuditor")

						if _, err := cmd.Output(); err != nil{
							fmt.Println("command: ", cmd)
							fmt.Println("Err: ", cmd.Stdout)
							fmt.Println("Err Code: ", err)
						} else {
							fmt.Println("command: ", cmd)
							fmt.Println(cmd.Stdout)
						}
					}

					fmt.Println("Updating UAA Users")

					UAAOrgManLen := len(Orgs.Org.OrgUsers.UAA.OrgManagers)

					for j := 0; j < UAAOrgManLen; j++ {

						cmd := exec.Command("cf", "set-org-role", Orgs.Org.OrgUsers.UAA.OrgManagers[j], Orgs.Org.Name, "OrgManager")

						if _, err := cmd.Output(); err != nil{
							fmt.Println("command: ", cmd)
							fmt.Println("Err: ", cmd.Stdout)
							fmt.Println("Err Code: ", err)
						} else {
							fmt.Println("command: ", cmd)
							fmt.Println(cmd.Stdout)
						}
					}

					UAAOrgAudLen := len(Orgs.Org.OrgUsers.UAA.OrgAuditors)

					for j := 0; j < UAAOrgAudLen; j++ {

						cmd := exec.Command("cf", "set-org-role", Orgs.Org.OrgUsers.UAA.OrgAuditors[j], Orgs.Org.Name, "OrgAuditor")

						if _, err := cmd.Output(); err != nil{
							fmt.Println("command: ", cmd)
							fmt.Println("Err: ", cmd.Stdout)
							fmt.Println("Err Code: ", err)
						} else {
							fmt.Println("command: ", cmd)
							fmt.Println(cmd.Stdout)
						}
					}

					fmt.Println("Updating SSO Users")

					SSOOrgManLen := len(Orgs.Org.OrgUsers.SSO.OrgManagers)

					for j := 0; j < SSOOrgManLen; j++ {

						cmd := exec.Command("cf", "set-org-role", Orgs.Org.OrgUsers.SSO.OrgManagers[j], Orgs.Org.Name, "OrgManager")

						if _, err := cmd.Output(); err != nil{
							fmt.Println("command: ", cmd)
							fmt.Println("Err: ", cmd.Stdout)
							fmt.Println("Err Code: ", err)
						} else {
							fmt.Println("command: ", cmd)
							fmt.Println(cmd.Stdout)
						}
					}

					SSOOrgAudLen := len(Orgs.Org.OrgUsers.SSO.OrgAuditors)

					for j := 0; j < SSOOrgAudLen; j++ {

						cmd := exec.Command("cf", "set-org-role", Orgs.Org.OrgUsers.SSO.OrgAuditors[j], Orgs.Org.Name, "OrgAuditor")

						if _, err := cmd.Output(); err != nil{
							fmt.Println("command: ", cmd)
							fmt.Println("Err: ", cmd.Stdout)
							fmt.Println("Err Code: ", err)
						} else {
							fmt.Println("command: ", cmd)
							fmt.Println(cmd.Stdout)
						}
					}
				} else {
					fmt.Println("command: ", guid)
					fmt.Println("Err: ", guid.Stdout)
					fmt.Println("Err Code: ", err)
					fmt.Println("Pulling Org Guid Id: ", guid.Stdout)
					fmt.Println("Please Create Org")
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
func CreateOrUpdateSpaceUsers(clustername string, cpath string) error {

	var Orgs Orglist
	var ProtectedOrgs ProtectedList
	var list List

	ListYml := cpath+"/"+clustername+"/OrgsList.yml"
	fileOrgYml, err := ioutil.ReadFile(ListYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileOrgYml), &list)
	if err != nil {
		panic(err)
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
	LenList := len(list.OrgList)

	for i := 0; i < LenList; i++ {

		var count, totalcount int
		fmt.Println("Org: ", list.OrgList[i])
		for p := 0; p < LenProtectedOrgs; p++ {
			fmt.Println("Protected Org: ", ProtectedOrgs.Org[p])
			if ProtectedOrgs.Org[p] == list.OrgList[i] {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}
		if totalcount == 0 {


			OrgsYml := cpath+"/"+clustername+"/"+list.OrgList[i]+"/Org.yml"
			fileOrgYml, err := ioutil.ReadFile(OrgsYml)

			if err != nil {
				fmt.Println(err)
			}

			err = yaml.Unmarshal([]byte(fileOrgYml), &Orgs)
			if err != nil {
				panic(err)
			}
			if list.OrgList[i] == Orgs.Org.Name {
				guid := exec.Command("cf", "org", Orgs.Org.Name, "--guid")

				if _, err := guid.Output(); err == nil {

					fmt.Println("command: ", guid)
					fmt.Println("Org exists: ", guid.Stdout)
					targetOrg := exec.Command("cf", "t", "-o", Orgs.Org.Name)
					if _, err := targetOrg.Output(); err == nil {
						fmt.Println("command: ", targetOrg)
						fmt.Println("Targeted Org: ", targetOrg.Stdout)
					} else {
						fmt.Println("command: ", targetOrg)
						fmt.Println("Err: ", targetOrg.Stdout)
						fmt.Println("Err Code: ", targetOrg.Stderr)
					}
					SpaceLen := len(Orgs.Org.Spaces)

					for j := 0; j < SpaceLen; j++ {

						guid = exec.Command("cf", "space", Orgs.Org.Spaces[j].Name, "--guid")
						if _, err := guid.Output(); err == nil {
							fmt.Println("command: ", guid)
							fmt.Println("Space exists: ", guid.Stdout)
							fmt.Println("Creating Space users")

							fmt.Println("Updating LDAP Users")

							LDAPSpaceManLen := len(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers)

							for k := 0; k < LDAPSpaceManLen; k++ {
								cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceManagers[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceManager")
								if _, err := cmd.Output(); err != nil{
									fmt.Println("command: ", cmd)
									fmt.Println("Err: ", cmd.Stdout)
									fmt.Println("Err Code: ", err)
								} else {
									fmt.Println("command: ", cmd)
									fmt.Println(cmd.Stdout)
								}
							}

							LDAPSpaceDevLen := len(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers)

							for k := 0; k < LDAPSpaceDevLen; k++ {
								cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceDevelopers[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceDeveloper")
								if _, err := cmd.Output(); err != nil{
									fmt.Println("command: ", cmd)
									fmt.Println("Err: ", cmd.Stdout)
									fmt.Println("Err Code: ", err)
								} else {
									fmt.Println("command: ", cmd)
									fmt.Println(cmd.Stdout)
								}
							}

							LDAPSpaceAuditLen := len(Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors)

							for k := 0; k < LDAPSpaceAuditLen; k++ {
								cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.LDAP.SpaceAuditors[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceAuditor")
								if _, err := cmd.Output(); err != nil{
									fmt.Println("command: ", cmd)
									fmt.Println("Err: ", cmd.Stdout)
									fmt.Println("Err Code: ", err)
								} else {
									fmt.Println("command: ", cmd)
									fmt.Println(cmd.Stdout)
								}
							}


							fmt.Println("Updating UAA Users")

							UAASpaceManLen := len(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers)

							for k := 0; k < UAASpaceManLen; k++ {
								cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceManagers[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceManager")
								if _, err := cmd.Output(); err != nil{
									fmt.Println("command: ", cmd)
									fmt.Println("Err: ", cmd.Stdout)
									fmt.Println("Err Code: ", err)
								} else {
									fmt.Println("command: ", cmd)
									fmt.Println(cmd.Stdout)
								}
							}

							UAASpaceDevLen := len(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers)

							for k := 0; k < UAASpaceDevLen; k++ {
								cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceDevelopers[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceDeveloper")
								if _, err := cmd.Output(); err != nil{
									fmt.Println("command: ", cmd)
									fmt.Println("Err: ", cmd.Stdout)
									fmt.Println("Err Code: ", err)
								} else {
									fmt.Println("command: ", cmd)
									fmt.Println(cmd.Stdout)
								}
							}

							UAASpaceAuditLen := len(Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors)

							for k := 0; k < UAASpaceAuditLen; k++ {
								cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.UAA.SpaceAuditors[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceAuditor")
								if _, err := cmd.Output(); err != nil{
									fmt.Println("command: ", cmd)
									fmt.Println("Err: ", cmd.Stdout)
									fmt.Println("Err Code: ", err)
								} else {
									fmt.Println("command: ", cmd)
									fmt.Println(cmd.Stdout)
								}
							}

							fmt.Println("Updating SSO Users")

							SSOSpaceManLen := len(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers)

							for k := 0; k < SSOSpaceManLen; k++ {
								cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceManagers[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceManager")
								if _, err := cmd.Output(); err != nil{
									fmt.Println("command: ", cmd)
									fmt.Println("Err: ", cmd.Stdout)
									fmt.Println("Err Code: ", err)
								} else {
									fmt.Println("command: ", cmd)
									fmt.Println(cmd.Stdout)
								}
							}

							SSOSpaceDevLen := len(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers)

							for k := 0; k < SSOSpaceDevLen; k++ {
								cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceDevelopers[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceDeveloper")
								if _, err := cmd.Output(); err != nil{
									fmt.Println("command: ", cmd)
									fmt.Println("Err: ", cmd.Stdout)
									fmt.Println("Err Code: ", err)
								} else {
									fmt.Println("command: ", cmd)
									fmt.Println(cmd.Stdout)
								}
							}

							SSOSpaceAuditLen := len(Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceAuditors)

							for k := 0; k < SSOSpaceAuditLen; k++ {
								cmd := exec.Command("cf", "set-space-role", Orgs.Org.Spaces[j].SpaceUsers.SSO.SpaceAuditors[k], Orgs.Org.Name, Orgs.Org.Spaces[j].Name, "SpaceAuditor")
								if _, err := cmd.Output(); err != nil{
									fmt.Println("command: ", cmd)
									fmt.Println("Err: ", cmd.Stdout)
									fmt.Println("Err Code: ", err)
								} else {
									fmt.Println("command: ", cmd)
									fmt.Println(cmd.Stdout)
								}
							}

						} else {
							fmt.Println("command: ",guid)
							fmt.Println("Err: ", guid.Stdout)
							fmt.Println("Err Code: ", err)
							fmt.Println("Space doesn't exists, Please create Space")
						}
					}
				} else {
					fmt.Println("command: ", guid)
					fmt.Println("Err: ", guid.Stdout)
					fmt.Println("Err Code: ", err)
					fmt.Println("Org doesn't exists, Please create Org")
				}
			} else {
				fmt.Println("Org Name does't match with folder name")
			}
		}
	}
	return err
}
func CreateOrUpdateSpacesASGs(clustername string, cpath string, ostype string) error {

	var Orgs Orglist
	var ProtectedOrgs ProtectedList
	var list List

	ListYml := cpath+"/"+clustername+"/OrgsList.yml"
	fileOrgYml, err := ioutil.ReadFile(ListYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileOrgYml), &list)
	if err != nil {
		panic(err)
	}

	var InitClusterConfigVals InitClusterConfigVals
	ConfigFile := cpath+"/"+clustername+"/config.yml"

	fileConfigYml, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileConfigYml), &InitClusterConfigVals)
	if err != nil {
		panic(err)
	}
	var ASGPath, OrgsYml string

	ProtectedOrgsYml := cpath+"/"+clustername+"/ProtectedResources.yml"
	fileProtectedYml, err := ioutil.ReadFile(ProtectedOrgsYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileProtectedYml), &ProtectedOrgs)
	if err != nil {
		panic(err)
	}

	LenList := len(list.OrgList)
	LenProtectedOrgs := len(ProtectedOrgs.Org)


	for i := 0; i < LenList; i++ {

		var count, totalcount int
		fmt.Println("Org: ", list.OrgList[i])
		for p := 0; p < LenProtectedOrgs; p++ {
			fmt.Println("Protected Org: ", ProtectedOrgs.Org[p])
			if ProtectedOrgs.Org[p] == list.OrgList[i] {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {
			fmt.Println("This is not Protected Org")

			if ostype == "windows" {
				ASGPath = cpath+"\\"+clustername+"\\"+list.OrgList[i]+"\\ASGs\\"
				OrgsYml = cpath+"\\"+clustername+"\\"+list.OrgList[i]+"\\Org.yml"
			} else {
				ASGPath = cpath+"/"+clustername+"/"+list.OrgList[i]+"/ASGs/"
				OrgsYml = cpath+"/"+clustername+"/"+list.OrgList[i]+"/Org.yml"
			}


			fileOrgYml, err := ioutil.ReadFile(OrgsYml)

			if err != nil {
				fmt.Println(err)
			}

			err = yaml.Unmarshal([]byte(fileOrgYml), &Orgs)
			if err != nil {
				panic(err)
			}
			if list.OrgList[i] == Orgs.Org.Name {
				guid := exec.Command("cf", "org", Orgs.Org.Name, "--guid")

				if _, err := guid.Output(); err == nil {

					fmt.Println("command: ", guid)
					fmt.Println("Org exists: ", guid.Stdout)
					SpaceLen := len(Orgs.Org.Spaces)

					TargetOrg := exec.Command("cf", "t", "-o", Orgs.Org.Name)
					if _, err := TargetOrg.Output(); err == nil {
						fmt.Println("command: ", TargetOrg)
						fmt.Println("Targeting: ", TargetOrg.Stdout)
					} else {
						fmt.Println("command: ", TargetOrg)
						fmt.Println("Err: ", TargetOrg.Stdout)
						fmt.Println("Err Code: ", err)
					}

					for j := 0; j < SpaceLen; j++ {

						fmt.Println("Creating Spaces ASGs")
						guid = exec.Command("cf", "space", Orgs.Org.Spaces[j].Name, "--guid")

						if _, err := guid.Output(); err == nil{

							fmt.Println("command: ", guid)
							fmt.Println("Space exists: ", guid.Stdout)
							fmt.Println("Creating or updating ASGs")
							if InitClusterConfigVals.ClusterDetails.EnableASG == true {
								fmt.Println("Enable ASGs: ", InitClusterConfigVals.ClusterDetails.EnableASG)
								CreateOrUpdateASGs(Orgs.Org.Name, Orgs.Org.Spaces[j].Name, ASGPath, ostype)
							} else {
								fmt.Println("Enable ASGs: ", InitClusterConfigVals.ClusterDetails.EnableASG)
								fmt.Println("ASGs not enabled")
							}
						} else {
							fmt.Println("command: ", guid)
							fmt.Println("Pulling Space Guid ID: ", guid.Stdout )
							fmt.Println("Space doesn't exist, please create space")
						}
					}
				} else {
					fmt.Println("command: ", guid )
					fmt.Println("Err: ", guid.Stdout)
					fmt.Println("Err Code: ", err)
					fmt.Println("Org doesn't exists, Please create Org")
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
func CreateOrUpdateASGs(Org string, Space string, asgpath string, ostype string) {

	ASGPath := asgpath
	ASGName := Org+"_"+Space+".json"
	path := ASGPath+ASGName

	//check := exec.Command("powershell", "-command","Get-Content", path)

	var check *exec.Cmd

	if ostype == "windows" {
		check = exec.Command("powershell", "-command","Get-Content", path)
		//check = exec.Command("type", path)
	} else {
		check = exec.Command("cat", path)
	}

	//check := exec.Command("cat", path)

	if _, err := check.Output(); err != nil {
		fmt.Println("command: ", check)
		fmt.Println("Err: ", check.Stdout)
		fmt.Println("Err Code: ", err)
		fmt.Println("No ASG defined for Org and Space combination")
	} else
	{
		fmt.Println("command: ", check)
		fmt.Println(check.Stdout)
		fmt.Println("Binding ASGs")

		checkcreate := exec.Command("cf", "security-group", ASGName)
		if _, err := checkcreate.Output(); err != nil {
			fmt.Println("command: ", checkcreate)
			fmt.Println("Err: ", checkcreate.Stdout)
			fmt.Println("Err Code: ", err)
			fmt.Println("ASG doesn't exist, Creating ASG")

			createasg := exec.Command("cf", "create-security-group", ASGName, path)
			if _, err := createasg.Output(); err != nil {
				fmt.Println("command: ", createasg)
				fmt.Println("Err: ", createasg.Stdout)
				fmt.Println("Err Code: ", err)
				fmt.Println("ASG creation failed")
			} else {
				fmt.Println("command: ", createasg)
				fmt.Println(createasg.Stdout)
			}
		} else {
			fmt.Println("command: ", checkcreate)
			fmt.Println(checkcreate.Stdout)
			fmt.Println("ASG exist, Updating ASG")
			updateasg := exec.Command("cf", "update-security-group", ASGName, path)
			if _, err := updateasg.Output(); err != nil {
				fmt.Println("command: ", updateasg)
				fmt.Println("Err: ", updateasg.Stdout)
				fmt.Println("Err Code: ", err)
				fmt.Println("ASG update failed")
			} else {
				fmt.Println("command: ", updateasg)
				fmt.Println(updateasg.Stdout)
			}
		}
		fmt.Println("Creating or Updating ASG finished, binding ASG")
		bindasg := exec.Command("cf", "bind-security-group", ASGName, Org, Space, "--lifecycle", "running")
		if _, err := bindasg.Output(); err != nil {
			fmt.Println("command: ", bindasg)
			fmt.Println("Err: ", bindasg.Stdout)
			fmt.Println("Err Code: ", err)
			fmt.Println("ASG binding failed")
		} else {
			fmt.Println("command: ", bindasg)
			fmt.Println(bindasg.Stdout)
		}
	}
	return
}
func CreateOrUpdateProtOrgAsg(clustername string, cpath string, ostype string) {

	var ProtectedOrgs ProtectedList
	var ASGpath string
	ProtectedOrgsYml := cpath+"/"+clustername+"/ProtectedResources.yml"
	fileProtectedYml, err := ioutil.ReadFile(ProtectedOrgsYml)

	var InitClusterConfigVals InitClusterConfigVals
	ConfigFile := cpath+"/"+clustername+"/config.yml"

	fileConfigYml, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileConfigYml), &InitClusterConfigVals)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(fileProtectedYml), &ProtectedOrgs)
	if err != nil {
		panic(err)
	}

	if ostype == "windows" {
		ASGpath = cpath+"\\"+clustername+"\\ProtectedOrgsASGs\\"
	} else {
		ASGpath = cpath+"/"+clustername+"/ProtectedOrgsASGs/"
	}

	LenProtectedOrgs := len(ProtectedOrgs.Org)
	var check *exec.Cmd
	ASGfile := ASGpath+ProtectedOrgs.DefaultRunningSecurityGroup+".json"
	if InitClusterConfigVals.ClusterDetails.EnableASG == true {
		fmt.Println("Enable ASGs: ", InitClusterConfigVals.ClusterDetails.EnableASG)

		if ostype == "windows" {
			check = exec.Command("powershell", "-command","Get-Content", ASGfile)
		} else {
			check = exec.Command("cat", ASGfile)
		}

		if _, err := check.Output(); err != nil {
			fmt.Println("ASG for Protected Orgs: ", ProtectedOrgs.DefaultRunningSecurityGroup)
			fmt.Println("command: ", check)
			fmt.Println("Err: ", check.Stdout)
			fmt.Println("Err Code: ", err)
			fmt.Println("No Default ASG file provided in path for Protected Orgs")
		} else {
			fmt.Println("command: ", check)
			fmt.Println(check.Stdout)
			fmt.Println("ASG for Protected Orgs: ", ProtectedOrgs.DefaultRunningSecurityGroup)
			checkdasg := exec.Command("cf", "security-group", ProtectedOrgs.DefaultRunningSecurityGroup)
			if _, err := checkdasg.Output(); err != nil {
				fmt.Println("command: ", checkdasg)
				fmt.Println("Err: ", checkdasg.Stdout)
				fmt.Println("Err Code: ", err)
				fmt.Println("Default ASG doesn't exist, Creating default ASG")
				createdasg := exec.Command("cf", "create-security-group", ProtectedOrgs.DefaultRunningSecurityGroup, ASGfile)
				if _, err := createdasg.Output(); err != nil {
					fmt.Println("command: ", createdasg)
					fmt.Println("Err: ", createdasg.Stdout)
					fmt.Println("Err Code: ", err)
					fmt.Println("Creating default ASG failed")
				} else {
					fmt.Println("command: ", createdasg)
					fmt.Println(createdasg.Stdout)
				}
			} else {
				fmt.Println("Default ASG exist, Updating default ASG")
				updatedefasg := exec.Command("cf", "update-security-group", ProtectedOrgs.DefaultRunningSecurityGroup, ASGfile)
				if _, err := updatedefasg.Output(); err != nil {
					fmt.Println("command: ", updatedefasg)
					fmt.Println("Err: ", updatedefasg.Stdout)
					fmt.Println("Err Code: ", err)
					fmt.Println("Default ASG not updated")
				} else {
					fmt.Println("command: ", updatedefasg)
					fmt.Println(updatedefasg.Stdout)
				}
			}
		}

		for p := 0; p < LenProtectedOrgs; p++ {
			fmt.Println("Protected Org: ", ProtectedOrgs.Org[p])
			fmt.Println("ASG for Protected Orgs: ", ProtectedOrgs.DefaultRunningSecurityGroup)
			bindasg := exec.Command("cf", "bind-security-group", ProtectedOrgs.DefaultRunningSecurityGroup, ProtectedOrgs.Org[p], "--lifecycle", "running")
			if _, err := bindasg.Output(); err != nil{
				fmt.Println("command: ", bindasg)
				fmt.Println("Err: ", bindasg.Stdout)
				fmt.Println("Err Code: ", err)
				fmt.Println("Failed to bind to protected Org")
			} else {
				fmt.Println("command: ", bindasg)
				fmt.Println(bindasg.Stdout)
			}
		}
	} else {
		fmt.Println("Enable ASGs: ", InitClusterConfigVals.ClusterDetails.EnableASG)
		fmt.Println("ASGs not enabled")
	}
}
func OrgsInit(clustername string, cpath string) error {

	var list List
	var ProtectedOrgs ProtectedList

	ListYml := cpath + "/" + clustername + "/OrgsList.yml"
	fileOrgYml, err := ioutil.ReadFile(ListYml)
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileOrgYml), &list)
	if err != nil {
		panic(err)
	}

	ProtectedOrgsYml := cpath + "/" + clustername + "/ProtectedResources.yml"
	fileProtectedYml, err := ioutil.ReadFile(ProtectedOrgsYml)

	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal([]byte(fileProtectedYml), &ProtectedOrgs)
	if err != nil {
		panic(err)
	}

	LenList := len(list.OrgList)
	LenProtectedOrgs := len(ProtectedOrgs.Org)

	for i := 0; i < LenList; i++ {
		var count, totalcount int
		fmt.Println("Org: ", list.OrgList[i])
		for p := 0; p < LenProtectedOrgs; p++ {
			fmt.Println("Protected Org: ", ProtectedOrgs.Org[p])
			if ProtectedOrgs.Org[p] == list.OrgList[i] {
				count = 1
			} else {
				count = 0
			}
			totalcount = totalcount + count
		}

		if totalcount == 0 {
			fmt.Println("This is not Protected Org")

			mgmtpath := cpath + "/" + clustername + "/" + list.OrgList[i]
			ASGPath := cpath + "/" + clustername + "/" + list.OrgList[i] + "/ASGs/"
			OrgsYml := cpath + "/" + clustername + "/" + list.OrgList[i] +"/Org.yml"
			JsonPath := cpath + "/" + clustername + "/" + list.OrgList[i] + "/ASGs/" + "test_test.json"

			_, err = os.Stat(mgmtpath)
			if os.IsNotExist(err) {

				fmt.Println("Creating <cluster>/<Org> folder")
				errDir := os.MkdirAll(mgmtpath, 0755)
				if errDir != nil {
					log.Fatal(err)
				}

				var OrgTmp = `---
Org:
  Name: ""
  Quota: ""
  OrgUsers:
    LDAP:
      OrgManagers:
        - User1
        - User2
        - User3
      OrgAuditors:
        - User1
        - User2
    SSO:
      OrgManagers:
        - User1
        - User2
        - User3
      OrgAuditors:
        - User1
        - User2
    UAA:
      OrgManagers:
        - User1
        - User2
        - User3
      OrgAuditors:
        - User1
        - User2
  Spaces:
    - Name: Space1
      IsolationSeg: "test-segment-1"
      SpaceUsers:
        LDAP:
          SpaceManagers:
            - User1
            - User2
            - User3
          SpaceDevelopers:
            - User1
            - User2
            - User3
          SpaceAuditors:
            - User1
            - User2
            - User3
    - Name: Space2
      IsolationSeg: "test-segment-2"
      SpaceUsers:
        LDAP:
          SpaceManagers:
            - User1
            - User2
            - User3
          SpaceDevelopers:
            - User1
            - User2
            - User3
          SpaceAuditors:
            - User1
            - User2
            - User3
SpaceAudit: List #Delete/Rename/List
UserAudit:  List #Unset/List
ASGAudit:   List #Delete/List`

				fmt.Println("Creating <cluster>/<Org> sample yaml files")
				err = ioutil.WriteFile(OrgsYml, []byte(OrgTmp), 0644)
				check(err)
			} else {
				fmt.Println("<cluster>/<Org> exists, please manually edit file to make changes or provide new cluster name")
			}
			_, err = os.Stat(ASGPath)
			if os.IsNotExist(err) {
				errDir := os.MkdirAll(ASGPath, 0755)
				if errDir != nil {
					log.Fatal(err)
					fmt.Println("<cluster>/<Org>/ASGs exist, please manually edit file to make changes or provide new cluster name")
				} else {
					fmt.Println("Creating <cluster>/<Org>/ASGs")
					var AsgTmp = `---
[
  {
    "protocol": "tcp",
    "destination": "10.x.x.88",
    "ports": "1443",
	"log": true,
	"description": "Allow DNS lookup by default."
  }
]`

					fmt.Println("Creating <cluster>/<Org>/ASGs sample json file")
					err = ioutil.WriteFile(JsonPath, []byte(AsgTmp), 0644)
					check(err)
				}
			}
		}
	}
	return nil
}
func Init(clustername string, endpoint string, user string, org string, space string, asg string, cpath string) (err error) {

	type ClusterDetails struct {
		EndPoint         string `yaml:"EndPoint"`
		User         string `yaml:"User"`
		Org            string `yaml:"Org"`
		Space string  `yaml:"Space"`
		EnableASG     string `yaml:"EnableASG"`
	}


	mgmtpath := cpath+"/"+clustername
	ASGPath := cpath+"/"+clustername+"/ProtectedOrgsASGs/"
	QuotasYml := cpath+"/"+clustername+"/Quota.yml"
	ProtectedResourcesYml := cpath+"/"+clustername+"/ProtectedResources.yml"
	ListOrgsYml := cpath+"/"+clustername+"/OrgsList.yml"


	_, err = os.Stat(mgmtpath)
	if os.IsNotExist(err) {

		fmt.Println("Creating <cluster> folder")
		errDir := os.MkdirAll(mgmtpath, 0755)


		var data = `---
ClusterDetails:
  EndPoint: {{ .EndPoint }}
  User: {{ .User }}
  Org: {{ .Org }}
  Space: {{ .Space }}
  EnableASG: {{ .EnableASG }}`

		// Create the file:
		err = ioutil.WriteFile(mgmtpath+"/config.tmpl", []byte(data), 0644)
		check(err)

		values := ClusterDetails{EndPoint: endpoint, User: user, Org: org, Space: space, EnableASG: asg}

		var templates *template.Template
		var allFiles []string

		if err != nil {
			fmt.Println(err)
		}

		filename := "config.tmpl"
		fullPath := mgmtpath + "/config.tmpl"
		if strings.HasSuffix(filename, ".tmpl") {
			allFiles = append(allFiles, fullPath)
		}

		fmt.Println(allFiles)
		templates, err = template.ParseFiles(allFiles...)
		if err != nil {
			fmt.Println(err)
		}

		s1 := templates.Lookup("config.tmpl")
		f, err := os.Create(mgmtpath + "/config.yml")
		if err != nil {
			panic(err)
		}

		fmt.Println("Initializing folder and config files")

		err = s1.Execute(f, values)
		defer f.Close() // don't forget to close the file when finished.
		if err != nil {
			panic(err)
		}

		var QuotasTmp = `---
quota:
  - Name: default
    memory_limit: 1024M
    allow_paid_plans: False
    app_instance_limit: 25
    service_instance_limit: 25
  - Name: small_quota
    memory_limit: 2048M
  - Name: medium_quota
    memory_limit: 2048M
  - Name: large_quota
    memory_limit: 2048M
Audit:  List`

		var ProtectedListTmp = `---
Org:
  - system
  - healthwatch
  - dynatrace
quota:
  - default
DefaultRunningSecurityGroup: default_security_group`

		var ListTmp = `---
OrgList:
  - Org-1
  - Org-2
  - Org-3
Audit: List`

		fmt.Println("Creating <cluster>/ sample yaml files")
		err = ioutil.WriteFile(QuotasYml, []byte(QuotasTmp), 0644)
		check(err)
		err = ioutil.WriteFile(ProtectedResourcesYml, []byte(ProtectedListTmp), 0644)
		check(err)
		err = ioutil.WriteFile(ListOrgsYml, []byte(ListTmp), 0644)
		check(err)

		if errDir != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("<cluster> exists, please manually edit file to make changes or provide new cluster name")
	}

	_, err = os.Stat(ASGPath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(ASGPath, 0755)
		if errDir != nil {
			log.Fatal(err)
			fmt.Println("<cluster>/ASGs exist, please manually edit file to make changes or provide new cluster name")
		} else {
			fmt.Println("Creating <cluster>/ASGs")
		}
	}

	return
}
func check(e error) {
	if e != nil {
		fmt.Println("<cluster>/ yamls exists, please manually edit file to make changes or provide new cluster name")
		panic(e)
	}
}
