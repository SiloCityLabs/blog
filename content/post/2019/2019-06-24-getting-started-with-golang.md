---
title: Getting Started with Golang
author: Luis Rodriguez
type: post
date: 2019-06-24
url: /2019/06/24/getting-started-with-golang
categories:
  - Projects
  - Golang
tags:
  - go
  - tutorial
  - beginner
  - vscode
  - functions
  - struct
  - git
  - golang
  - array
  - slices

---

## Go Applications

A Go app can vary from 3 to millions of lines of code. They should also be written with one or multiple files with the extension `.go` You can use any text editor to write a Go program like nano, but our favorite and most supported is Visual Studio Code.

## Setting up a Go Environment

For all of our tutorials we will be using VSCode to help us ease into the language. The reason for this is due to its ability to automatically handle imports and autocomplete. This will make it easy to learn specific function names quickly without having to know everything about Golang.

Download and install [Visual Studio Code](https://code.visualstudio.com/)

Download and install the latest 64-bit [Go installer](https://golang.org/dl/) from "Featured downloads." Easy for Windows you can just install the msi and your all set.

If you are a new developer it is also recommended to have [git installed](https://git-scm.com/). We wont be using it in tutorial but it may be used in future tutorials. You can also use it to make edits to this blog post from our repository.

### Installing on Unix/Linux/Mac OS X

Extract the archive into /usr/local/go using the following commands

```
tar -C /usr/local -xzf go*.tar.gz
```

Adding /usr/local/go/bin to the PATH environment variable.

**Linux**

```
export PATH=$PATH:/usr/local/go/bin
```

**Mac**

```
export PATH=$PATH:/usr/local/go/bin
```

**FreeBSD**

```
export PATH=$PATH:/usr/local/go/bin
```

Restart any open command prompts for the change to take effect. Verify the installation by running the following command:

```
go version
```


## Go and VSCode

We will start by making a quick hello world to verify that Go and VSCode are both setup properly. During this you will see notifications popup from VSCode to install go plugins for [linting](https://www.google.com/search?q=linter) and autocomplete.

On the first start screen hit the Open folder button and create a folder for your projects under `C:\Users\{username}\Go\src\` (Windows) or `/home/{username}/go/src` (Linux). Inside you will want to add a folder for your hello world project. We will be calling it `hello`. Hit new File and type `hello.go`. Almost Immediately I was presented with a plugin installer for Go.

![Golang Install Plugin](/uploads/2019/06/golang-extension.png)

Moments after it installs it recommends other plugins. I just hit `install all`. This will help you get started with go much quicker than ussual by adding many features lacking by default in the editor.

![Golang Install All](/uploads/2019/06/golang-plugins-all.png)

## Packages

Every go program is made from packages, Lets start a program using package main like all other programs.

```go
package main

func main() {
	
}

```

We will be calling a function instead of the package `fmt` that will be imported automatically by vscode. Type `fmt.Println("Hello World")` inside of your `func main` brackets then hit `ctrl + s`. As you typed that out im sure you noticed the autocomplete gave you 3 other print functions to use that belong to package `fmt`. Your code should now look like below:

```go
package main

import "fmt"

func main() {
	fmt.Println()
}

```

## Imports

To use other packages in go you have to import them into your code. As I mentioned above vscode will handle this for you. In certain instances you may need to know how to do it by hand. This is how imports can be written:

```go
import "fmt"
import "math"
```
or
```go
import(
 "fmt"
 "math
)
```

## Functions

Basic functions in Go can take zero or more arguments. The example below takes two strings and returns back one string. The [type](https://tour.golang.org/basics/11) has to be declared after each parameter variable name. Lets take the two strings and return them back combined as one.

```go
func functionName(parameter1 string, param2 string) (returned string) {
	// Body of the function, code goes here.
	return parameter1 + param2
}
```

You should be able to add this function to your hello world program and call it like so.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(functionName("Hello ", "Again"))
}

func functionName(parameter1 string, param2 string) (returned string) {
	// Body of the function, code goes here.
	return parameter1 + param2
}
```

Run your go program by opening a new terminal from the top menu bar. By default vscode selects powershell. You can change that by selecting the dropdown and clicking `select default shell`. You will want `Git Bash` for most operating systems. Now type the following to get your result:

```go
go run hello.go
```

Output:
```
Hello World
Hello Again
```

## Variables

A variable definition tells the compiler where and how much to create storage for the variable. It specifies the type and contains a list of one or more variables of that type like below:

```go
var variableName string
```

`Var` declares the list of variables and just like functions the type "string" goes last.

`Var` statements can be at package or function level. We see both in the example below.

```go
package main

import "fmt"

var golang, keys, passwords bool

func main() {
  var x int

  fmt.Println(x, golang, keys, passwords)
}
```

If you run this your result will be the defaults of string and bool:

```
0 false false false
```

## For Loop

A `for` loop is a repetition structure that allows you to efficiently write a loop that needs to execute a specific amount of times. The syntax is as follows:

```go
package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

```

Go only has one looping construct. There are no other looping contructs in go like other languages. The basic for loop has three components separated by semicolons:

 - init statement - executed before the first iteration
 - condition expression - evaluated before every iteration
 - post statement - executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement. The loop will stop iterating once the boolean condition evaluates to false.

The result of the code will be

```
10
```

You can see more loop examples [here](https://gobyexample.com/for).

## Arrays

Go also provides a data structure called `array`. Arrays can store a fixed-size collection of elements of the same type. An array is used to store a collection of data, but it is often more useful to think of an array as a collection of variables of the same type.

Lets declare an array of elements with the type string and a size of 10.

```go

var usernames [10] string

```

You can also initialize an array in Go with a list of items. Go will calculate the size automatically at compile time. Arrays cannot be resized later on but dont worry because it provides convenient ways of working with them.

```go
var usernames = []string{"gocrazy","gorungo","going2fast"}
```

Lets spit out usernames and see our result:

```
$ go run hello.go
[gocrazy gorungo going2fast]
```

## Slices

A slice is an abstraction over Go Array. As Go arrays allow you to define type of variables that can hold several data items of the same kind but it does not allow you to increase the size dynamically or get a sub array. It provides many helper functions required on Array and are widely used in Go across the community.

To define a slice you can declare it as an array without specifying size. You can also use the `make` function to create one.

```go
var numbers []int /* a slice with an unspecified size */
numbers2 := make([]int, 5, 5) /* a slice of length 5 and capacity 5 */
```

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

This expression creates a slice of the first five elements of the array a:

```go
a[0:5]
```

Here is an example:

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)
}
```

The result below:

```
$ go run hello.go
[2 3 5 7 11 13]
[3 5 7]
```

## Structure

A `structure` is a user defined type available in Go that allows you to combine data items of different kinds. A struct statement defines a new data type for your program and binds the name. Once a struct is defined it can be used like any other type. Lets defined one, populate it and print it out.

```go
package main

import "fmt"

type settings struct {
	Hostname string
	Port     string
}

func main() {
	mywebsite := settings{"silocitylabs.com", "80"}
	fmt.Println(mywebsite)

	var mysecondwebsite settings
	mysecondwebsite.Hostname = "blog.silocitylabs.com"
	mysecondwebsite.Port = "80"
	fmt.Println(mysecondwebsite)
}

```

Result

```
$ go run hello.go
{silocitylabs.com 80}
{blog.silocitylabs.com 80}
```

## Third party packages

Just like built in Go packages and your own packages. You can import packages from other places. Lets start by trying to find one online that generates random words. I searched "golang random words package" on Google and came up with the following https://github.com/tjarratt/babble. Before we begin to import this package and use it like its documentation states. We need to have it downloaded on our machine.


In your terminal type the following
```
go get github.com/tjarratt/babble
```

Now lets go ahead and use it just like the example below:

```go
package main

import (
	"fmt"

	"github.com/tjarratt/babble"
)

func main() {
	babbler := babble.NewBabbler()
	fmt.Println(babbler.Babble()) // excessive-yak-shaving (or other phrase)

	// optionally set your own separator
	babbler.Separator = " "
	fmt.Println(babbler.Babble()) // "hello world" (or other phrase)

	// optionally set the number of words you want
	babbler.Count = 1
	fmt.Println(babbler.Babble()) // antibiomicrobrial (or other word)
}
```

Running this will result in 3 separate random phrases every time. You can look at the source for their code on github from the same link to see how their packages work for further documentation.

## Homework

Try writing a go program by searching how to open a txt file. With a list in the file then println it out to the screen. We can help you review your code at our [meetup sessions](https://www.meetup.com/Buffalo-GoLang-Meetup-Group/).
