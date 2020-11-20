package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

const (
	dx = 500
	dy = 200
)

func create_image() {
	file, err := os.Create("test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	rgba := image.NewRGBA(image.Rect(0, 0, dx, dy))
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			rgba.Set(x, y, color.NRGBA{uint8(x % 256), uint8(y % 256), 0, 255})
		}
	}
	fmt.Println(rgba.At(400, 100))    //{144 100 0 255}
	fmt.Println(rgba.Bounds())        //(0,0)-(500,200)
	fmt.Println(rgba.Opaque())        //true，其完全透明
	fmt.Println(rgba.PixOffset(1, 1)) //2004
	fmt.Println(rgba.Stride)          //2000
	jpeg.Encode(file, rgba, nil)      //将image信息存入文件中
}

func add_water_mask() {
	//原始图片是test.jpg
	imgb, _ := os.Open("test.jpg")
	img, _ := jpeg.Decode(imgb)
	defer imgb.Close()

	wmb, _ := os.Open("btn_on.png")
	watermark, _ := png.Decode(wmb)
	defer wmb.Close()

	//把水印写到右下角，并向0坐标各偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-10, img.Bounds().Dy()-watermark.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	//生成新图片new.jpg，并设置图片质量..
	imgw, _ := os.Create("new_test.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})

	defer imgw.Close()

	fmt.Println("水印添加结束,请查看new_test.jpg图片...")
}

func main() {
	create_image()
	add_water_mask()
}
