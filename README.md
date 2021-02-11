# Muse
**Muse** is a telegram bot that forwards message from some Channels to the receivers.  
This project is still under development. 
More features will be added in the future.  
The table below shows the current status of the features: 

| Feature | Text | Photo / Video | Audio | File | Media Group
| :-----: | :---: | :---: | :---: | :---: | :---: |
| Forward | ✔ | ✔ | ✔ | ✔ | ✔ |
| Reply   | ❌ | ❌ | ❌ | ❌ | ❌ |
| Edit    | ❌ | ❌ | - | - | ❌ |

All of these features will be implemented.  

## Overview  
| Bot API Wrapper | Conf Manager | Log Manager |
| :-----: | :-------------------: | :--------: |
| [telegram-bot-api v5](https://github.com/go-telegram-bot-api/telegram-bot-api/tree/bot-api-5.0) | [spf13/viper](https://github.com/spf13/viper) | [sirupsen/logrus](https://github.com/sirupsen/logrus)

## Getting Started
0. Apply for a Telegram Bot from BotFather.
1. Download the source code files and compile it.   
2. Create a config file under the same directory of the binary file, and set it correctly.  
3. Run. Enjoy.  

## Configuration
A configuration file named `conf.yaml` needs to be placed under the same directory of the binary file.  
The configuration file needs to be in `YAML` format.
Here is an example:

```yaml
bot:
  token: TokenStingOfYourBot  # you need to get this from @BotFather
  debug: false                # decide whether the program is in debug mode
rule:                  # the ID and channel-username you want to forward from/to
  -1004998307033:      # sender
    - -1001699850137   # receiver
    - -10052614894123  # receiver
  5821739913:          # sender
    - -1001699850137   # receiver
```

The senders are always Telegram channels, and the receivers may vary:
channels, groups, or maybe some users.  
If you set sender as a user, or a group, then the program will simply ignore it.  

### Understanding the "rule" field in configuration file

The following `YAML` piece:  

```yaml
rule:
  1:
    - 2
    - 3
  2:
    - 3
  3: 
    - 1
    - 4
  5:
    - 6
  7:
    - 6
```

Means the following forward relation:  

![Explain relation using a graph](https://i.loli.net/2021/02/11/t7hUjJcfBAYKeQi.png)
