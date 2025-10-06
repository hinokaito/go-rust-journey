fn add(a: i32, b: i32) -> i32 {
    a + b
}

fn sub(a: i32, b: i32) -> i32 {
    a - b
}

fn mul(a: i32, b: i32) -> i32 {
    a * b
}

fn div(a: i32, b: i32) -> i32 {
    a / b
}

fn main() {
    loop {
        println!("数字を空白区切りで入力してください: ");
        let (a, b) = {
            let mut input = String::new();
            std::io::stdin()
                .read_line(&mut input)
                .expect("読み込みに失敗しました");
            
            let mut numbers = input
                .split_whitespace()
                .map(|s| s.parse::<i32>().expect("数値の解析に失敗しました"));
            
            (numbers.next().unwrap(), numbers.next().unwrap())
        };

        println!("演算子を入力してください(終了するにはqを入力): ");
        let operator = {
            let mut input = String::new();
            std::io::stdin()
                .read_line(&mut input)
                .expect("読み込みに失敗しました");
            input.trim().to_string()
        };

        match operator.as_str() {
            "+" => println!("{}", add(a, b)),
            "-" => println!("{}", sub(a, b)),
            "*" => println!("{}", mul(a, b)),
                "/" => println!("{}", div(a, b)),
                "q" => break,
                _ => println!("無効な演算子です"),
        }
    }

}
