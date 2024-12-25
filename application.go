package main

import "time"

const (
    StatusNew = "New"
    StatusInQueue = "InQueue"
    StatusInProgress = "InProgress"
    StatusNeedsInfo = "NeedsInfo"
    StatusApproved = "Approved"
    StatusRejected = "Rejected"
)

type EventType string

const (
    Wedding EventType = "Wedding"
    Corporate EventType = "Corporate"
    Birthday EventType = "Birthday"
)

type Application struct {
    ID            int
    ClientID      int
    EventType     EventType
    Status        string
    Priority      int
    Requirements  *Requirements
    CreatedAt     time.Time
}

type Requirements struct {
    ID              int
    GuestCount      int
    Budget          float64
    PreferredDate   time.Time
    Location        string
    SpecialRequests string
    Status          string
}
