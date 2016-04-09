package main
import(
    "fmt"
    "math/rand"
    "time"
    "sort"
)

func random(max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max)
}

func idxReturn(src [][]string, a int) string {
    b_len := len(src[a])
    b := random(b_len)  
    return src[a][b]
}

func userSearch(src []string, u string) bool {
	sort.Strings(src)
	i := sort.SearchStrings(src, u)
	if i < len(src) && src[i] == u {
		return true
	}
	return false
}

func main() {
    var res_group int

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
    
    yes_men := []string{
    	"joyce",
    	"joe",
    }
    
    user := "barry"
    
    if userSearch(yes_men, user) {
		res_group = 0
    } else {
	    res_group = random(len(res_list))
    }
    
    res_final := idxReturn(res_list, res_group)
    fmt.Println(res_final)
}
