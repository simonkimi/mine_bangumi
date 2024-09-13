package magnet

//
//import (
//	"context"
//	"github.com/anacrolix/torrent"
//	"github.com/simonkimi/minebangumi/pkg/errno"
//	"os"
//	"strings"
//	"time"
//)
//
//type FileInfo struct {
//	Name  string
//	Files []string
//}
//
//func ParseMagnet(ctx context.Context, magnet string) (*FileInfo, error) {
//	clientConfig := torrent.NewDefaultClientConfig()
//	clientConfig.NoUpload = true
//	clientConfig.DataDir = os.TempDir()
//	clientConfig.Seed = false
//	clientConfig.DisableUTP = true
//
//	client, err := torrent.NewClient(clientConfig)
//
//	if err != nil {
//		return nil, errno.NewApiErrorWithCausef(errno.InternalServerError, err, "failed to create torrent client")
//	}
//	defer client.Close()
//
//	torrentFile, err := client.AddMagnet(magnet)
//	if err != nil {
//		return nil, errno.NewApiErrorWithCausef(errno.InternalServerError, err, "failed to add magnet")
//	}
//	timeout, cancel := context.WithTimeout(ctx, 30*time.Second)
//	defer cancel()
//
//	select {
//	case <-ctx.Done():
//		return nil, errno.NewApiError(errno.ErrorCancel)
//	case <-timeout.Done():
//		return nil, errno.NewApiErrorf(errno.ErrorTimeout, "Timeout to get torrent info")
//	case <-torrentFile.GotInfo():
//		info := torrentFile.Info()
//		fileInfo := &FileInfo{Name: info.Name}
//		for _, file := range info.UpvertedFiles() {
//			for _, path := range file.Path {
//				if strings.Contains(path, "_____padding_file_0_") {
//					continue
//				}
//				fileInfo.Files = append(fileInfo.Files, path)
//			}
//		}
//		return fileInfo, nil
//	}
//}
