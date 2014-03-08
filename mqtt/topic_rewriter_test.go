package mqtt

import (
	"testing"

	"github.com/huin/mqtt"
	. "launchpad.net/gocheck"
)

func Test3(t *testing.T) { TestingT(t) }

type TopicRewriterSuite struct {
	topicRewriter TopicRewriter
}

var _ = Suite(&TopicRewriterSuite{})

func (s *TopicRewriterSuite) SetUpTest(c *C) {
	s.topicRewriter = &TopicPrefixRewriter{
		Prefix:      "$block",
		Replacement: "$cloud/123",
	}

}

func (s *TopicRewriterSuite) TestTopicName(c *C) {
	topicName := s.topicRewriter.RewriteTopicName("$block/test/123")
	c.Assert(topicName, Equals, "$cloud/123/test/123")
}

func (s *TopicRewriterSuite) TestTopicNames(c *C) {

	testTopicNames := []string{"$block/test/123", "$block/test/456", "$block/test/789"}
	expectedTopicNames := []string{"$cloud/123/test/123", "$cloud/123/test/456", "$cloud/123/test/789"}

	topicNames := s.topicRewriter.RenameTopicNames(testTopicNames)
	c.Assert(topicNames, DeepEquals, expectedTopicNames)
}

func (s *TopicRewriterSuite) TestTopics(c *C) {
	testTopics := []mqtt.TopicQos{
		mqtt.TopicQos{
			Topic: "$block/test/123",
			Qos:   1,
		},
	}
	expectedTopics := []mqtt.TopicQos{
		mqtt.TopicQos{
			Topic: "$cloud/123/test/123",
			Qos:   1,
		},
	}

	topics := s.topicRewriter.RewriteTopics(testTopics)
	c.Assert(topics, DeepEquals, expectedTopics)

}
