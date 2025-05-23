package api

import "time"

type Kill struct {
	RoundId             int       `json:"roundId"`
	CreatedAt           time.Time `json:"createdAt"`
	KillerId            int       `json:"killerId"`
	VictimId            int       `json:"victimId"`
	AssistantId         *int      `json:"assistantId"` // Используем указатель, если может быть null
	WeaponId            int       `json:"weaponId"`
	IsHeadshot          bool      `json:"isHeadshot"`
	IsWallbang          bool      `json:"isWallbang"`
	IsOneshot           bool      `json:"isOneshot"`
	IsAirshot           bool      `json:"isAirshot"`
	IsNoscope           bool      `json:"isNoscope"`
	KillerPositionX     int       `json:"killerPositionX"`
	KillerPositionY     int       `json:"killerPositionY"`
	VictimPositionX     int       `json:"victimPositionX"`
	VictimPositionY     int       `json:"victimPositionY"`
	KillerBlindedBy     *int      `json:"killerBlindedBy"`     // Используем указатель, если может быть null
	KillerBlindDuration *int      `json:"killerBlindDuration"` // Используем указатель, если может быть null
	VictimBlindedBy     int       `json:"victimBlindedBy"`
	VictimBlindDuration int       `json:"victimBlindDuration"`
	IsTeamkill          bool      `json:"isTeamkill"`
	TypeName            string    `json:"__typename"`
}

type GraphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]int `json:"variables"`
}

type PlayerStats struct {
	ID            int
	Nickname      string
	Kills         int
	Deaths        int
	Assists       int
	AverageDamage float64
	FirstDeath    int
	FirstKill     int
	Headshots     int
	Clutches      int
	Team          int
	Rounds        int
	Traded        int // разменял другого игрока
	Exchanged     int // Был разменят
	KPR           float64
	DPR           float64
	KASTScore     float64
	Impact        float64
	MultiKills    [5]int
	Rating        float64
}

type Stats struct {
	Players map[int]*PlayerStats
}

type GraphQLKillsResponse struct {
	Data struct {
		Kills []Kill `json:"kills"`
	} `json:"data"`
}

type GraphQLDamagesResponse struct {
	Data struct {
		Damages []Damage `json:"damages"`
	} `json:"data"`
}

type Damage struct {
	RoundId          int    `json:"roundId"`
	InflictorId      int    `json:"inflictorId"`
	VictimId         int    `json:"victimId"`
	WeaponId         int    `json:"weaponId"`
	HitboxGroup      int    `json:"hitboxGroup"`
	DamageReal       int    `json:"damageReal"`
	DamageNormalized int    `json:"damageNormalized"`
	Hits             int    `json:"hits"`
	TypeName         string `json:"__typename"`
}

type GetMatchStatsResponse struct {
	Data struct {
		Match Match `json:"match"`
	} `json:"data"`
}

type Match struct {
	ID                   int         `json:"id"`
	Type                 string      `json:"type"`
	Status               string      `json:"status"`
	BestOf               int         `json:"best_of"`
	GameID               int         `json:"game_id"`
	HasWinner            bool        `json:"has_winner"`
	StartedAt            string      `json:"started_at"`
	FinishedAt           *string     `json:"finished_at"`
	MaxRoundsCount       int         `json:"max_rounds_count"`
	ServerInstanceID     *int        `json:"server_instance_id"`
	CancellationReason   *string     `json:"cancellation_reason"`
	ReplayExpirationDate *string     `json:"replay_expiration_date"`
	Rounds               []Round     `json:"rounds"`
	Maps                 []MatchMap  `json:"maps"`
	GameMode             GameMode    `json:"game_mode"`
	Teams                []MatchTeam `json:"teams"`
	Members              []Member    `json:"members"`
	Typename             string      `json:"__typename"`
}

type GameMode struct {
	ID             int    `json:"id"`
	TeamSize       int    `json:"teamSize"`
	FirstTeamSize  int    `json:"firstTeamSize"`
	SecondTeamSize int    `json:"secondTeamSize"`
	TypeName       string `json:"__typename"`
}

type Round struct {
	ID             int    `json:"id"`
	WinReason      string `json:"win_reason"`
	StartedAt      string `json:"started_at"`
	FinishedAt     string `json:"finished_at"`
	MatchMapID     int    `json:"match_map_id"`
	SpawnedPlayers []int  `json:"spawned_players"` // Исправлено на slice
	WinMatchTeamID *int   `json:"win_match_team_id"`
	Typename       string `json:"__typename"`
}

type MatchMap struct {
	ID         int              `json:"id"`
	Number     int              `json:"number"`
	MapID      int              `json:"map_id"`
	StartedAt  string           `json:"started_at"`
	FinishedAt *string          `json:"finished_at"`
	GameStatus string           `json:"game_status"`
	Replays    []MatchMapReplay `json:"replays"`
	Map        Map              `json:"map"`
	Typename   string           `json:"__typename"`
}

type MatchMapReplay struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	CreatedAt string `json:"createdAt"`
	TypeName  string `json:"__typename"`
}

type Map struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Offset   *float64 `json:"offset"` // Исправлено на указатель
	Scale    *float64 `json:"scale"`  // Исправлено на указатель
	Preview  string   `json:"preview"`
	Topview  string   `json:"topview"`
	Overview string   `json:"overview"`
	FlipV    bool     `json:"flip_v"`
	FlipH    bool     `json:"flip_h"`
	Typename string   `json:"__typename"`
}

type MatchTeam struct {
	ID             int                `json:"id"`
	Name           string             `json:"name"`
	Size           int                `json:"size"`
	Score          int                `json:"score"`
	ChatID         *int               `json:"chat_id"` // Исправлено на *int
	IsWinner       bool               `json:"is_winner"`
	CaptainID      *int               `json:"captain_id"` // Исправлено на *int
	IsDisqualified bool               `json:"is_disqualified"`
	MapStats       []MatchTeamMapStat `json:"map_stats"`
	Typename       string             `json:"__typename"`
}

type Member struct {
	Hash        string             `json:"hash"`
	Role        string             `json:"role"`
	Ready       bool               `json:"ready"`
	Impact      *float64           `json:"impact"` // Исправлено на указатель
	Connected   bool               `json:"connected"`
	IsLeaver    bool               `json:"is_leaver"`
	RatingDiff  *float64           `json:"rating_diff"` // Исправлено на указатель
	MatchTeamID int                `json:"match_team_id"`
	Private     MatchMemberPrivate `json:"private"`
	Typename    string             `json:"__typename"`
}

type MatchMemberPrivate struct {
	Rating   int    `json:"rating"`
	PartyID  int    `json:"partyId"`
	User     User   `json:"user"`
	TypeName string `json:"__typename"`
}

type User struct {
	ID                            int           `json:"id"`
	Link                          *string       `json:"link"` // Исправлено на указатель
	Avatar                        string        `json:"avatar"`
	Online                        bool          `json:"online"`
	Verified                      bool          `json:"verified"`
	IsMobile                      bool          `json:"is_mobile"`
	NickName                      string        `json:"nickName"`
	AnimatedAvatar                *string       `json:"animated_avatar"` // Исправлено на указатель
	IsMedia                       bool          `json:"is_media"`
	DisplayMediaStatus            bool          `json:"display_media_status"`
	PrivacyOnlineStatusVisibility string        `json:"privacy_online_status_visibility"`
	Subscription                  *Subscription `json:"subscription"`
	Icon                          *ProfileIcon  `json:"icon"`
	Stats                         []UserStat    `json:"stats"`
	Typename                      string        `json:"__typename"`
}

type UserStat struct {
	Kills      int     `json:"kills"`
	Deaths     int     `json:"deaths"`
	Place      *int    `json:"place"` // Исправлено на указатель
	Rating     float64 `json:"rating"`
	WinRate    float64 `json:"win_rate"`
	GameModeID int     `json:"game_mode_id"`
	Typename   string  `json:"__typename"`
}

type MatchTeamMapStat struct {
	Score       int     `json:"score"`
	IsWinner    bool    `json:"isWinner"`
	MatchMapID  int     `json:"matchMapId"`
	MatchTeamID int     `json:"matchTeamId"`
	InitialSide *string `json:"initialSide"`
	TypeName    string  `json:"__typename"`
}

type Subscription struct {
	PlanID int `json:"planId"`
}

type ProfileIcon struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}
