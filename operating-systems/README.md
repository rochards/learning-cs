# Fundamentals of Operating Systems

My notes from [Fundamentals of Operating Systems](https://www.udemy.com/course/fundamentals-of-operating-systems) course.

## Table of content
- [Introduction](#section-1-and-section-2)
- [The anatomy of a process](#section-3-the-anatomy-of-a-process)
- [Terminal commands for linux used through the course](#terminal-commands-for-linux-used-through-the-course)

## Section 1 and Section 2

They are just an introduction and an overview of what the course is about and why we may need an Operating System.

## Section 3: The anatomy of a process

Process vs Program: A process is a program in motion/execution.

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


## Terminal commands for linux used through the course

To know more about any commands below, just use the `man <command-name>` in terminal. Ex.: `man uname`
- `uname -r`: shows the kernel version;
- `gcc -S <file-name.c> -o <file-name.s>`: to generate the assembly code in your current architecture from a `.c` file;
- `gcc -g <file-name.c> -o <file-name.s>`: to compile to machine code. The `-g` flag instruct the compile to include debugging information;
- `gdb <file-name>`: to debug a compiled program. And after you executed it, you may type:
  - `start`: to begin the debugging;
  - `n`: to execute the next code instruction;
  - `info registers`: to show the registers and current information holded by those. :bulb: if you see a register pc or rip, those are the program counter (PC) or (Register Instruction Pointer).

## Curiosity
Cost time to from the CPU perspective to read data from:
| Source    | :hourglass: |
| --------- | ----------- |
| registers | 1 ns        |
| L1 cache  | 2 ns        |
| L2 cache  | 7 ns        |
| L3 cache  | 15 ns       |
| RAM       | 100 ns      |
| SSD       | 150 ns      |
| HDD       | 10 ms       |

