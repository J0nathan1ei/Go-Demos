package main


func main() {
	//初始化两个slice
	s1 := make([]int, 3, 4)
	s2 := s1[: 2]

	s2[0] ++
	println(s1[0] == s2[0]) //true

	s1 = append(s1, 0)
	s2[0] ++
	println(s1[0] == s2[0]) //true

	s1 = append(s1, 0) //s1会扩容，并指向新空间
	s2[0] ++
	println(s1[0] == s2[0]) //false
}