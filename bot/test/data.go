package test

import "github.com/slainsama/msgr_server/models"

func newStartUpdate() models.TelegramUpdate {
	startUpdate := models.TelegramUpdate{}
	startUpdate.Message.Text = "/startUpload"
	startUpdate.Message.Chat.ID = 123
	startUpdate.Message.From.ID = 123
	startUpdate.Message.Entities = append(startUpdate.Message.Entities, struct {
		Offset int    "json:\"offset\""
		Length int    "json:\"length\""
		Type   string "json:\"type\""
	}{
		Offset: 0,
		Length: 12,
		Type:   "bot_command",
	})
	return startUpdate
}

func newHelloUpdate() models.TelegramUpdate {
	helloUpdate := models.TelegramUpdate{}
	helloUpdate.Message.Text = "/hello"
	helloUpdate.Message.Chat.ID = 123
	helloUpdate.Message.From.ID = 123
	helloUpdate.Message.Entities = append(helloUpdate.Message.Entities, struct {
		Offset int    "json:\"offset\""
		Length int    "json:\"length\""
		Type   string "json:\"type\""
	}{
		Offset: 0,
		Length: 6,
		Type:   "bot_command",
	})
	return helloUpdate
}

func newAnotherHelloUpdate() models.TelegramUpdate {
	anotherHelloUpdate := models.TelegramUpdate{}
	anotherHelloUpdate.Message.Text = "/another_hello"
	anotherHelloUpdate.Message.Chat.ID = 123
	anotherHelloUpdate.Message.From.ID = 123
	anotherHelloUpdate.Message.Entities = append(anotherHelloUpdate.Message.Entities, struct {
		Offset int    "json:\"offset\""
		Length int    "json:\"length\""
		Type   string "json:\"type\""
	}{
		Offset: 0,
		Length: 14,
		Type:   "bot_command",
	})
	return anotherHelloUpdate
}

func newEndUpdate() models.TelegramUpdate {
	endUpdate := models.TelegramUpdate{}
	endUpdate.Message.Text = "/endUpload"
	endUpdate.Message.Chat.ID = 123
	endUpdate.Message.From.ID = 123
	endUpdate.Message.Entities = append(endUpdate.Message.Entities, struct {
		Offset int    "json:\"offset\""
		Length int    "json:\"length\""
		Type   string "json:\"type\""
	}{
		Offset: 0,
		Length: 10,
		Type:   "bot_command",
	})
	return endUpdate
}
