package cli

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/james4k/rcon"
)

func Start(hostPort string, password string, in io.Reader, out io.Writer) {
	remoteConsole, err := rcon.Dial(hostPort, password)
	if err != nil {
		log.Fatal("Failed to connect to RCON server", err)
	}
	defer remoteConsole.Close()

	scanner := bufio.NewScanner(in)
	out.Write([]byte("> "))
	for scanner.Scan() {
		cmd := scanner.Text()
		reqId, err := remoteConsole.Write(cmd)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to send command:", err.Error())
			continue
		}

		resp, respReqId, err := remoteConsole.Read()
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Fprintln(os.Stderr, "Failed to read command:", err.Error())
			continue
		}

		if reqId != respReqId {
			fmt.Fprintln(out, "Weird. This response is for another request.")
		}

		fmt.Fprintln(out, resp)
		out.Write([]byte("> "))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func Execute(hostPort string, password string, out io.Writer, command ...string) {
	remoteConsole, err := rcon.Dial(hostPort, password)
	if err != nil {
		log.Fatal("Failed to connect to RCON server", err)
	}
	defer remoteConsole.Close()

	preparedCmd := strings.Join(command, " ")
	reqId, err := remoteConsole.Write(preparedCmd)

	resp, respReqId, err := remoteConsole.Read()
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Fprintln(os.Stderr, "Failed to read command:", err.Error())
		return
	}

	if reqId != respReqId {
		fmt.Fprintln(out, "Weird. This response is for another request.")
	}

	fmt.Fprintln(out, resp)
}
