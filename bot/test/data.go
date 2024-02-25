package test

import (
	"github.com/slainsama/msgr_server/bot/types"
)

func newStartUpdate() types.TelegramUpdate {
	startUpdate := types.TelegramUpdate{}

	startUpdate.Message = &types.Message{}
	startUpdate.Message.Text = "/startUpload"
	startUpdate.Message.Chat = &types.Chat{}
	startUpdate.Message.Chat.ID = 123
	startUpdate.Message.From = &types.User{}
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
func newStartWithoutPermissionUpdate() types.TelegramUpdate {
	startUpdate := types.TelegramUpdate{}

	startUpdate.Message = &types.Message{}
	startUpdate.Message.Text = "/startUpload"
	startUpdate.Message.Chat = &types.Chat{}
	startUpdate.Message.Chat.ID = 123
	startUpdate.Message.From = &types.User{}
	startUpdate.Message.From.ID = 426

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

func newHelloUpdate() types.TelegramUpdate {
	helloUpdate := types.TelegramUpdate{}
	helloUpdate.Message = &types.Message{}
	helloUpdate.Message.Text = "/hello"
	helloUpdate.Message.Chat = &types.Chat{}
	helloUpdate.Message.Chat.ID = 123
	helloUpdate.Message.From = &types.User{}
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

func newHelloWithoutPermissionUpdate() types.TelegramUpdate {
	helloUpdate := types.TelegramUpdate{}
	helloUpdate.Message = &types.Message{}
	helloUpdate.Message.Text = "/hello"
	helloUpdate.Message.Chat = &types.Chat{}
	helloUpdate.Message.Chat.ID = 123
	helloUpdate.Message.From = &types.User{}
	helloUpdate.Message.From.ID = 426
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

func newMultiCommandHelloUpdate() types.TelegramUpdate {
	multiCommandHelloUpdate := types.TelegramUpdate{}
	multiCommandHelloUpdate.Message = &types.Message{}
	multiCommandHelloUpdate.Message.Text = "/hello 1 2 3 /another_hello 1 2 3 4 5 6"
	multiCommandHelloUpdate.Message.Chat = &types.Chat{}
	multiCommandHelloUpdate.Message.Chat.ID = 123
	multiCommandHelloUpdate.Message.From = &types.User{}
	multiCommandHelloUpdate.Message.From.ID = 123
	multiCommandHelloUpdate.Message.Entities = append(multiCommandHelloUpdate.Message.Entities, struct {
		Offset int    "json:\"offset\""
		Length int    "json:\"length\""
		Type   string "json:\"type\""
	}{
		Offset: 0,
		Length: 6,
		Type:   "bot_command",
	})
	multiCommandHelloUpdate.Message.Entities = append(multiCommandHelloUpdate.Message.Entities, struct {
		Offset int    "json:\"offset\""
		Length int    "json:\"length\""
		Type   string "json:\"type\""
	}{
		Offset: 13,
		Length: 14,
		Type:   "bot_command",
	})
	return multiCommandHelloUpdate
}

func newAnotherHelloUpdate() types.TelegramUpdate {
	anotherHelloUpdate := types.TelegramUpdate{}
	anotherHelloUpdate.Message = &types.Message{}
	anotherHelloUpdate.Message.Text = "/another_hello"
	anotherHelloUpdate.Message.Chat = &types.Chat{}
	anotherHelloUpdate.Message.Chat.ID = 123
	anotherHelloUpdate.Message.From = &types.User{}
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

func newEndUpdate() types.TelegramUpdate {
	endUpdate := types.TelegramUpdate{}
	endUpdate.Message = &types.Message{}
	endUpdate.Message.Text = "/endUpload"
	endUpdate.Message.Chat = &types.Chat{}
	endUpdate.Message.Chat.ID = 123
	endUpdate.Message.From = &types.User{}
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

func newEndWithoutPermissionUpdate() types.TelegramUpdate {
	endUpdate := types.TelegramUpdate{}
	endUpdate.Message = &types.Message{}
	endUpdate.Message.Text = "/endUpload"
	endUpdate.Message.Chat = &types.Chat{}
	endUpdate.Message.Chat.ID = 123
	endUpdate.Message.From = &types.User{}
	endUpdate.Message.From.ID = 426
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
