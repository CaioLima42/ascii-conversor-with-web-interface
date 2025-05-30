package processimage

import(
	"image"
	"strings"
	envs "github.com/CaioLima42/ascii-conversor-with-web-interface/.envs"
)

func Gray2Ascii(img image.Gray16) string {
	var asciiImg strings.Builder
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			val := img.Gray16At(x, y).Y
			index := int(float32(val) * float32(envs.LENASCII) / float32(envs.MAXGRAYCOLOR))
			asciiImg.WriteRune(envs.ASCIILIST[index])
			asciiImg.WriteRune(' ')
		}
		asciiImg.WriteRune('\n')
	}
	return asciiImg.String()
}