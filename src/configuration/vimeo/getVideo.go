package vimeo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"go.uber.org/zap"
)

type GetVideoResponse struct {
	Title        string `json:"title"`
	Html	       string `json:"html"`
	ThumbnailUrl string `json:"thumbnail_url"`
}

type GetVideoParams struct {
	VideoID    string
	Width      int
	Autoplay   bool
	Muted      bool
	Background bool
}

func GetVideo(params GetVideoParams) (*GetVideoResponse, int, error) {
	baseURL := "https://vimeo.com/api/oembed.json"

	videoURL := fmt.Sprintf("https://vimeo.com/%s", params.VideoID)

	query := url.Values{}
	query.Set("autopause", "false")
	query.Set("url", videoURL)

	if params.Width > 0 {
		query.Set("width", strconv.Itoa(params.Width))
	}
	if params.Autoplay {
		query.Set("autoplay", "1")
	}
	if params.Muted {
		query.Set("muted", "1")
	}
	if params.Background {
		query.Set("background", "1")
	}

	url := baseURL + "?" + query.Encode()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error("Error trying NewRequest", err, zap.String("journey", "GetVideo"))
		return nil, 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("Error trying do request", err, zap.String("journey", "PutVideoFolder"))
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error trying ReadAll", err, zap.String("journey", "GetVideo"))
		return nil, 0, err
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, resp.StatusCode, nil
		}
		logger.Error("Error trying GetVideo", errors.New(string(body)), zap.String("journey", "GetVideo"))
		return nil, resp.StatusCode, errors.New("failed to get video")
	}

	var jsonResp GetVideoResponse
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		logger.Error("Error trying NewDecoder", err, zap.String("journey", "GetVideo"))
		return nil, resp.StatusCode, err
	}
	
	var thumbnailUrl string
	var newWidthStr string
	if params.Width >= 350 {
		newWidthStr = "640"
	}
	parts := strings.Split(jsonResp.ThumbnailUrl, "-d_")
	if len(parts) == 2 {
		base, rest := parts[0]+"-d_", parts[1]
		if i := strings.Index(rest, "x"); i != -1 {
			if len(rest) - 1 >= i + 1 {
				rest = rest[i+1:]
			}
			for i := 0; i < len(rest); i++ {
				_,err := strconv.Atoi(string(rest[i]))
				if err != nil {
					break
				}
				if len(rest) - 1 >= i + 1 {
					rest = rest[i+1:]
				} else {
					break
				}
			}
		}
		thumbnailUrl = base + newWidthStr + rest
	} else {
		thumbnailUrl = jsonResp.ThumbnailUrl
	}
	jsonResp.ThumbnailUrl = thumbnailUrl
	
	return &jsonResp, resp.StatusCode, nil
}