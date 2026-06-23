os.Args[0] と os.Executable() の違い(Windows)
=============================================

| 経由             | os.Args[0]              | os.Executable()    |
|------------------|-------------------------|--------------------|
|EXEファイル自体   |本体のパス               | 本体のパス         |
|ハードリンク      |ハードリンクのパス       | ハードリンクのパス |
|シンボリックリンク|シンボリックリンクのパス | **本体のパス**     |

- ハードリンクは
  `mklink /H hardlink.exe .\study-go-arg0.exe`
  で作成
- シンボリックリンクは
  `mklink    symlink.exe  .\study-go-arg0.exe`
  で作成

```main.go
package main

import (
    "os"
)

func main() {
    println("os.Args[0]:", os.Args[0])
    exeName, err := os.Executable()
    if err != nil {
        println(err.Error())
        os.Exit(1)
    }
    println("os.Executable():", exeName)
}
```

```
> study-go-arg0.exe
os.Args[0]: study-go-arg0.exe
os.Executable(): C:\Users\hymko\src\study-go-arg0\study-go-arg0.exe

> hardlink.exe
os.Args[0]: hardlink.exe
os.Executable(): C:\Users\hymko\src\study-go-arg0\hardlink.exe

> symlink.exe
os.Args[0]: symlink.exe
os.Executable(): C:\Users\hymko\src\study-go-arg0\study-go-arg0.exe
```
