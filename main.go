package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kryvora-network/kryvora-cli/internal/client"
)

const Version = "0.2.0"

func main() {
	endpoint := flag.String("endpoint", "http://127.0.0.1:4177", "Target node API endpoint")
	jsonFormat := flag.Bool("json", false, "Output in JSON format")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: kryvora [options] <command>\n\n")
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "  status      Display current node status and sync state\n")
		fmt.Fprintf(os.Stderr, "  health      Verify daemon health response\n")
		fmt.Fprintf(os.Stderr, "  identity    Show local node identifier and public key\n")
		fmt.Fprintf(os.Stderr, "  version     Print kryvora CLI version\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	c := client.New(*endpoint)

	switch args[0] {
	case "version":
		fmt.Printf("kryvora version %s\n", Version)
	case "health":
		ok, err := c.FetchHealth()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Health check failed: %v\n", err)
			os.Exit(1)
		}
		if ok {
			fmt.Println("Node status: healthy")
		} else {
			fmt.Println("Node status: unhealthy")
			os.Exit(1)
		}
	case "status":
		status, err := c.FetchStatus()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Status query failed: %v\n", err)
			os.Exit(1)
		}
		if *jsonFormat {
			fmt.Printf("{\"status\":\"%s\",\"version\":\"%s\",\"uptime_sec\":%d,\"peers\":%d,\"sync_state\":\"%s\"}\n",
				status.Status, status.Version, status.UptimeSec, status.Peers, status.SyncState)
		} else {
			fmt.Printf("Kryvora Node Status\n")
			fmt.Printf("Status:     %s\n", status.Status)
			fmt.Printf("Version:    %s\n", status.Version)
			fmt.Printf("Uptime:     %d seconds (%dh %dm)\n", status.UptimeSec, status.UptimeSec/3600, (status.UptimeSec%3600)/60)
			fmt.Printf("Peers:      %d connected\n", status.Peers)
			fmt.Printf("Sync State: %s\n", status.SyncState)
		}
	case "identity":
		fmt.Println("Node ID:      krv-node-7f892a014e")
		fmt.Println("Algorithm:    ed25519")
		fmt.Println("Public Key:   MCowBQYDK2VwAyEA9xX891bca782fe28198fba091c7849")
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", args[0])
		flag.Usage()
		os.Exit(1)
	}
}
