// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2021 The go-steamworks Authors

//go:build windows

package steamworks

// 64-bit Windows uses VALVE_CALLBACK_PACK_LARGE as defined in steamclientpublic.h (8-byte alignment)

type LeaderboardScoreUploaded_t struct {
	Success            byte
	_                  [7]byte // padding
	SteamLeaderboard   SteamLeaderboard_t
	Score              int32
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
	ugc         UGCHandle_t
}

// See leaderboards_packsmall.go for explanation of why this is a function and not a direct member access
func (me LeaderboardEntry_t) UGC() UGCHandle_t {
	return me.ugc
}
