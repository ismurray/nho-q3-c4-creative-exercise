package main

import "github.com/fogleman/gg"

const (
	imageSize       = 2000
	petalSize       = 2048
	faceRadius      = 300
	faceCenter      = imageSize / 2
	lineThickness   = 10
	lensWidth       = faceRadius / 2
	lensHeight      = faceRadius / 4
	lensStartY      = (imageSize / 2) - (faceRadius / 2.25)
	leftLensStartX  = (imageSize / 2) - (faceRadius / 1.75)
	rightLensStartX = (imageSize / 2) + (faceRadius / 10)
)

func drawBackground(dc *gg.Context) {
	// draw petal background
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), petalSize/2, petalSize/2)
		dc.DrawEllipse(petalSize/2, petalSize/2, petalSize*7/16, petalSize/8)
		dc.Fill()
		dc.Pop()
	}
}

func main() {
	dc := gg.NewContext(imageSize, imageSize)
	drawBackground(dc)

	// draw base of face
	dc.DrawEllipse(faceCenter, faceCenter, faceRadius*0.95, faceRadius)
	// flesh color
	dc.SetRGBA255(229, 207, 210, 255)
	dc.Fill()
	// outline
	dc.DrawEllipse(faceCenter, faceCenter, faceRadius*0.95, faceRadius)
	// black
	dc.SetRGBA255(0, 0, 0, 255)
	dc.SetLineWidth(lineThickness)
	dc.Stroke()

	// draw eyes
	// left iris
	dc.DrawCircle(leftLensStartX+(lensWidth/2), lensStartY+(lensHeight/2), faceRadius/10)
	// right iris
	dc.DrawCircle(rightLensStartX+(lensWidth/2), lensStartY+(lensHeight/2), faceRadius/10)
	// eye color (bottle green)
	dc.SetRGBA255(0, 106, 78, 255)
	dc.Fill()
	// left pupil
	dc.DrawCircle(leftLensStartX+(lensWidth/2), lensStartY+(lensHeight/2), faceRadius/25)
	// right pupil
	dc.DrawCircle(rightLensStartX+(lensWidth/2), lensStartY+(lensHeight/2), faceRadius/25)
	// pupil color (black)
	dc.SetRGBA255(0, 0, 0, 255)
	dc.Fill()

	// draw nose
	dc.MoveTo(faceCenter, lensStartY)
	dc.LineTo(faceCenter-(faceRadius/8), faceCenter+(faceRadius/4))
	dc.LineTo(faceCenter+(faceRadius/8), faceCenter+(faceRadius/4))
	dc.LineTo(faceCenter+(faceRadius/10), faceCenter+(faceRadius/10))
	dc.SetLineWidth(lineThickness)
	dc.Stroke()

	// draw septum
	dc.DrawEllipticalArc(faceCenter, faceCenter+(faceRadius/4), 15, 15, 0, 3.14)
	// black
	dc.SetRGBA255(0, 0, 0, 255)
	dc.SetLineWidth(lineThickness / 2)
	dc.Stroke()

	// draw glasses
	// left frame
	dc.DrawRoundedRectangle(leftLensStartX, lensStartY, lensWidth, lensHeight, faceRadius/20)
	// left arm
	dc.LineTo(leftLensStartX-(lensWidth*0.75), lensStartY+(lensHeight/2))
	// right frame
	dc.DrawRoundedRectangle(rightLensStartX, lensStartY, lensWidth, lensHeight, faceRadius/20)
	// right arm
	dc.MoveTo(rightLensStartX+lensWidth, lensStartY)
	dc.LineTo(rightLensStartX+lensWidth+(lensWidth*0.75), lensStartY+(lensHeight/2))
	// bridge
	dc.MoveTo(leftLensStartX+(lensWidth/2), lensStartY)
	dc.LineTo(rightLensStartX+(lensWidth/2), lensStartY)
	// black
	dc.SetRGBA255(0, 0, 0, 255)
	dc.SetLineWidth(lineThickness)
	dc.Stroke()

	// draw eyebrows
	// left eyebrow
	dc.DrawRoundedRectangle(leftLensStartX+lensWidth*0.15, lensStartY-faceRadius/10, lensWidth*0.75, lineThickness*1.5, lineThickness/2)
	// right eyebrow
	dc.DrawRoundedRectangle(rightLensStartX+lensWidth*0.15, lensStartY-faceRadius/10, lensWidth*0.75, lineThickness*1.5, lineThickness/2)
	// black
	dc.SetRGBA255(0, 0, 0, 255)
	dc.Fill()

	// draw hair
	// curls
	size := float64(faceRadius / 15)
	y := lensStartY - faceRadius/3.8
	for x := (faceCenter - faceRadius*0.75); x < (faceCenter + faceRadius*0.75); x += faceRadius / 8 {
		// blond hair curls
		dc.DrawCircle(x, y, size)
		// hair color (dirty blond)
		dc.SetRGBA255(226, 210, 158, 255)
		dc.Fill()
		// hair outlines
		dc.DrawCircle(x, y, size)
		// black
		dc.SetRGBA255(0, 0, 0, 255)
		dc.SetLineWidth(lineThickness / 4)
		dc.Stroke()
		size += faceRadius / 30 * 0.8
		y -= 7
	}

	// draw smile
	dc.DrawEllipticalArc(faceCenter, faceCenter+(faceRadius/2.75), 40, 35, 0, 3.14)
	// black
	dc.SetRGBA255(0, 0, 0, 255)
	dc.Fill()

	dc.SavePNG("ian-self-portrait.png")
}
