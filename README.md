# 介绍

gtool 是一个使用 go 标准库写的用来简化操作的库  
目前主要实现

- file 文件处理
- 泛型 map
- 泛型 slice
- 并发泛型 map
- 并发泛型 slice

# 安装

```
go get github.com/Leviathangk/go-tool@latest
```

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

# generic

使用泛型构建的一些包，不含并发包

## gmap

是泛型 map，避免了转换  
有以下方法

- Get
- GetOrSet
- Set
- Exists
- Delete
- Range

```
m := gmap.NewMap[string, string]()
m.Set("key", "value")
fmt.Println(m.Get("key"))

m.Range(func(key string, value string) {
    fmt.Printf("%v -> %v\n", key, value)
})

fmt.Println(m.GetOrSet("key2", "value2"))
```

## gslice

### 针对已有数据处理

针对泛型切片数组的一系列操作

- Append
- Insert
- Remove
- RemoveByIndex
- Pop
- Shift
- Copy
- Range
- IndexOf
- Exists

示例

```
res := gslice.Insert([]int{1, 2, 3, 4, 5, 6}, 0, 8)
fmt.Println(res)

res = gslice.Remove([]int{1, 1, 5, 67, 1, 2, 1, 3, 1}, 1, 0)
fmt.Println(res)

res2 := gslice.Exists([]int{1, 2, 3, 4, 5, 6}, 7)
fmt.Println(res2)
```

### 针对新建数据

针对新建数据，可以使用泛型切片结构体

- Append
- Insert
- Remove：支持指定数量或全部
- RemoveByIndex
- Pop
- Shift
- Copy
- Range
- IndexOf
- Exists
- Get：支持负索引
- Length

```
newS := gslice.NewSlice[int]()
newS.Append(0, 1, 2, 3, 2, 34, 2)
fmt.Println(newS.Get(1))

newS.Remove(2, 0)
fmt.Println(newS.Slice)
```

# gsync

处理并发的模块

## SyncMap

并发泛型 map，含有以下方法

- Get
- Set
- GetOrSet
- Exists
- Delete
- Range

```
m := gsync.NewMap[int, int]()
m.Set(1, 1)
m.Set(2, 2)

v, ok := m.Get(1)
if ok {
    fmt.Println(v)
}
```

## SyncSlice

并发泛型 slice，含有以下方法

- Append
- Insert
- Remove：支持指定数量或全部
- RemoveByIndex
- Pop
- Shift
- Copy
- Range
- IndexOf
- Exists
- Get：支持负索引
- Length

```
s := gsync.NewSlice[int]()
s.Append(1, 2, 3, 4)
fmt.Println(s.Slice)
```