package handler

import (
	"go_sample/app/domain/model"
	enum_connection "go_sample/app/interface/controller/ws/enum/connection"

	utility_error "go_sample/app/utility/error"
	"sync"

	"github.com/gorilla/websocket"
)

// connectionは１ユーザに1個、ログイン時に作成、ログアウト時に削除
// なのでユーザーIDをキーにconnectionのMAPを持つ
var connectionPool = ConnectionPool{
	connections: make(map[model.UserId]Connection),
}

type ConnectionPool struct {
	sync.RWMutex
	connections map[model.UserId]Connection
}

func (c *ConnectionPool) FilterConnectionsByObjective(o enum_connection.Objective, id string) []*websocket.Conn {
	c.RLock()
	defer c.RUnlock()

	var connections []*websocket.Conn
	for _, con := range c.connections {
		if con.objectives[o] == id {
			connections = append(connections, con.wsConnection)
		}
	}
	return connections
}

func (c *ConnectionPool) AddConnection(
	u model.UserId,
	o enum_connection.Objective,
	id string,
	con *websocket.Conn) error {
	c.Lock()
	defer c.Unlock()

	if v, ok := c.connections[u]; ok {
		return v.addObjective(o, id)
	} else {
		connection := Connection{
			wsConnection: con,
			objectives:   make(map[enum_connection.Objective]string),
		}
		connection.addObjective(o, id)
		c.connections[u] = connection
		return nil
	}
}

func (c *ConnectionPool) RemoveConnection(
	u model.UserId,
	o enum_connection.Objective,
) {
	c.Lock()
	defer c.Unlock()

	if v, ok := c.connections[u]; ok {
		v.removeObjective(o)

		if len(v.objectives) == 0 {
			delete(c.connections, u)
		}
	}
}

type Connection struct {
	// objectivesは
	// key: ConnectionObjective,
	// value:id(各IdにParseして使用するstring型)
	wsConnection *websocket.Conn
	objectives   map[enum_connection.Objective]string
}

func (c *Connection) addObjective(o enum_connection.Objective, id string) error {
	if _, ok := c.objectives[o]; ok {
		return utility_error.NewErrWsConnectionAlreadyExist(string(o), id)
	} else {
		c.objectives[o] = id
		return nil
	}
}

func (c *Connection) removeObjective(o enum_connection.Objective) {
	delete(c.objectives, o)
}

func sendMessageToConnections(connections []*websocket.Conn, v []byte) error {
	connectionPool.RLock()
	defer connectionPool.RUnlock()
	for _, connection := range connections {
		if err := connection.WriteMessage(websocket.TextMessage, v); err != nil {
			return err
		}
	}
	return nil
}
