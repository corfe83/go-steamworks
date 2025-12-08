//go:build windows

package steamworks

// 64-bit Windows uses VALVE_CALLBACK_PACK_LARGE as defined in steamclientpublic.h (8-byte alignment)

type LeaderboardFindResult_t struct {
	SteamLeaderboard SteamLeaderboard_t
	LeaderboardFound byte
	_                [7]byte // padding
}

type LeaderboardScoresDownloaded_t struct {
	SteamLeaderboard        SteamLeaderboard_t
	SteamLeaderboardEntries SteamLeaderboardEntries_t
	EntryCount              int32
	_                       [4]byte // padding
}

type LeaderboardScoreUploaded_t struct {
	Success            byte
	_                  [7]byte // padding
	SteamLeaderboard   SteamLeaderboard_t
	Score              int32
	ScoreChanged       byte
	_                  [3]byte // padding
	GlobalRankNew      int32
	GlobalRankPrevious int32
	_                  [4]byte // padding
}
type LeaderboardEntry_t struct {
	SteamIDUser CSteamID
	GlobalRank  int32
	Score       int32
	Details     int32
	ugc         UGCHandle_t
}

func (me LeaderboardEntry_t) UGC() UGCHandle_t {
	return me.ugc
}
