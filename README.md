## Background

Just like programs should do one thing well, users should have one program for each thing.

## Goal

It should be easy to do all your text editing in a single editor. Writing email, making Tent posts and creating vCards should be as simple as `liz email`, `liz status` and `liz contact`.

## Components

#### `liz` executable

A wrapper meant to provide sane defaults for different tasks. The actual programs it calls may change somewhat frequently.

#### Other executables

Whenever possible `liz` will call an existing program. If one doesn't exist it will be written and included here. Hopefully each such program will eventually get its own project.
