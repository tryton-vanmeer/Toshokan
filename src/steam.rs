use std::{env, path::PathBuf};

use anyhow::{Context, Result};
use steamlocate::{App, Library, SteamDir};

trait DisplayHomeAsTilde {
    fn display_home_as_tilde(&self) -> String;
}

impl DisplayHomeAsTilde for PathBuf {
    fn display_home_as_tilde(&self) -> String {
        let path = self.display().to_string();

        let home = env::var("HOME").unwrap();
        if path.starts_with(&home) {
            return path.replace(&home, "~");
        }

        path
    }
}

#[derive(Debug, Clone)]
pub struct Game {
    pub appid: u32,
    pub name: String,
    pub proton: bool,
    path: PathBuf,
}

impl Game {
    fn from_steamapp(app: &App, library: &Library) -> Self {
        let proton = match app.user_config.get("platform_override_source") {
            Some(platform) => platform == "windows",
            _ => false,
        };

        Self {
            appid: app.app_id,
            name: app.name.clone().unwrap(),
            proton,
            path: library.resolve_app_dir(app),
        }
    }

    pub fn path(&self) -> String {
        self.path.display_home_as_tilde()
    }

    pub fn proton_prefix(&self) -> String {
        let mut path = self.path.clone();

        path.pop(); // game name
        path.pop(); // common
        path.push("compatdata");
        path.push(PathBuf::from(self.appid.to_string()));

        path.display_home_as_tilde()
    }
}

fn should_filter(id: u32) -> bool {
    [
        1113280, // Proton 4.11
        1420170, // Proton 5.13
        1580130, // Proton 6.3
        1887720, // Proton 7.0
        2348590, // Proton 8.0
        1826330, // Proton EasyAntiCheat Runtime
        1493710, // Proton Experimental
        2180100, // Proton Hotfix
        1391110, // Steam Linux Runtime 2.0 (soldier)
        1628350, // Steam Linux Runtime 3.0 (sniper)
        228980,  // Steamworks Common Redistributables
    ]
    .contains(&id)
}

pub fn get_games() -> Result<Vec<Game>> {
    let steam = SteamDir::locate().context("unable to find steamdir")?;
    let mut games: Vec<Game> = Vec::new();

    for library in steam.libraries()? {
        match library {
            Err(err) => eprintln!("failed reading library: {err}"),
            Ok(library) => {
                for app in library.apps() {
                    match app {
                        Err(err) => eprintln!("failed reading app: {err}"),
                        Ok(app) => {
                            if !should_filter(app.app_id) {
                                games.push(Game::from_steamapp(&app, &library))
                            }
                        }
                    }
                }
            }
        }
    }

    games.sort_by_key(|game| game.name.to_string());

    Ok(games)
}

pub fn get_game(appid: u32) -> Result<Game> {
    let steam = SteamDir::locate().context("unable to find steamdir")?;
    let (app, library) = steam
        .find_app(appid)?
        .context(format!("unable to find game with appid {}", appid))?;

    Ok(Game::from_steamapp(&app, &library))
}
