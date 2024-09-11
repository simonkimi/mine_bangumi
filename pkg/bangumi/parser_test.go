package bangumi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTorrentName2(t *testing.T) {
	filename := "[Nekomoe kissaten][Helck][01][1080p][JPTC].mp4"
	bf := ParseBangumiSourceName(filename, "")
	fmt.Printf("%+v\n", bf)
}

func TestParseTorrentName(t *testing.T) {
	fileName := "[Lilith-Raws] Boku no Kokoro no Yabai Yatsu - 01 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4].mp4"
	bf := ParseBangumiSourceName(fileName, "")
	assert.Equal(t, "Lilith-Raws", bf.Group)
	assert.Equal(t, "Boku no Kokoro no Yabai Yatsu", bf.Title)
	assert.Equal(t, -1, bf.Season)
	assert.Equal(t, 1, bf.Episode)

	fileName = "[Sakurato] Tonikaku Kawaii S2 [01][AVC-8bit 1080p AAC][CHS].mp4"
	bf = ParseBangumiSourceName(fileName, "")
	assert.Equal(t, "Sakurato", bf.Group)
	assert.Equal(t, "Tonikaku Kawaii", bf.Title)
	assert.Equal(t, 2, bf.Season)
	assert.Equal(t, 1, bf.Episode)

	fileName = "[SweetSub&LoliHouse] Heavenly Delusion - 01 [WebRip 1080p HEVC-10bit AAC ASSx2].mkv"
	bf = ParseBangumiSourceName(fileName, "")
	assert.Equal(t, "SweetSub&LoliHouse", bf.Group)
	assert.Equal(t, "Heavenly Delusion", bf.Title)
	assert.Equal(t, -1, bf.Season)
	assert.Equal(t, 1, bf.Episode)

	fileName = "[SBSUB][CONAN][1082][V2][1080P][AVC_AAC][CHS_JP](C1E4E331).mp4"
	bf = ParseBangumiSourceName(fileName, "")
	assert.Equal(t, "SBSUB", bf.Group)
	assert.Equal(t, "CONAN", bf.Title)
	assert.Equal(t, -1, bf.Season)
	assert.Equal(t, 1082, bf.Episode)

	fileName = "海盗战记 (2019) S01E01.mp4"
	bf = ParseBangumiSourceName(fileName, "")
	assert.Equal(t, "海盗战记 (2019)", bf.Title)
	assert.Equal(t, 1, bf.Season)
	assert.Equal(t, 1, bf.Episode)

	fileName = "海盗战记/海盗战记 S01E01.mp4"
	bf = ParseBangumiSourceName(fileName, "")
	assert.Equal(t, "海盗战记", bf.Title)
	assert.Equal(t, 1, bf.Season)
	assert.Equal(t, 1, bf.Episode)

	fileName = "海盗战记 S01E01.zh-tw.ass"
	sf := ParseBangumiSourceName(fileName, "")
	assert.Equal(t, "海盗战记", sf.Title)
	assert.Equal(t, 1, sf.Season)
	assert.Equal(t, 1, sf.Episode)
	assert.Equal(t, "zh-Hant", sf.Language)

	fileName = "海盗战记 S01E01.SC.ass"
	sf = ParseBangumiSourceName(fileName, "")
	assert.Equal(t, "海盗战记", sf.Title)
	assert.Equal(t, 1, sf.Season)
	assert.Equal(t, 1, sf.Episode)
	assert.Equal(t, "zh-Hans", sf.Language)

	fileName = "水星的魔女(2022) S00E19.mp4"
	bf = ParseBangumiSourceName(fileName, "")
	assert.Equal(t, "水星的魔女(2022)", bf.Title)
	assert.Equal(t, 0, bf.Season)
	assert.Equal(t, 19, bf.Episode)

	fileName = "【失眠搬运组】放学后失眠的你-Kimi wa Houkago Insomnia - 06 [bilibili - 1080p AVC1 CHS-JP].mp4"
	bf = ParseBangumiSourceName(fileName, "")
	assert.Equal(t, "放学后失眠的你-Kimi wa Houkago Insomnia", bf.Title)
	assert.Equal(t, -1, bf.Season)
	assert.Equal(t, 6, bf.Episode)
}
