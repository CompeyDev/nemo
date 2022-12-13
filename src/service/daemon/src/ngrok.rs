use core::time;
use std::{env, fs::File, io::{self, ErrorKind}, process::{Command, Stdio}};
use reqwest;
use ngrok;
use url::Url;
use tar::Archive;
use flate2::read::GzDecoder; 

pub fn main() {
    const LINUX_DOWNLOAD_URL: &'static str = "https://bin.equinox.io/c/bNyj1mQVY4c/ngrok-v3-stable-linux-amd64.tgz";
    
    // TODO: fix trait bounds for resp.bytes()
    if env::consts::OS == "linux" {
        let client = reqwest::blocking::ClientBuilder::new().timeout(time::Duration::from_secs(120)).build().expect(logger::error_return("failed to build http client").as_str());
        logger::info("Installing ngrok runtime...", true);
        let resp = client.get(&*LINUX_DOWNLOAD_URL).send().expect(logger::error_return("failed to request ngrok runtime").as_str());
        let mut out = File::create("ngrok.tgz").expect(logger::error_return("failed to write ngrok runtime").as_str());
        io::copy(&mut resp.bytes().expect("hi"), &mut out).expect("failed to write ngrok runtime");
    }

    match untar_archive() {
        Ok(()) => (),
        Err(e) => match e.kind() {
            ErrorKind::PermissionDenied => (),
            other => panic!("{}", logger::error_return(format!("unknown error {:#?}", other).as_str()).as_str())
        }
    }

    initialize_tunnel().unwrap();   
}

fn initialize_tunnel() -> std::io::Result<()>{
    // Get the ngrok API key from the user's environment variables.

    let ngrok_api_key = match env::var("NGROK_API_KEY") {
        Ok(api_key) => api_key,
        Err(_e) => panic!("ngrok.rs :: Failed to fetch ngrok API key, do you have it set in environment variables?")
    };

    println!("ngrok.rs :: Using ngrok API token {ngrok_api_key}");
    logger::info(format!("Using ngrok API token {}", ngrok_api_key).as_str(), true);

    // Authenticate into the CLI, using the above API Key.
    Command::new("./ngrok").arg("config").arg("add-authtoken").arg(ngrok_api_key).stdout(Stdio::piped()).spawn().expect("ngrok.rs :: failed to authenticate with API key");  

    // Start the tunnel.
    let tunnel = ngrok::builder()
        .executable("./ngrok")
        .port(40043)
        .run()?;

    let public_url: Url = tunnel.https()?.to_owned();

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