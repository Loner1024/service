package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var service string

func init() {
	flag.StringVar(&service, "service", "", "filter which service to see")
}

func main() {
	flag.Parse()
	var b strings.Builder

	// Scan standard input for log data per line.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()

		// Convert the json to a map for processing.
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(s), &m)
		if err != nil {
			if service == "" {
				fmt.Println(s)
			}
			continue
		}

		// If a service filter was provided, check.
		if service != "" && m["service"] != service {
			continue
		}

		// Logs trace id.
		traceID := "00000000-0000-0000-0000-000000000000"
		if v, ok := m["traceID"]; ok {
			traceID = fmt.Sprintf("%v", v)
		}

		// Build out the known portions of the log in the order
		b.Reset()
		b.WriteString(fmt.Sprintf("%s: %s: %s: %s: %s: %s: ",
			m["service"],
			m["ts"],
			m["level"],
			traceID,
			m["caller"],
			m["msg"],
		))

		// Add the reset of the keys ignoring the ones we already
		// added for the log.
		for k, v := range m {
			switch k {
			case "service", "ts", "level", "traceid", "caller", "msg":
				continue
			}

			b.WriteString(fmt.Sprintf("%s[%v]: ", k, v))
		}

		// Write the new log format, removing the last
		out := b.String()
		fmt.Println(out[:len(out)-2])

	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
