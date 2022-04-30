mod steam;

fn main() {
    let games = steam::get_games();

    for game in games {
        println!("{}  {}", game.name, game.appid)
    }
}
