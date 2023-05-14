extern crate regex;


use std::process::Command;
use regex::Regex;
use std::error::Error;


fn main() -> Result<(), Box<dyn Error>> {
    let mut n = 5;
    println!("Hello, world!");
    
    let output = Command::new("go").arg("run").arg("main.go").output()?;
    if !output.status.success() {
        bail!("Command Executed with error code");
    }
    
    
    Ok(())
}
