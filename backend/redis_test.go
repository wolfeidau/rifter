package backend

import (
	"testing"
	"time"

	"github.com/mailgun/vulcan/timeutils"
	. "launchpad.net/gocheck"
)

func Test2(t *testing.T) { TestingT(t) }

type RedisBackendSuite struct {
	timeProvider *timeutils.FreezedTime
	backend      *RedisBackend
}

var _ = Suite(&RedisBackendSuite{})

func (s *RedisBackendSuite) SetUpTest(c *C) {
	start := time.Date(2012, 3, 4, 5, 6, 7, 0, time.UTC)
	s.timeProvider = &timeutils.FreezedTime{CurrentTime: start}
	config := &RedisConfig{
		Prefix: "rifter",
		Server: "localhost:6379",
	}
	backend, err := NewRedisBackend(config, s.timeProvider)
	c.Assert(err, IsNil)
	c.Assert(err, IsNil)
	s.backend = backend
}

func (s *RedisBackendSuite) TestUtcNow(c *C) {
	c.Assert(s.backend.UtcNow(), Equals, s.timeProvider.CurrentTime)
}

func (s *RedisBackendSuite) TestRedisBackendGetSet(c *C) {
	_, err := s.backend.DeleteCount("key1", time.Second)
	c.Assert(err, IsNil)
	counter, err := s.backend.GetCount("key1", time.Second)
	c.Assert(err, IsNil)
	c.Assert(counter, Equals, int64(0))

	err = s.backend.UpdateCount("key1", time.Second, 2)
	c.Assert(err, IsNil)

	counter, err = s.backend.GetCount("key1", time.Second)
	c.Assert(err, IsNil)
	c.Assert(counter, Equals, int64(2))
}
