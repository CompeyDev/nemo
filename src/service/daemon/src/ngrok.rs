use core::time;
use std::{env, str, fs::File, io::{self, ErrorKind, Write}, process::{Command, Stdio, exit}, sync::mpsc::channel, fmt::Error};
use reqwest;
use ngrok::{self, Tunnel};
use url::Url;
use ctrlc;
use tar::Archive;
use logger;
use flate2::read::GzDecoder; 

pub fn main() -> Result<String, std::io::Error> {
    const LINUX_DOWNLOAD_URL: &'static str = "https://bin.equinox.io/c/bNyj1mQVY4c/ngrok-v3-stable-linux-amd64.tgz";
    
    // TODO: fix trait bounds for resp.bytes()
    if env::consts::OS == "linux" {
        let client = reqwest::blocking::ClientBuilder::new().timeout(time::Duration::from_secs(120)).build().expect(logger::error_return("failed to build http client").as_str());
        logger::info("Installing ngrok runtime...", true);
        let mut resp = client.get(&*LINUX_DOWNLOAD_URL).send().expect(logger::error_return("failed to request ngrok runtime").as_str());
        let mut out = File::create("ngrok.tgz").expect(logger::error_return("failed to write ngrok runtime").as_str());
        let mut buf_writer = vec![];
        resp.copy_to(&mut buf_writer).unwrap();
        out.write(&buf_writer as &[u8]).unwrap();
        // io::copy(&mut resp.bytes().expect("hi"), &mut out).expect("failed to write ngrok runtime");
    }

    match untar_archive() {
        Ok(()) => (),
        Err(e) => match e.kind() {
            ErrorKind::PermissionDenied => (),
            other => panic!("{}", logger::error_return(format!("unknown error {:#?}", other).as_str()).as_str())
        }
    }

    let conn = initialize_tunnel().unwrap();
    
    // Handle errors first, then use the unchecked method to actually return a result

    conn.public_url()?.to_owned();

    let conn_uri = conn.public_url_unchecked();
    conn_uri.domain().ok_or_else(|| "failed to get tunnel domain");



    let (tx, rx) = channel();
    
    ctrlc::set_handler(move || { tx.send(()).unwrap(); }).expect(logger::error_return("failed to set exit handler").as_str());

    rx.recv().expect("failed to gracefully exit");
    // Hackily "clear" the line
    print!("\r  ");

    // Overwrite the empty spaces
    println!("\r{}", logger::info_return("Gracefully exiting daemon..."));

    drop(conn);
    
    exit(0);
}

fn initialize_tunnel() -> std::io::Result<Tunnel>{
    // Get the ngrok API key from the user's environment variables.

    let ngrok_api_key = match env::var("NGROK_API_KEY") {
        Ok(api_key) => api_key,
        Err(_e) => panic!("{}", logger::error_return("Failed to fetch ngrok API key, do you have it set in environment variables?"))
    };

    logger::info(format!("Using ngrok API token {}", ngrok_api_key).as_str(), true);

    // Authenticate into the CLI, using the above API Key.
    Command::new("./ngrok").arg("config").arg("add-authtoken").arg(ngrok_api_key).stdout(Stdio::piped()).spawn().expect("ngrok.rs :: failed to authenticate with API key");  

    // Start the tunnel.
    let tunnel = ngrok::builder()
        .executable("./ngrok")
        .https()
        .port(40043)
        .run()?;
    
    let r: Url = tunnel.public_url_unchecked().to_owned();
    let public_url = r.domain().unwrap();

    logger::info(format!("Tunnel is open at {:}", public_url).as_str(), true);
    
    Ok(tunnel)
}

fn untar_archive() -> Result<(), std::io::Error> {
    let tar = File::open("ngrok.tgz")?;
    let decoded = GzDecoder::new(tar);
    let mut extractor = Archive::new(decoded);

    extractor.unpack(".")?;

    Ok(())
}