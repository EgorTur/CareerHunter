package domain

import "time"

type Vacancy struct {
	ID          int64
	ExternalID  string // Уникальный ID с источника
	Title       string
	URL         string // Уникальный по ссылке
	Company     string
	PublishedAt time.Time
	CreatedAt   time.Time
	Source      string // 'habr', 'hh'
	Salary      int
}

// Job — задача для парсинга (очередь)
type Job struct {
	ID          int64
	Keyword     string
	Priorety    int    // 0=low, 1=normal, 2=high
	Status      string // pending, processing, done, failed
	RetryCount  int
	NextAttempt time.Time
	CreatedAt   time.Time
}
