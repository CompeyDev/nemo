mod lib;

fn main() {
    let connection_uri = unsafe { libc::strdup(b"0.0.0.0:40043\0".as_ptr() as _) as *const u8 };
    let payload_id = unsafe { libc::strdup(b"PAYLOAD_ID\0".as_ptr() as _) as *const u8 };
    let task_type = unsafe { libc::strdup(b"EXEC\0".as_ptr() as _) as *const u8 };
    let task = unsafe { libc::strdup(b"whoami\0".as_ptr() as _) as *const u8 };

    // TODO: deal with return values, maybe create a blocking system?
    lib::add_queue(connection_uri, 13, payload_id, 10, task_type, 4, task, 6);
}