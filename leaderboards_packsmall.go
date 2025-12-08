//go:build !windows

package steamworks

import "unsafe"

// Non-Windows platforms (Unix, macOS, FreeBSD) use VALVE_CALLBACK_PACK_SMALL as defined in steamclientpublic.h (4-byte alignment)

type LeaderboardFindResult_t struct {
	SteamLeaderboard SteamLeaderboard_t
	LeaderboardFound byte
	_                [3]byte // padding
}

type LeaderboardScoresDownloaded_t struct {
	SteamLeaderboard        SteamLeaderboard_t
	SteamLeaderboardEntries SteamLeaderboardEntries_t
	EntryCount              int32
}

type LeaderboardScoreUploaded_t struct {
	Success            byte
	_                  [3]byte // padding
	SteamLeaderboard   SteamLeaderboard_t
	Core               int32
	ScoreChanged       byte
	_                  [3]byte // padding
	GlobalRankNew      int32
	GlobalRankPrevious int32
}

type LeaderboardEntry_t struct {
	SteamIDUser CSteamID
	GlobalRank  int32
	Score       int32
	Details     int32
	ugc         [8]byte // if we use the actual UGCHandle_t here, struct size/alignment changes
}

func (entry LeaderboardEntry_t) UGC() UGCHandle_t {
	ugcPtr := unsafe.Pointer(&entry.ugc[0])
	return *(*UGCHandle_t)(ugcPtr)
}
