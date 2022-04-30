use std::{path::PathBuf, str::FromStr};

use steamlocate::{SteamDir, SteamApp};

pub struct Game {
    pub name: String,
    pub appid: u32,
    path: PathBuf
}

impl Game {
    pub fn from_steamapp(app: &SteamApp) -> Self {
        Self {
            name: app.name.as_ref().unwrap().to_string(),
            appid: app.appid,
            path: app.path.clone()
        }
    }

    pub fn get_path(self) -> String {
        self.path.into_os_string().into_string().unwrap()
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