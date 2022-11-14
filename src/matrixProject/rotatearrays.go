package main

import "fmt"


func rotate(nums []int, k int) {
    k = k % len(nums)
    temp := append(nums[k:], nums[0:k]...)
    copy(nums, temp) // this actually writes to where nums points to
}

func main() {
    nums := []int{1,2,3,4,5,6,7}
    rotate(nums ,3)
    fmt.Println(nums)
}