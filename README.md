## Description

Create and process all sorts of text documents from within vim using loosely coupled utilities. Liz is a rewrite of Mutt, but it's also Mutt for Tent, Mutt for vCard, etc. all rolled into one.

At my level of skill this is basically a moonshot project. At least it will be exciting!

## Components

#### `liz` executable

The only user facing part of the project. Run `liz status` to create and post a status post, `liz email` (not written yet) to create and send an email, etc.

All of these will be opinionated in how they interact with vim and how they handle the task (e.g. whether you're prompted to fill out the `subject` field of an email in vim or on the command line). Many people may use only a couple of these commands.

#### Other executables

None of the `liz-` programs such as `liz-tent` are meant to be called directly by the casual user. They handle parts of the project that can easily be abstracted away into seperate programs. Hopefully all these roles will be filled by dedicated projects someday.

#### Go packages

The rest of the directories in the project are Go packages for `liz`, none of which contain executables.
