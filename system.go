package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

type System struct {
	eventBus            *EventBus
	clientService       *ClientService
	buffer              *Buffer
	dispatcher          *ApplicationDispatcher
	statistics          *Statistics
	notificationService *NotificationService
	generator           *PoissonGenerator
	stepInterval        float64 // интервал одного шага в секундах

}

func NewSystem(bufferSize, managerCount, managerLoad int, lambda float64, stepInterval float64) *System {
	eventBus := NewEventBus()
	buffer := NewBuffer(bufferSize, eventBus)

	return &System{
		eventBus:            eventBus,
		clientService:       NewClientService(eventBus),
		buffer:              buffer,
		dispatcher:          NewApplicationDispatcher(managerCount, managerLoad, buffer, eventBus),
		statistics:          NewStatistics(eventBus),
		notificationService: NewNotificationService(eventBus),
		generator:           NewPoissonGenerator(lambda),
		stepInterval:        stepInterval,
	}
}

func (s *System) processNextStep() {
	// Обрабатываем заявки из буфера
	s.dispatcher.processBuffer()

	// Симулируем работу менеджеров
	s.dispatcher.simulateManagersWork()

	// Создаем новую заявку только если генератор Пуассона говорит "да"
	eventsCount := s.generator.GetEventsCountForInterval(s.stepInterval)

	// Создаем все сгенерированные заявки
	for i := 0; i < eventsCount; i++ {
		app := s.clientService.CreateApplication()
		s.eventBus.Publish(Event{
			Type:      "NewApplication",
			Data:      app,
			Timestamp: time.Now(),
		})
		s.dispatcher.ProcessApplication(app)
	}
}

func (s *System) RunStepMode() {
	screen.Clear()
	for {
		screen.MoveTopLeft()
		fmt.Println(time.Now())
		fmt.Println("\nPress Enter to continue or 'q' to quit...")

		var input string
		fmt.Scanln(&input)
		if input == "q" {
			break
		}

		screen.Clear()
		s.processNextStep()
		s.printSystemState()
	}
}

func (s *System) RunAutoMode() {
	for j := 0; j < 5; j++ {
		for i := 0; i < 100; i++ {
			s.processNextStep()
			time.Sleep(time.Second / 50)
		}
	}
	s.printFinalStatistics()
}

func (s *System) printSystemState() {
	visualizer := NewSystemVisualizer(s)
	visualizer.PrintSystemState()
}

func (s *System) printFinalStatistics() {
	s.statistics.PrintCurrentStats()
}
