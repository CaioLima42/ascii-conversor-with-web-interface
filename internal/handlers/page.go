package handlers

import (
	"net/http"
	"github.com/CaioLima42/ascii-conversor-with-web-interface/internal/utils"
)

func ReadVideo(w http.ResponseWriter, r *http.Request) {
	data, err := utils.ReadTemplate("index.html")
	if err != nil {
		w.Write([]byte("Erro ao processar o arquivo"))
		return
	}
	w.Write(data)
}
