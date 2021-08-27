## Src main 3rd libs:

### - Bot service

- https://github.com/bwmarrin/discordgo

- https://github.com/go-sql-driver/mysql

### - Web service

- https://github.com/gofiber/fiber/v2

- https://gorm.io/driver/mysql

- https://gorm.io/gorm

## Deploy wakeup & usage info (in progress):

- https://gopherbot0.herokuapp.com

## Adoption, tests, contrib:

- https://discord.io/go-latam

![](https://awebytes.files.wordpress.com/2021/08/gopherbot.png)

## TODO:

- [ ] Algoritmo de respuesta para challenges aleatorios dependiendo de: 

    1. **level**: ["easy", "medium", "hard"] 
    2. **challenge_type**: ["algorithm", "concurrency, "database", "web", "cli", "core"]
    3. **id** hint: get random id between total records count