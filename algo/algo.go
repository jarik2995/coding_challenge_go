package algo

import (
	"fmt"
	"regexp"
)

func Max(v1, v2 int) (r int) {
	if v1 > v2 {
		r = v1
	} else {
		r = v2
	}
	return
}

// Longest Word
// Have the function LongestWord(sen) take the sen parameter being passed and return the longest word in the string. If there are two or more words that are the same length, return the first word from the string with that length. Ignore punctuation and assume sen will not be empty. Words may also contain numbers, for example "Hello world123 567"
// Examples
// Input: "fun&!! time"
// Output: time
// Input: "I love dogs"
// Output: love
func LongestWord(s string) (r string) {
	re := regexp.MustCompile(`\w+`)
	words := re.FindAllString(s, -1)
	for _, w := range words {
		if len(w) > len(r) {
			r = w
		}
	}
	return r
}

// Given an array, rotate the array to the right by k steps, where k is non-negative.
// Example 1:
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
// Example 2:
// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]
func Rotate(nums []int, k int) []int {
	k = k % len(nums)
	if k == 0 {
		return nums
	}
	arr := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, arr)
	return nums
}

func Rotate1(nums []int, k int) []int {
	l := len(nums)
	k = k % l
	if k == 0 {
		return nums
	}
	for i := 0; i < k; i++ {
		for j := 1; j < l; j++ {
			nums[l-j], nums[l-j-1] = nums[l-j-1], nums[l-j]
		}
	}
	return nums
}

func Rotate2(nums []int, k int) {
	l := len(nums)
	arr := []int{}
	k = k % l
	for i := 0; i < l; i++ {
		arr = append(arr, nums[abs(l+i-k)%l])
	}
	for i := 0; i < l; i++ {
		nums[i] = arr[i]
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
// Example 1:
// Input: nums = [1,2,3,1]
// Output: true
// Example 2:
// Input: nums = [1,2,3,4]
// Output: false
// Example 3:
// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true
func ContainsDuplicate(nums []int) bool {
	// O(n)time O(n)space
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		_, exist := m[nums[i]]
		if exist {
			return true
		}
		m[nums[i]] = 1
	}
	return false
}

// if nums[n] <= n - 1
func ContainsDuplicateSpecialUseCase(nums []int) bool {
	// O(n)time O(1)space
	for i := 0; i < len(nums); i++ {
		if i == nums[i] {
			continue
		}
		if nums[nums[i]] == nums[i] {
			return true
		}
		tmp := nums[nums[i]]
		nums[nums[i]] = nums[i]
		nums[i] = tmp
		i--
	}
	return false
}

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.
// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Example 2:
// Input: nums = [0]
// Output: [0]
func MoveZeros1(nums []int) []int {
	ni := 0
	nz := 0
	count := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != 0 {
			ni++
		} else {
			if i > 0 && nums[i-1] == 0 {
				nz++
				continue
			}
			for j := 1; j <= ni; j++ {
				nums[i+j-1] = nums[i+j+nz]
				nums[i+j+nz] = 0
				count++
			}
			nz = 0
		}
	}
	fmt.Println(count)
	return nums
}

func MoveZeros2(nums []int) []int {
	nm := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nm++
			if nm > i {
				continue
			}
			nums[nm-1] = nums[i]
			nums[i] = 0
		}
	}
	return nums
}

func MoveZeros3(nums []int) {
	i, count := 0, 0
	for i < len(nums)-count {
		if nums[i] == 0 {
			nums = append(append(nums[:i], nums[i+1:len(nums)]...), 0)
			count++
			continue
		}
		i++
	}
}

// Desc: Given an integer array nums, find the subarray which has the largest sum and return its sum.
// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Input: nums = [1]
// Output: 1
// Input: nums = [5,4,-1,7,8]
// Output: 23
func MaxSubArray(nums []int) []interface{} {
	max := nums[0]
	cur := 0
	start := 0
	end := 0
	sub := nums
	for i := 0; i < len(nums); i++ {
		cur = cur + nums[i]
		if cur >= nums[i] && cur >= max {
			max = cur
			end = i
		}
		if nums[i] >= max && nums[i] >= cur {
			max = nums[i]
			cur = nums[i]
			start = i
			end = i
		}
		if nums[i] >= cur && nums[i] <= max {
			cur = nums[i]
			start = i
			end = i
		}
	}
	if end == len(nums)-1 {
		sub = nums[start:]
	} else {
		sub = nums[start : end+1]
	}
	return []interface{}{max, sub}
}

// MaxSubArray using Sum difference method
func MaxSubArraySumDiff(nums []int) int {
	sum := 0
	max := nums[0]
	p := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum-p > max {
			max = sum - p
		}
		if sum < p {
			p = sum
		}
	}
	return max
}

// MaxSubArray using Divide and Conquer method
func MaxSubArrayDivideAndConquer(nums []int) int {
	return maxSubArrayDivide(nums, 0, len(nums)-1)
}

func maxSubArrayConquer(nums []int, left, right int) int {
	mid := left + (right-left)/2
	leftSum := nums[mid]
	rightSum := nums[mid+1]

	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}
	return leftSum + rightSum
}

func maxSubArrayDivide(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	mid := left + (right-left)/2
	leftSum := maxSubArrayDivide(nums, left, mid)
	rightSum := maxSubArrayDivide(nums, mid+1, right)
	bothSum := maxSubArrayConquer(nums, left, right)
	return maxItemArray([]int{leftSum, rightSum, bothSum})
}

func maxItemArray(items []int) (max int) {
	max = items[0]
	for i := 0; i < len(items); i++ {
		if items[i] > max {
			max = items[i]
		}
	}
	return
}

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]
func SortedTwoSum(nums []int, target int) (arr []int) {
	nums = QuickSort(nums)
	for i := 0; i < len(nums); i++ {
		if nums[len(nums)-1]+nums[i] == target {
			arr = []int{len(nums) - 1, i}
			return
		}
		if nums[len(nums)-1]+nums[i] > target {
			arr = SortedTwoSum(nums[i:len(nums)-1], target)
			return
		}
	}
	return
}

func TwoSum(nums []int, target int) (arr []int) {
	inds := map[int]int{}
	for i := 0; i < len(nums); i++ {
		inds[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		_, exist := inds[diff]
		if exist {
			arr = []int{i, inds[diff]}
		}
	}
	return
}

func partition(nums []int) (arr []int, pos int) {
	pi := nums[len(nums)-1]
	i := 0
	for j := 0; j < len(nums)-1; {
		if nums[j] < pi {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			i++
		}
		j++
	}
	nums[len(nums)-1] = nums[i]
	nums[i] = pi
	arr = append(arr, nums...)
	pos = i
	return
}

// Quick sort algorithm
func QuickSort(nums []int) (arr []int) {
	if len(nums) <= 1 {
		arr = append(arr, nums...)
		return
	}
	arr, pos := partition(nums)
	arr1 := QuickSort(arr[:pos])
	arr2 := QuickSort(arr[pos+1:])
	arr = append(arr1, arr[pos])
	arr = append(arr, arr2...)
	return
}

// Merge two sorted arrays
func MergeSortedArrays(arr1, arr2 []int) (arr []int) {
	l1 := len(arr1)
	l2 := len(arr2)
	for i, j := 0, 0; i < l1 || j < l2; {
		if i >= l1 {
			arr = append(arr, arr2[j])
			j++
			continue
		}
		if j >= l2 {
			arr = append(arr, arr1[i])
			i++
			continue
		}
		if arr1[i] < arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}
	return
}

// Factorial
func Factorial(v int) int {
	if v == 1 {
		return v
	}
	return v * Factorial(v-1)
}

// Fibonacci O(N^2) very very slow
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Fibonacci O(n)
func FibonacciIterative(n int) int {
	arr := []int{0, 1}
	for i := 2; i <= n; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}
	return arr[n]
}

// ReverseString
func ReverseStringItr(s string) string {
	l := len(s)
	b := []byte(s)
	if l < 1 {
		return s
	}
	for i := 0; i < l/2; i++ {
		b[i], b[l-i-1] = b[l-i-1], b[i]
	}
	return string(b)
}

func ReverseStringRecursive(s string) string {
	if len(s) <= 1 {
		return s
	}
	return s[len(s)-1:] + ReverseStringRecursive(s[:len(s)-1])
}
