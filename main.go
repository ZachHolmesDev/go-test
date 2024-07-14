package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type App struct {
	AppID int    `json:"appid"`
	Name  string `json:"name"`
}

type AppList struct {
	Apps []App `json:"apps"`
}

type AppListData struct {
	AppList AppList `json:"applist"`
}

func main() {
	fmt.Print("Starting Go")
	// fmt.Print(quote.Go())
	json_data, err := os.ReadFile("./all_steam_games.json")
	if err != nil {
		log.Fatalf("err reading JSON file: %v")
	}

	var all_steam_games_data AppListData

	err = json.Unmarshal(json_data, &all_steam_games_data)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// outwave steam id
	// game_id := "2050310"
	game_name := "OutWave"

	// call the function
	appID, err := find_id_by_name(all_steam_games_data, game_name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("AppID is: %d\n", appID)
	}
}

func find_id_by_name(data AppListData, name string) (int, error) {
	for _, app := range data.AppList.Apps {
		if app.Name == name {
			return app.AppID, nil
		}
	}
	return 0, errors.New("app not found")
}
