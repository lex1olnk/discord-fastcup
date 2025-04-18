package api

func (stats *Stats) getMatchData(match Match) []int {
	var currentPlayers []int
	for _, player := range match.Members {
		user := player.Private.User
		if _, ok := stats.Players[user.ID]; !ok {
			stats.Players[user.ID] = &PlayerStats{}
			stats.Players[user.ID].ID = user.ID
			stats.Players[user.ID].Nickname = user.NickName
		}
		currentPlayers = append(currentPlayers, user.ID)
		stats.Players[user.ID].Rounds += len(match.Rounds)
	}
	return currentPlayers
}

func calculateTrade(index int, killer Kill, kills []Kill, stats *Stats) int {
	for i := index - 1; i > -1; i-- {
		difference := killer.CreatedAt.Sub(kills[i].CreatedAt)
		if difference.Seconds() > 5 {
			return 0
		}
		if killer.VictimId == kills[i].KillerId {
			stats.Players[killer.KillerId].Traded++
			stats.Players[kills[i].VictimId].Exchanged++
			return kills[i].VictimId
		}
	}
	return 0
	//difference := trader.CreatedAt.Second() - killer.CreatedAt.Second()
	//fmt.Println(difference)
	// Время убийства

}

type roundType struct {
	kills     int
	isDead    bool
	hasKill   bool
	hasTrade  bool
	hasAssist bool
}

func (stats *Stats) processKills(kills []Kill, currentPlayers []int) {
	roundKAST := make(map[int]*roundType)
	for i := range currentPlayers {
		if _, ok := roundKAST[currentPlayers[i]]; !ok {
			roundKAST[currentPlayers[i]] = &roundType{}
		}
	}

	roundId := 0
	for i, kill := range kills {
		if roundId != kill.RoundId {
			for a := range roundKAST {
				if !roundKAST[a].isDead || roundKAST[a].hasKill || roundKAST[a].hasAssist || roundKAST[a].hasTrade {
					stats.Players[a].KASTScore++
				}
				if roundKAST[a].kills > 0 {
					stats.Players[a].MultiKills[roundKAST[a].kills-1]++
					roundKAST[a].kills = 0
				}
				roundKAST[a].isDead = false
				roundKAST[a].hasKill = false
				roundKAST[a].hasAssist = false
				roundKAST[a].hasTrade = false
			}
		}
		stats.Players[kill.VictimId].Deaths++
		roundKAST[kill.VictimId].isDead = true

		stats.Players[kill.KillerId].Kills++
		roundKAST[kill.KillerId].kills++
		if kill.IsHeadshot {
			stats.Players[kill.KillerId].Headshots++
		}
		roundKAST[kill.KillerId].hasKill = true

		if kill.AssistantId != nil {
			stats.Players[*kill.AssistantId].Assists++
			roundKAST[*kill.AssistantId].hasAssist = true
		}

		idExchanged := calculateTrade(i, kill, kills, stats)
		if idExchanged != 0 {
			roundKAST[idExchanged].hasTrade = true
		}

		// Расчет первого убийства и смерти
		if roundId != kill.RoundId {
			stats.Players[kill.KillerId].FirstKill++
			stats.Players[kill.VictimId].FirstDeath++

			roundId = kill.RoundId
		}

	}
	for a := range roundKAST {
		if !roundKAST[a].isDead || roundKAST[a].hasKill || roundKAST[a].hasAssist || roundKAST[a].hasTrade {
			stats.Players[a].KASTScore++
		}
		if roundKAST[a].kills > 0 {
			stats.Players[a].MultiKills[roundKAST[a].kills-1]++
		}
		roundKAST[a].isDead = false
		roundKAST[a].hasKill = false
		roundKAST[a].hasAssist = false
		roundKAST[a].hasTrade = false
	}
}

func (stats *Stats) processDamage(Damages []Damage) {
	for _, damage := range Damages {
		if _, ok := stats.Players[damage.InflictorId]; !ok {
			stats.Players[damage.InflictorId] = &PlayerStats{}
		}
		stats.Players[damage.InflictorId].AverageDamage += float64(damage.DamageNormalized)
	}
}
