package processimage

import(
	"image/color"
	"image"
)


func RGB2GrayColorMean(c color.Color) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16((r + g + b) / 3)
}

func GrayScaleImage(img image.Image, conversor func(color.Color) uint16) image.Gray16 {
	newImage := image.NewGray16(image.Rectangle{img.Bounds().Min, img.Bounds().Max})
	for i := 0; i < img.Bounds().Max.X; i++ {
		for j := 0; j < img.Bounds().Max.Y; j++ {
			newImage.SetGray16(i, j, color.Gray16{conversor(img.At(i, j))})
		}
	}
	return *newImage
}