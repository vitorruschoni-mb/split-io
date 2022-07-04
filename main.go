package main

import (
	"fmt"
	"github.com/splitio/go-client/splitio/client"
	"github.com/splitio/go-client/splitio/conf"
	"github.com/splitio/go-toolkit/logging"
	"math/rand"
	"time"
)

func main() {
	cfg := conf.Default()
	cfg.TaskPeriods.SplitSync = 5
	cfg.TaskPeriods.EventsSync = 5
	cfg.TaskPeriods.TelemetrySync = 30
	cfg.LoggerConfig.LogLevel = logging.LevelDebug

	factory, err := client.NewSplitFactory("6buq76ob0v4up1h1itau80bahmena08h7olo", cfg)
	if err != nil {
		fmt.Printf("SDK init error: %s\n", err)
		return
	}

	splitClient := factory.Client()
	err = splitClient.BlockUntilReady(25)
	if err != nil {
		fmt.Printf("SDK init error: %s\n", err)
		return
	}

	for {
		number := rand.Intn(5 - 1) + 1
		getEnv(splitClient, number)
		time.Sleep(1 * time.Second)
	}

	splitClient.Destroy()
}

func getEnv(splitClient *client.SplitClient, userID int) {
	treatment := splitClient.Treatment(fmt.Sprintf("%d", userID), "FEATURE_FLAG_USER_ORDER_CREATE_SUS", map[string]interface{}{
		"USER_ID": fmt.Sprintf("%d", userID),
		"BASE": "BTC",
		"METHOD": "PLACE",
		"order_type": "market",
	})
	fmt.Println("User ID:", userID, "Value:", treatment)
}
