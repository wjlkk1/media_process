package mediainfo

import "time"

type Video struct {
	BVNumber    string
	Owner       string
	UploadTime  time.Time
	UpdatedTime time.Time
	videoInfo   map[string]interface{}
	audioInfo   map[string]interface{}
}
