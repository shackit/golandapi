package todo

import (
	"context"
	"errors"
	"fmt"
	"goland-course-api/internal/db"
	"strings"
)

type Item struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}

type Manager interface {
	InsertItem(ctx context.Context, item db.Item) error
	GetAllItems(ctx context.Context) ([]db.Item, error)
}

type Service struct {
	db Manager
}

func NewService(db Manager) *Service {
	return &Service{
		db: db,
	}
}

func (svc *Service) Add(todo string) error {
	items, err := svc.GetAll()
	if err != nil {
		return fmt.Errorf("failed to read from database %w", err)
	}

	for _, t := range items {
		if t.Task == todo {
			return errors.New("task already exists")
		}
	}

	if err := svc.db.InsertItem(context.Background(), db.Item{
		Task:   todo,
		Status: "TO_BE_STARTED",
	}); err != nil {
		return fmt.Errorf("failed to insert item into database: %w", err)
	}

	return nil
}

func (svc *Service) Search(query string) ([]string, error) {
	items, err := svc.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read from database %w", err)
	}
	var results []string
	for _, t := range items {
		if strings.Contains(strings.ToLower(t.Task), strings.ToLower(query)) {
			results = append(results, t.Task)
		}
	}

	return results, nil
}

func (svc *Service) GetAll() ([]Item, error) {
	var results []Item
	items, err := svc.db.GetAllItems(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to read from database %w", err)
	}
	for _, i := range items {
		results = append(results, Item{
			Task:   i.Task,
			Status: i.Status,
		})
	}
	return results, nil
}
