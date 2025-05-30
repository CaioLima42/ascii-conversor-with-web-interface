package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/CaioLima42/ascii-conversor-with-web-interface/internal/ffmpeg"
)

type requestBody struct {
	Video string `json:"video"`
}

func CreateVideo(w http.ResponseWriter, r *http.Request) {
	var req requestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}
	videoData, err := base64.StdEncoding.DecodeString(req.Video)
	if err != nil {
		http.Error(w, "Erro ao decodificar base64", http.StatusBadRequest)
		return
	}
	if err := ffmpeg.ExtractFramesStream(videoData, w); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao extrair frames: %v", err), http.StatusInternalServerError)
	}
}

func ExtractAudio(w http.ResponseWriter, r *http.Request) {
	var req requestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}
	videoData, err := base64.StdEncoding.DecodeString(req.Video)
	if err != nil {
		http.Error(w, "Erro ao decodificar base64", http.StatusBadRequest)
		return
	}
	ffmpeg.ExtractAudio(videoData, w)
}
