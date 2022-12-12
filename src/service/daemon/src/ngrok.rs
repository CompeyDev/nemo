use core::time;
use std::{env, fs::File, io::{self, ErrorKind}, process::{Command, Stdio}};
use reqwest;
use ngrok;
use url::Url;
use tar::Archive;
use flate2::read::GzDecoder;

pub fn main() -> Result<(), reqwest::Error> {
    const LINUX_DOWNLOAD_URL: &'static str = "https://bin.equinox.io/c/bNyj1mQVY4c/ngrok-v3-stable-linux-amd64.tgz";
    
    
    if env::consts::OS == "linux" {
        let client = reqwest::blocking::ClientBuilder::new().timeout(time::Duration::from_secs(120)).build()?;
        println!("ngrok.rs :: Installing ngrok runtime...");
        let resp = client.get(&*LINUX_DOWNLOAD_URL).send()?;
        let mut out = File::create("ngrok.tgz").expect("failed to write ngrok runtime");
        io::copy(&mut resp.bytes().expect("invalid ngrok body").as_ref(), &mut out).expect("failed to write ngrok runtime");
    }

    match untar_archive() {
        Ok(()) => (),
        Err(e) => match e.kind() {
            ErrorKind::PermissionDenied => (),
            other => panic!("ngrok.rs :: unknown error {:?}", other)
        }
    }

    initialize_tunnel().unwrap();   

    Ok(())
}

fn initialize_tunnel() -> std::io::Result<()>{
    // Get the ngrok API key from the user's environment variables.

    let ngrok_api_key = match env::var("NGROK_API_KEY") {
        Ok(api_key) => api_key,
        Err(_e) => panic!("ngrok.rs :: Failed to fetch ngrok API key, do you have it set in environment variables?")
    };

    // Authenticate into the CLI, using the above API Key.
    Command::new("./ngrok").arg("config").arg("add-authtoken").arg(ngrok_api_key).stdout(Stdio::piped()).spawn().expect("ngrok.rs :: failed to authenticate with API key");  

    // Start the tunnel.
    let tunnel = ngrok::builder()
        .executable("./ngrok")
        .http()
        .port(8000)
        .run()?;

    let public_url: Url = tunnel.http()?.to_owned();

    println!("Tunnel is open at {:?}", public_url);


    Ok(())
}

fn untar_archive() -> Result<(), std::io::Error> {
    let tar = File::open("ngrok.tgz")?;
    let decoded = GzDecoder::new(tar);
    let mut extractor = Archive::new(decoded);

    extractor.unpack(".")?;

    Ok(())
}