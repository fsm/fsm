<a href="https://github.com/fsm"><p align="center"><img src="https://user-images.githubusercontent.com/2105067/35464215-a014d512-02a9-11e8-8913-63a066f6064e.png" alt="FSM" width="350px" align="center;"/></p></a>

# FSM

[![Go Report Card](https://goreportcard.com/badge/github.com/fsm/fsm)](https://goreportcard.com/report/github.com/fsm/fsm) [![Gitter](https://img.shields.io/gitter/room/nwjs/nw.js.svg)](https://gitter.im/fsm/Lobby)

This package contains a simple interface for a [finite-state machine](https://en.wikipedia.org/wiki/Finite-state_machine) in Go.

> [A finite-state machine] is an abstract machine that can be in exactly one of a finite number of states at any given time. The FSM can change from one state to another in response to some external inputs; the change from one state to another is called a transition. An FSM is defined by a list of its states, its initial state, and the conditions for each transition.
>
> Wikipedia

# Getting Started

Check out the [getting started](https://github.com/fsm/getting-started) repository.  It provides an example and a README that is more specific to building a chatbot.

# Conversational Interfaces

While the interfaces provided can be used for a large varity of state machine needs, the supported tooling is focused on building conversational interfaces.

Building a conversational interface as a finite-state machine will reduce a ton of cognitive overhead, as at any given point you only have to concern yourself with the current step in the conversation.

With FSM you can build robust conversational interfaces that **run on any platform with a single codebase.**

If you dig into the source code of FSM, you'll find there's not really any code at all. The core of this library is entirely just interfaces. This is what buys the ability to deploy your conversational interface to any platform.

# Supported Platforms

Currently there is support for the following platforms:

- [Amazon Alexa](https://github.com/fsm/alexa)
- [Facebook Messenger](https://github.com/fsm/messenger)
- [Command Line](https://github.com/fsm/cli)

# Roadmap

We're soon planning to support the following platforms.

- [Texting via Twilio](https://github.com/fsm/twilio)
- [Google Home](https://developers.google.com/actions/)
- [Slack](https://api.slack.com/bot-users)
- [Twitter](https://developer.twitter.com/)
- [Cortana](https://www.microsoft.com/en-us/windows/cortana)
- Web Deployments (think something like [Product Hunt Ship](https://www.producthunt.com/ship) landing pages)

> These targets libraries are relatively easy to build, so when the next hot platform comes out, your conversational interface can be easily adapted to it!

# License

[MIT](LICENSE.md)
