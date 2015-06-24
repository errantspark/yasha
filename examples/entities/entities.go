package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
	//"reflect"

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
		var herostates = make(map[string][]HeroState)

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
					//var oldstate HeroState
					if len(herostates[pe.Name]) > 1 {
						oldstate := herostates[pe.Name][len(herostates[pe.Name])-1]
						hero = hero.Diff(oldstate)
					}
					herostates[pe.Name] = append(herostates[pe.Name], hero)
				}
				//spew.Println(string(data))
				//if _, ok := pe.Delta["DT_DOTA_BaseNPC.m_vecOrigin"]; ok {
				//coord := coordFromCell(pe)
				//fmt.Printf("%30s | X: %5.0f Y: %5.0f\n", pe.Name[18:len(pe.Name)], coord.X, coord.Y)
				//}
			}
			/*
				if offset > 0 && pe.Tick > offset+90 {
					parser.OnEntityPreserved = nil
				}*/
		}
		parser.Parse()
		data, err := json.MarshalIndent(herostates, "", "  ")
		if err != nil {
			panic(err)
		}
		spew.Println(string(data))
	}
}

type HeroState struct {
	Tick        int        `json:",omitempty"`
	Name        string     `json:",omitempty"`
	Health      uint       `json:",omitempty"`
	HealthMax   uint       `json:",omitempty"`
	Cord        Coordinate `json:",omitempty"`
	AgiBase     float64    `json:",omitempty"`
	Agi         float64    `json:",omitempty"`
	IntBase     float64    `json:",omitempty"`
	Int         float64    `json:",omitempty"`
	StrBase     float64    `json:",omitempty"`
	Str         float64    `json:",omitempty"`
	DmgBonus    int        `json:",omitempty"`
	DmgMax      int        `json:",omitempty"`
	DmgMin      int        `json:",omitempty"`
	HealthRegen float64    `json:",omitempty"`
	Mana        float64    `json:",omitempty"`
	ManaRegen   float64    `json:",omitempty"`
	ManaMax     float64    `json:",omitempty"`
	Level       int        `json:",omitempty"`
	XP          int        `json:",omitempty"`
	PlayerID    int        `json:",omitempty"`
	Rotation    float64    `json:",omitempty"`
	VisionNight int        `json:",omitempty"`
	VisionDay   int        `json:",omitempty"`
	//TypeOfXP    string
}

func (h HeroState) Diff(oldState HeroState) HeroState {
	var nullState = HeroState{}
	if h.Tick == oldState.Tick {
		h.Tick = nullState.Tick
	}
	if h.Name == oldState.Name {
		h.Name = nullState.Name
	}
	if h.Health == oldState.Health {
		h.Health = nullState.Health
	}
	if h.HealthMax == oldState.HealthMax {
		h.HealthMax = nullState.HealthMax
	}
	if h.Cord == oldState.Cord {
		h.Cord = nullState.Cord
	}
	if h.AgiBase == oldState.AgiBase {
		h.AgiBase = nullState.AgiBase
	}
	if h.Agi == oldState.Agi {
		h.Agi = nullState.Agi
	}
	if h.IntBase == oldState.IntBase {
		h.IntBase = nullState.IntBase
	}
	if h.Int == oldState.Int {
		h.Int = nullState.Int
	}
	if h.StrBase == oldState.StrBase {
		h.StrBase = nullState.StrBase
	}
	if h.Str == oldState.Str {
		h.Str = nullState.Str
	}
	if h.DmgBonus == oldState.DmgBonus {
		h.DmgBonus = nullState.DmgBonus
	}
	if h.DmgMax == oldState.DmgMax {
		h.DmgMax = nullState.DmgMax
	}
	if h.DmgMin == oldState.DmgMin {
		h.DmgMin = nullState.DmgMin
	}
	if h.HealthRegen == oldState.HealthRegen {
		h.HealthRegen = nullState.HealthRegen
	}
	if h.Mana == oldState.Mana {
		h.Mana = nullState.Mana
	}
	if h.ManaRegen == oldState.ManaRegen {
		h.ManaRegen = nullState.ManaRegen
	}
	if h.ManaMax == oldState.ManaMax {
		h.ManaMax = nullState.ManaMax
	}
	if h.Level == oldState.Level {
		h.Level = nullState.Level
	}
	if h.XP == oldState.XP {
		h.XP = nullState.XP
	}
	if h.PlayerID == oldState.PlayerID {
		h.PlayerID = nullState.PlayerID
	}
	if h.Rotation == oldState.Rotation {
		h.Rotation = nullState.Rotation
	}
	if h.VisionNight == oldState.VisionNight {
		h.VisionNight = nullState.VisionNight
	}
	if h.VisionDay == oldState.VisionDay {
		h.VisionDay = nullState.VisionDay
	}
	return h
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
		//	XP:          pe.Values["DT_DOTA_BaseNPC.m_iCurrentXP"].(int),
		XP:          pe.Values["DT_DOTA_BaseNPC_Hero.m_iCurrentXP"].(int),
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
	X, Y float64 `json:",omitempty"`
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
