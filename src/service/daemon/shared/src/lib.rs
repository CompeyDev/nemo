use std::collections::HashMap;
use std::error::Error;
use reqwest;
use logger;
use libc;
use foreign_string::FString;

// &'static str is buggy with async functions, so using &str instead
#[no_mangle]
pub async extern fn add_queue(connection_uri: *const u8, connection_uri_size: usize, payload_id: *const u8, payload_id_size: usize, task_type: *const u8, task_type_size: usize, task: *const u8, task_size: usize) -> Result<(), Box<dyn Error>> {
    // Convert to C compatible strings

    let sconn_uri = unsafe { FString::new(connection_uri, connection_uri_size) };
    let spayload_id = unsafe { FString::new(payload_id, payload_id_size) };
    let stask_type = unsafe { FString::new(task_type, task_type_size) };
    let stask = unsafe { FString::new(task, task_size) };

    // Prepare request body
    let mut body = HashMap::new();
    body.insert("TaskType", stask_type.as_str());
    body.insert("TaskName", stask.as_str());
    body.insert("PayloadID", spayload_id.as_str());

    // logger::info("adding new task to tasks pool", true);

    // Make the request
    reqwest::Client::new()
        .post(format!("{}/{}", sconn_uri.as_str(), "addQueue"))
            .json(&body)
            .send()
            .await
            .expect(logger::error_return("failed to add task to tasks pool").as_str());
    
    // logger::info("successfully added task to tasks pool", true);

    Ok(())
}