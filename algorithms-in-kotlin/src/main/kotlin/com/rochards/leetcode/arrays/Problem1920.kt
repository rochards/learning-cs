package main.kotlin.com.rochards.leetcode.arrays

/*
* Given a zero-based permutation nums (0-indexed), build an array ans of the same length where ans[i] = nums[nums[i]]
* for each 0 <= i < nums.length and return it.
* A zero-based permutation nums is an array of distinct integers from 0 to nums.length - 1 (inclusive).
*
* Full description to problem: https://leetcode.com/problems/build-array-from-permutation/description/
* */
fun buildArray(nums: IntArray): IntArray {
    val ans = IntArray(nums.size)
    for (i in nums.indices) {
        ans[i] = nums[nums[i]]
    }
    return ans
}

fun main() {
    buildArray(intArrayOf(0,2,1,5,3,4)).forEach { print("$it ") }
}