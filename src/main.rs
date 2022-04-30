#[macro_use] extern crate prettytable;
use prettytable::{Table, Row, Cell, Attr, format};

mod steam;

fn main() {
    let games = steam::get_games();

    let mut table = Table::new();
    table.set_format(*format::consts::FORMAT_CLEAN);

    table.set_titles(
        Row::new(vec![
            Cell::new("App ID")
            .with_style(Attr::Bold),
            Cell::new("Name")
            .with_style(Attr::Bold),
            Cell::new("Install Directory")
            .with_style(Attr::Bold)
        ])
    );

    for game in games {
        table.add_row(row![game.appid, game.name, game.path()]);
    }

    table.printstd();
}
