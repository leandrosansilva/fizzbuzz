# FizzBuzz Enterprise Edition - Go version

Then the other day I realized I had never implemented FizzBuzz in my whole life,
as it never came as a question in a job interview.

I then decided to implement it in idiomatic Go, using a naive algorithm, making it as abstract as possible.

Think on it as some kind of FizzBuzz enterprise edition, as verbose as it can get,
but without relying on the typical non-sense you find in other enterprise FizzBuzz.

It's still a work in progress, as there's no documentation and the code is not reusable.

Using the naive algorithm I've managed to improve the performance from 2MiB/s to 200MiB,
using a single core and without manual loop unrolling. The improvevents were mostly
related on bufferizing as much data as possible in order to invoke Write() as rerely as possible
on os.Stdout.

I've been tracking the performance with:

```sh
go run . 0 | pv
```

This is very far from the GiB/s that you can find at https://codegolf.stackexchange.com/questions/215216/high-throughput-fizz-buzz/

Well, I tried. It's enterprise.

As I spent already too much time in this toy, I'm finished changing it for this life :-)
