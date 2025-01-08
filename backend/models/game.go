package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const GameTypeValue string = "game"

type Game struct {
	BaseModel
	PlayerIds []uuid.UUID `gorm:"type:uuid[]" json:"playerIds"`
	State     int         `gorm:"type:int" json:"state"`
}

func GetAllGames() []Game {
	// TODO: Implement database query to fetch all users
	game := Game{
		BaseModel: BaseModel{
			Id:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Type:      GameTypeValue,
		},
		PlayerIds: []uuid.UUID{uuid.New()},
		State:     0,
	}
	fmt.Println(game)
	return []Game{game}
}
