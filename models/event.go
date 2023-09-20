package models

type Event struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Price       float64 `json:"price"`
    Description string `json:"description"`
    EventType   string `json:"eventType"`
    Date        string `json:"date"`
}