package shot

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"os"
	"time"
)

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func Shot() string {
	//获取所有活动屏幕
	n := screenshot.NumActiveDisplays()
	//全屏截取所有活动屏幕
	if n > 0 {
		for i := 0; i < n; i++ {
			img, err := screenshot.CaptureDisplay(i)
			if err != nil {
				panic(err)
			}
			save(img, fmt.Sprintf("screens/%d-screen-%s.png", i, time.Now().Format("2006-01-02 15-04-05")))
		}
		return fmt.Sprintf("screen-%s.png", time.Now().Format("2006-01-02 15-04-05"))
	}
	return ""
}

