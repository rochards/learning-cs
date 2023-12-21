package com.rochards.leetcode.arrays

/*
* Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that
* is missing from the array.
*
* The full description to the problem is on: https://leetcode.com/problems/missing-number/
* */

/*
* About the solution:
* - Time complexity is O(n)
* - Space complexity is O(n)
* */
fun missingNumber(nums: IntArray): Int {

    // it will represent a space complexity of O(n) because of the extra space needed
    val presentNumbers = mutableSetOf<Int>()

    // loop time complexity: O(n)
    for (num in nums) {
        presentNumbers.add(num) // ideally is this operation has a time complexity of O(1)
    }

    // loop time complexity: O(n)
    for (i in nums.indices) {
        // ideally the 'contains(i)' operation has a time complexity of O(1)
        if (!presentNumbers.contains(i)) {
            return i
        }
    }

    return nums.size // the worst case is when the missing number is the last one in the range [0, n]
}

fun main() {

    println("Test [3,0,1]: " + missingNumber(intArrayOf(3,0,1)))
    println("Test [9,6,4,2,3,5,7,0,1]: " + missingNumber(intArrayOf(9,6,4,2,3,5,7,0,1)))
}