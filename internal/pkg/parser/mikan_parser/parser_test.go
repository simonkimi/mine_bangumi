package mikan_parser

import (
	"github.com/joho/godotenv"
	"github.com/simonkimi/minebangumi/internal/pkg/config"
	"os"
	"testing"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	config.Setup()
	config.AppConfig.MikanProxy.Enable = os.Getenv("MIKAN_PROXY_ENABLED") == "1"
	config.AppConfig.MikanProxy.Scheme = os.Getenv("MIKAN_PROXY_SCHEME")
	config.AppConfig.MikanProxy.Host = os.Getenv("MIKAN_PROXY_HOST")
	config.AppConfig.MikanProxy.Port = os.Getenv("MIKAN_PROXY_PORT")
	config.AppConfig.MikanProxy.UseAuth = os.Getenv("MIKAN_PROXY_USE_AUTH") == "1"
	config.AppConfig.MikanProxy.Username = os.Getenv("MIKAN_PROXY_USERNAME")
	config.AppConfig.MikanProxy.Password = os.Getenv("MIKAN_PROXY_PASSWORD")
}

func TestParseBangumi(t *testing.T) {
	result, err := ParseFromUrl("https://mikanani.me/RSS/Bangumi?bangumiId=3386&subgroupid=615")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestParserMyBangumi(t *testing.T) {
	result, err := ParseFromUrl("https://mikanani.me/RSS/MyBangumi?token=b99ffHuTfy1nTftJ9H9DK0Kz6jyN18DgL6JhmSvtjXQ%3d")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)

}
