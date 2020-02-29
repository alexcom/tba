package bot

import (
	"github.com/alexcom/tba/telegram"
	"github.com/sirupsen/logrus"
	"time"
)

type AuthZFunction func(update *telegram.Update) (allowed bool, chatID int, message string)

type Options struct {
	APIToken           string
	LongPollingTimeout int
	FailRetryInterval  int
	Authorized         AuthZFunction
}

// Constructor for Bot
func NewBot(opts Options) (*Bot, error) {
	status, err := LoadStatus()
	if err != nil {
		return nil, err
	}
	authZFunc := opts.Authorized
	if authZFunc == nil {
		authZFunc = func(*telegram.Update) (bool, int, string) {
			return true, 0, ""
		}
	}
	bot := Bot{
		failRetryInterval: opts.FailRetryInterval,
		Telegram:          telegram.NewClient(opts.APIToken, opts.LongPollingTimeout),
		status:            *status,
		authorized:        authZFunc,
	}
	return &bot, nil
}

type Bot struct {
	authorized        AuthZFunction
	failRetryInterval int
	Telegram          *telegram.BaseClient
	status            Status
	updateHandlers    []UpdateHandler
}

func (bot Bot) Run() {
	for {
		updates, err := bot.Telegram.GetUpdates(telegram.GetUpdatesRequest{
			Offset: bot.status.LastUpdate() + 1,
		})
		if err != nil {
			logrus.WithError(err).Error("update receive failure")
			time.Sleep(time.Duration(bot.failRetryInterval) * time.Second)
			continue
		}
		for _, update := range *updates {
			if allowed, chatID, message := bot.authorized(&update); allowed {
				bot.processUpdate(&update)
			} else {
				bot.Report(chatID, WARN, message)
			}
			bot.status.SetUpdate(update.UpdateID)
		}
		if bot.status.Changed() {
			if err = bot.status.Save(); err != nil {
				logrus.WithError(err).Error("fail saving status file")
			}
		}
	}
}

func (bot *Bot) processUpdate(update *telegram.Update) {

	if len(bot.updateHandlers) == 0 {
		logrus.Warn("no handlers registered. Please use OnX(...) methods to register some")
		return
	}

	for _, handler := range bot.updateHandlers {
		stop, err := handler(bot, update)
		if err != nil {
			logrus.WithError(err).Error("handling update")
		}
		if stop {
			return
		}
	}
}

type UpdateHandler func(bot *Bot, update *telegram.Update) (breakChain bool, err error)
type MessageHandler func(bot *Bot, update *telegram.Message) (breakChain bool, err error)
type EditedMessageHandler func(bot *Bot, update *telegram.Message) (breakChain bool, err error)
type ChannelPostHandler func(bot *Bot, update *telegram.Message) (breakChain bool, err error)
type EditedChannelPostHandler func(bot *Bot, update *telegram.Message) (breakChain bool, err error)
type InlineQueryHandler func(bot *Bot, update *telegram.InlineQuery) (breakChain bool, err error)
type ChosenInlineResultHandler func(bot *Bot, update *telegram.ChosenInlineResult) (breakChain bool, err error)
type CallbackQueryHandler func(bot *Bot, update *telegram.CallbackQuery) (breakChain bool, err error)
type ShippingQueryHandler func(bot *Bot, update *telegram.ShippingQuery) (breakChain bool, err error)
type PreCheckoutQueryHandler func(bot *Bot, update *telegram.PreCheckoutQuery) (breakChain bool, err error)
type PollHandler func(bot *Bot, update *telegram.Poll) (breakChain bool, err error)

func (bot *Bot) OnUpdate(handler UpdateHandler) {
	bot.updateHandlers = append(bot.updateHandlers, handler)
}

func (bot *Bot) OnMessage(handler MessageHandler) {
	bot.updateHandlers = append(bot.updateHandlers, func(bot *Bot, update *telegram.Update) (breakChain bool, err error) {
		return handler(bot, update.Message)
	})
}

func (bot *Bot) OnEditedMessage(handler MessageHandler) {
	bot.updateHandlers = append(bot.updateHandlers, func(bot *Bot, update *telegram.Update) (breakChain bool, err error) {
		return handler(bot, update.Message)
	})
}

func (bot *Bot) OnChannelPost(handler MessageHandler) {
	bot.updateHandlers = append(bot.updateHandlers, func(bot *Bot, update *telegram.Update) (breakChain bool, err error) {
		return handler(bot, update.Message)
	})
}

func (bot *Bot) EditedChannelPostMessage(handler MessageHandler) {
	bot.updateHandlers = append(bot.updateHandlers, func(bot *Bot, update *telegram.Update) (breakChain bool, err error) {
		return handler(bot, update.Message)
	})
}

func (bot *Bot) OnInlineQuery(handler InlineQueryHandler) {
	bot.updateHandlers = append(bot.updateHandlers, func(bot *Bot, update *telegram.Update) (breakChain bool, err error) {
		return handler(bot, update.InlineQuery)
	})
}

func (bot *Bot) OnChosenInlineResult(handler ChosenInlineResultHandler) {
	bot.updateHandlers = append(bot.updateHandlers, func(bot *Bot, update *telegram.Update) (breakChain bool, err error) {
		return handler(bot, update.ChosenInlineResult)
	})
}

func (bot *Bot) OnCallbackQuery(handler CallbackQueryHandler) {
	bot.updateHandlers = append(bot.updateHandlers, func(bot *Bot, update *telegram.Update) (breakChain bool, err error) {
		return handler(bot, update.CallbackQuery)
	})
}

func (bot *Bot) OnShippingQuery(handler ShippingQueryHandler) {
	bot.updateHandlers = append(bot.updateHandlers, func(bot *Bot, update *telegram.Update) (breakChain bool, err error) {
		return handler(bot, update.ShippingQuery)
	})
}

func (bot *Bot) OnPreCheckoutQuery(handler PreCheckoutQueryHandler) {
	bot.updateHandlers = append(bot.updateHandlers, func(bot *Bot, update *telegram.Update) (breakChain bool, err error) {
		return handler(bot, update.PreCheckoutQuery)
	})
}

func (bot *Bot) OnPoll(handler PollHandler) {
	bot.updateHandlers = append(bot.updateHandlers, func(bot *Bot, update *telegram.Update) (breakChain bool, err error) {
		return handler(bot, update.Poll)
	})
}
