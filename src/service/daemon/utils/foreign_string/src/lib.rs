use libc;

// Reference: https://stackoverflow.com/questions/70719073/how-to-convert-const-u8-with-length-to-str-without-reallocation

pub struct FString {
    data: *const u8,
    length: usize,
}

impl FString {
    // safety: data must point to nul-terminated memory allocated with malloc()
    pub unsafe fn new(data: *const u8, length: usize) -> FString {
        // Note: no reallocation happens here, we use `str::from_utf8()` only to
        // check whether the pointer contains valid UTF-8.
        // If panic is unacceptable, the constructor can return a `Result`
        if std::str::from_utf8(std::slice::from_raw_parts(data, length)).is_err() {
            panic!("invalid utf-8")
        }
        FString { data, length }
    }

    pub fn as_str(&self) -> &str {
        unsafe {
            // from_utf8_unchecked is sound because we checked in the constructor
            std::str::from_utf8_unchecked(std::slice::from_raw_parts(self.data, self.length))
        }
    }
}

impl Drop for FString {
    fn drop(&mut self) {
        unsafe {
            libc::free(self.data as *mut _);
        }
    }
}