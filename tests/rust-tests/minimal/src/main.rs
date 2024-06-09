#![no_main]
sp1_zkvm::entrypoint!(main);

fn main() {
    println!("Hello, world!");
    let mut v = Vec::new();
    for i in 0..10 {
        v.push(i)
    }
    println!("Hello, vec! {:?}", v);
}
