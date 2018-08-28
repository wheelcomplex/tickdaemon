// main for tick daemon

package main

func earlyinit() {
	println("tick daemon earlyinit 2/2")
}

func earlyinit() {
	println("tick daemon earlyinit 1/2")
}

func init() {
	println("tick daemon init 1/2")
}

func init() {
	println("tick daemon init 2/2")
}

func main() {
	println("tick daemon main")
}
