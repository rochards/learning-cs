# Fundamentals of Operating Systems

My notes from [Fundamentals of Operating Systems](https://www.udemy.com/course/fundamentals-of-operating-systems) course.

## Table of content
* [Introduction](#section-1-and-section-2)
* [The anatomy of a process](#section-3-the-anatomy-of-a-process)

## Section 1 and Section 2

They are just an introduction and an overview of what the course is about and why we may need an Operating System.

## Section 3: The anatomy of a process

Process vs Program: A process is a program in motion/execution.

Program:
* Lives on disk;
* Code that is compiled and linked for a CPU;
* Only works on that CPU architecture when compiled;
* Produces executable file program, the so-called process;

Look at the format of a program in Linux systems:  
<p style="text-align: center;">
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/7/77/Elf-layout--en.svg/800px-Elf-layout--en.svg.png" alt="ELF (Executable and Linkable Format) program format">
</p>

Process:
* Lives on memory;

Look at the format of a process in an operating system:  
<p style="text-align: center;">
  <img src="images/process_anatomy.drawio.svg" alt="Anatomy of a process">
</p>


### Question to ask at the end:
- If a program is code compiled and linked for CPU, a python program is not a program?