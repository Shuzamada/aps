
package main

import (
    "time"
)
import "math/rand"

type ApplicationDispatcher struct {
    managers     []*Manager
    eventBus     *EventBus
    buffer       *Buffer
}

func NewApplicationDispatcher(managerCount, managerLoad int, buffer *Buffer, eventBus *EventBus) *ApplicationDispatcher {
    managers := make([]*Manager, managerCount)
    for i := 0; i < managerCount; i++ {
        managers[i] = NewManager(i+1, managerLoad)
    }

    return &ApplicationDispatcher{
        managers: managers,
        eventBus: eventBus,
        buffer:   buffer,
    }
}

func (d *ApplicationDispatcher) ProcessApplication(app *Application) {
    manager := d.findSuitableManager(app)
    if manager == nil {
        app.Status = StatusInQueue
        d.buffer.Add(app)
        return
    }

    d.assignToManager(app, manager)
}

func (d *ApplicationDispatcher) processBuffer() {
    if len(d.buffer.applications) == 0 {
        return
    }

    manager := d.findAvailableManager()
    if manager != nil {
        app := d.buffer.applications[0]
        d.buffer.applications = d.buffer.applications[1:]

        d.assignToManager(app, manager)

        d.eventBus.Publish(Event{
            Type:      "ApplicationTakenFromBuffer",
            Data:      app,
            Timestamp: time.Now(),
        })
    }
}

func (d *ApplicationDispatcher) simulateManagersWork() {
    for _, manager := range d.managers {
        if manager.CurrentLoad > 0 && rand.Float32() < 0.3 {
            if completedApp := manager.completeRandomApplication(); completedApp != nil {
                if completedApp.Status == StatusApproved {
                    d.eventBus.Publish(Event{
                        Type:      "ApplicationProcessed",
                        Data:      completedApp,
                        Timestamp: time.Now(),
                    })
                }
                if completedApp.Status == StatusRejected {
                    d.eventBus.Publish(Event{
                        Type:      "ApplicationRejected",
                        Data:      completedApp,
                        Timestamp: time.Now(),
                    })
                }
            }
        }
    }
}

func (d *ApplicationDispatcher) findSuitableManager(app *Application) *Manager {
    var bestManager *Manager
    bestScore := -1

    for _, manager := range d.managers {
        if manager.CurrentLoad >= manager.MaxLoad {
            continue
        }

        score := manager.calculateSuitability(app)
        if score > bestScore {
            bestScore = score
            bestManager = manager
        }
    }

    return bestManager
}

func (d *ApplicationDispatcher) findAvailableManager() *Manager {
    for _, manager := range d.managers {
        if manager.CurrentLoad < manager.MaxLoad {
            return manager
        }
    }
    return nil
}

func (d *ApplicationDispatcher) assignToManager(app *Application, manager *Manager) {
    manager.CurrentLoad++
    manager.Applications = append(manager.Applications, app)
    app.Status = StatusInProgress

    d.eventBus.Publish(Event{
        Type:      "ApplicationAssigned",
        Data:      app,
        Timestamp: time.Now(),
    })
}
