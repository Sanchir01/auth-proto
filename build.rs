use std::{env, path::PathBuf};
use dotenvy::dotenv;
fn main() -> Result<(), Box<dyn std::error::Error>> {
    dotenv().ok().expect("Could not load .env file");
    let out_dir = PathBuf::from(env::var("OUT_DIR").unwrap());
    println!("cargo:warning=OUT_DIR is: {}", out_dir.display());
    let file_descriptor_path = out_dir.join("descriptor.bin");
    tonic_build::configure()
        .file_descriptor_set_path(file_descriptor_path)
        .build_server(true)
        .out_dir(out_dir.to_str().unwrap())
        .compile_protos(&[
            "proto/metrics/metrics.proto",
            "proto/order/order.proto",
            "proto/auth/auth.proto"
        ], &["./"])?;
    
    Ok(())
}