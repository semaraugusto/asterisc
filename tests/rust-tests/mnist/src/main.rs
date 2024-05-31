#![no_main]
sp1_zkvm::entrypoint!(main);

// use anyhow::Error as E;
// use candle::
use candle::{DType, Device, Result, Tensor};
use candle_nn::{Linear, Module, VarBuilder};
// use rand::prelude::*;
use std::collections::HashMap;

struct LinearModel {
    linear: Linear,
}

const IMAGE_DIM: usize = 784;
const LABELS: usize = 10;

trait Model: Sized {
    fn new(vs: VarBuilder) -> Result<Self>;
    fn forward(&self, xs: &Tensor) -> Result<Tensor>;
}

impl Model for LinearModel {
    fn new(vs: VarBuilder) -> Result<Self> {
        let linear = linear_z(IMAGE_DIM, LABELS, vs)?;
        Ok(Self { linear })
    }

    fn forward(&self, xs: &Tensor) -> Result<Tensor> {
        self.linear.forward(xs)
    }
}

fn linear_z(in_dim: usize, out_dim: usize, vs: VarBuilder) -> Result<Linear> {
    let ws = vs.get_with_hints((out_dim, in_dim), "weight", candle_nn::init::ZERO)?;
    let bs = vs.get_with_hints(out_dim, "bias", candle_nn::init::ZERO)?;
    Ok(Linear::new(ws, Some(bs)))
}

fn main() {
    println!("Starting");
    let device = Device::Cpu;
    let dtype = DType::F32;

    let mut tensors: HashMap<String, Tensor> = HashMap::new();
    println!("empty tensor hashmap created");

    // let weight = Tensor::randn(0f32, 1., (IMAGE_DIM, LABELS), &device).unwrap();
    let weight = Tensor::randn(0f32, 1., (LABELS, IMAGE_DIM), &device).unwrap();
    println!("Weight loaded. Shape: {:?}", weight.shape());
    let bias = Tensor::randn(0f32, 1., LABELS, &device).unwrap();
    println!("Bias loaded. Shape: {:?}", bias.shape());
    tensors.insert("weight".to_string(), weight);
    tensors.insert("bias".to_string(), bias);

    let vb = VarBuilder::from_tensors(tensors, dtype, &device);

    // let input = Tensor::randn(0f32, 1., (IMAGE_DIM, 1), &device).unwrap();
    let input = Tensor::randn(0f32, 1., (1, IMAGE_DIM), &device).unwrap();
    println!("Input loaded. Shape: {:?}", input.shape());

    let model = LinearModel::new(vb).unwrap();

    let logits = model.forward(&input).unwrap();
    println!("Pred made. Shape: {:?}", logits.shape());
    println!("Pred made: {}", logits);
    let digit = logits.argmax(1).unwrap();
    println!("Digit: {}", digit);
}
