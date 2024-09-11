package mikan

import (
	"github.com/joho/godotenv"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/pkg/setup"
	"os"
	"testing"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	setup.SetupTest()
	config.AppConfig.ProxyConfig.Enable = os.Getenv("PROXY_ENABLED") == "1"
	config.AppConfig.ProxyConfig.Scheme = os.Getenv("PROXY_SCHEME")
	config.AppConfig.ProxyConfig.Host = os.Getenv("PROXY_HOST")
	config.AppConfig.ProxyConfig.Port = os.Getenv("PROXY_PORT")
	config.AppConfig.ProxyConfig.UseAuth = os.Getenv("PROXY_USE_AUTH") == "1"
	config.AppConfig.ProxyConfig.Username = os.Getenv("PROXY_USERNAME")
	config.AppConfig.ProxyConfig.Password = os.Getenv("PROXY_PASSWORD")
}

func TestParseBangumi(t *testing.T) {
	result, err := ParseRssData("https://mikanani.me/RSS/Bangumi?bangumiId=3386&subgroupid=615")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestParserMyBangumi(t *testing.T) {
	result, err := ParseRssData("https://mikanani.me/RSS/MyBangumi?token=b99ffHuTfy1nTftJ9H9DK0Kz6jyN18DgL6JhmSvtjXQ%3d")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)

}
