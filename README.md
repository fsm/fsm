# FSM

This package contains a simple interface for a [finite-state machine](https://en.wikipedia.org/wiki/Finite-state_machine) in Go.

## Conversational Interfaces

While the interfaces provided can be used for a large varity of state machine needs, the supported tooling is focused on building conversational interfaces.

Building a conversational interface as a finite-state machine will reduce a ton of cognitive overhead, as at any given point you only have to concern yourself with the current step in the conversation.

With FSM you can build robust conversational interfaces that **run on any platform with a single codebase.**

If you dig into the source code of FSM, you'll find there's not really any code at all. The core of this library is entirely just interfaces. This is what buys the ability to deploy your conversational interface to any platform.

## Supported Platforms

Currently there is support for the following platforms:

- [Amazon Alexa](https://github.com/fsm/alexa)
- [Facebook Messenger](https://github.com/fsm/messenger)
- [Command Line](https://github.com/fsm/cli)
- [Texting via Twilio](#)

Very soon we're also going to support:

- Google Home
- Slack
- Twitter
- Web Deployments

> These targets libraries are extremely easy to build, so when the next hot platform comes out, your bot can be easily adapted to it.

## License

[MIT](LICENSE.md)
