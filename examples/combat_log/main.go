package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dotabuff/yasha"
	"github.com/dotabuff/yasha/dota"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected a .dem file as argument")
	}

	for _, path := range os.Args[1:] {
		parser := yasha.ParserFromFile(path)

		var now time.Duration
		var gameTime, preGameStarttime float64

		parser.OnEntityPreserved = func(pe *yasha.PacketEntity) {
			if pe.Name == "DT_DOTAGamerulesProxy" {
				gameTime = pe.Values["DT_DOTAGamerules.m_fGameTime"].(float64)
				preGameStarttime = pe.Values["DT_DOTAGamerules.m_flPreGameStartTime"].(float64)
				now = time.Duration(gameTime-preGameStarttime) * time.Second
			}
		}
		parser.OnFileInfo = func(obj *dota.CDemoFileInfo) {
			fmt.Println(obj)
		}

		parser.OnCombatLog = func(entry yasha.CombatLogEntry) {
			switch log := entry.(type) {
			case *yasha.CombatLogMultikill:
				fmt.Println("multikill")
				fmt.Println(log)
			case *yasha.CombatLogKillStreak:
				fmt.Println("killstreak")
				fmt.Println(log)
			case *yasha.CombatLogTeamBuildingKill:
				fmt.Println("bdeath")
				fmt.Println(log)
			case *yasha.CombatLogDeath:
				fmt.Println("death")
				fmt.Println(log)
			}
			/*
				switch log := entry.(type) {
				case *yasha.CombatLogPurchase:
					fmt.Printf("%7s | %s bought a %s\n", now, log.Buyer, log.Item)
				case *yasha.CombatLogAbility:
					if log.Target == "dota_unknown" {
						fmt.Printf("%7s | %s casted %s\n", now, log.Attacker, log.Ability)
					} else {
						fmt.Printf("%7s | %s casted %s on %s\n", now, log.Attacker, log.Ability, log.Target)
					}
				case *yasha.CombatLogHeal:
					fmt.Printf("%7s | %s heals %s for %dHP\n", now, log.Source, log.Target, log.Value)
				case *yasha.CombatLogDamage:
					fmt.Printf("%7s | %s damages %s for %dHP\n", now, log.Source, log.Target, log.Value)
					fmt.Println(log)
				}
			*/
		}
		parser.Parse()
	}
}
