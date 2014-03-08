package mqtt

import "github.com/huin/mqtt"

type MsgRewriter struct {
	InRewriter  TopicRewriter
	OutRewriter TopicRewriter
}

func CreatMsgRewriter(inRewriter TopicRewriter, outRewriter TopicRewriter) *MsgRewriter {
	return &MsgRewriter{
		InRewriter:  inRewriter,
		OutRewriter: outRewriter,
	}
}

func (mr *MsgRewriter) RewriteIngress(msg mqtt.Message) mqtt.Message {
	return msg
}

func (mr *MsgRewriter) RewriteEgress(msg mqtt.Message) mqtt.Message {
	return msg
}
