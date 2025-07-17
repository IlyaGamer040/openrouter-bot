package main

import (
	"context"
	"log"

	"github.com/IlyaGamer040/openrouter-bot/internal/config"
	"github.com/IlyaGamer040/openrouter-bot/internal/service"
)

func main() {

	ctx := context.Background()

	cfg, _ := config.NewConfig()

	api := service.New(cfg)

	//  тестовое дерьмо message := "Что на картинке?"

	// imageurl := "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e7/Everest_North_Face_toward_Base_Camp_Tibet_Luca_Galuzzi_2006.jpg/1920px-Everest_North_Face_toward_Base_Camp_Tibet_Luca_Galuzzi_2006.jpg"

	//log.Print(api.SendImageMessage(ctx, imageurl, message))
}
