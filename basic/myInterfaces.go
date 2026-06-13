package main

import (
	"fmt"
)

type IHeadphone interface {
	Play()
	VolumeUp()
	VolumeDown()
}

// apple
type AppleHeadphone struct {}

func ( a AppleHeadphone) Play(){ fmt.Println("Apple: Play")}
func ( a AppleHeadphone) VolumeUp(){ fmt.Println("Apple: + Volume")}
func ( a AppleHeadphone) VolumeDown(){ fmt.Println("Apple: - Volume")}

// sony
type SonyHeadphone struct {}
func ( s SonyHeadphone) Play(){ fmt.Println("Sony: Play")}
func ( s SonyHeadphone) VolumeUp(){ fmt.Println("Sony: + Volume")}
func ( s SonyHeadphone) VolumeDown(){ fmt.Println("Sony: - Volume")}

// cheaper headphone
type SamsungHeadphone struct {}
func ( c SamsungHeadphone) Play(){ fmt.Println("Samsung: Play")}
func ( c SamsungHeadphone) VolumeUp(){ fmt.Println("Samsung: + Volume")}
func ( c SamsungHeadphone) VolumeDown(){ fmt.Println("Samsung: - Volume")}

// A function for all of them - It's important otherwise the interface is useless
// Use the interface
func Test(h IHeadphone){
	h.Play()
	h.VolumeUp()
	h.VolumeDown()
}

func MyInterfaces() {

	Test(AppleHeadphone{})
	fmt.Println()
	Test(SonyHeadphone{})
}