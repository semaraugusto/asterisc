use core::mem;
use core::slice;
// use core::str::;
use kona_common::io;
// extern crate alloc;
use alloc::format;

// const ENDIANNESS: bool = true; // big endian
const ENDIANNESS: bool = false; // little endian

pub trait FromBytes {
    fn from_le_bytes(a: &mut &[u8]) -> Self;
    fn from_be_bytes(a: &mut &[u8]) -> Self;
}

impl<const N: usize> FromBytes for [u8; N] {
    fn from_le_bytes(a: &mut &[u8]) -> [u8; N] {
        let (int_bytes, rest) = a.split_at(N);

        let mut me = [0u8; N];
        me.copy_from_slice(int_bytes);

        *a = rest;
        me
    }
    fn from_be_bytes(a: &mut &[u8]) -> [u8; N] {
        let (int_bytes, rest) = a.split_at(N);

        let mut me = [0u8; N];
        me.copy_from_slice(int_bytes);

        *a = rest;
        me
    }
}

impl FromBytes for u64 {
    fn from_le_bytes(a: &mut &[u8]) -> u64 {
        u64::from_le_bytes(FromBytes::from_le_bytes(a))
    }
    fn from_be_bytes(a: &mut &[u8]) -> u64 {
        u64::from_be_bytes(FromBytes::from_le_bytes(a))
    }
}
impl FromBytes for u32 {
    fn from_le_bytes(a: &mut &[u8]) -> u32 {
        u32::from_le_bytes(FromBytes::from_le_bytes(a))
    }
    fn from_be_bytes(a: &mut &[u8]) -> u32 {
        u32::from_be_bytes(FromBytes::from_le_bytes(a))
    }
}
impl FromBytes for u8 {
    fn from_le_bytes(a: &mut &[u8]) -> u8 {
        u8::from_le_bytes(FromBytes::from_le_bytes(a))
    }
    fn from_be_bytes(a: &mut &[u8]) -> u8 {
        u8::from_be_bytes(FromBytes::from_le_bytes(a))
    }
}

impl FromBytes for i64 {
    fn from_le_bytes(a: &mut &[u8]) -> i64 {
        i64::from_le_bytes(FromBytes::from_le_bytes(a))
    }
    fn from_be_bytes(a: &mut &[u8]) -> i64 {
        i64::from_be_bytes(FromBytes::from_le_bytes(a))
    }
}
impl FromBytes for i32 {
    fn from_le_bytes(a: &mut &[u8]) -> i32 {
        i32::from_le_bytes(FromBytes::from_le_bytes(a))
    }
    fn from_be_bytes(a: &mut &[u8]) -> i32 {
        i32::from_be_bytes(FromBytes::from_le_bytes(a))
    }
}

impl FromBytes for usize {
    fn from_le_bytes(a: &mut &[u8]) -> usize {
        usize::from_le_bytes(FromBytes::from_le_bytes(a))
    }
    fn from_be_bytes(a: &mut &[u8]) -> usize {
        usize::from_be_bytes(FromBytes::from_le_bytes(a))
    }
}
impl FromBytes for isize {
    fn from_le_bytes(a: &mut &[u8]) -> isize {
        isize::from_le_bytes(FromBytes::from_le_bytes(a))
    }
    fn from_be_bytes(a: &mut &[u8]) -> isize {
        isize::from_be_bytes(FromBytes::from_le_bytes(a))
    }
}

pub fn read_numeric<T: FromBytes>(addr: usize) -> T {
    read_numeric_endian::<T, ENDIANNESS>(addr)
}

pub fn read_numeric_endian<T: FromBytes, const IS_BIG_ENDIAN: bool>(addr: usize) -> T {
    let mut raw_bytes = unsafe { slice::from_raw_parts(addr as *const u8, mem::size_of::<T>()) };
    io::print(&format!("raw_bytes at: {:?} {:?}!\n", addr, &raw_bytes));
    match IS_BIG_ENDIAN {
        true => T::from_be_bytes(&mut raw_bytes),
        false => T::from_le_bytes(&mut raw_bytes),
    }
}
