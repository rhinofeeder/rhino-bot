# Rhino Bot

Rhino Bot is the Twitch bot used by [RhinoFeeder](https://twitch.tv/rhinofeeder). Rhino Bot serves as a replacement for
other well known bots and is capable of handling commands, sending messages on a regular time interval, and reacting to
chatters based on a custom set of criteria.

For discussing suggestions and changes, join the [RhinoFeeder Discord Server](https://discord.com/invite/mrzNnq6).

* [Installation](#installation)
* [Usage](#usage)
* [Contributing](#contributing)
    * [Behaviors](#behaviors)
    * [Commands](#commands)
        * [Adding a command](#adding-a-command)
    * [Conditionals](#conditionals)
        * [Adding a conditional](#adding-a-conditional)
    * [Timers](#timers)
        * [Adding a timer](#adding-a-timer)

## Installation

1. Clone this repository.
1. While logged into the Twitch account that you'll be using for the bot, go to https://twitchapps.com/tmi and get the
   oauth token.
1. Run `make install TOKEN=<token from previous step>`.
    1. If you do not provide a token, the file will still be created without any contents.

## Usage

In `main.go`, update the properties of the `bot.RhinoBot` struct to the appropriate values for your channel. Only the
following values should need to be changed:

- `Channel` - the name of the Channel that the bot will be joining. This must be in all lowercase.
- `Name` - the name your bot will use in chat. This should be the same as the Twitch account handle that you are using
  for the bot.

Then, run `make run` from your terminal.

To shut down Rhino Bot, kill the terminal process.

## Contributing

Pull requests are welcome. For help or guidance implementing your changes, you can join
the [RhinoFeeder Discord Server](https://discord.com/invite/mrzNnq6).

This project uses GitHub actions to run various checks on all Pull Requests. At this time, submitted pull request will
be run through a linter that will annotate their file diff. There is also an action to make sure that all tests still
run successfully.

The following sections go over Behaviors for Rhino Bot and how to add new Behaviors. These sections assume a base level
understanding of Twitch chat functionality and terminology.

### Behaviors

Behaviors are the term used for any action the bot takes. At this time, these include commands, conditionals, and
timers. Each of these will be explained further in sections below.

All behaviors must implement the following function(s):

- `Handle(string) (string, error)` - Returns the string that the bot will send to chat, or an error if one occurred.
  Errors should be used for internal, unrecoverable errors only. An example would be failing to open a file that was
  required to write to. This function is also responsible for keeping track of when the command was last called for
  calculating cooldown time.

Each type of Behavior will have specific functions that it will need to implement in addition to these. Check the
relevant section below to see what additional functions you'll need to add.

### Commands

A command is a behavior that is manually triggered by a chatter. Rhino Bot supports commands that can only be triggered
by chatters with moderator privileges. The syntax for triggering a command is `!<command_name> [input]`.

#### Adding a command

Start by creating a new file at `behavior/command/<command_name>.go`. Create a struct to represent your command. Use the
naming convention `<Command>Command` for your struct. Commands must implement the following functions:

- `Name() string` - returns the string that chatters will use to trigger it. For example, a command whose `Name`
  function returns "foo" will be triggered in chat with `!foo`.
- `RequiresMod() bool` - Returns whether a command requires a moderator to trigger it.
- `OnCooldown() bool` - Returns whether the command is on cooldown or not.

Write a few basic tests for any logic or functionality in your command.

Then add a pointer to your struct to the `commands` global variable in `main.go`. Keep this list in alphabetical order
by command name.

### Conditionals

Conditionals are behaviors that are triggered by a developer-defined set of criteria. This could be based on a random
chance, whether a message contains certain text, or anything the developer chooses. Rhino Bot will check all non-command
messages in chat to see if they match the criteria. Once one conditional is triggered, no more can be triggered for the
same message.

#### Adding a conditional

Start by creating a new file at `behavior/conditional/<conditional_name>.go`. Create a struct to represent your
conditional. Use the naming convention `<Conditional>Conditional` for your struct. Conditionals must implement the
following functions:

- `ShouldHandle(string) bool` - Returns whether the bot should trigger the conditional.

Write a few basic tests for any logic or functionality in your conditional.

Then add a pointer to your struct to the `conditionals` global variable in `main.go`. Keep this list in alphabetical
order by conditional name.

### Timers

Timers are behaviors that are triggered at a regular time interval, starting from when the bot first starts. Timers will
continue to run until the bot shuts down.

#### Adding a timer

Start by creating a new file at `behavior/timer/<timer_name>.go`. Create a struct to represent your timer. Use the
naming convention `<Timer>Timer` for your struct. Conditionals must implement the following functions:

Write a few basic tests for any logic or functionality in your timer.

Then add a pointer to your struct to the `timers` global variable in `main.go`. Keep this list in alphabetical order by
timer name.

## Contributors

Rhino Bot thrives off of contributions from the community. This section is a special thanks to those who have helped
Rhino Bot grow to what it is today!

[![](https://contrib.rocks/image?repo=rhinofeeder/rhino-bot)](https://github.com/rhinofeeder/rhino-bot/graphs/contributors)