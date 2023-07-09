package main

// this program can take the items in scaffold.yaml, and proceed to making scaffolding for ef core
import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Hello World!")
	cmd := exec.Command("powershell", "ls", "./")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
