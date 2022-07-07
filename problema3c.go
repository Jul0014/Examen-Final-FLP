package main

import (
	"os"; "image"; "time"; "log"; "image/jpeg";"sync"
)

func check(err error) {  
  if err != nil {  
    panic(err)  
  }  
}

func getImageFromFilePath(nombre string) (image.Image, error) {
    f, err := os.Open(nombre)
    check(err)
    defer f.Close()
    image, err := jpeg.Decode(f)
    return image, err

}

func main() {
	var r [256] int
  for i:= 0; i < 256; i ++{
    r[i] = 0    
  }

  var b [256] int
  for i:= 0; i < 256; i ++{
    b[i] = 0    
  }

  var g [256] int
  for i:= 0; i < 256; i ++{
    g[i] = 0    
  }

  img, err := getImageFromFilePath("im1.jpg")
  check(err)

  size := img.Bounds().Size()

  wg := new(sync.WaitGroup)
  
  start := time.Now()
  for y := 0; y < size.Y; y++ { 
    wg.Add(1)
    y:= y
    go func(){
      for x := 0; x < size.X; x++ { 
        rC, gC, bC, a := img.At(x, y).RGBA()
        rC, gC, bC, a = rC>>8, gC>>8, bC>>8, a>>8
  
        r[rC] += 1
        g[gC] += 1
        b[bC] += 1
      }
      defer wg.Done()//importante
    }()
  } 

  elapsed := time.Since(start)
  log.Printf("Blending %s", elapsed)

  for i := 0; i < 256; i ++{
    fmt.Printf("(%d) R: %d G: %d B: %d \n", i, r[i], g[i], b[i])
  }
}