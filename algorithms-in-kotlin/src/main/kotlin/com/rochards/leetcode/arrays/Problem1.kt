package main.kotlin.com.rochards.leetcode.arrays

/*
* Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to
* target.
*
* Full description of the problem: https://leetcode.com/problems/two-sum/
* */

/*
* About the solution:
* - Time complexity is O(n^2) because of the nested loop;
* - Space complexity is O(1) because the extra space is constant regardless of the input size
* */
fun twoSum(nums: IntArray, target: Int): IntArray {
    for (i in nums.indices) {
        for (j in i + 1 until nums.size) { // it iterates of n-1 elements, but still O(n) time complexity
            if (nums[i] + nums[j] == target) {
                return intArrayOf(i, j)
            }
        }
    }
    return intArrayOf(0, 0)
}

/*
* About the solution:
* - Time complexity is O(n) because is not a nested loop anymore
* - Space complexity is O(n) because the extra space required depends on the input size
* */
fun twoSum2(nums: IntArray, target: Int): IntArray {

    val numberAndIndices = mutableMapOf<Int, Int>() // extra space depends on the input size

    for (i in nums.indices) {
        numberAndIndices[nums[i]] = i
    }

    for (i in nums.indices) {
        /*
        * target = a + b
        * b = target - a
        * */
        val addend = target - nums[i]
        if (numberAndIndices.containsKey(addend) && numberAndIndices[addend] != i) {
            return intArrayOf(numberAndIndices[addend]!!, i)
        }
    }

    return intArrayOf(0, 0)
}

fun main() {
    print("Example1: [2,7,11,15] -> ")
    twoSum(intArrayOf(2, 7, 11, 15), 9).forEach { print("$it ") }
    println()
    print("Example2: [3,2,4] -> ")
    twoSum(intArrayOf(3, 2, 4), 6).forEach { print("$it ") }
}