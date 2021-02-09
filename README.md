# Muse
**Muse** is a telegram bot that forwards message from some Channels to the receivers.  
This project is still under development. 
More features will be added in the future.  
The table below shows the current status of the features: 

| Feature | Text | Photo / Video | Audio | File | Media Group
| :-----: | :---: | :---: | :---: | :---: | :---: |
| Forward | ✔ | ✔ | ✔ | ✔ | as separated messages |
| Reply   | ❌ | ❌ | ❌ | ❌ | ❌ |
| Edit    | ❌ | ❌ | - | - | ❌ |

All of these features will be implemented.  

# Overview  
| Bot API Wrapper | Configuration Manager |
| :-----: | :-------------------: |
| [telegram-bot-api v5](https://github.com/go-telegram-bot-api/telegram-bot-api/tree/bot-api-5.0) | [viper](https://github.com/spf13/viper) |

Imagine forwarding procedure as a *complete bipartite graph*. A complete bipartite graph `K_(3,5)` is shown below:  
![The complete bipartite graph K_(3,5)](https://i.loli.net/2021/01/31/kAvLZMEIKSbNPXG.png)  
The blue dots represent senders, while the red dots represent receivers. 
In the graph shown above, if one sender sends a message, every receiver shall receive it. 
In this project, the senders are always Telegram channels, and the receivers may vary: channels, groups, or maybe some users.  

# Configuration
A configuration file named `conf.yaml` needs to be placed under the same directory of the binary file.  
The configuration file needs to be in `YAML` format. 
Here is an example: 

```yaml
bot:
  token: TokenStingOfYourBot  # you need to get this from @BotFather
  debug: false                # decide whether the program is in debug mode
forward:                      # the ID and channel-username you want to forward from/to
  src: [-1004998307033, 5821739913]        # senders
  dest: [-1001699850137, -10052614894123]  # receivers
  # these configurations are also legal:
  # src: -1001699850137
  # dest: 5821739913
  # or
  # src: [5821739913]
  # dest: [-1001699850137]
```

In this project, the sender must be channels. 
If you put the `ID` of a user or a group into the `src` field, the program will ignore it.  
If you only need to forward from/to one source, you don't even need a YAML array.  

# Getting Started
0. Apply for a Telegram Bot from BotFather.
1. Download the source code files and compile it.   
2. Create a config file under the same directory of the binary file, and configure it correctly.  
3. Run. Enjoy.  
