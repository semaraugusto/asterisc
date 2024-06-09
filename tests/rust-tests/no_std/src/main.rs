#![no_std]
#![cfg_attr(any(target_arch = "mips", target_arch = "riscv64"), no_main)]

use kona_common::io;
use kona_common_proc::client_entry;

extern crate alloc;
use alloc::format;
pub mod utils;

#[client_entry(0xFFFFFFF)]
fn main() -> Result<()> {
    io::print("Hello, world!\n");

    let one_mem = utils::read_numeric::<u64>(0x1000);
    let ten_mem = utils::read_numeric::<u64>(0x1008);
    io::print(&format!("one_mem {:?}!\n", &one_mem));
    io::print(&format!("ten_mem {:?}!\n", &ten_mem));

    // let one = assert_eq!
    assert_eq!(one_mem, 1u64);
    assert_eq!(ten_mem, 10u64);
    Ok(())
}
