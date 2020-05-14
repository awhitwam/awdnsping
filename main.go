package main

import (
	"encoding/csv"
	"github.com/bogdanovich/dns_resolver"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

//aw test

import (
	"fmt"
	"strings"
)

func main() {
	var queryServer string = os.Args[1]
	var filename string = "DnsPing " + strings.Replace(queryServer, ".", "-", -1) + ".csv"
	server := []string{queryServer}
	resolver := dns_resolver.New(server)
	lines := LoadSites()
	file, err := os.Create(filename)
	cntrtotal := 0
	cntrbad := 0
	_ = err
	lastbad := ""
	tabout := new(tabwriter.Writer)
	tabout.Init(os.Stdout, 8, 8, 0, '\t', 0)

	fmt.Printf("%s", server)
	for {
		site := lines[rand.Intn(2000)][1]
		resolver.RetryTimes = 1
		start := time.Now()
		ip, err := resolver.LookupHost(site)
		duration := time.Since(start)
		cntrtotal = cntrtotal + 1
		if err != nil {
			cntrbad = cntrbad + 1
			lastbad = start.Format("15:04:05")
		}
		time.Sleep(200)

		//log.Printf("%s >%s< in %d--%s", server, site, duration.Milliseconds(), err)
		var myoutput string = fmt.Sprintf("%s,%s,%s,%d,%s", start.Format(("2006-01-02 15:04:05")), server, site, duration.Milliseconds(), err)
		//fmt.Printf("\r",myoutput)
		//fmt.Printf("\rTotal:%d Bad:%d   Working:%s  Last Fail:%s                              ",cntrtotal,cntrbad,site,lastbad)
		defer tabout.Flush()
		fmt.Fprintf(tabout, "\rTotal:%d\tBad:%d\tLfail:%s\tLookup:%s\t                 ", cntrtotal, cntrbad, lastbad, site)
		tabout.Flush()
		i, err := file.WriteString(myoutput + "\r\n")
		_ = i
		_ = err
		_ = ip

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
