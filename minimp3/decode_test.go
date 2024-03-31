package minimp3

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/gopxl/beep/speaker"
)

func TestMp3(t *testing.T) {

	resp, _ := http.Get("http://m10.music.126.net/20230826232756/a34d7413ef11640d3ad2a5f10b91a2dd/ymusic/466b/40f0/1794/414344cd36fa2d9b177464719cd6c8cd.mp3")

	streamer, format, _ := Decode(resp.Body)

	log.Printf("Convert audio sample rate: %d, channels: %d\n", format.SampleRate, format.NumChannels)

	_ = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))

	speaker.Play(streamer)

	select {}
}
