//go:build !windows

package steamworks

// Non-Windows platforms (Unix, macOS, FreeBSD) use VALVE_CALLBACK_PACK_SMALL as defined in steamclientpublic.h (4-byte alignment)

import "encoding/binary"

type leaderboardEntry_t struct {
	// m_steamIDUser: 8 bytes
	// m_nGlobalRank: 4 bytes
	// m_nScore:      4 bytes
	// m_cDetails:    4 bytes
	// m_hUGC:        8 bytes
	// total:         28 bytes
	data [28]byte
}

func (me leaderboardEntry_t) Read() leaderboardEntry {
	var result leaderboardEntry
	result.steamIDUser = CSteamID(binary.NativeEndian.Uint64(me.data[0:8]))
	result.globalRank = int32(binary.NativeEndian.Uint32(me.data[8:12]))
	result.score = int32(binary.NativeEndian.Uint32(me.data[12:16]))
	result.details = int32(binary.NativeEndian.Uint32(me.data[16:20]))
	result.UGC = UGCHandle_t(binary.NativeEndian.Uint64(me.data[20:28]))
	return result
}
