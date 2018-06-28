package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	getopt "github.com/kesselborn/go-getopt"
)

type Conf map[string][]string;

func main() {
	optionDefinition := getopt.Options{
		"服务器列表管理",
		getopt.Definitions{
			{"config|c", "yaml config file", getopt.IsConfigFile | getopt.Required, "./config.yaml"},
			{"project|p", "project name", getopt.Optional, "im"},
			{"index|i", "server index", getopt.Optional, 0}, 
		},
	}

	options, _, _, e := optionDefinition.ParseCommandLine()

	help, wantsHelp := options["help"]

	if e != nil || wantsHelp {
		exit_code := 0

		switch {
		case wantsHelp && help.String == "usage":
			fmt.Print(optionDefinition.Usage())
		case wantsHelp && help.String == "help":
			fmt.Print(optionDefinition.Help())
		default:
			fmt.Println("**** Error: ", e.Error(), "\n", optionDefinition.Help())
			exit_code = e.ErrorCode
		}
		os.Exit(exit_code)
	}

	content := getConfig(options["config"].String,
		options["project"].String,
		options["index"].Int, 
	);
	fmt.Println(content);
}

func getConfig(file string, project string, index int64) string {
	content, err := ioutil.ReadFile(file);
	if(err != nil){
		fmt.Printf("read file err :%s\n", err);
	}

	m := make(Conf);
	err = yaml.Unmarshal(content, &m);
	if(err != nil){
		fmt.Printf("yaml.unmarshal err: $s\n", err);
	}

	if(project == ""){
		return getProjectList(m)
	}

	if(m[project] != nil){
		if(index == 0){
			return getServerList(m, project);
		}else{
			l := int64(len(m[project]))
			if(index <= l){
				return m[project][index-1];
			}
		}
	}

	return ""
}

func getProjectList(conf Conf) string{
	//返回可选的project列表
	project_list := "avaiable project:\n"
	for i, _ := range conf {
		project_list += i + "\n";
	}
	return project_list;
}

func getServerList(conf Conf, project string) string{
	server_list := ""
	for i, v := range conf[project] {
		server_list += fmt.Sprintf("%d:%s\n", i+1, v);
	}
	return server_list;
}
