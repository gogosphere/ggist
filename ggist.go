package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/speedata/optionparser"
	"github.com/tjgq/clipboard"
	"golang.org/x/oauth2"
)

var (
	op        = optionparser.NewOptionParser()
	ggistUser = os.Getenv("ggistUser")
)

type State struct {
	get    string
	set    string
	match  string
	value  string
	token  string
	user   string
	prompt string
	debug  string
	c      *github.Client
}

func New() {
	var (
		options     = make(map[string]string)
		ggisttoken  = os.Getenv("ggistToken")
		ggistPrompt = os.Getenv("ggistPrompt")
	)

	op.On("-s", "--set", "send gist (exclusive)", options)
	op.On("-g", "--get", "get gist and display (exclusive)", options)
	op.On("-m", "--match", "get gist and display", options)
	op.On("-v", "--debug", "debuging output", options)
	op.On("-d", "--description VAL", "Twitterish length description of your gist", options)
	err := op.Parse()
	e(err)

	var s = &State{
		get:    options["get"],
		set:    options["set"],
		match:  options["match"],
		value:  options["description"],
		debug:  options["debug"],
		token:  ggisttoken,
		user:   ggistUser,
		prompt: ggistPrompt,
	}

	s.validate(options)

	params := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: s.token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, params)
	client := github.NewClient(tc)
	s.c = client

	if s.set == "true" {
		s.sett(s.value)
		os.Exit(0)
	}
	if s.match == "true" {
		s.gett()
		os.Exit(0)
	}
	s.gett()

}

func main() {
	New()
}

func (s *State) validate(o map[string]string) {
	testvalid := make(map[bool]int)
	s.debug = o["debug"]
	o["debug"] = ""
	for _, v := range o {
		a, err := strconv.ParseBool(v)
		if err == nil {
			testvalid[a]++
		}
	}
	if testvalid[true] > 1 {
		truth(false)
	}
	if s.token == "" || s.user == "" {
		truth(false)
	}
	if s.set == "true" && s.value == "" {
		truth(false)
	}
	if s.match == "true" && s.value == "" {
		truth(false)
	}
}

func (s *State) gett() {
	var match []*github.Gist
	gists, resp, err := s.c.Gists.List(s.user, nil)
	e(err)
	if s.debug == "true" {
		d(resp)
	}
	if s.match == "true" {
		r := regexp.MustCompile(`(?i)` + s.value)
		for _, v := range gists {
			m := r.MatchString(*v.Description)
			if m == true {
				match = append(match, v)
			}
		}
	}
	if s.match == "true" {
		pgist(match)
	} else {
		pgist(gists)
	}

}

func pgist(g []*github.Gist) {
	for _, v := range g {
		fmt.Println(*v.Description)
		color.Cyan("%v\n", *v.GitPullURL)
	}
}

func (s *State) sett(v string) {
	pub := "false"
	val, err := clipboard.Get()
	e(err)

	pauxfile := github.GistFile{Filename: hashstring(), Content: mkpointer(val)}
	mymap := make(map[github.GistFilename]github.GistFile)
	mymap["justfilename"] = pauxfile

	mygist := &github.Gist{Description: mkpointer(s.value), Files: mymap, Public: mkboolpointer(pub)}
	thisgist, resp, err := s.c.Gists.Create(mygist)
	e(err)
	if s.debug == "true" {
		d(resp)
	}
	fmt.Printf("URL: %s\n", *thisgist.GitPullURL)
}

func (s *State) matchh(v string) {}

func slct() {}

// Utility functions
func truth(a bool) {
	if a == false {
		err := errors.New("Incompatable values in truth")
		e(err)
	}
}

func d(resp *github.Response) {
	fmt.Println(resp)
}

func e(e error) {
	if e != nil {
		op.Help()
		log.Fatal("Die: ", e)
	}
}

func mkboolpointer(s string) *bool {
	v, err := strconv.ParseBool(s)
	e(err)
	z := &v
	return z
}
func mkpointer(s string) *string {
	v := &s
	return v
}
func hashstring() *string {
	size := 32
	rb := make([]byte, size)
	_, err := rand.Read(rb)
	e(err)
	rs := base64.URLEncoding.EncodeToString(rb)
	rsa := rs[1:15]
	strings.TrimRight(rsa, "1")
	return &rsa
}
