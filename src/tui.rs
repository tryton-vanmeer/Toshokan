use crate::steam;
use crate::steam::Game;

use cursive::{views::*, Cursive, theme::*, theme::{PaletteColor::*, Color::*, BaseColor::*}, traits::{Scrollable}};

fn set_theme(siv: &mut Cursive) {
    let mut theme = siv.current_theme().clone();
    let mut palette = Palette::default();

    palette[Background] = TerminalDefault;
    palette[View] = TerminalDefault;
    palette[Primary] = TerminalDefault;
    palette[Secondary] = TerminalDefault;

    theme.shadow = false;
    theme.borders = BorderStyle::Simple;
    theme.palette = palette;

    siv.set_theme(theme);
}

fn build_game_list() -> SelectView<Game> {
    let mut list = SelectView::new();

    for game in steam::get_games() {
        list.add_item(game.name.to_string(), game);
    }

    return list;
}

fn build_game_info() -> ListView {
    let mut info = ListView::new();

    info.add_child("App ID", TextView::new("123456"));

    return info;
}

pub fn run() {
    let mut siv = cursive::default();
    set_theme(&mut siv);

    siv.add_global_callback('q', Cursive::quit);

    let list = build_game_list();

    let info = build_game_info();

    let layout = LinearLayout::horizontal()
        .child(Panel::new(list.scrollable()))
        .child(Panel::new(info));

    siv.add_fullscreen_layer(layout);
    siv.run();
}
