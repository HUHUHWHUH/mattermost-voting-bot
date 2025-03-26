package settings

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/rs/zerolog"
	"mattermost-voting-bot/internal/storage"
)

// Application содержит зависимости для бота
type Application struct {
	Logger                    zerolog.Logger
	Config                    Config
	MattermostClient          *model.Client4
	MattermostWebSocketClient *model.WebSocketClient
	MattermostUser            *model.User
	MattermostTeam            *model.Team
	MattermostChannel         *model.Channel

	Storage *storage.Storage
}
