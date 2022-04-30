use crate::steam;
use crate::steam::Game;

use cursive::{views::{Dialog, SelectView}, Cursive};

fn build_game_list() -> SelectView<Game> {
    let mut list = SelectView::new();

    for game in steam::get_games() {
        list.add_item(game.name.to_string(), game);
    }

    return list;
}

pub fn run() {
    let mut siv = cursive::default();

    siv.add_global_callback('q', Cursive::quit);

    let list = build_game_list();

    let dialog = Dialog::new()
        .title("Games")
        .content(list);

    siv.add_layer(dialog);
    siv.run();
}
