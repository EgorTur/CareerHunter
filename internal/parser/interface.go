package parser

import (
	"context"

	"github.com/EgorTur/CareerHunter/internal/domain"
)

type VacancySource interface {
	Fetch(ctx context.Context, keywords []string) ([]domain.Vacancy, error)
}
