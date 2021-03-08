package cli

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/hamburghammer/rcon"
)

// Start a interactive connection. Uses the given reader and writer to inteact with the created connection.
func Start(hostPort string, password string, in io.Reader, out io.Writer) {
	remoteConsole, err := rcon.Dial(hostPort, password)
	if err != nil {
		log.Fatal("Failed to connect to RCON server", err)
	}
	defer remoteConsole.Close()

	scanner := bufio.NewScanner(in)
	out.Write([]byte("To quit the session type 'exit'.\n"))
	out.Write([]byte("> "))
	for scanner.Scan() {
		cmd := scanner.Text()
		if cmd == "exit" {
			return
		}

		reqID, err := remoteConsole.Write(cmd)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to send command:", err.Error())
			continue
		}

		resp, respID, err := remoteConsole.Read()
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Fprintln(os.Stderr, "Failed to read command:", err.Error())
			continue
		}

		if reqID != respID {
			fmt.Fprintln(out, "Weird. This response is for another request.")
		}

		fmt.Fprintln(out, resp)
		out.Write([]byte("> "))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

// Execute runs the given command against the remote console and writes the response into the writer.
func Execute(hostPort string, password string, out io.Writer, command ...string) {
	remoteConsole, err := rcon.Dial(hostPort, password)
	if err != nil {
		log.Fatal("Failed to connect to RCON server", err)
	}
	defer remoteConsole.Close()

	preparedCmd := strings.Join(command, " ")
	reqID, err := remoteConsole.Write(preparedCmd)

	resp, respID, err := remoteConsole.Read()
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Fprintln(os.Stderr, "Failed to read command:", err.Error())
		return
	}

	if reqID != respID {
		fmt.Fprintln(out, "Weird. This response is for another request.")
	}

	fmt.Fprintln(out, resp)
}
