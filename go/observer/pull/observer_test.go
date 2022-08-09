package pull

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestObserverDataReceiving(t *testing.T) {
	email := "*super urgent email*"

	emailReceiver := NewSingleEmailReceiver()

	msgSender := NewMessageSender(emailReceiver)
	notificationSender := NewNotificationSender(emailReceiver)

	emailReceiver.AddObserver(msgSender)
	emailReceiver.AddObserver(notificationSender)

	emailReceiver.ReceiveEmail(email)
	require.Equal(t, email, *msgSender.Email)
	require.Equal(t, email, *notificationSender.Email)
}
