// #![no_std]
#![cfg_attr(target_arch = "riscv64", no_main)]

// use alloc::format;
use alloc::vec::Vec;
use core::mem;
use core::slice;
// use kona_common::io;

// pub const ENDIANNESS: bool = true; // big endian
pub const ENDIANNESS: bool = false; // little endian

pub const INPUT_ADDR: usize = 0x1000;
pub const MODEL_ADDR: usize = 0x2000;
pub const OUTPUT_ADDR: usize = 0x3000;
pub const MAGIC_ADDR: usize = 0x4000;

pub trait FromBytes {
    fn from_le_bytes(a: &mut &[u8]) -> Self;
    fn from_be_bytes(a: &mut &[u8]) -> Self;
}

pub trait ToBytes {
    fn to_le_bytes<const N: usize>(a: Self) -> [u8; N];
    fn to_be_bytes<const N: usize>(a: Self) -> [u8; N];
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

impl ToBytes for u64 {
    fn to_le_bytes<const N: usize>(a: Self) -> [u8; N] {
        assert_eq!(N, 8);
        let val = u64::to_le_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
    }
    fn to_be_bytes<const N: usize>(a: Self) -> [u8; N] {
        assert_eq!(N, 8);
        let val = u64::to_be_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
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

impl ToBytes for u32 {
    fn to_le_bytes<const N: usize>(a: Self) -> [u8; N] {
        assert_eq!(N, 4);
        let val = u32::to_le_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
    }
    fn to_be_bytes<const N: usize>(a: Self) -> [u8; N] {
        assert_eq!(N, 4);
        let val = u32::to_be_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
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

impl ToBytes for u8 {
    fn to_le_bytes<const N: usize>(a: Self) -> [u8; N] {
        assert_eq!(N, 1);
        let val = u8::to_le_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
    }
    fn to_be_bytes<const N: usize>(a: Self) -> [u8; N] {
        assert_eq!(N, 1);
        let val = u8::to_be_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
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

impl ToBytes for i64 {
    fn to_le_bytes<const N: usize>(a: Self) -> [u8; N] {
        assert_eq!(N, 8);
        let val = i64::to_le_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
    }
    fn to_be_bytes<const N: usize>(a: Self) -> [u8; N] {
        assert_eq!(N, 8);
        let val = i64::to_be_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
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

impl ToBytes for i32 {
    fn to_le_bytes<const N: usize>(a: Self) -> [u8; N] {
        assert_eq!(N, 4);
        let val = i32::to_le_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
    }
    fn to_be_bytes<const N: usize>(a: Self) -> [u8; N] {
        assert_eq!(N, 4);
        let val = i32::to_be_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
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

impl ToBytes for usize {
    fn to_le_bytes<const N: usize>(a: Self) -> [u8; N] {
        #[cfg(target_pointer_width = "64")]
        assert_eq!(N, 8);
        #[cfg(target_pointer_width = "32")]
        assert_eq!(N, 4);

        let val = usize::to_le_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
    }
    fn to_be_bytes<const N: usize>(a: Self) -> [u8; N] {
        #[cfg(target_pointer_width = "64")]
        assert_eq!(N, 8);
        #[cfg(target_pointer_width = "32")]
        assert_eq!(N, 4);
        let val = usize::to_be_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
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
impl ToBytes for isize {
    fn to_le_bytes<const N: usize>(a: Self) -> [u8; N] {
        #[cfg(target_pointer_width = "64")]
        assert_eq!(N, 8);
        #[cfg(target_pointer_width = "32")]
        assert_eq!(N, 4);
        let val = isize::to_le_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
    }
    fn to_be_bytes<const N: usize>(a: Self) -> [u8; N] {
        #[cfg(target_pointer_width = "64")]
        assert_eq!(N, 8);
        #[cfg(target_pointer_width = "32")]
        assert_eq!(N, 4);

        let val = isize::to_be_bytes(a);
        unsafe { *val.as_chunks_unchecked::<N>().first().unwrap() }
    }
}

pub fn write_output(data: &[u8]) -> ! {
    write_output_(data)
}

pub fn write_output_(data: &[u8]) -> ! {
    let size = data.len();
    let output_addr = OUTPUT_ADDR as *mut usize;
    unsafe {
        // let output_addr = OUTPUT_ADDR as *mut usize;
        *output_addr = size;
        let output_addr = output_addr.add(1) as *mut u8;
        let output_slice = slice::from_raw_parts_mut(output_addr, size);
        output_slice.copy_from_slice(data);
    }
    halt(0)
}

pub fn halt(exit_code: u8) -> ! {
    #[cfg(any(target_arch = "riscv32", target_arch = "riscv64"))]
    unsafe {
        core::arch::asm!(
            "ecall",
            in("a7") 93u8,
            in("a0") exit_code
        );
        unreachable!()
    }
    core::panic!("[HOST] Halted with exit code: {}", exit_code)
}

pub fn read_byte_slice(addr: &'_ usize) -> &'_ [u8] {
    read_byte_slice_(addr)
}

pub fn read_byte_slice_(addr: &'_ usize) -> &'_ [u8] {
    let size = read_numeric::<usize>(*addr);
    unsafe { slice::from_raw_parts(*addr as *const u8, size) }
}

pub fn read_byte_vec(addr: usize) -> Vec<u8> {
    read_byte_vec_(addr)
}

pub fn read_byte_vec_(addr: usize) -> Vec<u8> {
    let size = read_numeric::<usize>(addr);
    unsafe { Vec::from_raw_parts(addr as *mut u8, size, size) }
}

pub fn read_numeric_slice_<
    T: FromBytes + Copy + Sized,
    const IS_BIG_ENDIAN: bool,
    const SIZE: usize,
>(
    addr: usize,
) -> [T; SIZE] {
    // TODO: check if this is ok - it's a bit fishy but might be faster than the loop
    // unsafe { slice::from_raw_parts(addr as *mut T, size) }

    let zero = unsafe { mem::zeroed::<T>() };
    let mut slc = [zero; SIZE];
    for (i, val) in slc.iter_mut().enumerate() {
        *val = read_numeric::<T>(addr + i * mem::size_of::<T>());
    }
    slc
}

pub fn read_numeric_vec_<T: FromBytes>(addr: usize, size: usize) -> Vec<T> {
    // TODO: check if this is ok - it's a bit fishy but might be faster than the loop
    // unsafe { Vec::from_raw_parts(addr as *mut T, size, size) }

    (0..size).fold(Vec::with_capacity(size), |mut vec, i| {
        vec.push(read_numeric::<T>(addr + i * mem::size_of::<T>()));
        vec
    })
}

pub fn read_numeric<T: FromBytes>(addr: usize) -> T {
    read_numeric_::<T, ENDIANNESS>(addr)
}

pub fn numeric_to_bytes<T: ToBytes, const N: usize>(value: T) -> [u8; N] {
    numeric_to_bytes_::<T, N, ENDIANNESS>(value)
}

pub fn numeric_to_bytes_<T: ToBytes, const N: usize, const IS_BIG_ENDIAN: bool>(
    value: T,
) -> [u8; N] {
    match IS_BIG_ENDIAN {
        true => T::to_be_bytes(value),
        false => T::to_le_bytes(value),
    }
}

pub fn read_numeric_<T: FromBytes, const IS_BIG_ENDIAN: bool>(addr: usize) -> T {
    let mut raw_bytes = unsafe { slice::from_raw_parts(addr as *const u8, mem::size_of::<T>()) };
    // io::print(&format!("raw_bytes at: {:?} {:?}!\n", addr, &raw_bytes));
    match IS_BIG_ENDIAN {
        true => T::from_be_bytes(&mut raw_bytes),
        false => T::from_le_bytes(&mut raw_bytes),
    }
}
