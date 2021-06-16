package main


import (
    "fmt" 
    "github.com/bwmarrin/discordgo"

)
func main(){
    bot, err := discordgo.New("Bot " + ("token"))

    if err != nil {
        panic(err)
    }

    
    bot.AddHandler(ready)
    bot.AddHandler(messageCreate)

    err = bot.Open()

	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}
    
	fmt.Println("Bot online. Cierra la consola para apagarme")
	for {}

	bot.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready){
    s.UpdateStatus(0, "presencia")
    fmt.Println("logged in as user " +string(s.State.User.ID))
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
  

    if m.Content == "ping"{
        s.ChannelMessageSend(m.ChannelID, "pong")
    }
    
}
