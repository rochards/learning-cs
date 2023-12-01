package com.rochards.search

/**
* @return the position of the number in the array or -1 when not found
* */
fun binarySearch(numbers: IntArray, target: Int): Int {

    var firstPosition = 0
    var lastPosition = numbers.lastIndex

    while (firstPosition <= lastPosition) {
        val midPosition = (firstPosition + lastPosition) / 2

        when {
            target == numbers[midPosition] -> return midPosition
            target < numbers[midPosition] -> lastPosition = midPosition - 1
            else -> firstPosition = midPosition + 1
        }
    }

    return -1
}

fun main() {
    val numbers = intArrayOf(1, 2, 5, 10, 20, 51, 1092)
    println(binarySearch(numbers, 20)) // should print 4
    println(binarySearch(numbers, 90)) // should print -1
}