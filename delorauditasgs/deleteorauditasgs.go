package deleteorauditasgs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
)

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


func DeleteOrAuditASGs(Org string, Space string, asgpath string, ostype string, audit string, ASGName string, MasterASGAudit bool) {

	ASGPath := asgpath
	//ASGName := Org+"_"+Space+".json"
	path := ASGPath+ASGName
	ASG := Org+"_"+Space

	//var asglist ASGListJson

	var check *exec.Cmd

	if ostype == "windows" {
		check = exec.Command("powershell", "-command","Get-Content", path)
	} else {
		check = exec.Command("cat", path)
	}

	if _, err := check.Output(); err != nil {
		fmt.Println("command: ", check)
		fmt.Println("Err: ", check.Stdout, err)
		fmt.Println("No running ASG defined for Org/Space combination", Org, Space)
		fmt.Println("Checking if ASG has been already binded to Org/Space combination")

		var checkasg *exec.Cmd
		if ostype == "windows" {
			path := "\""+"/v3/security_groups?names="+ASG+"\""
			checkasg = exec.Command("powershell", "-command","cf", "curl", path, "--output", "DeleteOrAuditASGs_asg.json")
		} else {
			path := "/v3/security_groups?names="+ASG
			checkasg = exec.Command("cf", "curl", path, "--output", "DeleteOrAuditASGs_asg.json")
		}

		if _, err := checkasg.Output(); err != nil {
			fmt.Println("command: ", checkasg)
			fmt.Println("Err: ", checkasg.Stdout, err)
		} else {

			fmt.Println("command: ", checkasg)
			fileAsgJson, err := ioutil.ReadFile("DeleteOrAuditASGs_asg.json")
			if err != nil {
				fmt.Println(err)
			}

			var asglist ASGListJson
			if err := json.Unmarshal(fileAsgJson, &asglist); err != nil {
				panic(err)
			}

			if len(asglist.Resources) == 0 {
				fmt.Println("Running ASG",ASG," is not binded for deleting")
			} else {
				if audit == "delete" {
					fmt.Println("DELETE!DELETE!")
					fmt.Println("Unbinding running ASG: ", ASG)
					if MasterASGAudit == true {
						unbind := exec.Command("cf", "unbind-security-group", ASG, Org, Space, "--lifecycle", "running")
						if _, err := unbind.Output(); err != nil {
							fmt.Println("command: ", unbind)
							fmt.Println("Err: ", unbind.Stdout, err)
						} else {
							fmt.Println("Deleting running ASG: ", ASG)
							delete := exec.Command("cf", "delete-security-group", ASG, "-f")
							if _, err := delete.Output(); err != nil {
								fmt.Println("command: ", delete)
								fmt.Println("Err: ", delete.Stdout, err)
							} else {
								fmt.Println("command: ", delete)
								fmt.Println(delete.Stdout)
							}
						}
					} else {
						fmt.Println("AuditASGs is not enabled")
					}

				} else if audit == "list" {
					fmt.Println("DELETE!DELETE!")
					fmt.Println("ASG to be deleted, Org, Space: ",ASG, Org, Space)
				} else {
					fmt.Println("Provide Valid Input")
				}

				results := exec.Command("cf", "security-groups")
				if _, err := results.Output(); err != nil{
					fmt.Println("command: ", results)
					fmt.Println("Err: ", results.Stdout, err)
				} else {
					//fmt.Println("command: ", results)
					fmt.Println(results.Stdout)
				}
			}
		}
	} else {
		//fmt.Println("command: ", check)
		//fmt.Println(check.Stdout)
		fmt.Println("Running ASG defined for Org, Space combination", ASG)
	}
	return
}

