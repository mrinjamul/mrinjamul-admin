package message

import (
	"context"
	"errors"
	"strconv"
	"sync"

	"github.com/mrinjamul/mrinjamul-admin/firebases"
	"github.com/rs/xid"
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

// Add will add a new messege based on http post request
func Add(message Message) string {
	msg, err := createMessege(message)
	if err != nil {
		return "0"
	}
	err = sendDataFirestore(msg)
	if err != nil {
		return "0"
	}
	return msg.ID
}

func sendDataFirestore(msg Message) error {
	docName := xid.New().String()
	doc := make(map[string]interface{})
	// doc[docName] = map[string]interface{}{
	// 	"id":      msg.ID,
	// 	"name":    msg.Name,
	// 	"email":   msg.Email,
	// 	"subject": msg.Subject,
	// 	"message": msg.Message,
	// 	"read":    false,
	// }
	doc["id"] = msg.ID
	doc["name"] = msg.Name
	doc["email"] = msg.Email
	doc["subject"] = msg.Subject
	doc["message"] = msg.Message
	doc["read"] = msg.Read
	ctx := context.Background()
	app, err := firebases.GetFirebaseApp()
	if err != nil {
		return err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}
	_, err = client.Collection("github-messages").Doc(docName).Set(ctx, doc)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		return err
	}
	return nil
}

func createMessege(msg Message) (Message, error) {
	var id int
	list, err := getFireStoreData()
	if err != nil {
		return Message{}, err
	}
	for _, m := range list {
		if m.ID == "" {
			m.ID = "0"
		}
		mId, err := strconv.Atoi(m.ID)
		if err != nil {
			return Message{}, err
		}
		if mId > id {
			id = mId
		}
	}

	msg = Message{
		ID:      strconv.Itoa(id + 1),
		Name:    msg.Name,
		Email:   msg.Email,
		Subject: msg.Subject,
		Message: msg.Message,
		Read:    false,
	}
	if msg.Name == "" {
		return Message{}, errors.New("error: name can not be nil")
	}
	if msg.Email == "" {
		return Message{}, errors.New("error: email address can not be nil")
	}
	if msg.Subject == "" {
		return Message{}, errors.New("error: subject can not be nil")
	}
	if msg.Message == "" {
		return Message{}, errors.New("error: messege body can not be nil")
	}

	return msg, nil
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
