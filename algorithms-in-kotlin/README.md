# The content listed here is a summary of the course [Algorithms and Data Structures Tutorial - Full Course for Beginners](https://www.youtube.com/watch?v=8hly31xKli0)

- [Implementation of binary search](src/main/kotlin/com/rochards/search)

## About data structures

### Arrays and Lists

An **array** is one of the most basic data structure that it is probably implemented in every programming language.  
There are a lot of resources available to learn about arrays. You can find more information about arrays in Kotlin [here](https://kotlinlang.org/docs/arrays.html). However, keep in mind the following information:
- Arrays are not resizable;
- Accessing an element by its index has a constant time complexity, O(1);
- Changing the value at a specific index it's also O(1).

A **list** is a more complex data structure but works in a way similar to an array in Kotlin. Click [here](https://kotlinlang.org/docs/collections-overview.html#list) to learn more.  
Here are some key points to remember about lists in Kotlin:
- There are **mutable** and **immutable** implementations of list:
  - mutable: in contrast to arrays, these kind of lists are resizable, and you're allowed to change any element within the list and add new elements. In Kotlin, this is achieved through the `MutableList` interface and its default implementation, `ArrayList`. At the end of the day, an `ArrayList` is implemented as an array with the capability of being resizable.
  - immutable: the `List` interface represents an immutable list. Like an array, it's non-resizable but in contrast you can't change any element within it. In other words, it is a read-only list.
- You also can access elements by index;
- Accessing any element has a time complexity of O(1);
- Changing the value at a specific index it's also O(1);
- Adding a new element to the end of the list (only applicable for the mutable ones) has a time complexity of O(1). Occasionally, resizing the list to fit new elements may be necessary, which takes O(n) because it involves creating a new array and copying the existing elements to it. However, this resizing operation happens so rarely that it's considered an amortized time of O(1). Head to [this link](https://medium.com/@satorusasozaki/amortized-time-in-the-time-complexity-of-an-algorithm-6dd9a5d38045) for a more in-depth explanation;
- (todo): talk about inserting element in any position