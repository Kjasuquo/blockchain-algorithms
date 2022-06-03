# Golang technical tests

Hello young padawan, weclome to this Golang technical test.
Your work today will be to solve different exercises and enigma 
that will prove both your level in code but also your ability to think and adapt.
Initiative is encouraged, you can add some bonuses as long as you explain your choices in comment.


## Rot 13

Google it.  
Clue: `man ascii` can help you  

Rules:
 - Implements the function in rot13/rot13.go
 - Only alphabetical characters should be impacted.
 - Run your code with `go run rot13/rot13.go`
 - Decode the sentence: `Pbatenghyngvba, lbh'er ba gur tbbq jnl gb wbva Chyfne NV !`


## Concurrency

Well done, you have resolved the first enigma.
Now we will check your real skills in Golang by using goroutine, waitgroup & channel, Ready ?

All this task have to be done in the file concurrency/concurrency.go and runned by `go run concurrency/concurrency.go`.

In first you have to implements the function `findPrimeNumber`.
The aim of this function is to find every prime number in the given interval and send the result in a channel.
As remember of your mathematics course, a Prime number is a number that is divisible only by itself or by 1.
For exemple: 5 is prime because none of these following divisions produce an integer number : 5/2=2.5, 5/3=1.66, 5/4=1.25,
but 6 is not prime because it can be divided by 2 or 3.
More info : https://en.wikipedia.org/wiki/Prime_number

If you were smart you will see that a large set of number that is easily predictable will never be prime,
this can divide your execution time by two or even three (don't loose your time on this sentense if you don't know what I mean, 
this is not the main purpose of this exercise).

This function take 4 params:
 - from & to, both are uint32 and represent the research interval
 - channel, a channel to which send each prime number you found

When you have done with this function you should fulfill the `main` with a goroutine to print every number received though the channel.

Rules:
 - You program should wait for every goroutine to finish (use waitgroup) and shouldn't deadlock or panic
 - You are not allowed to change the `findPrimeNumber` signature
 - The three `findPrimeNumber` calls have to remain in separate goroutine (but you can add code in the goroutine)

Good luck

## Mini Graph

This project have it's own README under the folder `graph`: [MiniGraph README.md](graph/README.md)