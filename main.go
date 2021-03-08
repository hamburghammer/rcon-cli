package main

import (
	"fmt"
	"net"
	"os"

	"github.com/itzg/rcon-cli/cli"
	flags "github.com/spf13/pflag"
)

const (
	defaultHost     = "localhost"
	defaultPort     = "25575"
	defaultPassword = ""

	envVarPrefix = "RCON_CLI_"
)

// vars holding the parsed configuration
var (
	host     string
	port     string
	password string
	help     bool
)

func main() {
	commands, _ := parseArgs(os.Args[1:])
	loadEnvIfNotSet()
	if help {
		printHelp()
		os.Exit(0)
		return
	}

	run(commands...)
}

func parseArgs(args []string) ([]string, error) {
	flags.StringVar(&host, "host", defaultHost, "RCON server's hostname.")
	flags.StringVar(&port, "port", defaultPort, "RCON server's port.")
	flags.StringVar(&password, "password", defaultPassword, "RCON server's password.")
	flags.BoolVarP(&help, "help", "h", false, "Prints this help message and exits.")

	err := flags.CommandLine.Parse(args)
	if err != nil {
		return []string{}, err
	}

	command := flags.Args()

	return command, nil
}

func loadEnvIfNotSet() {
	envHost := os.Getenv(fmt.Sprintf("%sHOST", envVarPrefix))
	if envHost != "" && host == defaultHost {
		host = envHost
	}

	envPort := os.Getenv(fmt.Sprintf("%sPORT", envVarPrefix))
	if envPort != "" && port == defaultPort {
		port = envPort
	}

	envPassword := os.Getenv(fmt.Sprintf("%sPASSWORD", envVarPrefix))
	if envPassword != "" && password == defaultPassword {
		password = envPassword
	}
}

func run(args ...string) {
	hostPort := net.JoinHostPort(host, port)

	if len(args) == 0 {
		cli.Start(hostPort, password, os.Stdin, os.Stdout)
	} else {
		cli.Execute(hostPort, password, os.Stdout, args...)
	}
}

func printHelp() {
	fmt.Println(`rcon-cli is a CLI to interact with a RCON server.
It can be run in an interactive mode or to execute a single command.

USAGE:
	rcon-cli [FLAGS] [RCON command ...]
	
FLAGS:`)
	flags.PrintDefaults()
	fmt.Printf(`
ENVIRONMENT VARIABLE:
	All flags can be set through the flag name in capslock with the %s prefix (see examples).
	Flags have allways priority over env vars!

EXAMPLES:
	rcon-cli --host 127.0.0.1 --port 25575
	rcon-cli --password admin123 stop
	%sPORT=25575 rcon-cli stop
`, envVarPrefix, envVarPrefix)
}
