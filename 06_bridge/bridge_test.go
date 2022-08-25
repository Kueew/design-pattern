package bridge

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTelephoneSender_Notify(t *testing.T) {
	telephoneSender := NewTelephoneMsgSender([]string{"18000000000"})
	n := NewNormalNotification(telephoneSender)
	err := n.Notify("msg")
	require.Nil(t, err)
}
func TestEamilSender_Notify(t *testing.T) {
	eamilSender := NewEmailMsgSender([]string{"test@qq.com"})
	n := NewNormalNotification(eamilSender)
	err := n.Notify("msg")
	require.Nil(t, err)
}
