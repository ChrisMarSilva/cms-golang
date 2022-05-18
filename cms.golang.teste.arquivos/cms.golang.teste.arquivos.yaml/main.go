package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
	//"gopkg.in/yaml.v3"
)

type myData struct {
	Conf struct {
		Hits      int64
		Time      int64
		CamelCase string `yaml:"camelCase"`
	}
}

type Config1 struct {
	Info string
	Data struct {
		Source      string
		Destination string
	}
	Run []struct {
		Id     string
		Exe    string
		Output string
	}
}

type Config struct {
	Firewall_network_rules map[string][]string
}

type Options struct {
	Src string
	Dst string
}

type Yaml struct {
	Schema     string
	ID         string
	Version    string
	Dependency []Dependency
}

type Dependency struct {
	Name     string
	Type     string
	CWD      string
	Install  []Install
	Provides []Provide
}

type Install struct {
	Name       string
	Group      string
	Type       string
	Properties Properties
}

type Properties struct {
	Name string
	URL  string
}

type Provide struct {
	Name       string
	Properties Properties
}

type Service struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
		Labels    struct {
			RouterDeisIoRoutable string `yaml:"router.deis.io/routable"`
		} `yaml:"labels"`
		Annotations struct {
			RouterDeisIoDomains string `yaml:"router.deis.io/domains"`
		} `yaml:"annotations"`
	} `yaml:"metadata"`
	Spec struct {
		Type     string `yaml:"type"`
		Selector struct {
			App string `yaml:"app"`
		} `yaml:"selector"`
		Ports []struct {
			Name       string `yaml:"name"`
			Port       int    `yaml:"port"`
			TargetPort int    `yaml:"targetPort"`
			NodePort   int    `yaml:"nodePort,omitempty"`
		} `yaml:"ports"`
	} `yaml:"spec"`
}

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main() {

	filename := "conf.yaml"

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	//-------------

	// c := &myData{}
	// err = yaml.Unmarshal(buf, c)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf(*c)

	//-------------

	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", config.Firewall_network_rules)

	//-------------

	var service Service
	err = yaml.Unmarshal(buf, &service)
	if err != nil {
		panic(err)
	}
	fmt.Println(service.Metadata.Name)

	//-------------

	var conf Config1
	reader, _ := os.Open(filename)
	buf1, _ := ioutil.ReadAll(reader)
	yaml.Unmarshal(buf1, &conf)
	fmt.Printf("%+v\n", conf)

	//-------------

	fmt.Println("FIM")
}
