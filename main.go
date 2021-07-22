package main

import (
	"math/rand"

	RLBot "github.com/Trey2k/RLBotGo"
)

var lastTouch float32

// Tick handler takes in a GameState which contains the gameTickPacket, ballPredidctions, fieldInfo and matchSettings
func tick(gameState *RLBot.GameState, socket *RLBot.Socket) {
	// Send a random controller state every time the ball is touched
	if gameState.GameTick.Ball.LatestTouch.GameSeconds != 0 && gameState.GameTick.Ball.LatestTouch.GameSeconds != lastTouch {
		PlayerInput := &RLBot.PlayerInput{
			PlayerIndex: 0,
			ControllerState: RLBot.ControllerState{
				Throttle:  float32(rand.Int31n(10 - -10)+-10) / 100,
				Steer:     float32(rand.Int31n(10 - -10)+-10) / 100,
				Yaw:       float32(rand.Int31n(10 - -10)+-10) / 100,
				Pitch:     float32(rand.Int31n(10 - -10)+-10) / 100,
				Roll:      float32(rand.Int31n(10 - -10)+-10) / 100,
				Jump:      rand.Int31n(1) == 1,
				Boost:     rand.Int31n(1) == 1,
				Handbrake: rand.Int31n(1) == 1,
				UseItem:   rand.Int31n(1) == 1,
			},
		}

		lastTouch = gameState.GameTick.Ball.LatestTouch.GameSeconds
		socket.SendMessage(RLBot.DataType_PlayerInput, PlayerInput)
	}

}

func main() {

	// connect to RLBot
	socket, err := RLBot.InitConnection(23234)
	if err != nil {
		panic(err)
	}
	// Prepare ready message
	readyMsg := &RLBot.ReadyMessage{
		WantsBallPredictions: true,
		WantsQuickChat:       true,
		WantsGameMessages:    true,
	}

	// Send ready message
	err = socket.SendMessage(RLBot.DataType_ReadyMessage, readyMsg)
	if err != nil {
		panic(err)
	}

	// Set our tick handler
	socket.SetTickHandler(tick)

}
