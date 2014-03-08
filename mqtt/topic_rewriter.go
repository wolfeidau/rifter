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

// rewriter which checks for the prefix and replaces just that.
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
	if strings.HasPrefix(topic, tpw.Prefix) {
		return strings.Replace(topic, tpw.Prefix, tpw.Replacement, 1)
	} else {
		return topic
	}
}

func (tpw *TopicPrefixRewriter) RenameTopicNames(topicNames []string) []string {
	for i, _ := range topicNames {
		topicNames[i] = tpw.RewriteTopicName(topicNames[i])
	}
	return topicNames
}

func (tpw *TopicPrefixRewriter) RewriteTopics(topics []mqtt.TopicQos) []mqtt.TopicQos {
	for i, _ := range topics {
		topics[i].Topic = tpw.RewriteTopicName(topics[i].Topic)
	}
	return topics
}
