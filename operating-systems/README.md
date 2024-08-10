# Fundamentals of Operating Systems

My notes from [Fundamentals of Operating Systems](https://www.udemy.com/course/fundamentals-of-operating-systems) course.

## Table of content
- [Sections 1 and 2: Introduction](#section-1-and-section-2)
- [Section 3: The anatomy of a process](#section-3-the-anatomy-of-a-process)
- [Section 4: Memory Management](#section-4-memory-management)
- [Section 5: Inside the CPU](#section-5-inside-the-cpu)
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
- Has an Instruction Pointer, aka Program Counter, that points to the memory address for the next instruction to be executed;
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
- The ELF header in the program tells the OS where to start the program;
- For simplicity, we're not considering the L* caches;

<div align="center">
  <img src="images/simplified-process-execution.gif" alt="Anatomy of a process">
</div>
the execution of the remaining instructions happen the same way.

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
- The stack grows from high to low;
- It's used for function calls;
- Stack variables die quickly;
- The following registers where covered: *stack pointer (sp)*, *base pointer (bp)* and *link register (lr)*

**Curiosity section :nerd_face:**
- When a function returns, that portion of memory is not cleaned by the compiler. It all remains there and will be overwritten when another function is called;
- When a program compiles, the compiler goes through all the code, looks at all the functions, and 'decides' how much memory is required. It then tells the process to allocate this memory when it runs. Many optimizations are made, such as whether to allocate a variable in memory or keep it in registers (if there are enough);
- The kernel sets a default stack limit for every program. Of course, you can set compiler options to override this default when compiling your code. If the program exceeds this stack limit, it will encounter the infamous **stack overflow** error.


### The Data Section

- Fixed size portion in memory;
- Portion of the memory dedicated to:
  - Fixed size variables;
  - Global variables;
  - Static variables;
- All functions can access the variables from this region.

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
- There are some languages, like Erlang, that allow you to change the text/code and data sections at runtime.


### The Heap

- Characteristics:
  - Large dynamic place for memory allocations;
  - All dynamic allocation happen here;
  - Memory must be explicitly freed by your program;
  - It grows from low to high addresses (it grows up);
- All functions can access data from the heap using **pointers**.
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
- It happens when you, as a programmer, forget to free memory that you allocated. The function returns without calling the `free` syscall. If you don't explicitly tell the kernel that you don't need that portion of memory anymore, it remains allocated to your process.

**Dangling pointers**
-  A dangling pointer is a pointer that references a memory location that is no longer valid. Example:
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
- Due to the kernel mode switch (syscall involved) for allocation operations, it costs more to allocate memory in the heap than in the stack;
- Allocation in the heap is random, which leads to more frequent cache misses in the `L*` caches. This results in more trips to memory to fetch new data, which is also more costly.

**Escape analysis**
- If an object does not escape its scope (objects that are created and used within a single function or thread and do not "escape" that context), the compiler can allocate it **on the stack instead of the heap**, reducing the overhead of heap allocation and garbage collection. This make a lot of sense in languages like Java and Go.

**Program break**
- Is a pointer that represents the current end of the process's heap;
- When the process needs more heap, the program break is moved to a higher address to increase the heap size.

**Summary**
So, the heap:
- Stores large data;
- Remain until explicitly removed;
- All functions can access.

**Curiosity section :nerd_face:**
- Using c language we can:
  - Allocate memory in heap using the `malloc` function;
  - Deallocate memory using the `free` function;
  - Those functions make system calls, so the OS can do the memory allocation for your processes;
- Garbage collection in languages like Go and Java try to solve the problem of allocating and deallocating memory for the programmer;
- It's a good prevention strategy to set the pointer to `NULL` after freeing it to ensure it doesn't point to a freed memory location. Ex.: `free(ptr); ptr = NULL;`.

## Section 4: Memory Management

### The Anatomy of Memory

**Characteristics of memory**:
- Stores data;
- Volatile:
  - RAM - Random Access Memory
- Non-Volatile:
  - ROM - Read Only Memory

**Types of RAM**:
- SRAM - Static RAM:
  - 1 bit -> 1 flip flop circuit (5 or 6 transistors);
  - Access is always fast;
  - It's also used as caches in CPU and SSDs.
- DRAM - Dynamic RAM:
  - 1 bit -> 1 capacitor + 1 transistor;
  - Slower than SRAM;
  - SDRAM (Synchronous DRAM): the clock of CPU and RAM are synchronized;
  - DDR (Double Data Rate) :question: do some research for better understanding.

Take a look at a simplified example of reading from a **DDR4** SDRAM:
<div align="center">
  <img src="images/anatomy-of-memory-ddr4.svg" alt="DD4 memory">
</div>

- Each io pin gives 1 byte (8 bits) of information -> *64 pins x 1 byte = 64 bytes*;
- Because of the costing of reading, the CPU gets a total of 64 bytes (the CPU's cache line size) in response;

Take a look at a simplified example of reading from a **DDR5** SDRAM:
<div align="center">
  <img src="images/anatomy-of-memory-ddr5.svg" alt="DD5 memory">
</div>

- Each io pin gives 2 bytes (16 bits) of information -> *32 pins x 2 bytes = 64 bytes*;
- Because of the costing of reading, the CPU gets a total of 64 bytes (the CPU's cache line size) in response;
- :bulb: because of this idea of channels, as one CPU core is reading from channel A, the other core can read from channel B.

### Reading and Writing from and to Memory

Take a look at simplified example of reading from memory:
<div align="center">
  <img src="images/reading-from-memory.gif" alt="Reading from memory">
</div>

- Note that after getting the 64 bytes, all the next instructions are already in the L caches, so there is no need to fetch from RAM again for a while;
- Also, note that if any instruction is a call to a function, it will be necessary to read from memory again, causing us to lose what we already had in the L cache. Because of that, calling too many functions in the code can lead to a decrease in code execution performance.

Take a look at simplified example of writing to memory:
<div align="center">
  <img src="images/write-to-memory.gif" alt="Write to memory">
</div>

### Virtual Memory

**Fragmentation**

It happens when free memory are scattered in small blocks across the RAM, making difficult to allocate contiguous blocks of memory;
<div align="center">
  <img src="images/anatomy-of-memory-fragmentation.gif" alt="Memory fragmentation">
</div>

- **External fragmentation**: the image above provides a great example of external fragmentation. There is free space available between blocks, but yet the memory can't fit a new process;
- **Internal fragmentation**: it happens when the blocks allocated are larger than needed, resulting in potentially free space available (because it is not being used by the process) inside the block. For example, if a process requests 18 bytes and the allocator rounds up to the nearest 32 bytes, 14 bytes are wasted within that block.

**Virtual Memory and Fragmentation**

- Let's use fixed block size and call it **page**;
- A **page size** is often 4 kB;
- Each process has virtual address space;
<div align="center">
  <img src="images/anatomy-of-memory-vm.png" alt="Virtual memory mapping">
</div>

- Note that when looking at the virtual memory, all the processes are alike;
- `A.P2` and `C.P1` show internal fragmentation, where that portion of the process doesn't need the whole page, but it is allocated anyway. However, this solves the external fragmentation problem.
- We map logical pages (virtual/logical addresses) to physical pages (physical addresses), and this information is stored in the **process page table**;
  - Page table is stored in memory;
  - Each process has it own page table that resides in somewhere in the kernel space;
<div align="center">
  <img src="images/anatomy-of-memory-page-table.png" alt="Page table">
</div>

-
  - In the image above, VMA stands for Virtual Memory Address, and PA stands for Page Address, which is the physical address. Note that the page table only holds the initial address for each page. Considering that each page is 4 kB in size, the Operating System does the math to determine the ending address.

**Shared Memory**

- Let's consider a problem: a program generates 5 processes. There will be many regions in memory with duplicated text/code. So, with virtual memory, we load the text/code once and map all virtual addresses to the same physical addresses.
<div align="center">
  <img src="images/anatomy-of-memory-shared-memory.png" alt="Shared memory">
</div>

- Another example: when a needed library is loaded for a program, if another program comes in and needs the same library, the OS just maps to the page where this library resides in memory for the new program's page table.

**Isolation**
- With direct access to physical addresses, processes could attempt to load an address they aren't supposed to. Virtual memory solves this issue by completely isolating each process.

**Swap**
- Large processes can't fit entirely in memory, so some virtual pages are put onto disk;
- This can also be used if there are too many processes and there is not enough RAM available;
- The OS will decide which pages to offload to disk;
- There is a bit in the mapping that indicates whether that particular page is on disk or in RAM;
- Remember, for the CPU, the data and code must be in RAM. So if the page isn't loaded in memory at that particular time, the OS will issue a page fault and load it from the swap file. The OS will also allocate new memory and update the new physical address in the mapping (because the page can be loaded anywhere in memory).

**Summary**
- Virtual memory addresses a lot of issues, like sharing memory, loading large processes, isolating processes, and external fragmentation;
- Virtual memory comes with the cost of using page tables and having page faults (a kernel mode switch is needed here);
- There is also an additional layer of translation with the MMU (Memory Management Unit), because the CPU can't read virtual addresses.

**Curiosity section :nerd_face:**
- To avoid the constant cost of reading from RAM for address translations, there is a TLB (Translation Lookaside Buffer) in CPU, which is a cache for the page table.

**My Q&A of this section**:
- Does fragmentation only occur in physical memory? **A**: No! It also happens in virtual memory.
- What is the advantage of using virtual memory? **A:** We get rid of external fragmentation.
- Why is hard to share memory with physical memory? **A** This was easier accomplished using virtual memory.
- Memory allocation must be contiguous: should the whole process fit in a contiguous space, or can the Kernel divide it? **A**: Actually, the kernel divides the process into pages. From the perspective of virtual memory, the process appears as contiguous in memory, but it is divided in physical memory.
- What is the cost of using virtual memory in terms of performance? **A**: It's usually the same as reading from memory (~ `50 ns - 100 ns`), since the page table resides in RAM.
- Why processes need contiguous slots of memory? **A**: It's just a design choice that all existing OS agreed for simplicity. But remember that the process is contiguous only when we look at the virtual memory, but in physical memory the pages can be dispersed.
- Is the entire process divided into pages, including the stack and heap? Do they also appear as contiguous in virtual memory but can actually be scattered in physical memory? **A**: Yes for all questions.
- If the stack are scattered in physical memory, does this mean that the base pointer and stack pointer registers look at the virtual addresses? **A**: Yes everything points to virtual memory. When the actual execution is required the MMU is consulted to do the mapping to get the content. 
- Who decided that a page must be 4 kB? Is it a convention for all operating systems? **A**: Yes, it is a convention.

### DMA (Direct Memory Access)

**Peripherals read**:  
Data from network/disk must pass through the CPU:
- Network -> CPU -> RAM;
- Disk -> CPU -> RAM;
- Keyboard -> CPU -> RAM;

It's slow for large transfers.

**With DMA**:
- Allows direct access from/to network/disk to/from RAM;
- The DMA controller initializes the operation;
- It only understands/reads physical addresses.

**Pros and Cons of DMA**:
- Efficient for large transfers;
- No virtual memory management;
- Less CPU overhead;
- However, there is an initialization cost;
- There are security concerns and complexities involved;
- Cannot be used for interrupts (keyboard/mouse).

**My Q&A of this section**:
- Why do we need DMA, or what benefits does it bring? **A**: The benefit is usually attribute when there is a large data transfer involved.

### Lab

Using the `top`command in Linux we're presented with information like this:
```bash
MiB Mem :  924.2 total,  590.7 free,   111.0 used,   222.4 buff/cache
MiB Swap:  100.0 total,  100.0 free,     0.0 used.   745.1 avail Mem

  PID USER PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
18978 pi   20   0   10400   2980   2468 R   1.0   0.3  14:39.22 chrome
```
- The **Mem** line is in fact the physical memory available:
  - In general, the **buff/cache** is a portion of memory that the O.S uses for optimizations like I/O management.
- The **Swap** line represents the disk space used to extend the amount of memory:
  - The **avail Mem** info can be tricky to understand. It represents the amount of memory available before the need for swapping.
- The **VIRT** represents in `KiB` the amount of Virtual Memory used by the process;
- The **RES (Resident Set Size)** represents in `KiB` the portion of virtual memory that is currently held in RAM;
- The **SHR** represents in `KiB` the Shared Memory portion used by the process, such as shared libraries.

In the folder `code-section4`, there are two code examples that show memory allocation: `alloc-example-1.c` and `alloc-example-2.c`. After compiling them, using `time ./<compiled-file-name>` command, we see that the `alloc-example-1.c` takes longer to run. This is because of the 1,000,000 syscalls involved. Remember: syscall involves a kernel mode switch, which takes time.

## Section 5: Inside the CPU

### CPU components and Architecture

Take a look at the image below:
<div align="center">
  <img src="images/inside-the-cpu-basic-components.svg" alt="Basic components of a CPU">
</div>

as we know, a CPU can be composed of multiple cores, in other words, multiple CPUs whose basic components are:
- The **ALU (Arithmetic Logic Unit)** performs the arithmetic computations: addition, multiplication, subtraction, etc., and logic operations: AND, OR, NOT, etc;
- The **CU (Control Unit)** fetches instructions from memory, registers, decodes instructions, and coordinates the actions of the ALU;
- The **MMU (Memory Management Unit)** translates virtual addresses to physical addresses and handles cache operations. The TLB (Translation Lookaside Buffer) resides here;
- The **Registers** are fast units of storage for data and instructions. The ALU only executes what is here:
  - We have general-purpose registers for storing data;
  - We also have specific ones, like the PC (Program Counter), SP (Stack Pointer), IR (Instruction Register), and so on.
- The **Caches (L1, L2, L3)** store frequently accessed data to speed up memory access:
  - L1 is local to core;
  - L2 is local to core, but it used to be the one shared across CPUs in absence of L3;
  - L3 is shared between multiple cores.

**CPU Architecture**:
- RISC - Reduced Instruction Set:
  - Instructions are simpler compared to CISC;
  - One instruction is executed in one single CPU cycle;
  - Lower power consume;
  - ARM.
  - Ex.: to execute `a = a + b` it would take 4 instructions
  ```arm
  LDR r0, a
  LDR r1, b
  ADD r0, r1
  STR a1, r0
  ```
- CISC - Complex Instruction Set:
  - One instruction takes more than one CPU cycles to be executed;
  - Requires more power to be executed;
  - x86 (Intel/AMD);
  - Ex.: to execute `a = a + b` it would take 1 instruction
  ```x86
  ADD a, b # a and b are memory addresses
  ```

**Clock Speed**
- It measures how fast a CPU can process an operations in cycles;
  - 1 operation can be: fetching instructions from memory.
- Ex.: 1 GHz is equals 1 billion cycles per second;
- In RISC 1 cycle could mean 1 instruction executed.

### Instruction Life Cycle

- Fetch from memory (if it is not in L1 cache):
- Decode: decode is like understanding the instruction. The CU does that;
- Execute: the ALU does that;

Look at this example:
<div align="center">
  <img src="images/inside-the-cpu-execution-example.gif" alt="Program execution example">
</div>

All the virtual translation were skipped for simplicity.
Walking through the example:
1. We have a program already loaded in memory, so it's a process;
2. The CPU needs to execute the next instruction, so the CU "asks" the MMU for this instruction;
3. The `pc` register is holding a virtual address that gets translated to the physical address 640 by the MMU;
4. The MMU doesn't find this instruction in any caches, so it goes to RAM;
5. The RAM returns 64 bytes, so the MMU updates the L caches;
6. Now the MMU gives the CU the instruction;
7. The CU starts decoding the instruction by converting the `sub` operation to a code understood by the ALU and also gets the actual value of the `sp` register;
8. The CU now gives the instruction to the ALU;
9. The ALU executes the instruction;
10. The CU writes the result back to `sp`;
11. The `pc` gets updated;
12. The CPU needs to execute the next instruction, so the CU "asks" the MMU for this instruction;
13. The `pc` register is holding a virtual address that gets translated to the physical address 644 by the MMU;
14. The MMU finds this instruction in cache and the cycles repeats.

### Pipelining and Parallelism

**Without pipelining**: Using the model proposed below, the CPU would spend most of the time idle, mainly because of the `fetch` cycle, which can take hundreds of nanoseconds for the instruction to be returned from the RAM:
```
Instruction 1: fetch -> decode -> execute -> write results
Instruction 2: ------------------------------------------- fetch -> decode -> execute -> write results
Instruction n: --------------------------------------------------------------------------------------- ...
```

**With pipelining**: this can be improved doing multiple things in parallel, so a single core/CPU can do:
```
Instruction 1: fetch -> decode -> execute -> write results
Instruction 2: -------- fetch  -> decode  -> execute -> write results
Instruction 3: ------------------ fetch   -> decode  -> execute -> write results
Instruction n: ----------------------------------------------------------------- ...
```

**Parallelism**: we are talking about having more the one CPU in a single computer, which is very common nowadays.

**Hyper-threading**: hyper-threading exposes a single core as multiple logical cores. The manufacturers found a way to present a single core as two logical cores to the OS. Look at the image below:
<div align="center">
  <img src="images/inside-the-cpu-hyper-threading.png" alt="Hyper threading">
</div>

before hyper-threading, there was no other way rather than to queue the processes and to execute the 2nd by the performing a context switch, as shown in the left side of the image. But with hyper-threading, as shown in the right side of the image with a single core, we get the illusion that two processes are being executed simultaneously. In fact, both are being executed without the need for a context switch. This can be achieved by:
- Givin each process dedicated registers for only the needed ones, like `pc`, `bp`, `sp`, etc.;
- Sharing other things like L caches at the same time;
- The pipelining executing mode discussed above.


**Curiosity section :nerd_face:**

- Take a look at the image below:
<div align="center">
  <img src="images/inside-the-cpu-memory-and-bus.png" alt="Basic components of a CPU">
</div>

DIMM1 and DIMM2 are two different slots of memory on the motherboard. Note that in this design architecture, the manufacturer decided to dedicate each DIMM to distinct CPU cores. The DSM (Distribute Shared Memory), though, allows cores to access memory across buses, which of course is slower. This is important to know because the process can be placed on any core, and due to the travel required to access data in memory, it can sometimes execute slower or faster.
- To avoid the constant cost of reading from RAM for address translations, there is a TLB (Translation Lookaside Buffer) in CPU, which is a cache for the page table. When there is the need for context switch, it means another process wins the CPU, the TLB must be flushed;
- In a Linux machine, type in terminal `getconf -a | grep CACHE` to check the caches size of your CPU. You will see a result like this in bytes:
```bash
LEVEL1_ICACHE_SIZE                 32768 # = 32 KiB for Instructions
LEVEL1_DCACHE_SIZE                 49152 # = 48 KiB for Data
LEVEL2_CACHE_SIZE                  1310720 # = 1280 KiB only for data
LEVEL3_CACHE_SIZE                  12582912 # = 12 MiB only of data
```

### Lab

Using the `top`command in Linux we're presented with information like this:
```bash
%Cpu(s):  2.4 us,  0.9 sy,  0.0 ni, 96.7 id,  0.1 wa,  0.0 hi,  0.0 si,  0.0 st
```
Let's discuss some of the metrics above:  
- The **us (user space)** value: is the percentage of time the CPU is spending running the user's started processes;
- The **sy (system space)** value: is the percentage of time the CPU is spending running the system/kernel processes;
- The **id (idle)** value: is the percentage of time the CPU is processing anything;
- The **wa (I/O wait)** value: is the percentage of time the CPU is waiting for I/O operations to complete;


## Terminal commands for linux used through the course

To know more about any commands below, just use the `man <command-name>` in terminal. Ex.: `man uname`
- `uname -r`: shows the kernel version;
- `gcc -S <file-name.c> -o <file-name.s>`: to generate the assembly code in your current architecture from a `.c` file;
- `gcc -g <file-name.c> -o <file-name.s>`: to compile to machine code. The `-g` flag instruct the compile to include debugging information;
- `gdb <file-name>`: to debug a compiled program. And after you executed it, you may type:
  - `start`: to begin the debugging;
  - `n`: to execute the next code instruction;
  - `info registers`: to show the registers and current information held by those. :bulb: if you see a register pc or rip, those are the program counter (PC) or (Register Instruction Pointer);
- `top`: to show the executing processes in you computer with some statistics;
- `time ./<compiled-file-name>`: to get statistic info about the program execution;
- `getconf -a | grep CACHE`: to get info about CPU's caches;
- `uname -m`: to get the architecture of your CPU;
- `lscpu`: to get a lot of info about your CPU;

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

