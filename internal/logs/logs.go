package logs

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GenerateLog(file, layout string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	sum := 0
	rand.Seed(time.Now().UnixNano())
	for {
		sum++
		n := 0
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			n++
			line := strings.Replace(scanner.Text(), "PREFIX", time.Now().Format(layout), 1)
			fmt.Printf("%s line=^_^%d-%d^_^\n", line, sum, n)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
		f.Seek(0, 0)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
