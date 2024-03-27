package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type MyJsonName struct {
	Offset  int64 `json:"offset"`
	Preview bool  `json:"preview"`
	Result  struct {
		Raw                               string   `json:"_raw"`
		Time                              string   `json:"_time"`
		CriblBreaker                      string   `json:"cribl_breaker"`
		CriblPipe                         string   `json:"cribl_pipe"`
		Event                             []string `json:"event"`
		Eventtype                         []string `json:"eventtype"`
		Fields_container_id               string   `json:"fields.container.id"`
		Fields_container_image_name       string   `json:"fields.container.image.name"`
		Fields_container_image_tag        string   `json:"fields.container.image.tag"`
		Fields_deployment_environment     string   `json:"fields.deployment.environment"`
		Fields_k8s_cluster_name           string   `json:"fields.k8s.cluster.name"`
		Fields_k8s_container_name         string   `json:"fields.k8s.container.name"`
		Fields_k8s_container_restartCount string   `json:"fields.k8s.container.restart_count"`
		Fields_k8s_namespace_name         string   `json:"fields.k8s.namespace.name"`
		Fields_k8s_node_name              string   `json:"fields.k8s.node.name"`
		Fields_k8s_pod_name               string   `json:"fields.k8s.pod.name"`
		Fields_k8s_pod_uid                string   `json:"fields.k8s.pod.uid"`
		Fields_landscape                  string   `json:"fields.landscape"`
		Fields_log_iostream               string   `json:"fields.log.iostream"`
		Fields_logtag                     string   `json:"fields.logtag"`
		Fields_os_type                    string   `json:"fields.os.type"`
		Host                              []string `json:"host"`
		Index                             []string `json:"index"`
		Linecount                         string   `json:"linecount"`
		Port                              string   `json:"port"`
		Source                            []string `json:"source"`
		Sourcetype                        []string `json:"sourcetype"`
		SplunkServer                      string   `json:"splunk_server"`
		TimeEpoch                         []string `json:"time"`
	} `json:"result"`
}

// take in a file with json data process it according to the myJsonName struct
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename")
		return
	}

	filename := os.Args[1]
	searchString := os.Args[2]
	data, err := readFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Now you can use the data
	for _, item := range data {
		if strings.Contains(item.Result.Event[0], searchString) {
			line := fmt.Sprintf("%s %s %s %s %s", item.Result.Time, item.Result.Event[0], item.Result.Fields_k8s_pod_name, item.Result.Fields_k8s_cluster_name, item.Result.Fields_deployment_environment)
			fmt.Println(line)
		}
	}
}

func readFile(fileName string) ([]MyJsonName, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var data []MyJsonName
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var myJsonName MyJsonName
		err := json.Unmarshal([]byte(scanner.Text()), &myJsonName)
		if err != nil {
			return nil, err
		}
		data = append(data, myJsonName)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
