#![no_main]
use std::error::Error;
use reqwest;

async fn add_queue() -> Result<(), Box<dyn Error>> {
    const CONNECTION_URI: &'static str = "http://0.0.0.0:40043";
    let client = reqwest::Client::new(); 
    
    client.post(CONNECTION_URI.to_owned() + "/addQueue")
        .send()
        .await?;
    Ok(())
}
