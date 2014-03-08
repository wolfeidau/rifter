package mqtt

import (
	"strings"

	"github.com/huin/mqtt"
)

type TopicRewriter interface {
	RewriteTopicName(topic string) string
	RenameTopicNames(topicNames []string) []string
	RewriteTopics(topics []mqtt.TopicQos) []mqtt.TopicQos
}

type TopicPrefixRewriter struct {
	Prefix      string
	Replacement string
}

func NewTopicPrefixRewriter(prefix string, replacement string) *TopicPrefixRewriter {
	return &TopicPrefixRewriter{
		Prefix:      prefix,
		Replacement: replacement,
	}
}

func (tpw *TopicPrefixRewriter) RewriteTopicName(topic string) string {
	return strings.Replace(topic, tpw.Prefix, tpw.Replacement, 1)
}

func (tpw *TopicPrefixRewriter) RenameTopicNames(topicNames []string) []string {
	for i, _ := range topicNames {
		topicNames[i] = strings.Replace(topicNames[i], tpw.Prefix, tpw.Replacement, 1)
	}
	return topicNames
}

func (tpw *TopicPrefixRewriter) RewriteTopics(topics []mqtt.TopicQos) []mqtt.TopicQos {
	for i, _ := range topics {
		topics[i].Topic = strings.Replace(topics[i].Topic, tpw.Prefix, tpw.Replacement, 1)
	}
	return topics

}
