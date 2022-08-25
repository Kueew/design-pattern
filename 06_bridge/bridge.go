package bridge

type IMsgSender interface {
	Send(message string) error
}

type TelephoneMsgSender struct {
	telephones []string
}

func (sender *TelephoneMsgSender) Send(message string) error {
	//todo:发送短信逻辑
	return nil
}

func NewTelephoneMsgSender(telephones []string) *TelephoneMsgSender {
	return &TelephoneMsgSender{telephones: telephones}
}

type EmailMsgSender struct {
	eamils []string
}

func (sender *EmailMsgSender) Send(message string) error {
	//todo:发送邮件逻辑
	return nil
}

func NewEmailMsgSender(eamils []string) *EmailMsgSender {
	return &EmailMsgSender{eamils: eamils}
}

type INotification interface {
	Notify(message string) error
}

type NormalNotification struct {
	sender IMsgSender
}

func (n *NormalNotification) Notify(message string) error {
	return n.sender.Send(message)
}

func NewNormalNotification(sender IMsgSender) *NormalNotification {
	return &NormalNotification{sender: sender}
}
