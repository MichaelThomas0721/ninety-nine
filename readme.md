# Ninety-Nine (Card Game) Strategy Simulator

This is a simulator for the card game Ninety-Nine. I was interested in finding the best strategy for the game and wanted an excuse to learn GoLang. I've made some strategies in the strategies folder. The strategies are in the form of json files, they work by looking at the current score and prioritizing some cards over others. I have also added the ability to use a priority list only if a dependency is filled, in this case the only dependency I have added it number of special cards in your hand. I have some plans for some more dependencies such as the current turn or how many turns it has been since the score went over a certain threshold. I also have plans to eventually make a self adjusting algorithm. This is the first project I have done in GoLang so my apologies if it is a little messy or suboptimal. Currently it takes about 375us to simulate 1 game with 12 players, it gets shorter if you have less players and longer if you have more. I am always looking to improve the efficiency so if you see something that could help improve the speed then let me know :)


# Setup

This project requires `GoLang 1.21` or higher so be sure you have it installed
As long as you have GoLang 1.21 installed then you should be able to just download it and run it with the `go run *.go` command