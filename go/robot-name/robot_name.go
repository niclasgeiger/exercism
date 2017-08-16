package robotname

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const testVersion = 1

var (
	alphabet = strings.Split("A B C D E F G H I J K L M N O P Q R S T U V W X Y Z", " ")
	r        = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Robot struct {
	old   string
	named bool
	name  string
}

func (r *Robot) Name() string {
	name := r.name
	if !r.named {
		name = getRandomName()
		for r.name != r.old {
			name = getRandomName()
		}
		r.name = name
	}
	r.named = true
	return name
}

func (r *Robot) Reset() {
	r.old = r.name
	r.named = false
}

func getRandomName() string {
	letters := fmt.Sprintf("%s%s", alphabet[r.Intn(len(alphabet))], alphabet[r.Intn(len(alphabet))])
	return fmt.Sprintf("%s%d%d%d", letters, r.Intn(10), r.Intn(10), r.Intn(10))
}
