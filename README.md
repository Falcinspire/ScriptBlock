# ScriptBlock

NOTE: This software is in beta. Development efforts were halted upon contributor's return to college. Continued efforts are still planned, however.

Scripting Language for Minecraft: Java Edition. Instead of writing a bunch of .mcfunction files, ScriptBlock allows you to write something that looks more like traditional code.

## Installing

This project is not yet available by executable or installer. The only option is to clone the golang project (usually to the go home on your system), and to install it. The master branch is functional, but this tutorial is written for the dev branch, so please check that out before installing. *The master branch uses a different qualified name format; the url prefix is dropped.*
```
git clone https://github.com/Falcinspire/ScriptBlock.git
git checkout dev
go install
```
NOTE: This command sequence is untested.

## Usage

If the bin directory of the go workspace is in the system path, then use is simple:
```
scriptblock <project> 
```
i.e.
```
scriptblock github.com/Falcinspire/predator
```
If the bin directory of the go workspace is not in the system path, use the same command as above, but with the full path to the scriptblock executable:
```
C:/Program Files/go/bin/scriptblock.exe <project> 
```
NOTE: This form is untested.

## Getting Started

The language workspace is modelled after Golang. There is a central workspace that contains the source code and the outputs. The workspace will be set with the environment variable "ScriptBlockPath". 
Projects will have a name and a path. Together, the name and the path make up the qualified name, which will be used for `import` statements and project structure. The qualified name is planned to eventually include a version number/code as well. This is to eventually make using libraries trivial. 
An example qualified name is `github.com/Falcinspire/predator`. This project would have the name "predator" and be stored at  `https:/github.com/Falcinspire/predator`. In a local workspace, this would be stored at `%ScriptBlockPath%/src/github.com/Falcinspire/predator`.
Project builds are stored in the bin folder, which is adjacent to the src folder. 

## What to do with the output 

The output of the ScriptBlock compiler is most of what is necessary for a data pack. Further instructions will be included later as the project matures.

## What it looks like

[Screenshot](https://github.com/Falcinspire/ScriptBlock/blob/master/screenshot.png)
