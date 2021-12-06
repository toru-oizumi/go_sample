package handler

import (
	"go_sample/app/domain/model"

	"sync"

	"github.com/gorilla/websocket"
)

// connectionは１ユーザに1個、ログイン時に作成、ログアウト時に削除
// なのでユーザーIDをキーにconnectionを持つ
var connectionPool = ConnectionPool{
	connections: make(map[model.UserID]*websocket.Conn),
}

type ConnectionPool struct {
	sync.RWMutex
	connections map[model.UserID]*websocket.Conn
}

func (c *ConnectionPool) FilterConnectionsByUserIDs(ids []model.UserID) []*websocket.Conn {
	c.RLock()
	defer c.RUnlock()

	var connections []*websocket.Conn
	for _, userID := range ids {
		if v, ok := c.connections[userID]; ok {
			connections = append(connections, v)
		}
	}
	return connections
}

func (c *ConnectionPool) AddConnection(id model.UserID, con *websocket.Conn) {
	c.Lock()
	defer c.Unlock()

	// TODO: 既にconnectionがある場合は、無視すべきか、エラーにすべきか、上書きすべきか
	if _, ok := c.connections[id]; !ok {
		c.connections[id] = con
	}
}

func (c *ConnectionPool) RemoveConnection(id model.UserID) {
	c.Lock()
	defer c.Unlock()

	delete(c.connections, id)
}

func sendMessageToConnections(connections []*websocket.Conn, v []byte) error {
	for _, connection := range connections {
		if err := connection.WriteMessage(websocket.TextMessage, v); err != nil {
			return err
		}
	}
	return nil
}
