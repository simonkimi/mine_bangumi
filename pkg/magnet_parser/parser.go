package magnet_parser

import (
	"context"
	"github.com/anacrolix/torrent"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"os"
)

type FileInfo struct {
	Name  string
	Files []string
}

func ParseMagnet(magnet string, ctx context.Context) (*FileInfo, error) {
	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.NoUpload = true
	clientConfig.DataDir = os.TempDir()
	clientConfig.Seed = false
	clientConfig.DisableUTP = true

	client, err := torrent.NewClient(clientConfig)

	if err != nil {
		return nil, errno.NewApiErrorWithCausef(errno.InternalServerError, err, "failed to create torrent client")
	}
	defer client.Close()

	torrentFile, err := client.AddMagnet(magnet)
	if err != nil {
		return nil, errno.NewApiErrorWithCausef(errno.InternalServerError, err, "failed to add magnet")
	}
	select {
	case <-torrentFile.GotInfo():
	case <-ctx.Done():
		return nil, errno.NewApiError(errno.ErrorCancel)
	}
	info := torrentFile.Info()

	fileInfo := &FileInfo{Name: info.Name}

	for _, file := range info.UpvertedFiles() {
		for _, path := range file.Path {
			fileInfo.Files = append(fileInfo.Files, path)
		}
	}

	return fileInfo, nil
}
