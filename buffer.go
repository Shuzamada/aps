package main

import (
    "sort"
    "time"
)

type Buffer struct {
    maxSize       int
    applications  []*Application
    eventBus      *EventBus
}

func NewBuffer(size int, eventBus *EventBus) *Buffer {
    return &Buffer{
        maxSize:      size,
        applications: make([]*Application, 0),
        eventBus:     eventBus,
    }
}

func (b *Buffer) Add(app *Application) bool {
    if len(b.applications) >= b.maxSize {
        b.removeOldestApplication()
    }

    b.applications = append(b.applications, app)
    b.sortByPriority()
    
    b.eventBus.Publish(Event{
        Type:      "ApplicationBuffered",
        Data:      app,
        Timestamp: time.Now(),
    })
    return true
}

func (b *Buffer) removeOldestApplication() {
    if len(b.applications) == 0 {
        return
    }

    oldestApp := b.applications[0]
    b.applications = b.applications[1:]

    b.eventBus.Publish(Event{
        Type:      "ApplicationRemoved",
        Data:      oldestApp,
        Timestamp: time.Now(),
    })
}

func (b *Buffer) sortByPriority() {
    sort.Slice(b.applications, func(i, j int) bool {
        return b.applications[i].Priority > b.applications[j].Priority
    })
}
