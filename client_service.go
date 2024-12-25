package main

import (
    "math/rand"
    "time"
)

type ClientService struct {
    eventBus *EventBus
    nextID   int
}

func NewClientService(eventBus *EventBus) *ClientService {
    return &ClientService{
        eventBus: eventBus,
        nextID:   1,
    }
}

func (cs *ClientService) CreateApplication() *Application {
    eventTypes := []EventType{Wedding, Corporate, Birthday}
    randomPriority := rand.Intn(5) + 1 // Priority from 1 to 5

    app := &Application{
        ID:        cs.nextID,
        ClientID:  cs.nextID,
        EventType: eventTypes[rand.Intn(len(eventTypes))],
        Status:    "New",
        Priority:  randomPriority,
        CreatedAt: time.Now(),
        Requirements: &Requirements{
            ID:           cs.nextID,
            GuestCount:   50 + rand.Intn(200),
            Budget:       5000 + float64(rand.Intn(15000)),
            PreferredDate: time.Now().AddDate(0, rand.Intn(6), rand.Intn(30)),
            Status:       "New",
        },
    }
    cs.nextID++
    return app
}
