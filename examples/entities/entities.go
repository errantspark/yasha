package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/dotabuff/yasha"
)

const MAX_COORDINATE float64 = 16384

// Example of printing any updates to hero coordinates

//&{4232
//DT_DOTAGamerules.m_fGameTime:143.66839599609375
//DT_DOTAGamerules.m_flPreGameStartTime:53.9661865234375

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected a .dem file as argument")
	}

	for _, path := range os.Args[1:] {
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
			//if strings.HasPrefix(pe.Name, "DT_DOTA_Unit_Hero_") {
			if offset > 0 && pe.Tick > offset && pe.Tick < offset+90 {
				if strings.HasPrefix(pe.Name, "DT_DOTA_Unit_Hero_") {
					var hero = parseHeroState(pe)
					data, err := json.MarshalIndent(hero, "", "  ")
					if err != nil {
						panic(err)
					}
					spew.Println(string(data))
				}
				//spew.Println(string(data))
				//if _, ok := pe.Delta["DT_DOTA_BaseNPC.m_vecOrigin"]; ok {
				//coord := coordFromCell(pe)
				//fmt.Printf("%30s | X: %5.0f Y: %5.0f\n", pe.Name[18:len(pe.Name)], coord.X, coord.Y)
				//}
			}
			if offset > 0 && pe.Tick > offset+90 {
				parser.OnEntityPreserved = nil
			}
		}
		parser.Parse()
	}
}

type HeroState struct {
	Tick        int
	Name        string
	Health      uint
	HealthMax   uint
	Cord        Coordinate
	AgiBase     float64
	Agi         float64
	IntBase     float64
	Int         float64
	StrBase     float64
	Str         float64
	DmgBonus    int
	DmgMax      int
	DmgMin      int
	HealthRegen float64
	Mana        float64
	ManaRegen   float64
	ManaMax     float64
	Level       int
	XP          int
	PlayerID    int
	Rotation    float64
	VisionNight int
	VisionDay   int
}

func parseHeroState(pe *yasha.PacketEntity) HeroState {
	if !strings.HasPrefix(pe.Name, "DT_DOTA_Unit_Hero_") {
		panic("not a hero packet")
	}
	coord := coordFromCell(pe)
	return HeroState{
		HealthMax:   pe.Values["DT_DOTA_BaseNPC.m_iMaxHealth"].(uint),
		Health:      pe.Values["DT_DOTA_BaseNPC.m_iHealth"].(uint),
		Tick:        pe.Tick,
		Name:        pe.Name,
		Cord:        coord,
		AgiBase:     pe.Values["DT_DOTA_BaseNPC_Hero.m_flAgility"].(float64),
		Agi:         pe.Values["DT_DOTA_BaseNPC_Hero.m_flAgilityTotal"].(float64),
		IntBase:     pe.Values["DT_DOTA_BaseNPC_Hero.m_flIntellect"].(float64),
		Int:         pe.Values["DT_DOTA_BaseNPC_Hero.m_flIntellectTotal"].(float64),
		StrBase:     pe.Values["DT_DOTA_BaseNPC_Hero.m_flStrength"].(float64),
		Str:         pe.Values["DT_DOTA_BaseNPC_Hero.m_flStrengthTotal"].(float64),
		DmgBonus:    pe.Values["DT_DOTA_BaseNPC.m_iDamageBonus"].(int),
		DmgMax:      pe.Values["DT_DOTA_BaseNPC.m_iDamageMax"].(int),
		DmgMin:      pe.Values["DT_DOTA_BaseNPC.m_iDamageMin"].(int),
		HealthRegen: pe.Values["DT_DOTA_BaseNPC.m_flHealthThinkRegen"].(float64),
		Mana:        pe.Values["DT_DOTA_BaseNPC.m_flMana"].(float64),
		ManaRegen:   pe.Values["DT_DOTA_BaseNPC.m_flManaThinkRegen"].(float64),
		ManaMax:     pe.Values["DT_DOTA_BaseNPC.m_flMaxMana"].(float64),
		Level:       pe.Values["DT_DOTA_BaseNPC.m_iCurrentLevel"].(int),
		XP:          pe.Values["DT_DOTA_BaseNPC.m_iCustomXPValue"].(int),
		PlayerID:    pe.Values["DT_DOTA_BaseNPC_Hero.m_iPlayerID"].(int),
		Rotation:    pe.Values["DT_DOTA_BaseNPC.m_angRotation[1]"].(float64),
		VisionNight: pe.Values["DT_DOTA_BaseNPC.m_iNightTimeVisionRange"].(int),
		VisionDay:   pe.Values["DT_DOTA_BaseNPC.m_iDayTimeVisionRange"].(int),
	}
}

func getTickOffset(tick int, pretime float64, gametime float64) int {
	var nowtime = gametime - pretime - 90
	var tickoffset = float64(tick) - (nowtime * 30)
	return int(tickoffset)
}

type Coordinate struct {
	X, Y float64
}

func coordFromCell(pe *yasha.PacketEntity) Coordinate {
	cellbits, ok := pe.Values["DT_BaseEntity.m_cellbits"].(int)
	if !ok {
		return Coordinate{X: 0, Y: 0}
	}
	cellWidth := float64(uint(1) << uint(cellbits))

	var cX, cY, vX, vY float64

	if vO2, ok := pe.Values["DT_DOTA_BaseNPC.m_vecOrigin"].(*yasha.Vector2); ok {
		cX = float64(pe.Values["DT_DOTA_BaseNPC.m_cellX"].(int))
		cY = float64(pe.Values["DT_DOTA_BaseNPC.m_cellY"].(int))
		vX, vY = vO2.X, vO2.Y
	} else {
		vO3 := pe.Values["DT_BaseEntity.m_vecOrigin"].(*yasha.Vector3)
		cX = float64(pe.Values["DT_BaseEntity.m_cellX"].(int))
		cY = float64(pe.Values["DT_BaseEntity.m_cellY"].(int))
		vX, vY = vO3.X, vO3.Y
	}

	x := ((cX * cellWidth) - MAX_COORDINATE) + vX
	y := ((cY * cellWidth) - MAX_COORDINATE) + vY

	return Coordinate{X: x, Y: y}
}
