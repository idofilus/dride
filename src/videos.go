package main

import (
    "fmt"
	"io/ioutil"
	"log"
	"regexp"
)

const VIDEOS_DIR = "recordings"

type VideoClip struct {
	Name string `json:"name"`
	VideoSrc string `json:"videoSrc"`
	ThumbSrc string `json:"thumbSrc"`
	Checksum string `json:"checksum"`
}

type VideoEvent struct {
	Name string `json:"name"`
	Date int `json:"date"`
}

type VideosClip struct {
	Id string `json:"id"`
	Timestamp string `json:"timestamp"`
	Clips []VideoClip `json:"clips"`
	Gpx string `json:"gpx"`
	HaveEmergencyEvent bool `json:"haveEmergencyEvent"`
	Event []VideoEvent `json:"event"`
	OnSubDir bool `json:"onSubDir"`
}

type VideosClipsResult struct {
    Videos []VideosClip `json:"videos"`
}

func getVideosClips(page int, limit int, fromTimestamp int) VideosClipsResult {

	videos := []VideosClip{}
	clips := []VideoClip{}

	// Just a dummy regex to filter our specific videos
	re := regexp.MustCompile(`BigBuckBunny_h264_\d+.mp4`)
	mp4_re := regexp.MustCompile(`(.*).(?:mp4)$`)

	index := 0
	videos_files, err := ioutil.ReadDir("./" + VIDEOS_DIR)
    if err != nil {
        log.Fatal(err)
    }
 
    for _, video_file := range videos_files {
		if !video_file.IsDir() {
			continue
		}

		clips_files, err := ioutil.ReadDir("./" + VIDEOS_DIR  + "/" + video_file.Name())
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range clips_files {
			if !re.Match([]byte(f.Name())) {
				continue
			}
	
			index++
	
			// Offset to our page
			if (page - 1) * limit >= index {
				continue
			}
			
			clip := VideoClip {
				Name: f.Name(),
				VideoSrc: "http://localhost/" + VIDEOS_DIR + "/" + f.Name(),
				ThumbSrc: "http://localhost/" + VIDEOS_DIR + "/" + mp4_re.ReplaceAllString(f.Name(), `$1.jpg`),
				Checksum: "5f039b4ef0058a1d652f13d612375a5b",
			}
			
			clips = append(clips, clip)
	
			if len(clips) >= limit {
				break
			}
		}

		video_events := []VideoEvent{}
		video_events = append(video_events, VideoEvent{
			Name: "buttonPress",
			Date: 1629730257,
		})

		videos = append(videos, VideosClip {
			Id: video_file.Name(),
			Timestamp: video_file.Name(),
			Clips: clips,
			Gpx: "http://localhost/" + VIDEOS_DIR + "/route.geojson",
			HaveEmergencyEvent: false,
			Event: video_events,
			OnSubDir: false,
		})
    }

	result := VideosClipsResult { Videos: videos }

	return result
}