package main

import (
	"fmt"
	"math/rand"
)

const (
	grayKeyName        = "AAA"
	grayKeyAnotherName = "AAA"
)

func IsEnableV3DirectConstantDefineTest() {
	print("hello")

}

func IsEnableV3BasicLitTest() {
	return

}

func IsEnableV3ConstantDefineTest() {
	return

}

func randomBool() bool {
	if rand.Int() > 10 {
		return true
	}
	return false
}

func BoolSimplifyAlwaysFalseTest() {
	// false && other => false

	fmt.Printf("something")
}

func BoolSimplifyAlwaysTrueTest() {
	// true || other => true
	return

}
