package main

import (
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"

	"../../../yasha"
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
		var nowRaw float64

		parser.OnFileInfo = func(obj *dota.CDemoFileInfo) {
			fmt.Println(obj)
		}

		var dumpedMap bool
		parser.OnActiveModifierDelta = func(modmap map[int]*yasha.StringTableItem, modbuf yasha.ModifierBuffs) {
			if !dumpedMap {
				fmt.Println("modmap")
				spew.Dump(modmap)
				dumpedMap = true
			}
			if nowRaw > 85 && nowRaw < 100 {
				fmt.Println("modbuf")
				spew.Dump(modbuf)

				if nowRaw > 100 {
					parser.OnActiveModifierDelta = nil
				}
			}
		}

		var lul = func(entry yasha.CombatLogEntry) {
			if nowRaw > 85 && nowRaw < 100 {
				switch log := entry.(type) {
				case *yasha.CombatLogModifierAdd:
					fmt.Println(now)
					fmt.Println("combatmod")
					spew.Dump(log)
				}
			}
			if nowRaw > 100 {
				parser.OnCombatLog = nil
			}
			/*
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
		parser.OnEntityPreserved = func(pe *yasha.PacketEntity) {
			if pe.Name == "DT_DOTAGamerulesProxy" {
				gameTime = pe.Values["DT_DOTAGamerules.m_fGameTime"].(float64)
				preGameStarttime = pe.Values["DT_DOTAGamerules.m_flPreGameStartTime"].(float64)
				now = time.Duration(gameTime-preGameStarttime) * time.Second
				nowRaw = gameTime - preGameStarttime
				if nowRaw > 85 {
					parser.OnCombatLog = lul
				}
			}
		}
		parser.Parse()
	}
}
