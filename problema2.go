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
  //x
  xC := 0.75;
  //Obtenemos la imagen 1
  imgPath := "im1.jpg"
  f, err := os.Open(imgPath) 
  check(err) 
  defer f.Close() 

  img, _, err := image.Decode(f)

  size := img.Bounds().Size()  
  rect := image.Rect(0, 0, size.X, size.Y)  
  wImg := image.NewRGBA(rect)

  //Obtenemos la imagen 2
  imgPath2 := "im2.jpg"
  f2, err2 := os.Open(imgPath2) 
  check(err2) 
  defer f2.Close() 

  img2, _, err2 := image.Decode(f2)
  
  start := time.Now()
  // loop though all the x  
  for x := 0; x < size.X; x++ { 
    // and now loop thorough all of this x's y 
    for y := 0; y < size.Y; y++ { 
      r, g, b, a := img.At(x, y).RGBA()
      r, g, b, a = r>>8, g>>8, b>>8, a>>8

      r2, g2, b2, a2 := img2.At(x, y).RGBA()
      r2, g2, b2, a2 = r2>>8, g2>>8, b2>>8, a2>>8 

      nR := xC * float64(r) + (1-xC) * float64(r2)
      nG := xC * float64(g) + (1-xC) * float64(g2)
      nB := xC * float64(b) + (1-xC) * float64(b2)

      if uint8(nR) > 255 {
        nR = 255
      }
      if uint8(nG) > 255 {
        nG = 255
      }
      if uint8(nB) > 255 {
        nB = 255
      }
        
      c := color.RGBA{ 
        R: uint8(nR), G: uint8(nG), B: uint8(nB), A: uint8(a), 
      } 
      wImg.Set(x, y, c)
    } 
  } 

  elapsed := time.Since(start)
  log.Printf("Blending %s", elapsed)

  ext := filepath.Ext(imgPath)  
  name := strings.TrimSuffix(filepath.Base(imgPath), ext)  

  sxC := fmt.Sprintf("%.2f", xC)
  newImagePath := fmt.Sprintf("%s/%s_%s_blended%s", filepath.Dir(imgPath), name, sxC, ext)  
 
  fg, err := os.Create(newImagePath)  
  defer fg.Close()  
  check(err)  
  err = jpeg.Encode(fg, wImg, nil)  
  check(err)  
}