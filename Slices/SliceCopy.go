package main

import (
	"fmt"

)

func linked(int, int, int){
	s1 := []int{1,2,3,4,5}
	s2 := s1
	s3 := s1[:]
	s1[3] = 99

	return s1[3], s2[3], s3[3]
}

func noLink() (int, int){
	s1 := []int{1,2,3,4,5}
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capLinked() (int, int){
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1,2,3,4,5
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capNoLink() (int, int){
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s3[0], 
}