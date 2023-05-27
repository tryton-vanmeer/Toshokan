use std::fmt::Display;

use anyhow::Result;
use colored::Colorize;
use reqwest::{blocking::Client, header::USER_AGENT};
use serde::Deserialize;

#[derive(Debug, Deserialize)]
pub struct Summary {
    pub tier: String,
}

impl Summary {
    pub fn from_appid(appid: u32) -> Result<Self> {
        let url = format!(
            "https://www.protondb.com/api/v1/reports/summaries/{}.json",
            appid
        );

        let response = Client::new()
            .get(url)
            .header(
                USER_AGENT,
                format!("{}/{}", env!("CARGO_PKG_NAME"), env!("CARGO_PKG_VERSION")),
            )
            .send()?;

        let summary: Summary = serde_json::from_str(&response.text()?)?;

        Ok(summary)
    }
}

impl Display for Summary {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self.tier.as_str() {
            "borked" => write!(f, "{}", &self.tier.red()),
            "bronze" => write!(f, "{}", &self.tier.yellow()),
            "silver" => write!(f, "{}", &self.tier.white()),
            "gold" => write!(f, "{}", &self.tier.yellow().bold()),
            "platinum" => write!(f, "{}", &self.tier.white().bold()),
            _ => f.write_str(&self.tier),
        }
    }
}
