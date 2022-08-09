package push

import "fmt"

// Observer is the incoming data consumer
type Observer interface {
	update(value string)
}

// EmailNotifier notifies about the incoming email
type EmailNotifier interface {
	Observer
	Notify() string
}

type MessageSender struct{ email string }
type NotificationSender struct{ email string }

var _ EmailNotifier = (*NotificationSender)(nil)
var _ EmailNotifier = (*MessageSender)(nil)

func (p *NotificationSender) update(email string) {
	p.email = email
}

func (p *NotificationSender) Notify() string {
	return fmt.Sprintf("sending email to phone, %s", p.email)
}

func (m *MessageSender) update(email string) {
	m.email = email
}

// Notify returns a
func (m *MessageSender) Notify() string {
	return fmt.Sprintf("sending email to notification, %s", m.email)
}

// Subject is the main data producer which holds and transfers data to observers
type Subject interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	notifyObservers()
}

// EmailHolder is the data holder for incoming emails
type EmailHolder interface {
	Subject
	ReceiveEmail(email string)
}

// EmailInbox is the data structure for email inbox
type EmailInbox struct {
	readEmails   []string
	unreadEmails []string
	observers    map[Observer]struct{}
}

var _ EmailHolder = (*EmailInbox)(nil)

// NewEmailInbox returns an EmailInbox construct
func NewEmailInbox() *EmailInbox {
	return &EmailInbox{
		readEmails:   make([]string, 0),
		unreadEmails: make([]string, 0),
		observers:    make(map[Observer]struct{}),
	}
}

// AddObserver adds an observer to subscribe
func (e *EmailInbox) AddObserver(observer Observer) {
	e.observers[observer] = struct{}{}
}

// RemoveObserver removes an observer from subscribing
func (e *EmailInbox) RemoveObserver(observer Observer) {
	delete(e.observers, observer)
}

// NotifyObservers notifies all the observers about unread emails
func (e *EmailInbox) notifyObservers() {
	for _, em := range e.unreadEmails {
		for obs := range e.observers {
			obs.update(em)
		}
	}
	e.readEmails = append(e.readEmails, e.unreadEmails...)
	e.unreadEmails = make([]string, 0)
}

func (e *EmailInbox) ReceiveEmail(email string) {
	e.unreadEmails = append(e.unreadEmails, email)
	e.notifyObservers()
}
