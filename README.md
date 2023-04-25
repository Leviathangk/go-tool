# 介绍

gtool 是一个使用 go 标准库写的用来简化操作的库

# file

文件处理库

## 文件读取

### 一次性读取

```
content, err := file.ReadAll("D:\\ddd.txt")
if err != nil {
    log.Fatalln(err)
}
fmt.Println(string(content))
```

### 逐行读取

```
err := file.ReadLine("D:\\ddd.txt", func(line []byte) {
    fmt.Println(string(line))
})
if err != nil {
    log.Fatalln(err)
}
```

### 逐字节读取

针对大文件，每次只读取指定字节长度

```
err := file.ReadByte("D:\\ddd.txt", 1024, func(b []byte) {
	fmt.Println(string(b))
})
if err != nil {
	log.Fatalln(err)
}
```

## 文件写入

### 一次性写入

```
err := file.WriteAll("D:\\ddd2.txt", []byte("demo"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
if err != nil {
	log.Fatalln(err)
}
```

### 逐字节写入

针对大文件的写入，每次写入都会刷新到硬盘

```
writer, err := file.WriteByte("D:\\ddd2.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
if err != nil {
	log.Fatalln(err)
}
defer writer.Close()

_ = writer.Write([]byte("1"))
_ = writer.Write([]byte("2"))
_ = writer.Write([]byte("3"))
```

## 文件删除

这只算是统一，os 包也很方便

### 删除文件

```
_ := file.RemoveFile("")
```

### 删除文件夹

```
_ := file.RemoveDir("")
```

## 文件移动

不分文件、文件夹，父路径不存在则创建（0666 模式）

```
err := file.MoveTo("D:\\dddd", "D:\\dddd2")
fmt.Println(err)
```
