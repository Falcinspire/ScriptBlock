# ScriptBlock

NOTE: This software was a personal experiment in language design, and was never built with the intention of being a production-ready piece of software. However, I am proud of it, and it does work on the small projects that I have tested it with.

Scripting Language for Minecraft: Java Edition. Instead of writing a bunch of .mcfunction files, ScriptBlock allows you to write something that looks more like traditional code.

## Installing

This project is not yet available by executable or installer. The only option is to clone the golang project (usually to the go home on your system), and to install it.
```
git clone https://github.com/Falcinspire/ScriptBlock.git
go install github.com/Falcinspire/scriptblock/cmd/scriptblock
go install github.com/Falcinspire/scriptblock/cmd/scriptblock-new
```
NOTE: This command sequence is untested.

## Usage

`scriptblock-new` should be used to create a new project.

If the `bin` directory of the go workspace is in the system path, then building is simple:
```
scriptblock <project> 
```
i.e.
```
scriptblock github.com/Falcinspire/predator LATEST
```
If the `bin` directory of the go workspace is not in the system path, use the same command as above, but with the full path to the scriptblock executable:
```
"C:/Program Files/go/bin/scriptblock.exe" <project> 
```
NOTE: This form is untested.

## Getting Started

The language workspace is modelled after Golang. There is a central workspace that contains the source code and the outputs. The workspace will be set with the environment variable "ScriptBlockPath". 
Projects have a qualified name and a version. The qualified name is the location of the project both on github.com and on the local disk. The version will be included in the path on the local disk, and will be a release on github. The only exception is the version "LATEST", which signifies a developmental stage. The github representation for this is no release; the master branch, and the local disk representation is the qualified name followed by /LATEST.
An example qualified name is `github.com/Falcinspire/predator` and version is`'1.0.0`. This project would have the name "predator" and be stored at  `https:/github.com/Falcinspire/predator` with the tag `1.0.0`. In a local workspace, this would be stored at `%ScriptBlockPath%/src/github.com/Falcinspire/predator/1.0.0`.
Project builds are stored in the bin folder, which is adjacent to the src folder. 
Example:

    %ScriptBlockPath%
        ├── src                                            
        |   ├─ github.com                                  
        |       ├── falcinspire                          
        |            ├── predator 
        |                ├── LATEST
        |                   ├── test.sb                 
        |                   ├── helper.sb                   
        |                   ├── zombie.json                   
        ├── bin                                           
            ├─ github.com                
                ├── falcinspire
                     ├── predator 
                         ├── LATEST
                            ├── (namespace)
                                ├── functions
                                    ├── *.mcfunction
                            ├── (namespace)
                                ├── functions
                                    ├── *.mcfunction

There is also a module description file, module.yaml. This file contains basic information about the project, as well as a list of dependencies. See /examples for format. Right now the Name field is useless.

## What to do with the output 

The output of the ScriptBlock compiler is the part of the data pack needed for functions. Merge this folder into the /data folder of a datapack for it to work.

## What it looks like

![Screenshot](https://github.com/Falcinspire/ScriptBlock/blob/dev/screenshot.png)

## Example

If the repository is cloned and the environment variable "ScriptBlockPath" is set to the repository's /example directory, then running `scriptblock github.com/Falcinspire/predator LATEST` will demonstrate the software working. 
