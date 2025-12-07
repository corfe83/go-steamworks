//go:build !windows

package steamworks

// Non-Windows platforms (Unix, macOS, FreeBSD) use VALVE_CALLBACK_PACK_SMALL as defined in steamclientpublic.h (4-byte alignment)

type LeaderboardFindResult_t struct {
	SteamLeaderboard SteamLeaderboard_t
	LeaderboardFound byte
	_ [3]byte // padding
}

type LeaderboardScoresDownloaded_t struct {
	SteamLeaderboard SteamLeaderboard_t
	SteamLeaderboardEntries SteamLeaderboardEntries_t
	EntryCount int32
}

type LeaderboardScoreUploaded_t struct {
	Success                  byte
	_                        [3]byte // padding
	SteamLeaderboard         SteamLeaderboard_t
	core                     int32
	ScoreChanged             byte
	_                        [3]byte // padding
	lobalRankNew             int32
	lobalRankPrevious        int32
}

type LeaderboardEntry_t struct {
	SteamIDUser CSteamID
	GlobalRank  int32
	Score       int32
	Details     int32
	UGC         UGCHandle_t
}
