package main

import (
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/dotabuff/yasha"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected a .dem file as argument")
	}

	var path = os.Args[1]
	parser := yasha.ParserFromFile(path)

	var now time.Duration
	var gameTime, preGameStarttime float64
	var offset int

	parser.OnEntityPreserved = func(pe *yasha.PacketEntity) {
		if pe.Name == "DT_DOTAGamerulesProxy" {
			gameTime = pe.Values["DT_DOTAGamerules.m_fGameTime"].(float64)
			preGameStarttime = pe.Values["DT_DOTAGamerules.m_flPreGameStartTime"].(float64)
			if gameTime > 0 && preGameStarttime > 0 {
				offset = getTickOffset(pe.Tick, preGameStarttime, gameTime)
			}
			now = time.Duration(gameTime-preGameStarttime) * time.Second
		}
		if offset > 0 && pe.Tick > offset && pe.Tick < offset+90 {
			//if strings.HasPrefix(pe.Name, "DT_DOTA_Unit_Hero_") {
			fmt.Println(offset)
			/*
				data, err := json.MarshalIndent(pe, "", "  ")
				if err != nil {
					panic(err)
				}
			*/
			spew.Dump(pe)
			//spew.Println(string(data))
		}
		//if _, ok := pe.Delta["DT_DOTA_BaseNPC.m_vecOrigin"]; ok {
		//coord := coordFromCell(pe)
		//fmt.Printf("%30s | X: %5.0f Y: %5.0f\n", pe.Name[18:len(pe.Name)], coord.X, coord.Y)
		//}
	}
	parser.Parse()
}

func getTickOffset(tick int, pretime float64, gametime float64) int {
	var nowtime = gametime - pretime - 90
	var tickoffset = float64(tick) - (nowtime * 30)
	return int(tickoffset)
}
