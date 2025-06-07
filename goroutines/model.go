//================================= ตัวอย่าง การใช้งานร่วมกับ fiber
// package main

// import (
// 	"sync"
// )

// // Message represents a simple message structure
// type Message struct {
// 	Content string
// }

// // Subscriber represents a subscriber to a topic
// type Subscriber struct {
// 	ID      uint `gorm:"primaryKey"`
// 	Channel chan Message
// }

// // PubSub represents the publish/subscribe mechanism
// type PubSub struct {
// 	mu           sync.RWMutex
// 	Subscribers  map[uint]*Subscriber
// 	subscriberID uint
// }

// func (ps *PubSub) Subscribe() *Subscriber {
// 	ps.mu.Lock()
// 	defer ps.mu.Unlock()

// 	ps.subscriberID++
// 	subscriber := &Subscriber{
// 		ID:      ps.subscriberID,
// 		Channel: make(chan Message),
// 	}
// 	ps.Subscribers[subscriber.ID] = subscriber
// 	return subscriber
// }

// func (ps *PubSub) Unsubscribe(id uint) {
// 	ps.mu.Lock()
// 	defer ps.mu.Unlock()

// 	close(ps.Subscribers[id].Channel)
// 	delete(ps.Subscribers, id)
// }

// func (ps *PubSub) Publish(message Message) {
// 	ps.mu.RLock()
// 	defer ps.mu.RUnlock()

// 	for _, subscriber := range ps.Subscribers {
// 		subscriber.Channel <- message
// 	}
// }

//================================= ตัวอย่าง การใช้งานร่วมกับ gorm

package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

// ExampleModel represents a database model
type ExampleModel struct {
	gorm.Model
	Name string
}

// task is the function to be executed on a schedule
func task(db *gorm.DB) {
	now := time.Now()
	log.Println("Task is being run...", now)

	// Example database operation
	newRecord := ExampleModel{Name: fmt.Sprintf("Record at %s", now.Format(time.RFC3339))}
	db.Create(&newRecord)
}
