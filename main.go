package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	gololapi "github.com/Djipyy/GoLolApi"
	"github.com/yanzay/tbot"
)

// Configuration : Stores the config tokens for Telegram and League
type Configuration struct {
	TelegramToken string `json:"telegramToken"`
	LeagueToken   string `json:"leagueToken"`
}

var conf Configuration

func main() {

	// Load config.json and get Telegram and League tokens

	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		fmt.Println("error:", err)
	}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&conf)

	token := conf.TelegramToken
	// Create new telegram bot server using token
	bot, err := tbot.NewServer(token)
	if err != nil {
		log.Fatal(err)
	}

	// Handles the command /user
	bot.HandleFunc("/user {username}", UserHandler)

	// Start listening for messages
	err = bot.ListenAndServe()
	log.Fatal(err)
}

// UserHandler : Handles the command /user
func UserHandler(message *tbot.Message) {
	name := message.Vars["username"]
	str := GetSummonerInfo(name)
	fmt.Println(str)
	message.Reply(str)
}

// GetSummonerInfo : Gets summoner info
func GetSummonerInfo(username string) (str string) {
	// Access the Riot API in the NA server
	api := gololapi.NewAPI(gololapi.NA, conf.LeagueToken, 0.8)
	summ := api.GetSummonerByName(username)
	list := summ.GetChampionMasteries()
	str += ("Data for " + summ.Name + ":\n\n")
	for i := 0; i < 5; i++ {
		element := list[i]
		id := strconv.Itoa(element.ID)
		points := strconv.Itoa(element.Points)
		level := strconv.Itoa(element.Level)
		str += ("Champion " + id + ": Level " + level + "\n" + points + " mastery points" + "\n\n")
	}
	return
}
