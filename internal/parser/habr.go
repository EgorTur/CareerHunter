package parser

import (
	"context"
	"net/http"
	"time"
)

type HabrCareerParser struct {
	Client        *http.Client
	MaxConcurrent int
}

// type ParserRules struct {
// 	CountPages int
// 	Grade string
// 	Salary string
// 	Search string
// 	isRemote bool
// }

func NewHabrCareerParser() *HabrCareerParser {
	return &HabrCareerParser{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		MaxConcurrent: 10,
	}
}

func (p *HabrCareerParser) Fetch(ctx context.Context, keywords []string) ([]domain.Vacancy, error) {

	var vacancies []domain.Vacancy

	sem := make(chan struct{}, p.MaxConcurrent)

	for _, keyword := range keywords {
		
	}
	return vacancies, nil
}
