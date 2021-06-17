package message

import (
	"context"
	"errors"
	"sync"

	"github.com/mrinjamul/mrinjamul-admin/firebases"
	"google.golang.org/api/iterator"
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
	list = []Message{}
}

func getFireStoreData() ([]Message, error) {
	ctx := context.Background()
	app, err := firebases.GetFirebaseApp()
	if err != nil {
		return []Message{}, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return []Message{}, err
	}
	iter := client.Collection("github-messages").Documents(ctx)
	messages := []Message{}
	for {
		msg := Message{}
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return []Message{}, err
		}
		doc.DataTo(&msg)
		messages = append(messages, msg)
	}
	defer client.Close()
	return messages, nil
}

// Message data structure
type Message struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
	Read    bool   `json:"read"`
}

// Get retrieves all elements from the Message list
func Get() []Message {
	var err error
	list, err = getFireStoreData()
	if err != nil {
		panic(err)
	}
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
