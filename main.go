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
	cfg.LoggerConfig.LogLevel = logging.LevelError

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
		time.Sleep(10 * time.Millisecond)
	}

	splitClient.Destroy()
}

func getEnv(splitClient *client.SplitClient, userID int) {
	treatment := splitClient.Treatment(userID, "FEATURE_FLAG_USER_ORDER_CREATE_SUS", map[string]interface{}{
		"USER_ID": fmt.Sprintf("%d", userID),
	})
	fmt.Println("User ID:", userID, "Value:", treatment)
}
