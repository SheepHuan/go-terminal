package remote_term

import (
	"bufio"
	"fmt"
	"go-terminal/pkg"
	"os"
	"time"
)

func interactive(terminal pkg.RemoteTerminal) {
	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	for cmd != "!exit()" {
		sout, serr, _ := terminal.Execute(cmd)
		fmt.Println(fmt.Sprintf("%s\n%s", sout, serr))
		cmd, _ = reader.ReadString('\n')
	}

}

func Test() {

	ip, port := "127.0.0.1", 8080
	go pkg.LaunchRemoteTerminalService(ip, port)
	time.Sleep(2e9)
	rem_term := pkg.RemoteTerminal{}
	rem_term.Init(ip, port)
	interactive(rem_term)
}
