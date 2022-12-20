// #![no_main]
use std::error::Error;
use reqwest;
mod ngrok; 
use logger;

async fn add_queue(task_type: String, task: String) -> Result<(), Box<dyn Error>> {
    logger::info("adding new task to tasks pool", true);
    const CONNECTION_URI: &'static str = "http://0.0.0.0:40043";
    let client = reqwest::Client::new(); 
    
    client.post(CONNECTION_URI.to_owned() + "/addQueue")
        .send()
        .await?;
    
    Ok(())
}

fn main() {
    ngrok::main();
}


pub fn get_tunnel() -> Result<String, std::io::Error> {
    let tunnel_uri = ngrok::main()?;
    
    Ok(tunnel_uri)
}
