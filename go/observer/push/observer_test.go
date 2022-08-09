package push

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestObserverPassedValues(t *testing.T) {
	sentEmail := "*some testing email*"

	phoneSender := NotificationSender{}
	messageSender := MessageSender{}

	inbox := NewEmailInbox()
	inbox.AddObserver(&phoneSender)
	inbox.AddObserver(&messageSender)

	inbox.ReceiveEmail(sentEmail)

	require.Equal(t, fmt.Sprintf("sending email to phone, %s", sentEmail), phoneSender.Notify())
	require.Equal(t, fmt.Sprintf("sending email to sms, %s", sentEmail), messageSender.Notify())
}
