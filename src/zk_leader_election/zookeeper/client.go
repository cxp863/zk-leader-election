package zookeeper

import (
	paths "path"
	"strings"
	"time"

	"github.com/go-zookeeper/zk"
)

const (
	AnyVersion   int32  = -1
	GlobalPrefix string = "/election_demo"
)

type Config struct {
	Nodes   []string
	Timeout time.Duration
}

type Client interface {
	Create(relativePath string, data []byte, ephemeral, sequence bool) (string, error)
	Read(relativePath string) ([]byte, *zk.Stat, error)
	Update(relativePath string, data []byte, version int32) (*zk.Stat, error)
	Delete(relativePath string, version int32) error

	Exists(relativePath string) (bool, *zk.Stat, error)
	Children(relativePath string) ([]string, *zk.Stat, error)
	WithPrefix(relativePath string) Client

	Close()
}

type client struct {
	conn   *zk.Conn
	prefix string
}

func NewClient(config *Config) (Client, error) {
	if conn, _, err := zk.Connect(config.Nodes, config.Timeout, zk.WithLogInfo(false)); err != nil {
		return nil, err
	} else {
		return &client{conn, GlobalPrefix}, nil
	}
}

func (c *client) Create(relativePath string, data []byte, ephemeral, sequence bool) (string, error) {
	var flags int32
	if ephemeral { // 临时节点，conn.Close之后节点就无了
		flags += zk.FlagEphemeral
	}
	if sequence { // 序列节点
		flags += zk.FlagSequence
	}

	absolutePath := paths.Join(c.prefix, relativePath)
	if !strings.HasSuffix(absolutePath, "/") {
		absolutePath += "/"
	}

	return c.conn.Create(absolutePath, data, flags, zk.WorldACL(zk.PermAll))
}

func (c *client) Read(relativePath string) ([]byte, *zk.Stat, error) {
	absolutePath := paths.Join(c.prefix, relativePath)
	return c.conn.Get(absolutePath)
}

func (c *client) Update(relativePath string, data []byte, version int32) (*zk.Stat, error) {
	absolutePath := paths.Join(c.prefix, relativePath)
	return c.conn.Set(absolutePath, data, version)
}

func (c *client) Delete(relativePath string, version int32) error {
	absolutePath := paths.Join(c.prefix, relativePath)
	return c.conn.Delete(absolutePath, version)
}

func (c *client) Exists(relativePath string) (bool, *zk.Stat, error) {
	absolutePath := paths.Join(c.prefix, relativePath)
	return c.conn.Exists(absolutePath)
}

func (c *client) Children(relativePath string) ([]string, *zk.Stat, error) {
	absolutePath := paths.Join(c.prefix, relativePath)
	return c.conn.Children(absolutePath)
}

func (c *client) WithPrefix(prefix string) Client {
	return &client{c.conn, paths.Join(c.prefix, prefix)}
}

func (c *client) Close() {
	c.conn.Close()
}
