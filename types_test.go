package tgbotapi_test

import (
	"testing"
	"time"

	"github.com/arashrahimi46/bale-bot-go-api"
)

func TestUserStringWith(t *testing.T) {
	user := tgbotapi.User{
		ID:           0,
		FirstName:    "Test",
		LastName:     "Test",
		UserName:     "",
		LanguageCode: "en",
		IsBot:        false,
	}

	if user.String() != "Test Test" {
		t.Fail()
	}
}

func TestUserStringWithUserName(t *testing.T) {
	user := tgbotapi.User{
		ID:           0,
		FirstName:    "Test",
		LastName:     "Test",
		UserName:     "@test",
		LanguageCode: "en",
	}

	if user.String() != "@test" {
		t.Fail()
	}
}

func TestMessageTime(t *testing.T) {
	message := tgbotapi.Message{Date: 0}

	date := time.Unix(0, 0)
	if message.Time() != date {
		t.Fail()
	}
}

func TestMessageIsCommandWithCommand(t *testing.T) {
	message := tgbotapi.Message{Text: "/command"}
	message.Entities = &[]tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: 8}}

	if message.IsCommand() != true {
		t.Fail()
	}
}

func TestIsCommandWithText(t *testing.T) {
	message := tgbotapi.Message{Text: "some text"}

	if message.IsCommand() != false {
		t.Fail()
	}
}

func TestIsCommandWithEmptyText(t *testing.T) {
	message := tgbotapi.Message{Text: ""}

	if message.IsCommand() != false {
		t.Fail()
	}
}

func TestCommandWithCommand(t *testing.T) {
	message := tgbotapi.Message{Text: "/command"}
	message.Entities = &[]tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: 8}}

	if message.Command() != "command" {
		t.Fail()
	}
}

func TestCommandWithEmptyText(t *testing.T) {
	message := tgbotapi.Message{Text: ""}

	if message.Command() != "" {
		t.Fail()
	}
}

func TestCommandWithNonCommand(t *testing.T) {
	message := tgbotapi.Message{Text: "test text"}

	if message.Command() != "" {
		t.Fail()
	}
}

func TestCommandWithBotName(t *testing.T) {
	message := tgbotapi.Message{Text: "/command@testbot"}
	message.Entities = &[]tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: 16}}

	if message.Command() != "command" {
		t.Fail()
	}
}

func TestCommandWithAtWithBotName(t *testing.T) {
	message := tgbotapi.Message{Text: "/command@testbot"}
	message.Entities = &[]tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: 16}}

	if message.CommandWithAt() != "command@testbot" {
		t.Fail()
	}
}

func TestMessageCommandArgumentsWithArguments(t *testing.T) {
	message := tgbotapi.Message{Text: "/command with arguments"}
	message.Entities = &[]tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: 8}}
	if message.CommandArguments() != "with arguments" {
		t.Fail()
	}
}

func TestMessageCommandArgumentsWithMalformedArguments(t *testing.T) {
	message := tgbotapi.Message{Text: "/command-without argument space"}
	message.Entities = &[]tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: 8}}
	if message.CommandArguments() != "without argument space" {
		t.Fail()
	}
}

func TestMessageCommandArgumentsWithoutArguments(t *testing.T) {
	message := tgbotapi.Message{Text: "/command"}
	if message.CommandArguments() != "" {
		t.Fail()
	}
}

func TestMessageCommandArgumentsForNonCommand(t *testing.T) {
	message := tgbotapi.Message{Text: "test text"}
	if message.CommandArguments() != "" {
		t.Fail()
	}
}

func TestMessageEntityParseURLGood(t *testing.T) {
	entity := tgbotapi.MessageEntity{URL: "https://www.google.com"}

	if _, err := entity.ParseURL(); err != nil {
		t.Fail()
	}
}

func TestMessageEntityParseURLBad(t *testing.T) {
	entity := tgbotapi.MessageEntity{URL: ""}

	if _, err := entity.ParseURL(); err == nil {
		t.Fail()
	}
}

func TestChatIsPrivate(t *testing.T) {
	chat := tgbotapi.Chat{ID: 10, Type: "private"}

	if chat.IsPrivate() != true {
		t.Fail()
	}
}

func TestChatIsGroup(t *testing.T) {
	chat := tgbotapi.Chat{ID: 10, Type: "group"}

	if chat.IsGroup() != true {
		t.Fail()
	}
}

func TestChatIsChannel(t *testing.T) {
	chat := tgbotapi.Chat{ID: 10, Type: "channel"}

	if chat.IsChannel() != true {
		t.Fail()
	}
}

func TestChatIsSuperGroup(t *testing.T) {
	chat := tgbotapi.Chat{ID: 10, Type: "supergroup"}

	if !chat.IsSuperGroup() {
		t.Fail()
	}
}

func TestMessageEntityIsMention(t *testing.T) {
	entity := tgbotapi.MessageEntity{Type: "mention"}

	if !entity.IsMention() {
		t.Fail()
	}
}

func TestMessageEntityIsHashtag(t *testing.T) {
	entity := tgbotapi.MessageEntity{Type: "hashtag"}

	if !entity.IsHashtag() {
		t.Fail()
	}
}

func TestMessageEntityIsBotCommand(t *testing.T) {
	entity := tgbotapi.MessageEntity{Type: "bot_command"}

	if !entity.IsCommand() {
		t.Fail()
	}
}

func TestMessageEntityIsUrl(t *testing.T) {
	entity := tgbotapi.MessageEntity{Type: "url"}

	if !entity.IsUrl() {
		t.Fail()
	}
}

func TestMessageEntityIsEmail(t *testing.T) {
	entity := tgbotapi.MessageEntity{Type: "email"}

	if !entity.IsEmail() {
		t.Fail()
	}
}

func TestMessageEntityIsBold(t *testing.T) {
	entity := tgbotapi.MessageEntity{Type: "bold"}

	if !entity.IsBold() {
		t.Fail()
	}
}

func TestMessageEntityIsItalic(t *testing.T) {
	entity := tgbotapi.MessageEntity{Type: "italic"}

	if !entity.IsItalic() {
		t.Fail()
	}
}

func TestMessageEntityIsCode(t *testing.T) {
	entity := tgbotapi.MessageEntity{Type: "code"}

	if !entity.IsCode() {
		t.Fail()
	}
}

func TestMessageEntityIsPre(t *testing.T) {
	entity := tgbotapi.MessageEntity{Type: "pre"}

	if !entity.IsPre() {
		t.Fail()
	}
}

func TestMessageEntityIsTextLink(t *testing.T) {
	entity := tgbotapi.MessageEntity{Type: "text_link"}

	if !entity.IsTextLink() {
		t.Fail()
	}
}

func TestFileLink(t *testing.T) {
	file := tgbotapi.File{FilePath: "test/test.txt"}

	if file.Link("token") != "https://bot.bale.ai/file/bottoken/test/test.txt" {
		t.Fail()
	}
}
