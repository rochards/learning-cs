package main.kotlin.com.rochards.leetcode.arrays

/*
* Given an integer array nums, return true if any value appears at least twice in the array, and return false if every
* element is distinct.
*
* Full description of the problem: https://leetcode.com/problems/contains-duplicate/description/


* About the solution:
* - Time complexity is O(n) because of the for loop
* - Space complexity is O(n) because the extra space required depends on the input size
* */
fun containsDuplicate(nums: IntArray): Boolean {
    val existingNumbers = mutableSetOf<Int>()
    for (n in nums) {
        if (existingNumbers.contains(n))
            return true

        existingNumbers.add(n)
    }

    return false
}

fun main() {
    println("Test #1: [1,2,3,1]. Result: ${containsDuplicate(intArrayOf(1,2,3,1))}")
    println("Test #2: [1,2,3,4]. Result: ${containsDuplicate(intArrayOf(1,2,3,4))}")
    println("Test #3: [1,1,1,3,3,4,3,2,4,2]. Result: ${containsDuplicate(intArrayOf(1,1,1,3,3,4,3,2,4,2))}")
}