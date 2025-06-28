package service

import (
	"math/rand"
	"strconv"
	"time"

	"task-workmate/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type service struct {
	repo repository.Repository
}

type Service interface {
	CreateTask(ctx *fiber.Ctx) error
	DeleteTask(ctx *fiber.Ctx) error
	GetTask(ctx *fiber.Ctx) error
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateTask(ctx *fiber.Ctx) error {
	var req *repository.Task
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// назначаем время создания и указываем его
	req.CreatedAt = time.Now()

	// указываем статус in_progress
	req.Status = "in_progress"

	// горутина в которой будет сделан запрос с задачей, пока через sleep
	go func(task *repository.Task) {
		// моделируем ситуацию задержки от 3 до 5 минут
		seconds := rand.Intn(120) + 180
		duration := time.Duration(seconds) * time.Second
		time.Sleep(duration)

		// по завершению моделируем получение статуса
		task.Status = "done"

		// фиксируем время окончания, считаем время выполнения и указываем его
		timeEnd := time.Now()
		processTime := int64(timeEnd.Sub(task.CreatedAt).Seconds())
		task.ProcessTime = processTime

	}(req)

	s.repo.CreateTask(req)

	return ctx.Status(fiber.StatusOK).JSON(req)
}

func (s *service) DeleteTask(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	if id < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	if err := s.repo.DeleteTask(id); err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "task deleted"})
}

func (s *service) GetTask(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	if id < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	task, err := s.repo.GetTask(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}
