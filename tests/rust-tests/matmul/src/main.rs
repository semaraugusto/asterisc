#![no_main]
sp1_zkvm::entrypoint!(main);

// use candle_core::{Device, Tensor};
use candle::{Device, Tensor};

// fn main() -> Result<(), Box<dyn std::error::Error>> {
fn main() {
    println!("Starting execution");
    let device = Device::Cpu;

    let a = Tensor::randn(0f32, 1., (2, 3), &device).unwrap();
    let b = Tensor::randn(0f32, 1., (3, 4), &device).unwrap();
    println!("a = {a}");
    println!("b = {b}");

    let c = a.matmul(&b).unwrap();
    println!("RESULT = {c}");
}
