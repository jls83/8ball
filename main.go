package main
import(
    "fmt"
    "math/rand"
    "time"
)

func random(max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max)
}

func resReturn(src [][]string, a int) string {
    b_len := len(src[a])
    b := random(b_len)  
    return src[a][b]
}

func main() {
    res_list := [][]string{
        { // Positive responses
        "It is certain",
        "It is decidedly so",
        "Without a doubt",
        "Yes, definitely",
        "You may rely on it",
        "As I see it, yes",
        "Most likely",
        "Outlook good",
        "Yes",
        "Signs point to yes",
        },
        { // Neutral responses
        "Reply hazy try again",
        "Ask again later",
        "Better not tell you now",
        "Cannot predict now",
        "Concentrate and ask again",
        },
        { // Negative responses
        "Don't count on it",
        "My reply is no",
        "My sources say no",
        "Outlook not so good",
        "Very doubtful",
        },
    }
    res_group := random(len(res_list))
    res_final := resReturn(res_list, res_group)
    fmt.Println(res_final)
}
