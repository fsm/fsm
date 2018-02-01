<a href="https://github.com/fsm"><p align="center"><img src="https://user-images.githubusercontent.com/2105067/35464215-a014d512-02a9-11e8-8913-63a066f6064e.png" alt="FSM" width="350px" align="center;"/></p></a>

# FSM

This package contains a simple interface for a [finite-state machine](https://en.wikipedia.org/wiki/Finite-state_machine) in Go.

> [A finite-state machine] is an abstract machine that can be in exactly one of a finite number of states at any given time. The FSM can change from one state to another in response to some external inputs; the change from one state to another is called a transition. An FSM is defined by a list of its states, its initial state, and the conditions for each transition.
>
> Wikipedia

## Getting Started

Check out our [example repository](https://github.com/fsm/example). It provides an example and a README that is more specific to building a chatbot.

After looking at the example, you can get started building your own chatbot by cloning down the [getting started](https://github.com/fsm/getting-started) repo as a base for your project.

## Conversational Interfaces

While the interfaces provided can be used for a large varity of state machine needs, the supported tooling is focused on building conversational interfaces.

Building a conversational interface as a finite-state machine will reduce a ton of cognitive overhead, as at any given point you only have to concern yourself with the current step in the conversation.

With FSM you can build robust conversational interfaces that **run on any platform with a single codebase.**

If you dig into the source code of FSM, you'll find there's not really any code at all. The core of this library is entirely just interfaces. This is what buys the ability to deploy your conversational interface to any platform.

## FSM Components

The best way to explain the components of this library is by a concrete example.

Here we have a partial state machine of how a customer may interact with a bank teller.

![statemachine](https://user-images.githubusercontent.com/2105067/35538170-c049b938-0501-11e8-8064-1ba3d9b576be.png)

### [StateMachine](https://github.com/fsm/fsm/blob/master/fsm.go#L3-L4)

The entire diagram above is a StateMachine.  A StateMachine is an array of all States in the State Machine.

### [State](https://github.com/fsm/fsm/blob/master/fsm.go#L15-L21)

Each of the yellow circles in the diagram is a State.

### [Emitter](https://github.com/fsm/fsm/blob/master/fsm.go#L23-L27)

An Emitter is a definition of how to output data.

This is the 1/2 of what buys us the ability to deploy our conversational interfaces to multiple platforms.

### [Traverser](https://github.com/fsm/fsm/blob/master/fsm.go#L36-L47)

A Traverser is the abstract element that is traversing the state machine.

This is effectively a model for the user who is communicating with your conversational interface.

### [BuildState](https://github.com/fsm/fsm/blob/master/fsm.go#L11-L13)

A StateMachine is actually comprised of BuildState, which is a function that returns a State.

The reason for this function is the fact that this function also gives our State access to an Emitter and Traverser.

### [Store](https://github.com/fsm/fsm/blob/master/fsm.go#L29-L34)

There is also a Store that allows you to store arbitrary data for each Traverser.

This is the component that allows us to keep track of how much cash the traverser has in our example.

## Supported Platforms

Currently there is support for the following platforms:

- [Amazon Alexa](https://github.com/fsm/alexa)
- [Facebook Messenger](https://github.com/fsm/messenger)
- [Command Line](https://github.com/fsm/cli)
- [Texting via Twilio](https://github.com/fsm/twilio)

Very soon we're also going to support:

- Google Home
- Slack
- Twitter
- Web Deployments

> These targets libraries are extremely easy to build, so when the next hot platform comes out, your conversational interface can be easily adapted to it.

## License

[MIT](LICENSE.md)
