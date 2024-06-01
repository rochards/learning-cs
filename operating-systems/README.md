# Fundamentals of Operating Systems

My notes from [Fundamentals of Operating Systems](https://www.udemy.com/course/fundamentals-of-operating-systems) course.

## Table of content
- [Introduction](#section-1-and-section-2)
- [The anatomy of a process](#section-3-the-anatomy-of-a-process)
- [Terminal commands for linux used through the course](#terminal-commands-for-linux-used-through-the-course)

## Section 1 and Section 2

They are just an introduction and an overview of what the course is about and why we may need an Operating System.

## Section 3: The anatomy of a process

### Process vs Program

A process is a program in motion/execution.

**Program**:
- Lives on disk;
- Code that is compiled and linked for a CPU;
- Only works on that CPU architecture when compiled;
- Produces executable file program, the so-called process;

Look at the format of a program in Linux systems:  
<div align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/7/77/Elf-layout--en.svg/800px-Elf-layout--en.svg.png" alt="ELF (Executable and Linkable Format) program format" width="300">
</div>
ELF stands for Executable and Linkable Format.

:bulb: there is a simple code written in C in `code-section3` folder with the assembly `.s` file of this program compiled for a **x86_64** machine.

**Process**:
- Lives in memory;
- Is the program in execution;
- Uniquely identified by an ID, aka Process ID or just PID;
- Has a Instruction Pointer, aka Program Counter, that points to the memory address for the current instruction to be executed;
- Has PCB (Process Control Block), a region in kernel's memory that contains metadata information about the process.

Take a look on how a process lives in memory:  
<div align="center">
  <img src="images/process_anatomy.drawio.svg" alt="Anatomy of a process">
</div>
the arrows above mean that the Stack "grows down" and the Heap "grows up".


### Simple Process Execution

Initial considerations:
- In the image below we're considering that each instruction is at most 4 bytes long (32 bits architecture);
- `mov`, `add`, `str` are instructions;
- `r*`, `pc` (program counter), `ir` (instruction register) are CPU registers;
- the ELF header in the program tells the OS where to start the program;
- for simplicity, we're not considering the L* caches;

<div align="center">
  <img src="images/simplified-process-execution.gif" alt="Anatomy of a process">
</div>
the the execution of the remaining instructions happen the same way.

### The Stack

- The **stack pointer (sp)** is a register in the CPU the always points to the end of the stack.
<div align="center">
  <img src="images/the-stack-stack.svg" alt="Anatomy of a process">
</div>

- The **base pointer (bp)** is also a register in the CPU to mark the start of the function; 
- *bp* always points to the last function in the stack. In the image above, *bp* would be pointing to the beginning of `function 3`;
- Moving up and down we must allocate and deallocate memory
<div align="center">
  <img src="images/the-stack-allocation.svg" alt="The stack allocation">
</div>
to deallocate it's just made the opposite of above. Sometimes the OS can make optimizations to save time and won't deallocate just one variable if the function is still in use.

Example on how memory is allocated in stack to call a new function:
<div align="center">
  <img src="images/the-stack.gif" alt="The stack allocation for new functions">
</div>

the **link register (lr)** also holds the main's return address for fast execution when `func1` returns. If `func1` calls another function, then *lr* is updated

**Summary**:
- the stack grows from high to low;
- it's used for function calls;
- stack variables die quickly;
- the following registers where covered: *stack pointer (sp)*, *base pointer (bp)* and *link register (lr)*

**Curiosity section :nerd_face:**
- when a function returns, that portion of memory is not cleaned by the compiler. It all remains there and will be overwritten when another function is called;
- when a program compiles, the compiler goes through all the code, looks at all the functions, and 'decides' how much memory is required. It then tells the process to allocate this memory when it runs. Many optimizations are made, such as whether to allocate a variable in memory or keep it in registers (if there are enough);
- the kernel sets a default stack limit for every program. Of course, you can set compiler options to override this default when compiling your code. If the program exceeds this stack limit, it will encounter the infamous **stack overflow** error.


### The Data Section

- fixed size portion in memory;
- portion of the memory dedicated to:
  - fixed size variables;
  - global variables;
  - static variables;
- all functions can access the variables from this region.

Take this code as an example:
```c
int a = 10;
int b = 20;

int main() {
  int sum = a + b;
  return 0;
}
```
the assembly code of that sum line would be something like this (read from bottom to up):
```arm
str r2, r1, r0
ldr r1 [#DATA, -4] ; load B
ldr r0 [#DATA, 0]  ; load A
```
think of `#DATA` as a pointer to the Data section.

**Curiosity section :nerd_face:**
- there are some languages, like Erlang, that allow you to change the text/code and data sections at runtime.


### The Heap

- characteristics:
  - large dynamic place for memory allocations;
  - all dynamic allocation happen here;
  - memory must be explicitly freed by your program;
  - it grows from low to high addresses (it grows up);
- all functions can access data from the heap using **pointers**.
<div align="center">
  <img src="images/the-heap-heap.svg" alt="The heap">
</div>
variables in the DATA area can also point to values in the heap.

Code example in C language allocating memory in the heap:
```c
#include <stdlib.h>

int main() {
  int *ptr = malloc(sizeof(int)); // allocate memory

  *ptr = 10;
  *ptr += 1;

  free(ptr); // free the memory
}
```
I'm using a x86_64 architecture machine, so if I compile the code above to assembly, you'd be seeing system calls like `call malloc@PLT` to allocate memory and `call free@PLT` to free memory. Take a look at the `/code-section3/allocation-example.s`file

**Memory leaks**
- it happens when you, as a programmer, forget to free memory that you allocated. The function returns without calling the `free` syscall. If you don't explicitly tell the kernel that you don't need that portion of memory anymore, it remains allocated to your process.

**Dangling pointers**
-  a dangling pointer is a pointer that references a memory location that is no longer valid. Example:
```c
#include <stdio.h>
#include <stdlib.h>

void causeDanglingPointer() {
    int *ptr = (int *)malloc(sizeof(int));
    *ptr = 42;
    free(ptr); // ptr becomes a dangling pointer
    printf("%d\n", *ptr); // Undefined behavior: accessing memory after free
}

int main() {
    causeDanglingPointer();
    return 0;
}
```

**Performance**
- due to the kernel mode switch (syscall involved) for allocation operations, it costs more to allocate memory in the heap than in the stack;
- allocation in the heap is random, which leads to more frequent cache misses in the `L*` caches. This results in more trips to memory to fetch new data, which is also more costly.

**Curiosity section :nerd_face:**
- using c language we can:
  - allocate memory in heap using the `malloc` function;
  - deallocate memory using the `free` function;
  - those functions make system calls, so the OS can do the memory allocation for your processes;
- garbage collection in languages like Go and Java try to solve the problem of allocating and deallocating memory for the programmer;
- it's a good prevention strategy to set the pointer to `NULL` after freeing it to ensure it doesn't point to a freed memory location. Ex.: `free(ptr); ptr = NULL;`.

## Terminal commands for linux used through the course

To know more about any commands below, just use the `man <command-name>` in terminal. Ex.: `man uname`
- `uname -r`: shows the kernel version;
- `gcc -S <file-name.c> -o <file-name.s>`: to generate the assembly code in your current architecture from a `.c` file;
- `gcc -g <file-name.c> -o <file-name.s>`: to compile to machine code. The `-g` flag instruct the compile to include debugging information;
- `gdb <file-name>`: to debug a compiled program. And after you executed it, you may type:
  - `start`: to begin the debugging;
  - `n`: to execute the next code instruction;
  - `info registers`: to show the registers and current information held by those. :bulb: if you see a register pc or rip, those are the program counter (PC) or (Register Instruction Pointer).

## Curiosity
Average cost time from the CPU perspective to read data from:
| Source    | :hourglass: |
| --------- | ----------- |
| Registers | 1 ns        |
| L1 cache  | 2 ns        |
| L2 cache  | 7 ns        |
| L3 cache  | 15 ns       |
| RAM       | 100 ns      |
| SSD       | 150 ns      |
| HDD       | 10 ms       |

