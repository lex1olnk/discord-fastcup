package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GraphQLRequest структура для GraphQL-запроса

func NewStats() *Stats {
	return &Stats{
		Players: make(map[int]*PlayerStats),
	}
}

// GraphQLResponse структура для GraphQL-ответа

func sendGraphQLRequest(w http.ResponseWriter, query string, variables map[string]int, responseBody interface{}) bool {
	requestBody := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	// Кодируем тело запроса в JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, "Error encoding request body", http.StatusInternalServerError)
		return false
	}

	resp, err := http.Post("https://hasura.fastcup.net/v1/graphql", "application/json", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		http.Error(w, "Error fetching match data", http.StatusInternalServerError)
		return false
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		http.Error(w, "Error decoding response body", http.StatusInternalServerError)
		return false
	}

	return true
}

func getMatchMembers(w http.ResponseWriter, matchID int, stats *Stats) []int {
	query := fullMatchQuery

	variables := map[string]int{
		"matchId": matchID,
		"gameId":  3,
	}
	var responseBody GetMatchStatsResponse
	if !sendGraphQLRequest(w, query, variables, &responseBody) {
		return nil
	}

	return stats.getMatchData(responseBody.Data.Match)
}

func getMatchKills(w http.ResponseWriter, matchID int, stats *Stats, currentPlayers []int) bool {
	query := matchKillsQuery

	variables := map[string]int{
		"matchId": matchID,
		//"userId":  0, // Замените на нужный userId, если требуется
	}

	var responseBody GraphQLKillsResponse
	if !sendGraphQLRequest(w, query, variables, &responseBody) {
		return false
	}

	stats.processKills(responseBody.Data.Kills, currentPlayers)
	return true
}

func getMatchDamages(w http.ResponseWriter, matchID int, stats *Stats) bool {
	query := matchDamageQuery

	variables := map[string]int{
		"matchId": matchID,
		//"userId":  0, // Замените на нужный userId, если требуется
	}

	// Декодируем JSON-ответ
	var responseBody GraphQLDamagesResponse
	if !sendGraphQLRequest(w, query, variables, &responseBody) {
		return false
	}

	stats.processDamage(responseBody.Data.Damages)
	return true
}

// MatchHandler обрабатывает запросы к маршруту /match/{id}
func MatchHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем номер матча из URL
	matchID, err := strconv.Atoi(r.URL.Path[len("/match/"):])
	if err != nil || matchID <= 0 {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}
	// Формируем GraphQL-запрос
	stats := NewStats()
	currentPlayers := getMatchMembers(w, matchID, stats)
	if currentPlayers == nil {
		http.Error(w, "players haven't downloaded", http.StatusBadRequest)
		return
	}

	if !getMatchKills(w, matchID, stats, currentPlayers) {
		http.Error(w, "kills error", http.StatusBadRequest)
		return
	}

	if !getMatchDamages(w, matchID, stats) {
		http.Error(w, "damages error", http.StatusBadRequest)
		return
	}
	// Отправляем HTML-таблицу в ответе
	tableHTML := `
        <table border="1">
            <tr>
                <th>Player ID</th>
				<th>Nickname</th>
                <th>Kills</th>
                <th>Deaths</th>
                <th>Assists</th>
                <th>ADR</th>
                <th>Headshots</th>
				<th>FirstKill</th>
				<th>FirstDeath</th>
				<th>Traded</th>
				<th>Exchanged</th>
				<th>KAST</th>
				<th>Rounds</th>
            </tr>
    `

	for playerID, playerStats := range stats.Players {
		tableHTML += fmt.Sprintf(`
            <tr>
                <td>%d</td>
				<td>%s</td>
                <td>%d</td>
                <td>%d</td>
                <td>%d</td>
                <td>%f</td>
                <td>%d</td>
				<td>%d</td>
				<td>%d</td>
				<td>%d</td>
				<td>%d</td>
				<td>%f</td>
				<td>%d</td>
            </tr>
        `, playerID, playerStats.Nickname, playerStats.Kills, playerStats.Deaths, playerStats.Assists, playerStats.AverageDamage, playerStats.Headshots, playerStats.FirstKill, playerStats.FirstDeath, playerStats.Traded, playerStats.Exchanged, playerStats.KASTScore, playerStats.Rounds)
	}

	tableHTML += `</table>`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, tableHTML)
}
