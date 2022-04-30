use crate::steam;
use crate::steam::Game;

use cursive::{views::*, Cursive};

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

    let mut info = ListView::new();
    info.add_child("App ID", TextView::new("123456"));

    let layout = LinearLayout::horizontal()
        .child(Panel::new(build_game_list()))
        .child(Panel::new(info));

    siv.add_layer(layout);
    siv.run();
}
