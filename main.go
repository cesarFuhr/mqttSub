package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	err := bootstrap()
	if err != nil {
		fmt.Println(err)
	}

	subscriber := Subscriber{}
	subscriber.Connect(os.Getenv("MQTT_BROKER_URL"))

	if err := subscriber.Subscribe("testTopic"); err != nil {
		fmt.Println(err)
		return
	}

	exit := pleaseLeave()

	<-exit
	if err := subscriber.Unsubscribe("testTopic"); err != nil {
		fmt.Println(err)
		return
	}
}

func bootstrap() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func pleaseLeave() chan struct{} {
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		close(done)
	}()
	return done
}
