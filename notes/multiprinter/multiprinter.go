package main

import (
	"github.com/pterm/pterm"
	"io"
	"time"
)

func main() {
	multi := pterm.DefaultMultiPrinter
	multi.UpdateDelay = time.Millisecond * 100
	writer1 := multi.NewWriter()
	writer2 := multi.NewWriter()
	writer3 := multi.NewWriter()
	multi.Start()
	l := ss{&s{-150, &writer1}, &s{-350, &writer2}, &s{-450, &writer3}}
	i := 1
	by := func(s s) { pterm.Fprinto(*s.w, pterm.Sprintf("%d,\t%d", s.i, i)) }
	//by := func(s s) {}
	for i <= 100 {
		time.Sleep(time.Millisecond * 50)
		l.increment()
		l.visualize(by)
		i++
	}
	multi.Stop()
}

type s struct {
	i int
	w *io.Writer
}
type ss []*s

func (l *ss) increment() {
	for _, o := range *l {
		o.increment()
	}
}

func (v *s) increment() {
	v.i++
}

func (l *ss) visualize(by func(s s)) {
	for _, o := range *l {
		by(*o)
	}
}

func custom() {
	multi := pterm.DefaultMultiPrinter
	w1 := multi.NewWriter()
	w2 := multi.NewWriter()
	w3 := multi.NewWriter()
	multi.Start()
	for i := 0; i < 100; i++ {
		pterm.Fprinto(w1, i)
		pterm.Fprinto(w2, i+1000)
		pterm.Fprinto(w3, i+10000)
		time.Sleep(time.Millisecond * 500)
	}
}

func fromWebSource() {
	// Create a multi printer for managing multiple printers
	multi := pterm.DefaultMultiPrinter

	// Create two spinners with their own writers
	spinner1, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 1")
	spinner2, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 2")

	// Create five progress bars with their own writers and a total of 100
	pb1, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 1")
	pb2, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 2")
	pb3, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 3")
	pb4, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 4")
	pb5, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 5")

	// Start the multi printer
	multi.Start()

	// Increment progress bars and spinners based on certain conditions
	for i := 1; i <= 100; i++ {
		pb1.Increment() // Increment progress bar 1 every iteration

		if i%2 == 0 {
			pb2.Add(3) // Add 3 to progress bar 2 every even iteration
		}

		if i%5 == 0 {
			pb3.Increment() // Increment progress bar 3 every 5th iteration
		}

		if i%10 == 0 {
			pb4.Increment() // Increment progress bar 4 every 10th iteration
		}

		if i%3 == 0 {
			pb5.Increment() // Increment progress bar 5 every 3rd iteration
		}

		if i%50 == 0 {
			spinner1.Success("Spinner 1 is done!") // Mark spinner 1 as successful every 50th iteration
		}

		if i%60 == 0 {
			spinner2.Fail("Spinner 2 failed!") // Mark spinner 2 as failed every 60th iteration
		}

		time.Sleep(time.Millisecond * 50) // Sleep for 50 milliseconds between each iteration
	}

	// Stop the multi printer
	multi.Stop()
}
