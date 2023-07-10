package main

// this program can take the items in scaffold.yaml, and proceed to making scaffolding for ef core
import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
)

type Scaffold struct {
	ProjectLocation  string `yaml:"projectLocation"`
	ConnectionString string `yaml:"connectionString"`
	Provider         string `yaml:"provider"`
	Tables           []string
}

func main() {
	dat, err := os.ReadFile("./scaffold.yaml")
	fmt.Print(string(dat))
	if err != nil {
		fmt.Println(err)
	}

	scaffold := Scaffold{}
	err = yaml.Unmarshal(dat, &scaffold)
	fmt.Println(scaffold.ConnectionString)
	if err != nil {
		fmt.Println(err.Error())
	}
	cmd := exec.Command("powershell")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(out))
}
