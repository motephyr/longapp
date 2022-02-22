package modelhelper

import (
	"log"
	"strconv"

	self "github.com/motephyr/longapp/app"
	"github.com/motephyr/longapp/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type older struct{}

var Older older

func (older) AfterUpdateHook(exec boil.Executor, p *models.Older) error {
	// Do stuff
	log.Println("AfterUpdateHook")
	userOlders, _ := models.UserOlders(Where("older_id = ?", p.ID)).OneG()
	room := "/user/" + strconv.Itoa(userOlders.UserID.Int) + "/older"

	self.Http.Websocket.Server.BroadcastToRoom("", room, "update", p.ID)

	return nil
}

type userOlder struct{}

var UserOlder userOlder

func (userOlder) AfterInsertHook(exec boil.Executor, p *models.UserOlder) error {
	// Do stuff
	log.Println("AfterInsertHook")

	userOlders, _ := models.UserOlders(Where("id = ?", p.ID)).OneG()
	room := "/user/" + strconv.Itoa(userOlders.UserID.Int) + "/older"
	self.Http.Websocket.Server.BroadcastToRoom("", room, "insert", userOlders.OlderID.Int)

	return nil
}

func (userOlder) AfterDeleteHook(exec boil.Executor, p *models.UserOlder) error {
	// Do stuff
	log.Println("AfterDeleteHook")

	room := "/user/" + strconv.Itoa(p.UserID.Int) + "/older"

	self.Http.Websocket.Server.BroadcastToRoom("", room, "delete", p.OlderID.Int)

	return nil
}
