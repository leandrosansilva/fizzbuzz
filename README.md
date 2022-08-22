# FizzBuzz Enterprise Edition - Go version

Then the other day I realized I had never implemented FizzBuzz in my whole life,
as it never came as a question in a job interview.

I then decided to implement it in idiomatic Go, making it as abstract as possible.

Think on it as some kind of FizzBuzz enterprise edition, as verbose as it can get,
but without relying on the typical non-sense you find in other enterprise FizzBuzz.

It's still a work in progress, as there's no documentation and the code is not reusable.

One think to notice is that it's very slow, and I could reach 16MiB/s in my laptop.

The main reason for such slowness is that each number requires a call to Write() in an object
of type io.Writer.

This is very far from the GiB/s that you can find at https://codegolf.stackexchange.com/questions/215216/high-throughput-fizz-buzz/

Well, I tried. It's enterprise.
