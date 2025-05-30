package processimage

import( 
	"image"
	"image/color"
)

func NearestNeighborScaling(img image.Image, w, h int) image.Image {
	widthimg := img.Bounds().Max.X
	heightimg := img.Bounds().Max.Y
	MinPoint := image.Point{0, 0}
	MaxPoint := image.Point{w, h}
	scalonateImg := image.NewRGBA(image.Rectangle{MinPoint, MaxPoint})
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			srcX := i * widthimg / w
			srcY := j * heightimg / h
			r, g, b, a := img.At(srcX, srcY).RGBA()
			newPix := color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}
			scalonateImg.SetRGBA(i, j, newPix)
		}
	}
	return scalonateImg
}