package main

import (
	"fmt"

	RLBot "github.com/Trey2k/RLBotGo"
)

var lastTouch float32
var totalTouches int = 0

// getInput takes in a GameState which contains the gameTickPacket, ballPredidctions, fieldInfo and matchSettings
// it also takes in the RLBot object. And returns a PlayerInput

func getInput(gameState *RLBot.GameState, rlBot *RLBot.RLBot) *RLBot.ControllerState {
	PlayerInput := &RLBot.ControllerState{}

	// Count ball touches up to 10 and on 11 clear the messages and jump
	wasjustTouched := false
	if gameState.GameTick.Ball.LatestTouch.GameSeconds != 0 && lastTouch != gameState.GameTick.Ball.LatestTouch.GameSeconds {
		totalTouches++
		lastTouch = gameState.GameTick.Ball.LatestTouch.GameSeconds
		wasjustTouched = true
	}

	if wasjustTouched && totalTouches <= 10 {
		// DebugMessage is a helper function to let you quickly get debug text on screen. it will autmaicly place it so text will not overlap
		rlBot.DebugMessageAdd(fmt.Sprintf("The ball was touched %d times", totalTouches))
		PlayerInput.Jump = false
	} else if wasjustTouched && totalTouches > 10 {
		rlBot.DebugMessageClear()
		totalTouches = 0
		PlayerInput.Jump = true
	}
	return PlayerInput

}

func main() {

	// connect to RLBot
	rlBot, err := RLBot.Connect(23234)
	if err != nil {
		panic(err)
	}

	// Send ready message
	err = rlBot.SendReadyMessage(true, true, true)
	if err != nil {
		panic(err)
	}

	// Set our tick handler
	err = rlBot.SetGetInput(getInput)
	fmt.Println(err.Error())

}
