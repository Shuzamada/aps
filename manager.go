
package main

import (
    "math/rand"
)

type Manager struct {
    ID              int
    CurrentLoad     int
    MaxLoad         int
    Applications    []*Application
    Specializations []EventType
}

func NewManager(id int, maxLoad int) *Manager {
    specializations := []EventType{Wedding, Corporate, Birthday}
    rand.Shuffle(len(specializations), func(i, j int) {
        specializations[i], specializations[j] = specializations[j], specializations[i]
    })

    return &Manager{
        ID:              id,
        MaxLoad:         maxLoad,
        Applications:    make([]*Application, 0),
        Specializations: specializations[:1+rand.Intn(2)], // 1-2 specializations
    }
}

func (m *Manager) calculateSuitability(app *Application) int {
    score := 0
    
    for _, spec := range m.Specializations {
        if spec == app.EventType {
            score += 5
        }
    }
    
    workloadScore := 10 - (m.CurrentLoad * 2)
    if workloadScore > 0 {
        score += workloadScore
    }
    
    return score
}

func (m *Manager) completeRandomApplication() *Application {
    if len(m.Applications) == 0 {
        return nil
    }

    randomIndex := rand.Intn(len(m.Applications))
    app := m.Applications[randomIndex]

    m.processApplication(app)

    m.Applications = append(m.Applications[:randomIndex], m.Applications[randomIndex+1:]...)
    m.CurrentLoad--

    return app
}

func (m *Manager) processApplication(app *Application) {
    if app.Requirements.Status == "" {
        if rand.Float32() < 0.2 {
            app.Status = StatusNeedsInfo
            return
        }
    }
    
    if rand.Float32() < 0.8 {
        app.Status = StatusApproved
    } else {
        app.Status = StatusRejected
    }
}
