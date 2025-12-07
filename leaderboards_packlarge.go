//go:build windows

package steamworks

import "unsafe"

// 64-bit Windows uses VALVE_CALLBACK_PACK_LARGE as defined in steamclientpublic.h (8-byte alignment)

type LeaderboardFindResult_t struct {
	SteamLeaderboard SteamLeaderboard_t
	LeaderboardFound byte
	_ [7]byte // padding
}

// Expect size 16. If this line fails to compile, the struct size is incorrect according to steam headers.
var _ = [1]int{unsafe.Sizeof(LeaderboardFindResult_t{}) - 16: 0}

type LeaderboardScoresDownloaded_t struct {
	SteamLeaderboard SteamLeaderboard_t
	SteamLeaderboardEntries SteamLeaderboardEntries_t
	EntryCount int32
	_ [4]byte // padding
}

// Expect size 24. If this line fails to compile, the struct size is incorrect according to steam headers.
var _ = [1]int{unsafe.Sizeof(LeaderboardScoresDownloaded_t{}) - 24: 0}

type LeaderboardScoreUploaded_t struct {
	Success                  byte
	_                        [7]byte // padding
	SteamLeaderboard         SteamLeaderboard_t
	Score                    int32
	ScoreChanged             byte
	_                        [3]byte // padding
	GlobalRankNew            int32
	GlobalRankPrevious       int32
	_                        [4]byte // padding
}

var _ = [1]int{unsafe.Sizeof(unsafe.Offsetof(LeaderboardScoreUploaded_t{}.SteamLeaderboard)) - 8: 0}

// Expect size 40. If this line fails to compile, the struct size is incorrect according to steam headers.
var _ = [1]int{unsafe.Sizeof(LeaderboardScoreUploaded_t{}) - 40: 0}

type LeaderboardEntry_t struct {
	SteamIDUser CSteamID
	GlobalRank  int32
	Score       int32
	Details     int32
	_           [4]byte
	UGC         UGCHandle_t
}

// Expect size 40. If this line fails to compile, the struct size is incorrect according to steam headers.
var _ = [1]int{unsafe.Sizeof(LeaderboardEntry_t{}) - 32: 0}