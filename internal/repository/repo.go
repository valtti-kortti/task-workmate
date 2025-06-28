package repository

import (
	"errors"
	"sync"
)

type repository struct {
	repo map[int]*Task
	mu   sync.RWMutex
	id   int
}

type Repository interface {
	CreateTask(task *Task) *Task
	GetTask(id int) (*Task, error)
	DeleteTask(id int) error
}

func NewRepository() Repository {
	return &repository{repo: map[int]*Task{}, id: 0}
}

func (r *repository) CreateTask(task *Task) *Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.id++
	task.ID = r.id
	r.repo[r.id] = task
	return task
}

func (r *repository) GetTask(id int) (*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	task, ok := r.repo[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (r *repository) DeleteTask(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.repo[id]; !ok {
		return errors.New("task not found")
	}

	delete(r.repo, id)
	return nil
}
