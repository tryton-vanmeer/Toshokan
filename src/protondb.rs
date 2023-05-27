use anyhow::Result;
use reqwest::{blocking::Client, header::USER_AGENT};
use serde::Deserialize;

#[derive(Debug, Deserialize)]
struct APIResult {
    tier: String,
}

pub fn get_tier(appid: String) -> Result<String> {
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

    let response: APIResult = serde_json::from_str(&response.text()?)?;

    Ok(response.tier)
}
