package main
import "fmt"
import "os"
import "os/exec"
import "strings"

func main() {
if len(os.Args) < 2 {
	fmt.Println("usage: murder <port>")
} else {
	port := os.Args[1]
	cmd1 := exec.Command("fuser", port+"/tcp")
	pid, err1 := cmd1.Output()
	s_pid := strings.TrimSpace(string(pid))
	cmd2 := exec.Command("kill", "-9", s_pid)
	_, err2 := cmd2.Output()
	if err1 != nil {
		fmt.Println("no ports to murder")
	} else {
	if err2 != nil {
			fmt.Println("the murder of", "port", port, "was unsuccessful", "-", err2)
	} else {
		fmt.Println("port", port, "has been murdered!")
	}
	}
}
}