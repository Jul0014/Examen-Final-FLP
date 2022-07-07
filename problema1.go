package main

import (  
  "os"; "image"; "image/jpeg"; "image/color" ; "time"; "log";  
  "path/filepath"; "fmt"; "strings"  
)

func check(err error) {  
  if err != nil {  
    panic(err)  
  }  
}

func main() { 

  //Obtenemos la imagen 1
  imgPath1 := "im1.jpg"
  f1, err1 := os.Open(imgPath1) 
  check(err1) 
  defer f1.Close() 

  img1, _, err := image.Decode(f1)

  size := img1.Bounds().Size()  
  rect := image.Rect(0, 0, size.X, size.Y)  
  wImg := image.NewRGBA(rect)

  //Obtenemos la imagen 2
  imgPath2 := "im2.jpg"
  f2, err2 := os.Open(imgPath2) 
  check(err2) 
  defer f2.Close() 

  img2, _, err2 := image.Decode(f2)
  
  start := time.Now()

  for x := 0; x < size.X; x++ { 
    for y := 0; y < size.Y; y++ { 
      r1, g1, b1, a1 := img1.At(x, y).RGBA()
      r1, g1, b1, a1 = r1>>8, g1>>8, b1>>8, a1>>8

      r2, g2, b2, a2 := img2.At(x, y).RGBA()
      r2, g2, b2, a2 = r2>>8, g2>>8, b2>>8, a2>>8 

      nR := (float64(r1) + float64(r2))/2
      nG := (float64(g1) + float64(g2))/2
      nB := (float64(b1) + float64(b2))/2

      if nR > 255 {
        nR = 255
      }
      if nG > 255 {
        nG = 255
      }
      if (nB) > 255 {
        nB = 255
      }
      c := color.RGBA{ 
        R: uint8(nR), G: uint8(nG), B: uint8(nB), A: uint8(a1), 
      } 
      wImg.Set(x, y, c)
    } 
  } 

  elapsed := time.Since(start)
  log.Printf("Tiempo que demora:  %s", elapsed)

  ext := filepath.Ext(imgPath1)  
  name := strings.TrimSuffix(filepath.Base(imgPath1), ext)     
  newImagePath := fmt.Sprintf("%s/%s_mezcla2%s",     filepath.Dir(imgPath1), name, ext)  

 
  fg, err := os.Create(newImagePath)  
  defer fg.Close()  
  check(err)  
  err = jpeg.Encode(fg, wImg, nil)  
  check(err)  
}