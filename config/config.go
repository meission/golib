/*------------------------
name        config
describe    wmntec config library
author      ailn(z.ailn@wmntec.com)
create      2016-04-18
version     1.0
------------------------*/
package config

import (
	//golang official package
	"io/ioutil"
	"os"
	"strings"
)

//read config from ini style config file
//comChar like '#' in *nix or ';' in win
func ReadINI(filePath, comChar string) (map[string]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buff, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	conf := make(map[string]string)
	lines := strings.Split(string(buff), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, comChar) { //ignore comment line
			if len(line) >= 2 {
				if line[0] == "["[0] && line[len(line)-1] == "]"[0] { //section
					// ignore section line
				} else {
					kv := strings.Split(line, "=")
					if len(kv) == 2 {
						conf[kv[0]] = kv[1]
					}
				}
			}
		}
	}
	return conf, nil
}
