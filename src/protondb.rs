use anyhow::Result;
use reqwest::{blocking::Client, header::USER_AGENT};
use serde::Deserialize;

#[derive(Debug, Deserialize)]
pub struct Summary {
    tier: String,
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
