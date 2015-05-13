package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"os"
	"strings"
)

const (
	fp string = "config.ini"
)

var ini config.ConfigContainer

func main() {
	ini, _ = config.NewConfig("ini", fp)
	args := os.Args
	section := ""
	key := ""
	value := ""
	for k, v := range args {
		if k == 0 {
			continue
		}
		switch v {
		case "-h":
			fmt.Println("Usage:the doc is needed")
		case "-s":
			if len(args) >= k+2 {
				section = args[k+1]
			}
		case "-k":
			if len(args) >= k+2 {
				key = args[k+1]
			}
		case "-a":
			if len(args) >= k+2 {
				arg := strings.Split(args[k+1], ".")
				if len(arg) > 0 {
					section = arg[0]
					key = arg[1]
				}
			}
		case "-v":
			if len(args) >= k+2 {
				value = args[k+1]
			}
		}
	}

	if section != "" {
		if key != "" {
			if value != "" {
				add_key(section, key, value)
			} else {
				val := read(section, key)
				if val == "" {
					fmt.Println("you had not setted tips for:", section+"."+key)
				} else {
					fmt.Println(section + "." + key + ":" + val)
				}
			}
		} else {
			keys := kvs(section)
			if len(keys) > 0 {
				for _, v2 := range keys {
					val := read(section, v2)
					fmt.Println(section+"."+v2+": ", val)
				}
			} else {
				fmt.Println("you had not setted tips for:", section)
			}

		}
	} else {
		sections := sections()
		if len(sections) > 0 {
			for _, v1 := range sections {
				keys := kvs(v1)
				if len(keys) > 0 {
					for _, v2 := range keys {
						val := read(v1, v2)
						fmt.Println(v1+"."+v2+": ", val)
					}
				} else {
					continue
				}
			}
		} else {
			fmt.Println("you had not setted tips")
		}
	}
}

/**
* @Title:read
* @author:caozhipan
* @param
* @return
 */
func read(section, key string) string {
	return ini.String(section + "::" + key)
}

/**
* @Title:sessions
* @author:caozhipan
* @distruction:获取所有section
* @param
* @return
 */
func sections() []string {
	sections := ini.Strings("sections::section")
	return sections
}

/**
* @Title:kvs
* @distruction:获取section下所有键名
* @author:caozhipan
* @param
* @return
 */
func kvs(section string) []string {
	kvs := ini.Strings(section + "::kvs")
	return kvs
}

/**
* @Title:enable_section
* @author:caozhipan
* @param
* @return
 */
func enable_section(section string) int {
	sections := sections()
	ns := ""
	for _, s := range sections {
		if s == section {
			return 0
		}
		if s != "" {
			ns = ns + s + ";"
		}
	}
	ns = ns + section
	ini.Set("sections::section", ns)
	ini.SaveConfigFile(fp)
	return 1
}

/**
* @Title:enable_key
* @author:caozhipan
* @param
* @return
 */
func enable_key(section, key string) int {
	kvs := kvs(section)
	nk := ""
	for _, v := range kvs {
		if v == key {
			return 0
		}
		if v != "" {
			nk = nk + v + ";"
		}
	}
	nk = nk + key
	ini.Set(section+"::kvs", nk)
	ini.SaveConfigFile(fp)
	return 1
}

/**
* @Title:add_key
* @author:caozhipan
* @param
* @return
 */
func add_key(section, key, value string) {
	enable_section(section)
	enable_key(section, key)
	ini.Set(section+"::"+key, value)
	ini.SaveConfigFile(fp)
}
