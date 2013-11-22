## Background

Just like programs should do one thing well, users should have one program for each thing.

## Goal

Make it easy to do all your text editing in your favourite editor.

## Components

#### `liz` executable

A wrapper meant to provide sane defaults for different tasks. The actual programs it calls may change somewhat frequently.

The only command implemented so far is `liz status` to make a Tent status post. Examples of planned commands include `liz email` to write an email and `liz contact` to create a vcard.

#### Other executables

Whenever possible `liz` will call an existing program. If one doesn't exist it will be written and included here. Hopefully each such program will eventually get its own project.
