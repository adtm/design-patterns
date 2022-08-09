package pull

// Observer is the data consumer
type Observer interface {
	update()
}

// MessageSender is the struct which holds required information for sending a message
type MessageSender struct {
	SingleEmailReceiver
	// Email is public only for demonstration and testing purposes here
	Email *string
}

// NewMessageSender creates a new message sender
func NewMessageSender(emailReceiver SingleEmailReceiver) *MessageSender {
	return &MessageSender{
		Email:               nil,
		SingleEmailReceiver: emailReceiver,
	}
}

// NotificationSender is the struct which holds required information for sending a notification
type NotificationSender struct {
	SingleEmailReceiver
	// Email is public only for demonstration and testing purposes here
	Email *string
}

// NewNotificationSender creates a new not sender
func NewNotificationSender(emailReceiver SingleEmailReceiver) *NotificationSender {
	return &NotificationSender{
		Email:               nil,
		SingleEmailReceiver: emailReceiver,
	}
}

var _ Observer = (*MessageSender)(nil)
var _ Observer = (*NotificationSender)(nil)

func (m *MessageSender) update() {
	m.Email = m.SingleEmailReceiver.GetEmail()
}

func (n *NotificationSender) update() {
	n.Email = n.SingleEmailReceiver.GetEmail()
}

// Subject is the construct holding and transferring data to observers
type Subject interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	// notifies all the observers about incoming data changes
	notifyObservers()
}

// EmailInbox is the receiver for new emails
type EmailInbox struct {
	Email     *string
	observers map[Observer]struct{}
}

// SingleEmailReceiver is the interface for receiving emails
type SingleEmailReceiver interface {
	Subject
	ReceiveEmail(email string)
	GetEmail() *string
}

// NewSingleEmailReceiver is a new email construct
func NewSingleEmailReceiver() SingleEmailReceiver {
	return &EmailInbox{
		// we assume that our inbox only can store only one email
		Email:     nil,
		observers: make(map[Observer]struct{}),
	}
}

var _ SingleEmailReceiver = (*EmailInbox)(nil)

// GetEmail returns the stored email
func (e *EmailInbox) GetEmail() *string {
	return e.Email
}

// ReceiveEmail receives a new email
func (e *EmailInbox) ReceiveEmail(email string) {
	e.Email = &email
	e.notifyObservers()
}

func (e *EmailInbox) notifyObservers() {
	for ob := range e.observers {
		ob.update()
	}
}

// AddObserver adds a new observer consumer
func (e *EmailInbox) AddObserver(observer Observer) {
	e.observers[observer] = struct{}{}
}

// RemoveObserver removes a observer consumer
func (e *EmailInbox) RemoveObserver(observer Observer) {
	delete(e.observers, observer)
}
