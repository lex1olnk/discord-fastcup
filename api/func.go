package api

func (stats *Stats) processKills(kills []Kill) {
	for _, kill := range kills {
		if _, ok := stats.Players[kill.KillerId]; !ok {
			stats.Players[kill.KillerId] = &PlayerStats{}
		}
		stats.Players[kill.KillerId].Kills++
		if kill.IsHeadshot {
			stats.Players[kill.KillerId].Headshots++
		}

		if _, ok := stats.Players[kill.VictimId]; !ok {
			stats.Players[kill.VictimId] = &PlayerStats{}
		}
		stats.Players[kill.VictimId].Deaths++

		if kill.AssistantId != nil {
			if _, ok := stats.Players[*kill.AssistantId]; !ok {
				stats.Players[*kill.AssistantId] = &PlayerStats{}
			}
			stats.Players[*kill.AssistantId].Assists++
		}
	}
}

func (stats *Stats) processDamage(Damages []Damage) {
	for _, damage := range Damages {
		if _, ok := stats.Players[damage.InflictorId]; !ok {
			stats.Players[damage.InflictorId] = &PlayerStats{}
		}
		stats.Players[damage.InflictorId].AverageDamage += damage.DamageNormalized
	}
}
