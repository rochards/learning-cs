package com.rochards.search

/**
 * @return true if the number is in the array or false otherwise
 */
fun recursiveBinarySearch(numbers: IntArray, target: Int): Boolean {

    if (numbers.isEmpty())
        return false

    val midPosition = numbers.size / 2
    return when {
        target == numbers[midPosition] -> true
        target < numbers[midPosition] -> recursiveBinarySearch(numbers.sliceArray(0 until midPosition), target)
        else -> recursiveBinarySearch(numbers.sliceArray(midPosition + 1..numbers.lastIndex), target)
    }
}
fun main() {
    val numbers = intArrayOf(1, 2, 5, 10, 20, 51, 1092)
    println(recursiveBinarySearch(numbers, 20)) // should print true
    println(recursiveBinarySearch(numbers, 90)) // should print false
}