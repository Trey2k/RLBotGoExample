package main

import (
	"fmt"

	RLBot "github.com/Trey2k/RLBotGo"
)

var lastTouch float32
var totalTouches int = 0

// Tick handler takes in a GameState which contains the gameTickPacket, ballPredidctions, fieldInfo and matchSettings
func getInput(gameState *RLBot.GameState, rlBot *RLBot.RLBot) *RLBot.PlayerInput {
	// Send a random controller state every time the ball is touched

	PlayerInput := &RLBot.PlayerInput{}
	PlayerInput.PlayerIndex = 0

	wasjustTouched := false
	if gameState.GameTick.Ball.LatestTouch.GameSeconds != 0 && lastTouch != gameState.GameTick.Ball.LatestTouch.GameSeconds {
		totalTouches++
		lastTouch = gameState.GameTick.Ball.LatestTouch.GameSeconds
		wasjustTouched = true
	}

	if wasjustTouched && totalTouches <= 10 {
		rlBot.DebugMessageAdd(fmt.Sprintf("The ball was touched %d times", totalTouches))
		PlayerInput.ControllerState.Jump = false
	} else if wasjustTouched && totalTouches > 10 {
		rlBot.DebugMessageClear()
		totalTouches = 0
		PlayerInput.ControllerState.Jump = true
	}
	return PlayerInput

}

func main() {

	// connect to RLBot
	rlBot, err := RLBot.InitConnection(23234)
	if err != nil {
		panic(err)
	}

	// Send ready message
	err = rlBot.SendReadyMessage(true, true, true)
	if err != nil {
		panic(err)
	}

	// Set our tick handler
	rlBot.SetGetInput(getInput)

}
