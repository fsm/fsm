<a href="https://github.com/fsm"><p align="center"><img src="https://user-images.githubusercontent.com/2105067/35464215-a014d512-02a9-11e8-8913-63a066f6064e.png" alt="FSM" width="350px" align="center;"/></p></a>
<p align="center">
  <a href="https://github.com/fsm/fsm/releases"><img src="https://img.shields.io/github/tag/fsm/fsm.svg" alt="Version"></img></a>
  <a href="https://github.com/fsm/fsm/blob/master/LICENSE.md"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="MIT License"></img></a>
  <a href="https://travis-ci.org/fsm/fsm"><img src="https://travis-ci.org/fsm/fsm.svg?branch=master" alt="Build Status"></img></a>
  <a href="https://codecov.io/gh/fsm/fsm"><img src="https://codecov.io/gh/fsm/fsm/branch/master/graph/badge.svg" alt="codecov"></img></a>
  <a href="https://goreportcard.com/report/github.com/fsm/fsm"><img src="https://goreportcard.com/badge/github.com/fsm/fsm" alt="Go Report Card"></img></a>
  <a href="https://spectrum.chat/fsm"><img alt="Join the community on Spectrum" src="https://withspectrum.github.io/badge/badge.svg"/></a>
</p>

# FSM

This package contains a simple interface for a [finite-state machine](https://en.wikipedia.org/wiki/Finite-state_machine) in Go.

# What's a Finite-state Machine?

> [A finite-state machine] is an abstract machine that can be in exactly one of a finite number of states at any given time. The FSM can change from one state to another in response to some external inputs; the change from one state to another is called a transition. An FSM is defined by a list of its states, its initial state, and the conditions for each transition.
>
> Wikipedia

The FSM library was built as purely against the definition of a formal finite-state machine as possible, so the Wikipedia definition holds true for this library.

# Why Use a Finite-state Machine?

FSM was specifically made to build light-weight / generic conversational interfaces (think chatbots), so if that's what your trying to do you'll find this library quite nice.

Building a conversational interface as a finite-state machine will reduce a ton of cognitive overhead, as at any given point you only have to concern yourself with the current step in the conversation.

What is most attractive about FSM is the fact that you can build robust conversational interfaces that **run on any platform with a single codebase.**

# Getting Started

If you're looking to build a chatbot, check out the [getting started](https://github.com/fsm/getting-started) repository.

This README is more high-level and won't get into the details you'll want to start building.

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
