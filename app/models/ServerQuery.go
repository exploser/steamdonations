// ServerQuery.go
package models

import (
	"time"

	r "github.com/dancannon/gorethink"
	. "github.com/dobegor/steamdonations/app/util"
	"github.com/revel/revel/cache"
)

type ServerInfo_PlayersOnly struct {
	Timestamp interface{}
	ServerID  byte
	Players   byte
}

type ServerStatic struct {
	ServerID      byte
	ServerName    string
	MaxPlayers    byte
	ReservedSlots byte
	Map           string
	Mode          string
	IP            string
	FancyIP       string
	Graph         string
}

func GetServerStaticInfo(id int) (ServerStatic, error) {
	s := ServerStatic{}

	if err := cache.Get("serverStaticInfo_"+string(id), &s); err == nil {
		return s, nil
	}

	cursor, err := r.DB("db").Table("server_info").
		Filter(r.Row.Field("ServerID").Eq(id)).
		Limit(1).
		Run(DB)

	if err != nil {
		return ServerStatic{}, err
	}

	defer cursor.Close()

	if err := cursor.One(&s); err != nil {
		return ServerStatic{}, err
	}

	go cache.Set("serverStaticInfo_"+string(id), s, 1*time.Hour)
	return s, nil
}

func GetServerGraph(id int) ([]int, error) {
	var result []int
	if err := cache.Get("serverInfoGraph_"+string(id), &result); err == nil {
		return result, nil
	}

	cursor, err := r.DB("db").Table("server_status").
		OrderBy(r.OrderByOpts{Index: r.Desc("Timestamp")}).
		Filter(map[string]int{"ServerID": id}).
		Limit(10).
		OrderBy(r.Asc("Timestamp")).Field("Players").
		Run(DB)

	if err != nil {
		return nil, err
	}

	defer cursor.Close()

	cursor.All(&result)
	go cache.Set("serverInfoGraph_"+string(id), result, 5*time.Minute)
	return result, err
}
