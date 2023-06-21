package main

import (
	"image"
	"image/color"
)

func imageFrom4bpp(decompressedData []byte) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, 128, 128), color.Palette{
		color.RGBA{
			R: 43,
			G: 18,
			B: 39,
			A: 255,
		},
		color.RGBA{
			R: 94,
			G: 26,
			B: 56,
			A: 255,
		},
		color.RGBA{
			R: 56,
			G: 61,
			B: 183,
			A: 255,
		},
		color.RGBA{
			R: 45,
			G: 43,
			B: 40,
			A: 255,
		},
		color.RGBA{
			R: 85,
			G: 11,
			B: 232,
			A: 255,
		},
		color.RGBA{
			R: 184,
			G: 188,
			B: 113,
			A: 255,
		},
		color.RGBA{
			R: 22,
			G: 22,
			B: 22,
			A: 255,
		},
		color.RGBA{
			R: 57,
			G: 58,
			B: 41,
			A: 255,
		},
		color.RGBA{
			R: 17,
			G: 13,
			B: 147,
			A: 255,
		},
		color.RGBA{
			R: 241,
			G: 242,
			B: 203,
			A: 255,
		},
		color.RGBA{
			R: 84,
			G: 188,
			B: 35,
			A: 255,
		},
		color.RGBA{
			R: 82,
			G: 43,
			B: 173,
			A: 255,
		},
		color.RGBA{
			R: 97,
			G: 40,
			B: 104,
			A: 255,
		},
		color.RGBA{
			R: 7,
			G: 0,
			B: 2,
			A: 255,
		},
		color.RGBA{
			R: 110,
			G: 155,
			B: 125,
			A: 255,
		},
		color.RGBA{
			R: 95,
			G: 173,
			B: 103,
			A: 255,
		}},
	)

	pixelsForAllSprites := makeZeroArray()

	// Loop every 32 bytes. 32 bytes make up an 8x8 graphic.
	index := 0
	for index < 4 {
		DDs := make([]byte, 8)
		CCs := make([]byte, 8)
		BBs := make([]byte, 8)
		AAs := make([]byte, 8)

		offset := index * 32

		for i := 0; i < 8; i++ {
			DDs[i] = decompressedData[offset+(2*i)]
			CCs[i] = decompressedData[offset+1+(2*i)]
			BBs[i] = decompressedData[offset+16+(2*i)]
			AAs[i] = decompressedData[offset+16+1+(2*i)]
		}

		for i := 0; i < 64; i++ {
			pixelsForAllSprites[index][i] = (DDs[i/8] >> (7 - (i % 8))) & 1
			pixelsForAllSprites[index][i] |= ((CCs[i/8] >> (7 - (i % 8))) & 1) << 1
			pixelsForAllSprites[index][i] |= ((BBs[i/8] >> (7 - (i % 8))) & 1) << 2
			pixelsForAllSprites[index][i] |= ((AAs[i/8] >> (7 - (i % 8))) & 1) << 3
		}

		index++
	}

	// fmt.Println(pixelsForAllSprites)

	for i := 0; i < len(pixelsForAllSprites); i++ {
		tileRow := i / 16
		tileColumn := i % 16

		pixelsForSprite := pixelsForAllSprites[i]
		for j := 0; j < len(pixelsForSprite); j++ {
			xOffset := j / 8
			yOffset := j % 8

			img.SetColorIndex(tileColumn*16+xOffset, tileRow*16+yOffset, pixelsForSprite[j])
		}
	}

	// each pixel should contain these bits:
	// abcd

	// The layout of bytes is as follows (XX is one byte of that letter):
	// DD CC DD CC DD CC DD CC DD CC DD CC DD CC DD CC BB AA BB AA BB AA BB AA BB AA BB AA BB AA BB AA
	// DD CC DD CC DD CC DD CC DD CC DD CC DD CC DD CC BB AA BB AA BB AA BB AA BB AA BB AA BB AA BB AA

	return img
}

func makeZeroArray() [][]uint8 {
	a := make([][]uint8, 256)
	for i := range a {
		a[i] = make([]uint8, 64)

		for j := range a[i] {
			a[i][j] = 0
		}
	}

	return a
}
