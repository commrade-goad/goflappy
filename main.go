package main 

func main() {
    game := Game{};

    game.width = 720;
    game.height = 720;

    game.init_rl();
    defer game.close_rl();
    game.init_game_objects();
    game.game_loop();
}
