package main

import "fmt"

type NotificationService struct {
    eventBus *EventBus
}

func NewNotificationService(eventBus *EventBus) *NotificationService {
    ns := &NotificationService{eventBus: eventBus}
    
    eventChan := make(chan Event)
    eventBus.Subscribe("ApplicationBuffered", eventChan)
    eventBus.Subscribe("ApplicationRemoved", eventChan)
    eventBus.Subscribe("ApplicationProcessed", eventChan)
    
    go ns.handleEvents(eventChan)
    return ns
}

func (ns *NotificationService) handleEvents(eventChan chan Event) {
    for event := range eventChan {
        switch event.Type {
        case "ApplicationBuffered":
            ns.notifyClientAboutQueue(event.Data.(*Application))
        case "ApplicationRemoved":
            ns.notifyClientAboutRemoval(event.Data.(*Application))
        case "ApplicationProcessed":
            ns.notifyClientAboutResult(event.Data.(*Application))
        }
    }
}

func (ns *NotificationService) notifyClientAboutQueue(app *Application) {
    fmt.Printf("     Notification to Client %d: Your event planning request is in queue\n", app.ClientID)
}

func (ns *NotificationService) notifyClientAboutRemoval(app *Application) {
    fmt.Printf("     Notification to Client %d: Your event planning request was removed from queue\n", app.ClientID)
}

func (ns *NotificationService) notifyClientAboutResult(app *Application) {
    fmt.Printf("     Notification to Client %d: Your event planning request status is %s\n", app.ClientID, app.Status)
}
