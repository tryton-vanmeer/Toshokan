use crate::steam;
use crate::steam::Game;

use cursive::traits::Nameable;
use cursive::utils::markup::StyledString;
use cursive::{views::*, Cursive, theme::*, theme::{PaletteColor::*, Color::TerminalDefault}, traits::{Scrollable}};

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
    let mut list = SelectView::new().autojump();

    list.set_on_select(|s, item| {
        let content = build_game_info(item);
        s.call_on_name("info", |v: &mut TextView| {
            v.set_content(content);
        })
        .unwrap();
    });

    for game in steam::get_games() {
        list.add_item(game.name.to_string(), game);
    }

    return list;
}

fn build_game_info(game: &Game) -> StyledString {
    let mut styled = StyledString::styled("App ID ", Effect::Bold);
    styled.append(StyledString::plain(game.appid.to_string()));

    return styled;
}

pub fn run() {
    let mut siv = cursive::default();
    set_theme(&mut siv);

    siv.add_global_callback('q', Cursive::quit);

    let info = TextView::new("").with_name("info");
    let list = build_game_list();

    let layout = LinearLayout::horizontal()
        .child(Panel::new(list.scrollable()))
        .child(Panel::new(info));

    siv.add_fullscreen_layer(layout);
    siv.run();
}