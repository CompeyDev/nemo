use libc::{c_char, size_t};
use std::{string::String, ffi::CStr, mem::forget, collections::HashMap};
use ureq;
use std::time::Duration;
use daemon;

// Initial strategy: https://play.rust-lang.org/?version=stable&mode=debug&edition=2021&gist=b4bd989d2765453faf81424c9b8a4769

#[no_mangle]
pub extern "C" fn add_queue(connection_uri: *const c_char, connection_uri_size: size_t, payload_id: *const c_char, payload_id_size: size_t, task_type: *const c_char, task_type_size: size_t, task: *const c_char, task_size:size_t) {
    // Convert raw C strings into Rust strings
    logger::info("start", true);
    let sconnection_uri = unsafe { String::from_raw_parts(connection_uri as *mut u8, connection_uri_size, 20) };
    let spayload_id = unsafe { String::from_raw_parts(payload_id as *mut u8, payload_id_size, 20) };
    let stask_type = unsafe { String::from_raw_parts(task_type as *mut u8, task_type_size, 20) };
    let stask = unsafe { String::from_raw_parts(task as *mut u8, task_size, 20) };
    
    println!("{sconnection_uri}");

    logger::info("strings", true);
    // Prepare request body
    
    let tt = spayload_id.as_str();
    let tn = stask.as_str();
    let pi = spayload_id.as_str();

    logger::info("as strings", true);

    // TODO: coerce json_str to &str

    let json_str = f "\"TaskType\": {stask_type}, \"Task\": {stask}, \"PayloadID\": {spayload_id}";
    println!("{:#?}", json_str);

    if Some(json_str) == None {
       ureq::post(&sconnection_uri)
        .set("content-type", "application/json")
        .send_string(&json_str)
        .expect("failed to parse into string"); 
    }
    
    logger::info("creates body", true);
        
    // Prevent Rust from deallocating these pointers

    forget(sconnection_uri);
    forget(spayload_id);
    forget(stask_type);
    forget(stask);  
    logger::info("success!", true);
}

#[no_mangle]
pub extern "C" fn get_tunnel() {
    daemon::main();
}
