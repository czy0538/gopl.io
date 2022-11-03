// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

const (
	year  = 24 * 365
	month = 24 * 30
)

// !+
func main() {
	issueBody := github.IssueBody{
		Title: "test",
		Body:  time.Now().String(),
	}
	err := github.CreateIssue(os.Args[1:], &issueBody)
	if err != nil {
		log.Fatalln(err)
	}

	//result, err := github.SearchIssues(os.Args[1:])
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%d issues:\n", result.TotalCount)
	//for _, item := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %s %.55s\n",
	//		item.Number, item.User.Login, item.CreatedAt, item.Title)
	//}
	//
	//sortByTimeResult := SortByTime(result)
	//SortByTimePrint(sortByTimeResult, "less_than_one_month")
	//SortByTimePrint(sortByTimeResult, "less_than_one_year")
	//SortByTimePrint(sortByTimeResult, "more_than_one_year")

}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/

func SortByTime(result *github.IssuesSearchResult) *map[string][]*github.Issue {
	now := time.Now()
	issuesTimeSortMap := make(map[string][]*github.Issue)
	for _, item := range result.Items {
		dt := now.Sub(item.CreatedAt)
		if dt.Hours() < month {
			issuesTimeSortMap["less_than_one_month"] = append(issuesTimeSortMap["less_than_one_month"], item)
		} else if dt.Hours() < year {
			issuesTimeSortMap["less_than_one_year"] = append(issuesTimeSortMap["less_than_one_year"], item)
		} else {
			issuesTimeSortMap["more_than_one_year"] = append(issuesTimeSortMap["more_than_one_year"], item)
		}
	}

	return &issuesTimeSortMap
}

func SortByTimePrint(result *map[string][]*github.Issue, name string) bool {
	items, ok := (*result)[name]
	if ok {
		fmt.Println(name)
		for _, item := range items {
			fmt.Printf("#%-5d %-9.9s %s %.55s\n",
				item.Number, item.User.Login, item.CreatedAt, item.Title)
		}
		fmt.Println("")
	}

	return ok
}
