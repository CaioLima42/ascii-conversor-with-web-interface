package ffmpeg

import (
	"bufio"
	"fmt"
	"image/png"
	"io"
	"net/http"
	"regexp"
	"time"
	"os/exec"
	processImage "github.com/CaioLima42/ascii-conversor-with-web-interface/pkg/processImage"
)

type requestBody struct {
	Video string `json:"video"`
}

type VideoMetadata struct {
	Width     int
	Height    int
	FrameRate float64
}

func extractVideoMetadata(videoBytes []byte) (VideoMetadata, error) {
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "null", "-")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return VideoMetadata{}, fmt.Errorf("erro criando stdin: %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return VideoMetadata{}, fmt.Errorf("erro criando stderr: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return VideoMetadata{}, fmt.Errorf("erro iniciando ffmpeg: %v", err)
	}

	go func() {
		defer stdin.Close()
		stdin.Write(videoBytes)
	}()

	scanner := bufio.NewScanner(stderr)
	var meta VideoMetadata

	reResolution := regexp.MustCompile(`, (\d+)x(\d+)[,\s]`)
	reFps := regexp.MustCompile(`, (\d+(?:\.\d+)?) fps`)

	for scanner.Scan() {
		line := scanner.Text()
		if matches := reResolution.FindStringSubmatch(line); matches != nil {
			fmt.Sscanf(matches[1], "%d", &meta.Width)
			fmt.Sscanf(matches[2], "%d", &meta.Height)
		}
		if matches := reFps.FindStringSubmatch(line); matches != nil {
			fmt.Sscanf(matches[1], "%f", &meta.FrameRate)
		}
	}
	cmd.Wait()

	if meta.Width == 0 || meta.Height == 0 || meta.FrameRate == 0 {
		return meta, fmt.Errorf("metadados não encontrados")
	}

	return meta, nil
}

func ExtractFramesStream(videoBytes []byte, w http.ResponseWriter) error {
	meta, err := extractVideoMetadata(videoBytes)
	if err != nil {
		fmt.Println("Aviso: não foi possível extrair metadados:", err)
	}
	command := "ffmpeg -i pipe:0 -f image2pipe -vcodec png -vsync 0 pipe:1"
	stdout, cmd, err := ffmpegRunner(command, videoBytes)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/plain")
	flusher, canFlush := w.(http.Flusher)
	for {
		start := time.Now()
		img, err := png.Decode(stdout)
		if err != nil {
			if err == io.EOF || err.Error() == "unexpected EOF" {
				break
			}
			return fmt.Errorf("erro decodificando frame: %v", err)
		}
		scaled := processImage.NearestNeighborScaling(img, 90, 90)
		grayscale := processImage.GrayScaleImage(scaled, processImage.RGB2GrayColorMean)
		ascii := processImage.Gray2Ascii(grayscale)
		w.Write([]byte(ascii))
		if canFlush {
			flusher.Flush()
		}
		elapsed := time.Since(start)
		delay := time.Duration(1000.0/meta.FrameRate)*time.Millisecond - elapsed
		if delay > 0 {
			time.Sleep(delay)
		}
	}

	cmd.Wait()
	return nil
}

func ExtractAudio(videoData []byte, w http.ResponseWriter) {

	command := "ffmpeg -i pipe:0 -f mp3 -vn pipe:1"
	stdout, cmd, err := ffmpegRunner(command, videoData)
	if err != nil{
		fmt.Printf("Erro ao extrair o audio: %v", err)
	}
	w.Header().Set("Content-Type", "audio/mpeg")
	io.Copy(w, stdout)
	cmd.Wait()
}