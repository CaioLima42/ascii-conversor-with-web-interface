package ffmpeg

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func ffmpegRunner(command string, videoBytes []byte) (io.ReadCloser, *exec.Cmd, error) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("erro criando stdin: %v", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("erro criando stdout: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("erro iniciando ffmpeg: %v", err)
	}

	go func() {
		defer stdin.Close()
		stdin.Write(videoBytes)
	}()

	return stdout, cmd, nil
}
