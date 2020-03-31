[![GitHub](https://img.shields.io/github/license/Rushifaaa/go-tsukasa-bot?style=flat-square)](./LICENSE)
[![GitHub contributors](https://img.shields.io/github/contributors/Rushifaaa/go-tsukasa-bot?style=flat-square)](https://github.com/Rushifaaa/go-tsukasa-bot/graphs/contributors)
[![GitHub release (latest SemVer including pre-releases)](https://img.shields.io/github/v/release/Rushifaaa/go-tsukasa-bot?include_prereleases&sort=semver&style=flat-square)](https://github.com/Rushifaaa/go-tsukasa-bot/releases)
[![GitHub issues](https://img.shields.io/github/issues/Rushifaaa/go-tsukasa-bot?style=flat-square)](https://github.com/Rushifaaa/go-tsukasa-bot/issues)
[![Discord](https://img.shields.io/discord/508727953350328320?style=flat-square)](https://discord.gg/kFqWZtv)



<p align="center">
  <a href="https://github.com/Rushifaaa/go-tsukasa-bot">
    <!-- LOGO -->
  </a>

  <h1 align="center">Tsukasa - The GO version</h1>
</p>



## Table of Contents
- [Table of Contents](#table-of-contents)
- [About the Project](#about-the-project)
- [Why "Tsukasa"?](#why-%22tsukasa%22)
- [TODOS](#todos)
- [Getting Started](#getting-started)
- [Changelog](#changelog)
- [License](#license)
- [Contact](#contact)


## About the Project
Hey, again.. this is the GO version of the [Tsukasa Bot](https://github.com/Rushifaaa/tsukasa-bot).


## Why "Tsukasa"?
Because I like the name :P

## TODOS
* Repository
  - [ ] Wiki
  
* Commands
  - [x] Ping
  - [ ] Git (Shows this repository)
  - [ ] Join
  - [ ] Disconnect
  - [ ] Play with Queue - (Youtube URL)
  - [ ] Pause
  - [ ] Skip
  - [ ] Stop
  - [ ] Volume
  - [ ] Resume
  - [ ] Terminate (Developer/Hoster Only)
  
* Moderation Stuff
  - [ ] Auto Role (User joined Server)

  * Commands
    - [ ] Admin (grant permission to a role to manage the bot)
    - [ ] Add/Remove Role
    - [ ] Kick
    - [ ] Ban
    - [x] Mute
    - [ ] Clear Chat

  
## Getting Started


**Before you start**
[**Want to try the bot?**](https://discordapp.com/oauth2/authorize?&client_id=564526337377959956&scope=bot&permissions=8)

**NEED TO KNOW**  
*You need [GO](https://golang.org/)!*

So if you decided to host your own Bot you need a Bot Token **-I assume you know how to get it, if not google it or duck it :P**

First clone this repository.  
`git clone https://github.com/Rushifaaa/tsukasa-bot`

After that just go into the Directory.  
`cd tsukasa-bot`

Now you can choose if you want the stable version or the development version.  
You are on default in the `master` branch -> **Stable Version**  
`git checkout develop` -> **Development verison**

After you selected one option you need to install the dependencies with `go get`

Now, you can start the bot with `go run src/main.go`

If you are starting for the first time it will create a `config.json` where you need to enter the data that is required.

If you try to run it one more time it will start :D

Have fun!

ps. you need to invite your bot on your server! (google it :P)
No worry i will make a Wiki where I explain it in detail - currently no time -


## Changelog

The change log is located [here](./CHANGELOG.md).


## License

Distributed under the GPL-3.0 License. See [`LICENSE`](./LICENSE) for more information.


## Contact

You can contact me on [Discord](https://discord.gg/kFqWZtv).
