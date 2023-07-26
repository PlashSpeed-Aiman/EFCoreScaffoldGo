package main

// this program can take the items in scaffold.yaml, and proceed to making scaffolding for ef core
import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"strings"
)

type Scaffold struct {
	ProjectLocation  string   `yaml:"projectLocation"`
	ConnectionString string   `yaml:"connectionString"`
	Provider         string   `yaml:"provider"`
	Tables           []string `yaml:"tables"`
	Args             []string `yaml:"args"`
}

// Will be used soon
type ScaffoldV2 struct {
	ProjectLocation  string `yaml:"projectLocation"`
	ConnectionString struct {
		DataSource string `yaml:"dataSource"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Database   string `yaml:"database"`
	} `yaml:"connectionString"`
	ConnectionStringOptions []string `yaml:"connectionStringOptions"`
	Provider                string   `yaml:"provider"`
	Tables                  []string `yaml:"tables"`
	Args                    []string `yaml:"args"`
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
	res := JoinString(&scaffold)
	fmt.Println(res)
	cmd := exec.Command("powershell", "cd "+scaffold.ProjectLocation+";", "dotnet ef dbcontext scaffold ", res)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(out))
	fmt.Println("Press Any Key To Continue")
	fmt.Scanf(`%s`)
}

func JoinString(scaffold *Scaffold) string {
	var sb strings.Builder
	sb.WriteString(scaffold.ConnectionString)
	sb.WriteString(" ")
	sb.WriteString(scaffold.Provider)
	sb.WriteString(" ")
	for _, str := range scaffold.Tables {
		sb.WriteString("-t ")
		sb.WriteString(str)
	}
	for _, str := range scaffold.Args {
		sb.WriteString(" ")
		sb.WriteString(str)
	}
	return sb.String()

}
func JoinStringV2(scaffold *ScaffoldV2) string {
	var sb strings.Builder
	sb.WriteString(JoinConnectionString(scaffold))
	sb.WriteString(" ")
	sb.WriteString(scaffold.Provider)
	sb.WriteString(" ")
	for _, str := range scaffold.Tables {
		sb.WriteString("-t ")
		sb.WriteString(str)
	}
	for _, str := range scaffold.Args {
		sb.WriteString(" ")
		sb.WriteString(str)
	}

	return sb.String()

}
func JoinConnectionString(v2 *ScaffoldV2) string {
	var sb strings.Builder
	sb.WriteString("\"")
	sb.WriteString("data source=" + v2.ConnectionString.DataSource + ";")
	sb.WriteString("user id=" + v2.ConnectionString.Username + ";")
	sb.WriteString("password=" + v2.ConnectionString.Password + ";")
	sb.WriteString("database=" + v2.ConnectionString.Database + ";")
	for _, str := range v2.ConnectionStringOptions {
		sb.WriteString(str + ";")
	}
	sb.WriteString("\"")
	return sb.String()
}
