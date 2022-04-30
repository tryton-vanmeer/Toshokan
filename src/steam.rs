use steamlocate::{SteamDir, SteamApp};

pub struct Game {
    pub name: String,
    pub appid: u32
}

impl Game {
    pub fn from_steamapp(app: &SteamApp) -> Self {
        Self {
            name: "test".to_string(),
            appid: 123456
        }
    }
}

pub fn get_games() -> Vec<Game> {
    let mut games = Vec::new();

    let mut steamdir = SteamDir::locate().unwrap();
    let apps = steamdir.apps();

    for (_, app) in apps {
        match app {
            Some(app) => games.push(Game::from_steamapp(app)),
            None => ()
        }
    }

    return games;
}