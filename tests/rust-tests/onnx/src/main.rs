#![no_main]
sp1_zkvm::entrypoint!(main);

use anyhow::Result;
use candle::{DType, Device, Tensor};

// use clap::{Parser, Subcommand};

// #[derive(Subcommand, Debug, Clone)]
#[derive(Debug, Clone)]
enum Command {
    Print {
        // #[arg(long)]
        model_bytes: String,
    },
    SimpleEval {
        // #[arg(long)]
        model_bytes: String,

        // #[arg(long)]
        input_bytes: String,
    },
}

// #[derive(Parser, Debug)]
#[derive(Debug)]
// #[command(author, version, about, long_about = None)]
pub struct Args {
    // #[command(subcommand)]
    command: Command,
}

#[derive(Debug)]
enum FlattenedVec {
    U8(Vec<u8>),
    U32(Vec<u32>),
    I64(Vec<i64>),
    F32(Vec<f32>),
    F64(Vec<f64>),
}

fn flatten_to_bytes(value: &Tensor) -> Result<String, anyhow::Error> {
    let data = match value.dtype() {
        DType::U8 => value.flatten_all()?.to_vec1::<u8>().map(FlattenedVec::U8),
        DType::U32 => value.flatten_all()?.to_vec1::<u32>().map(FlattenedVec::U32),
        DType::I64 => value.flatten_all()?.to_vec1::<i64>().map(FlattenedVec::I64),
        DType::BF16 => value.flatten_all()?.to_vec1::<f32>().map(FlattenedVec::F32), // is this ok?
        DType::F16 => value.flatten_all()?.to_vec1::<f32>().map(FlattenedVec::F32),  // is this ok?
        DType::F32 => value.flatten_all()?.to_vec1::<f32>().map(FlattenedVec::F32),
        DType::F64 => value.flatten_all()?.to_vec1::<f64>().map(FlattenedVec::F64),
    }?;

    let bytes = match data {
        FlattenedVec::U8(data) => data
            .iter()
            .flat_map(|&x| x.to_ne_bytes())
            .collect::<Vec<_>>(),
        FlattenedVec::U32(data) => data
            .iter()
            .flat_map(|&x| x.to_ne_bytes())
            .collect::<Vec<_>>(),
        FlattenedVec::I64(data) => data
            .iter()
            .flat_map(|&x| x.to_ne_bytes())
            .collect::<Vec<_>>(),
        FlattenedVec::F32(data) => data
            .iter()
            .flat_map(|&x| x.to_ne_bytes())
            .collect::<Vec<_>>(),
        FlattenedVec::F64(data) => data
            .iter()
            .flat_map(|&x| x.to_ne_bytes())
            .collect::<Vec<_>>(),
    };

    Ok(format!("0x{}", hex::encode(bytes)))
}

pub fn main() {
    let model_bytes = include_str!("../model.bytes").trim().to_string();
    let input_bytes = include_str!("../input.bytes").trim().to_string();
    println!("model_bytes: {model_bytes}");
    println!("input_bytes: {input_bytes}");
    let args = Args {
        // command: Command::Print {
        //     // model_bytes: "0x0803120c6261636b656e642d746573743a4a0a120a01781201791a0474657374220452656c75120a53696e676c6552656c755a130a0178120e0a0c080112080a0208010a02080262130a0179120e0a0c080112080a0208010a02080242021006".to_string(),
        //     // input_bytes: "0x000080bf0000803f".to_string(),
        //     model_bytes,
        // },
        command: Command::SimpleEval {
            // model_bytes: "0x0803120c6261636b656e642d746573743a4a0a120a01781201791a0474657374220452656c75120a53696e676c6552656c755a130a0178120e0a0c080112080a0208010a02080262130a0179120e0a0c080112080a0208010a02080242021006".to_string(),
            // input_bytes: "0x000080bf0000803f".to_string(),
            model_bytes,
            input_bytes,
        },
    };
    println!("args: {args:?}");
    match args.command {
        Command::Print { model_bytes } => {
            let buf = hex::decode(&model_bytes[2..]).unwrap(); // skip 0x prefix
            let model = candle_onnx::read_model_bytes(buf).unwrap();
            println!("{model:?}");
            let graph = model.graph.unwrap();
            for node in graph.node.iter() {
                println!("{node:?}");
            }
        }
        Command::SimpleEval {
            model_bytes,
            input_bytes,
        } => {
            let buf = hex::decode(&model_bytes[2..]).unwrap(); // skip 0x prefix
            let model = candle_onnx::read_model_bytes(buf).unwrap();
            let graph = model.graph.as_ref().unwrap();
            let constants: std::collections::HashSet<_> =
                graph.initializer.iter().map(|i| i.name.as_str()).collect();
            let input_buf = hex::decode(&input_bytes[2..]).unwrap(); // skip 0x prefix
            let mut inputs: std::collections::HashMap<String, Tensor> =
                std::collections::HashMap::new();

            for input in graph.input.iter() {
                use candle_onnx::onnx::tensor_proto::DataType;
                if constants.contains(input.name.as_str()) {
                    continue;
                }

                let type_ = input.r#type.as_ref().expect("no type for input");
                // println!("{type_:?}");
                let type_ = type_.value.as_ref().expect("no type.value for input");
                // println!("{type_:?}");
                let value = match type_ {
                    candle_onnx::onnx::type_proto::Value::TensorType(tt) => {
                        let dt = match DataType::try_from(tt.elem_type) {
                            Ok(dt) => match candle_onnx::dtype(dt) {
                                Some(dt) => dt,
                                None => {
                                    println!(
                                        "unsupported 'value' data-type {dt:?} for {}",
                                        input.name
                                    );
                                    return;
                                }
                            },
                            type_ => {
                                println!("unsupported input type {type_:?}");
                                return;
                            }
                        };
                        let shape = tt.shape.as_ref().expect("no tensortype.shape for input");
                        println!("{:?}", shape.dim);
                        // let dims = shape.dim.iter().map(|dim| {
                        //     println!("{:?}", dim);
                        //     // match dim.value.as_ref().expect("no dim value") {
                        //     //     candle_onnx::onnx::tensor_shape_proto::dimension::Value::DimValue(v) => Ok(*v as usize),
                        //     //     candle_onnx::onnx::tensor_shape_proto::dimension::Value::DimParam(_) => Ok(42),
                        //     // }
                        //     42
                        // });
                        // // for dim in shape.dim.iter() {
                        let mut dims = Vec::new();
                        for dim in 0..2 {
                            println!("{:?}", dim);
                            dims.push(42);
                        }
                        // .collect::<Result<Vec<usize>>>();
                        // .collect::<Vec<usize>>();
                        // println!("{:?}", dims);
                        return;
                        // Tensor::from_raw_buffer(input_buf.as_slice(), dt, &dims, &Device::Cpu)
                        //     .unwrap()
                    }
                    type_ => {
                        println!("unsupported input type {type_:?}");
                        return;
                    }
                };
                // println!("input {}: {value:?}", input.name);
                // inputs.insert(input.name.clone(), value);
            }
            // // reject inputs with length > 1
            // if inputs.len() > 1 || inputs.is_empty() {
            //     panic!("only one input is supported");
            // }
            // let outputs = candle_onnx::simple_eval(&model, inputs).unwrap();
            // if outputs.len() > 1 || outputs.is_empty() {
            //     panic!("only one output is supported");
            // }
            // let (name, value) = outputs.iter().next().unwrap();
            // let output = flatten_to_bytes(value).unwrap();
            // println!("output {name}: {output:?} {value}");
        }
    }
}
