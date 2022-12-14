use std::collections::HashMap;
use std::error::Error;
use reqwest;
use logger;

#[no_mangle]
async fn add_queue(connection_uri: &'static String, payload_id: String, task_type: String, task: String) -> Result<(), Box<dyn Error>> {
    // Prepare request body
    let mut body = HashMap::new();
    body.insert("TaskType", task_type);
    body.insert("TaskName", task);
    body.insert("PayloadID", payload_id);

    logger::info("adding new task to tasks pool", true);

    // Make the request
    reqwest::Client::new()
        .post(connection_uri.to_owned() + &"/addQueue".to_string())
            .json(&body)
            .send()
            .await
            .expect(logger::error_return("failed to add task to tasks pool").as_str());
    
    logger::info("successfully added task to tasks pool", true);

    Ok(())
}