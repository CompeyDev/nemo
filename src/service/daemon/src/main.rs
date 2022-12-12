// #![no_main]
use std::error::Error;
use reqwest;
mod ngrok; 

async fn add_queue(task_type: String, task: String) -> Result<(), Box<dyn Error>> {
    println!("daemon.rs :: added new task to tasks pool");
    const CONNECTION_URI: &'static str = "http://0.0.0.0:40043";
    let client = reqwest::Client::new(); 
    
    client.post(CONNECTION_URI.to_owned() + "/addQueue")
        .send()
        .await?;
    
    Ok(())
}

fn main() {
    ngrok::main().unwrap();
}
