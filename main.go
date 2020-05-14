package main

import (
	"encoding/csv"
	"github.com/bogdanovich/dns_resolver"
	"log"
	"math/rand"
	"os"
	"time"
)

//aw test

import (
	"encoding/csv"
	"fmt"
	"github.com/bogdanovich/dns_resolver"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	var queryServer string = os.Args[1]
	var filename string = "DnsPing " + strings.Replace(queryServer, ".", "-", -1) + ".csv"
	server := []string{queryServer}
	resolver := dns_resolver.New(server)
	lines := LoadSites()
	file, err := os.Create(filename)
	_ = err

	for {
		site := lines[rand.Intn(2000)][1]
		resolver.RetryTimes = 1
		start := time.Now()
		ip, err := resolver.LookupHost(site)
		duration := time.Since(start)

		time.Sleep(200)
		_ = ip

		log.Printf("%s >%s< in %d--%s", server, site, duration.Milliseconds(), err)
		var myoutput string = fmt.Sprintf("%s,%s,%s,%d,%s", start.Format(("2006-01-02 15:04:05")), server, site, duration.Milliseconds(), err)
		i, err := file.WriteString(myoutput + "\r\n")
		_ = i
		_ = err

	}
	file.Close()
}
func LoadSites() [][]string {
	lines, err := ReadCsv("top2000.csv")
	if err != nil {
		panic(err)
	}
	return lines
}

func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
