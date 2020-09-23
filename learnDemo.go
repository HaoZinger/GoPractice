package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	listAll("./", 1)
	listAll_pro("./", 0)
}

func listAll(dir string, line int) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			for count := line; count > 0; count-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
			listAll(dir+"/"+info.Name(), line+1)
		} else {
			for count := line; count > 0; count-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}

	}

}
func listAll_pro(path string, curHier int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), "\\")
			listAll_pro(path+"/"+info.Name(), curHier+1)
		} else {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}
