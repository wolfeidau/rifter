package mqtt

import "github.com/huin/mqtt"

type MsgRewriter struct {
	IngressRewriter TopicRewriter
	EgressRewriter  TopicRewriter
}

func CreatMsgRewriter(ingressRewriter TopicRewriter, egressRewriter TopicRewriter) *MsgRewriter {
	return &MsgRewriter{
		IngressRewriter: ingressRewriter,
		EgressRewriter:  egressRewriter,
	}
}

func (mr *MsgRewriter) RewriteIngress(msg mqtt.Message) mqtt.Message {

	switch msg := msg.(type) {
	case *mqtt.Publish:
		msg.TopicName = mr.IngressRewriter.RewriteTopicName(msg.TopicName)
	case *mqtt.Subscribe:
		msg.Topics = mr.IngressRewriter.RewriteTopics(msg.Topics)
	case *mqtt.Unsubscribe:
		msg.Topics = mr.IngressRewriter.RenameTopicNames(msg.Topics)
	}
	return msg
}

func (mr *MsgRewriter) RewriteEgress(msg mqtt.Message) mqtt.Message {
	switch msg := msg.(type) {
	case *mqtt.Publish:
		msg.TopicName = mr.EgressRewriter.RewriteTopicName(msg.TopicName)
	}
	return msg
}
