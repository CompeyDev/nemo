mod lib;
use libc;
use std::ffi::{CString, CStr};


/// Test the C module internally.
fn main() {
    lib::add_queue(CString::new("data").expect("CString::new failed").as_ptr(), 5, CString::new("data").expect("CString::new failed").as_ptr(), 5, CString::new("data").expect("CString::new failed").as_ptr(), 5, CString::new("data").expect("CString::new failed").as_ptr(), 5);
}
