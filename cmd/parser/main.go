package main

import (
	"net/http"
	"time"
)

func main() {
	
}

type Parser struct {
	Client *http.Client
	MaxConcurrent int
}

type ParserRules struct {
	CountPages int
	Grade string
	Salary string
	Search string
	isRemote bool
}

func NewParser() *Parser {
	return &Parser{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		}, MaxConcurrent: 10,
	}
}

func (p *Parser) Parse(rules ParserRules) (error) {
	var countPages = rules.CountPages
	var Grade = rules.Grade
	var Salary = rules.Salary
	var Search = rules.Search
	var isRemote = rules.isRemote
	sem := make(chan struct{}, p.MaxConcurrent)
	
}