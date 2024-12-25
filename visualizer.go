package main

import (
    "fmt"
)

type SystemVisualizer struct {
    system *System
}

func NewSystemVisualizer(system *System) *SystemVisualizer {
    return &SystemVisualizer{system: system}
}

func (v *SystemVisualizer) PrintSystemState() {
    fmt.Println("\n=== Event Planning System State ===")
    v.printBuffer()
    v.printManagers()
    v.printStatistics()
    fmt.Println("================================")
}

func (v *SystemVisualizer) printBuffer() {
    fmt.Print("\nRequest Queue: ")
    fmt.Printf("%d/%d\n", len(v.system.buffer.applications), v.system.buffer.maxSize)
    if len(v.system.buffer.applications) > 0 {
        for i, app := range v.system.buffer.applications {
            fmt.Printf("[%d] Request ID: %d, Type: %s, Priority: %d, Status: %s\n",
                i+1, app.ID, app.EventType, app.Priority, app.Status)
        }
    }
}

func (v *SystemVisualizer) printManagers() {
    fmt.Println("\nEvent Managers State:")
    for _, manager := range v.system.dispatcher.managers {
        fmt.Printf(" - Manager %d: Load %d/%d, Specializations: %v\n",
            manager.ID,
            manager.CurrentLoad,
            manager.MaxLoad,
            manager.Specializations)

        if len(manager.Applications) > 0 {
            fmt.Println("   Current requests:")
            for _, app := range manager.Applications {
                fmt.Printf("     - ID: %d, Type: %s, Status: %s\n",
                    app.ID, app.EventType, app.Status)
            }
        }
    }
}

func (v *SystemVisualizer) printStatistics() {
    fmt.Println("\nCurrent Statistics:")
    v.system.statistics.PrintCurrentStats()
}

