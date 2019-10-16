---
title: Arm 64bit vs Xeon x86_64 performance on AWS
author: Luis Rodriguez
type: post
date: 2019-10-16T22:36:25+00:00
url: /post/2019/10/16/arm-vs-xeon-aws/
categories:
  - Projects
tags:
  - internet
  - aws
  - xeon
  - arm
  - arm_64
  - x86_64
  - benchmark
  - performance
  - comparison

---

The reason for this test was to determine if arm was a viable option on aws. I noticed that arm had higher network throughputs and faster disk speeds on the aws pages. So I decided I would compare their performance using golang tests. The entire test cost me $0.68 using a quick spot instance.


The comparison was done on two very similar instances of aws servers. On paper they would almost seem to function the same but arm was just a few cents more, but lets take a look on how they performed.

### Linux AMD64 (xeon)
```
BenchmarkSprint-2       20000000                74.8 ns/op             5 B/op          1 allocs/op
BenchmarkConcat-2       20000000                86.6 ns/op            16 B/op          1 allocs/op
BenchmarkMath-2         10000000               131 ns/op              16 B/op          2 allocs/op
BenchmarkCompare-2      2000000000               0.37 ns/op            0 B/op          0 allocs/op
BenchmarkIntString-2    10000000               120 ns/op              16 B/op          2 allocs/op
BenchmarkReplace-2      10000000               142 ns/op              64 B/op          2 allocs/op
BenchmarkSplit-2        10000000               150 ns/op              64 B/op          1 allocs/op
BenchmarkJoin-2         20000000                87.1 ns/op            32 B/op          1 allocs/op
BenchmarkJSON-2           300000              4019 ns/op             576 B/op         17 allocs/op
BenchmarkLargeJSON-2         100          14204481 ns/op         2980990 B/op      52214 allocs/op
BenchmarkByte-2          2000000               619 ns/op             112 B/op          3 allocs/op
BenchmarkRegex-2          200000              8330 ns/op           39038 B/op         30 allocs/op
BenchmarkMd5-2           2000000               664 ns/op              80 B/op          3 allocs/op
BenchmarkSha1-2          2000000               742 ns/op             112 B/op          3 allocs/op
BenchmarkSha256-2        2000000               975 ns/op             144 B/op          3 allocs/op
BenchmarkUUIDv4-2        1000000              1643 ns/op             160 B/op          4 allocs/op
```

### Linux ARM64
```
BenchmarkSprint-2       10000000               144 ns/op               5 B/op          1 allocs/op
BenchmarkConcat-2       10000000               180 ns/op              16 B/op          1 allocs/op
BenchmarkMath-2          5000000               270 ns/op              16 B/op          2 allocs/op
BenchmarkCompare-2      2000000000               0.87 ns/op            0 B/op          0 allocs/op
BenchmarkIntString-2     5000000               246 ns/op              16 B/op          2 allocs/op
BenchmarkReplace-2       3000000               427 ns/op              64 B/op          2 allocs/op
BenchmarkSplit-2         5000000               245 ns/op              64 B/op          1 allocs/op
BenchmarkJoin-2         10000000               159 ns/op              32 B/op          1 allocs/op
BenchmarkJSON-2           200000              9944 ns/op             576 B/op         17 allocs/op
BenchmarkLargeJSON-2          50          25567425 ns/op         2980985 B/op      52214 allocs/op
BenchmarkByte-2          1000000              1084 ns/op             112 B/op          3 allocs/op
BenchmarkRegex-2          100000             13893 ns/op           39036 B/op         30 allocs/op
BenchmarkMd5-2           1000000              1147 ns/op              80 B/op          3 allocs/op
BenchmarkSha1-2          1000000              1189 ns/op             112 B/op          3 allocs/op
BenchmarkSha256-2        1000000              1393 ns/op             144 B/op          3 allocs/op
BenchmarkUUIDv4-2         500000              2541 ns/op             160 B/op          4 allocs/op
```

As you can see arm actually performed much slower on most operations. Not even a negligable amount to compensate for the network speed gain you would get from using arm.

### Install script (run.sh)
```
#!/bin/bash

add-apt-repository ppa:gophers/archive
apt-get update
apt-get install golang-1.11-go iperf3 hdparm
/usr/lib/go-1.11/bin/go
/usr/lib/go-1.11/bin/go get -u github.com/google/uuid

echo '' > results.txt

echo "Running benchmarks on cpu/memory" >> results.txt
echo '' >> results.txt

/usr/lib/go-1.11/bin/go test -benchtime=1m -bench=. -benchmem >> results.txt


echo "Running benchmarks on network" >> results.txt
echo '' >> results.txt

for (( c=1; c<=5; c++ ))
do  
   iperf3 -c iperf.he.net >> results.txt
done

echo 'Running benchmarks on disk' >> results.txt
echo '' >> results.txt

hdparm -Tt /dev/xvda1 >> results.txt
```

### Test script (sprint_test.go)
```
package benchmark

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello" + " hello")
	}
}

func BenchmarkMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		added := 123123 + 1298685
		fmt.Sprint(added)
	}
}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if 312312 == 354351 {
			fmt.Sprint("dont remove me compiler")
		}
		if "312312" == "adadsad334" {
			fmt.Sprint("dont remove me compiler")
		}
		if false == true {
			fmt.Sprint("dont remove me compiler")
		}
	}
}

func BenchmarkIntString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf(strconv.Itoa(81312391))
	}
}

func BenchmarkReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Replace("aefaewdsfasdfasdf", "as", "jfiuwieiihoijpj", 1)
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Split("aefaew\ndsfa\nsdf\nasdf", "\n")
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join([]string{"asdasdas", "asdasdasd", "asdasdasd"}, "\n")
	}
}

type response1 struct {
	Page   int
	Fruits []string
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {

		res1D := &response1{
			Page:   1,
			Fruits: []string{"apple", "peach", "pear"}}
		res1B, _ := json.Marshal(res1D)

		res := response1{}
		json.Unmarshal(res1B, &res)
		fmt.Sprint(res)

	}
}

func BenchmarkLargeJSON(b *testing.B) {
	jString, _ := ioutil.ReadFile("test.json")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res := response1{}
		json.Unmarshal(jString, &res)
		fmt.Sprint(res)
	}
}

func BenchmarkByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var blah = "asdfsdnfanjkdsfn"
		fmt.Sprint([]byte(blah))
	}
}

func BenchmarkRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
		fmt.Sprint(matched, err)
	}
}

func BenchmarkMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5 := md5.Sum([]byte("Foo"))
		fmt.Sprintf("%x\n", md5)
	}
}

func BenchmarkSha1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha1 := sha1.Sum([]byte("Foo"))
		fmt.Sprintf("%x\n", sha1)
	}
}

func BenchmarkSha256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha256 := sha256.Sum256([]byte("Foo"))
		fmt.Sprintf("%x\n", sha256)
	}
}

func BenchmarkUUIDv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuidv4 := uuid.New()
		fmt.Sprintf("%x\n", uuidv4)
	}
}
```

### Other test file (test.json)
```
{
  "page": 1,
  "fruits": [
    "culpa",
    "mollit",
    "id",
    "nulla",
    "ullamco",
    "proident",
    ...  26000 lines concatenated for browser ...
    "ullamco",
    "dolor",
    "id",
    "minim",
    "velit",
    "ad",
    "non"
  ]
}
```

With that being said, I dont think I would be switching to arm anytime soon. The only good reason to use it would be for static file hosting but it would be much cheaper to use an s3 bucket or something similar.