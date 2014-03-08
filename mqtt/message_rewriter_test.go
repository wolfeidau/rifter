package mqtt

import (
	"testing"

	"github.com/huin/mqtt"
	. "launchpad.net/gocheck"
)

func Test2(t *testing.T) { TestingT(t) }

type MessagRewriterSuite struct {
	msgRewriter *MsgRewriter
}

var _ = Suite(&MessagRewriterSuite{})

func (s *MessagRewriterSuite) SetUpTest(c *C) {

}

func (s *MessagRewriterSuite) TestIngressMsgs(c *C) {

}

func (s *MessagRewriterSuite) TestEgressMsgs(c *C) {

}

func createPublish(topic string) mqtt.Message {
	return &mqtt.Publish{
		Header: mqtt.Header{
			DupFlag:  false,
			QosLevel: mqtt.QosAtMostOnce,
			Retain:   false,
		},
		TopicName: topic,
		Payload:   mqtt.BytesPayload{1, 2, 3},
	}
}
