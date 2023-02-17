package utils

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
)

// CoverGenerator 调用ffmpeg获取第一帧为封面，bug（路径错误）未解决
func CoverGenerator(videoDest, coverDest string) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoDest).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).Run()
	if err != nil {
		log.Println("生成失败：", err)
		return "", err
	}
	image, err := imaging.Decode(buf)
	if err != nil {
		log.Println("生成失败：", err)
		return "", err
	}
	err = imaging.Save(image, coverDest+".png")
	if err != nil {
		log.Println("生成失败：", err)
		return "", err
	}

	return coverDest + ".png", nil
}
