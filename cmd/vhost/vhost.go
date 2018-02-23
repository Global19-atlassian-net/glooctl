package vhost

import (
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/solo-io/gloo-api/pkg/api/types/v1"
	"github.com/solo-io/gloo-storage/file"
	"github.com/solo-io/gloo/pkg/protoutil"
)

func parseFile(filename string) (*v1.VirtualHost, error) {
	var v v1.VirtualHost
	err := file.ReadFileInto(filename, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func printJSON(v *v1.VirtualHost) {
	b, err := protoutil.Marshal(v)
	if err != nil {
		fmt.Println("unable to convert to JSON ", err)
		return
	}
	fmt.Println(string(b))
}

func printYAML(v *v1.VirtualHost) {
	jsn, err := protoutil.Marshal(v)
	if err != nil {
		fmt.Println("uanble to marshal ", err)
		return
	}
	b, err := yaml.JSONToYAML(jsn)
	if err != nil {
		fmt.Println("unable to convert to YAML", err)
		return
	}
	fmt.Println(string(b))
}

func printJSONList(vhosts []*v1.VirtualHost) {
	for _, v := range vhosts {
		printJSON(v)
	}
}

func printYAMLList(vhosts []*v1.VirtualHost) {
	for _, v := range vhosts {
		printYAML(v)
	}
}

func printSummaryList(v []*v1.VirtualHost) {
	for _, entry := range v {
		fmt.Println(entry.Name)
	}
}