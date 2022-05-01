use directories::UserDirs;
use steamlocate::{SteamDir, SteamApp};

pub struct Game {
    pub name: String,
    pub appid: u32,
    path: String
}

impl Game {
    pub fn from_steamapp(app: &SteamApp) -> Self {
        Self {
            name: app.name.as_ref().unwrap().to_string(),
            appid: app.appid,
            path: app.path.clone().into_os_string().into_string().unwrap()
        }
    }

    pub fn path(&self) -> String {
        let dirs = UserDirs::new().unwrap();
        let home = dirs.home_dir().to_str().unwrap();

        if self.path.starts_with(home) {
            return self.path.replace(home, "~");
        }

        self.path.clone()
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