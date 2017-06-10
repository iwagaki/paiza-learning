
fn main() {
    use std::io;
    let mut line = String::new();
    let _ = io::stdin().read_line(&mut line);
    //let split : Vec<&str> = line.as_str().split(' ').collect();
    let list : Vec<&str> = line.split_whitespace().collect();
    let a : i32 = list[0].trim().parse::<i32>().unwrap();
    let b : i32 = list[1].trim().parse::<i32>().unwrap();

    println!("{}", a + b);
}
