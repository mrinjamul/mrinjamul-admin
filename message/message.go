package message

import (
	"errors"
	"sync"
)

var (
	list []Message
	mtx  sync.RWMutex
	once sync.Once
)

func init() {
	once.Do(initialiseList)
}

func initialiseList() {
	list = []Message{
		{"1", "Injamul", "mrinjamul@gmail.com", "Hello", "Hi How are you", false},
		{"2", "Inja", "mrinja@gmail.com", "Hello", "Hi, 2", false},
		{"3", "Injam", "mrinjam@gmail.com", "Hello", "Hi, 3", false},
		{"4", "Inj", "mrinj@gmail.com", "Hello", "Hi, 4", false},
	}
}

// Message data structure
type Message struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Read    bool   `json:"read"`
}

// Get retrieves all elements from the Message list
func Get() []Message {
	return list
}

// Delete will remove a Messenge from the Messenge list
func Delete(id string) error {
	location, err := findMessageLocation(id)
	if err != nil {
		return err
	}
	removeElementByLocation(location)
	return nil
}

// MarkAsRead will set the read boolean to true
func MarkAsRead(id string) error {
	location, err := findMessageLocation(id)
	if err != nil {
		return err
	}
	setMessageReadByLocation(location)
	return nil
}

func findMessageLocation(id string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, t := range list {
		if isMatchingID(t.ID, id) {
			return i, nil
		}
	}
	return 0, errors.New("could not find message based on id")
}

func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}

func setMessageReadByLocation(location int) {
	mtx.Lock()
	list[location].Read = true
	mtx.Unlock()
}

func isMatchingID(a string, b string) bool {
	return a == b
}
