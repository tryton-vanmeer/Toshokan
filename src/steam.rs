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

fn should_filter(name: &str) -> bool {
    ![
        "Proton 4.11",
        "Proton 6.3",
        "Proton 7.0",
        "Proton 8.0",
        "Proton EasyAntiCheat Runtime",
        "Proton Experimental",
        "Proton Hotfix",
        "Steamworks Common Redistributables",
        "Steam Linux Runtime - Sniper",
        "Steam Linux Runtime - Soldier",
    ]
    .contains(&name)
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
                            if should_filter(app.name.as_ref().unwrap()){
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
