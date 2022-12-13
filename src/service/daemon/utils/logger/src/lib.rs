use colored::*;
use chrono;

// pub struct Options {
//     print_log: Option<bool>,
//     custom_color: Option<&str>,
// }

#[export_name = "logger::info"]
pub fn info(log: &str, should_print: bool) -> Option<String> {
    // if options == None {
    //     options.print_log = Some(true)
    // }

    let now = chrono::Local::now().to_string();
    let formatted_time_vec: Vec<&str> = now.split(" +").collect();
    let formatted_time = formatted_time_vec[0];
    let to_log = format!("{} {} :: {}", formatted_time.truecolor(128, 128, 128), "info".green(), log);
    if should_print == true {
        println!("{to_log}");
        return None;
    } else {
        return Some(to_log);
    }
}

#[export_name = "logger::info_return"]
pub fn info_return(log: &str) -> String {
    let now = chrono::Local::now().to_string();
    let formatted_time_vec: Vec<&str> = now.split(" +").collect();
    let formatted_time = formatted_time_vec[0];
    let to_log = format!("{} {} :: {}", formatted_time.truecolor(128, 128, 128), "info".green(), log);
    return to_log;
}

#[export_name = "logger::error"]
pub fn error(log: &str, should_panic: bool) -> Option<String>{
    let now = chrono::Local::now().to_string();
    let formatted_time_vec: Vec<&str> = now.split(" +").collect();
    let formatted_time = formatted_time_vec[0];
    let to_log = format!("{} {} :: {}", formatted_time.truecolor(128, 128, 128), "error".red(), log);
    if should_panic == true {
        panic!("{to_log}");
    } else {
        return Some(to_log);
    }
}

#[export_name = "logger::error_return"]
pub fn error_return(log: &str) -> String {
    let now = chrono::Local::now().to_string();
    let formatted_time_vec: Vec<&str> = now.split(" +").collect();
    let formatted_time = formatted_time_vec[0];
    let to_log = format!("{} {} :: {}", formatted_time.truecolor(128, 128, 128), "error".red(), log);
    return to_log;
}