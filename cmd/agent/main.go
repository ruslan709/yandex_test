package main

import (
	"os"
	"strconv"

	"github.com/ruslan709/yandex_calculate/yandex_calculate/internal/application/agent"
)

func main() {
	computingPower, _ := strconv.Atoi(os.Getenv("COMPUTING_POWER"))
	client := agent.NewClient("http://localhost:8080")

	for i := 0; i < computingPower; i++ {
		go agent.NewWorker(client).Start()
	}

	select {}
}
