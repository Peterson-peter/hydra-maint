package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseOutput(command *exec.Cmd) []string {
	out, err := command.Output()
	check(err)
	return bytesToString(out)
}

func bytesToString(data []byte) []string {
	//cast to []string
	return strings.Split(string(data[:]), "\n")
}

func verifyLink(link string) string {
	//command := exec.Command("readlink", "-f", link)

	//destion, err := os.Readlink(command)
	//(err)
	return link
}

func main() {
	dockerpids := parseOutput(exec.Command("/bin/bash", "-c", "pgrep dockerd"))
	location := "/proc/" + dockerpids[0] + "/fd/" //"ls -lh '/proc/" + dockerpids[0] + "/fd/'"
	fmt.Println(location)
	files, err := ioutil.ReadDir(location)
	check(err)
	for _, element := range files {
		if element.IsDir() == false {
			fi, err := os.Readlink(location + element.Name())
			check(err)
			if strings.Contains(fi, "deleted") {
				bugfile, _ := filepath.EvalSymlinks(location + element.Name())
				os.Truncate(location+element.Name(), 0)
				fmt.Println(bugfile)
			}

		}
	}

}
