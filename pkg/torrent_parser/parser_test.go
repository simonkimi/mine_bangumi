package torrent_parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTorrentName(t *testing.T) {
	fileName := "[Lilith-Raws] Boku no Kokoro no Yabai Yatsu - 01 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4].mp4"
	bf := ParseTorrentName(fileName, "", 0, TorrentFileTypeVideo).(*TorrentEpisodeFile)
	assert.Equal(t, bf.Group, "Lilith-Raws")
	assert.Equal(t, bf.Title, "Boku no Kokoro no Yabai Yatsu")
	assert.Equal(t, bf.Season, 1)
	assert.Equal(t, bf.Episode, 1)

	fileName = "[Sakurato] Tonikaku Kawaii S2 [01][AVC-8bit 1080p AAC][CHS].mp4"
	bf = ParseTorrentName(fileName, "", 0, TorrentFileTypeVideo).(*TorrentEpisodeFile)
	assert.Equal(t, bf.Group, "Sakurato")
	assert.Equal(t, bf.Title, "Tonikaku Kawaii")
	assert.Equal(t, bf.Season, 2)
	assert.Equal(t, bf.Episode, 1)

	fileName = "[SweetSub&LoliHouse] Heavenly Delusion - 01 [WebRip 1080p HEVC-10bit AAC ASSx2].mkv"
	bf = ParseTorrentName(fileName, "", 0, TorrentFileTypeVideo).(*TorrentEpisodeFile)
	assert.Equal(t, bf.Group, "SweetSub&LoliHouse")
	assert.Equal(t, bf.Title, "Heavenly Delusion")
	assert.Equal(t, bf.Season, 1)
	assert.Equal(t, bf.Episode, 1)

	fileName = "[SBSUB][CONAN][1082][V2][1080P][AVC_AAC][CHS_JP](C1E4E331).mp4"
	bf = ParseTorrentName(fileName, "", 0, TorrentFileTypeVideo).(*TorrentEpisodeFile)
	assert.Equal(t, bf.Group, "SBSUB")
	assert.Equal(t, bf.Title, "CONAN")
	assert.Equal(t, bf.Season, 1)
	assert.Equal(t, bf.Episode, 1082)

	fileName = "海盗战记 (2019) S01E01.mp4"
	bf = ParseTorrentName(fileName, "", 0, TorrentFileTypeVideo).(*TorrentEpisodeFile)
	assert.Equal(t, bf.Title, "海盗战记 (2019)")
	assert.Equal(t, bf.Season, 1)
	assert.Equal(t, bf.Episode, 1)

	fileName = "海盗战记/海盗战记 S01E01.mp4"
	bf = ParseTorrentName(fileName, "", 0, TorrentFileTypeVideo).(*TorrentEpisodeFile)
	assert.Equal(t, bf.Title, "海盗战记")
	assert.Equal(t, bf.Season, 1)
	assert.Equal(t, bf.Episode, 1)

	fileName = "海盗战记 S01E01.zh-tw.ass"
	sf := ParseTorrentName(fileName, "", 0, TorrentFileTypeSubtitle).(*TorrentSubtitleFile)
	assert.Equal(t, sf.Title, "海盗战记")
	assert.Equal(t, sf.Season, 1)
	assert.Equal(t, sf.Episode, 1)
	assert.Equal(t, sf.Language, "zh-Hant")

	fileName = "海盗战记 S01E01.SC.ass"
	sf = ParseTorrentName(fileName, "", 0, TorrentFileTypeSubtitle).(*TorrentSubtitleFile)
	assert.Equal(t, sf.Title, "海盗战记")
	assert.Equal(t, sf.Season, 1)
	assert.Equal(t, sf.Episode, 1)
	assert.Equal(t, sf.Language, "zh-Hans")

	fileName = "水星的魔女(2022) S00E19.mp4"
	bf = ParseTorrentName(fileName, "", 0, TorrentFileTypeVideo).(*TorrentEpisodeFile)
	assert.Equal(t, bf.Title, "水星的魔女(2022)")
	assert.Equal(t, bf.Season, 0)
	assert.Equal(t, bf.Episode, 19)

	fileName = "【失眠搬运组】放学后失眠的你-Kimi wa Houkago Insomnia - 06 [bilibili - 1080p AVC1 CHS-JP].mp4"
	bf = ParseTorrentName(fileName, "", 1, TorrentFileTypeVideo).(*TorrentEpisodeFile)
	assert.Equal(t, bf.Title, "放学后失眠的你-Kimi wa Houkago Insomnia")
	assert.Equal(t, bf.Season, 1)
	assert.Equal(t, bf.Episode, 6)
}
