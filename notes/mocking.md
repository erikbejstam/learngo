### Mocking

We'll now write a function that will use printing and the time lib to simulate a countdown. This is a little more advanced, which is why we should divide the steps up more. First, let's make the functionality of printing "3", and then all the numbers, and then implement timing.

---

After everything is implemented and timing is included, we realize that there are better ways to do things. We depend on the actual timer, that takes super long time to finish. We can't have tests running that long, especially not if the countdown is supposed to be longer for instance. So let's make Countdown have a Sleeper interface injected, and we can when testing inject a sleeper that only "fake sleeps". 

In this case we can use a ***sleeper*** which is a type of mock object. It records how many times a dependency has been called, which is handy in this case.

---

The thing now though is, what if the sleeps occurred out of order? we only know from our test that the mock sleeper was called three times, but when? Let's try to break this. So you can sleep 3 times in a row, then do the prints, and the tests will pass. So we have to have a better test.
Our Countdown function has two dependencies, and we essentially have to see that these work in the correct order, so we should create some type of mock that "swallows" them both. So we create an object `SpyCountdownOperations` that has a list of strings that records every operation sleep/write in the order they were made.

Jag beskriver bara återigen hur detta gjordes för mig själv. Vi har alltså en funktion som beror på två saker. Vi behöver se till att den kallar på dessa dependencies i rätt ordning. Men vi hinner inte använda de riktiga funktionerna för dessa, det tar för lång tid. Ok. Då gör vi en mock dependency som vår riktiga funktion kan använda istället för BÅDA dependencies. Då kan den notera att allt sker i rätt ordning genom att spara varje kallelse i en lista.

---

Another thing that would be nice is to have Sleeper have a cofigurable sleep time. Let's create a `ConfigurableSleeper`.

Now we start passing functions and stuff, which always kinda confuses me a little bit.

But as an answer to the above, now that I've had time to reflect a little bit on it. What they are showing in the tutorial I think is two different ways of "mocking". You could have a `Sleeper` interface, and inject MockSleeper.Sleep vs time.Sleep, I guess. But what they are doing, which is maybe a little cleaner? is they're mocking Time. And then you just pass different "time" modules instead of different "Sleepers" into your sleeper struct. Something like that. 

The `ConfigurableSleeper` struct almost becomes like an interface (but it isn't) in that it can be used in different ways, i.e. you can inject different sleeper methods into it. From what I understand, this type of DI is maybe preferable in smaller cases, but using an interface preferable in larger programs. I atleast think it *feels like* they're trying to show different ways of doing things. I don't think everything in this lessons is supposed to be "used" together.

--- 

I've also had trouble this lesson with when to use pointers etc. Especially because Go sometimes automatically dereferences pointers and stuff I think, which gets confusing. But basically you should think about whether the methods on the struct have pointer receivers or not (if they do, pass pointer). You should not have variable names like `bufPtr` or something, because of this automatic dereferencing and stuff. GPT says "In Go you shouldn't be too autistic about pointers" basically. In general the following kind of seems to be the most used way to do things:

```
{
    s := &SpyCountdownOperations{}
    buf := &bytes.Buffer{}
    Countdown(buf, s)
}
```

That is, when creating a struct, immediately save it in the variable as a pointer to it. The Countdown function doesn't explicitly accept pointers but it does the automatic dereferencing, so you can just pass the pointer straight into the function.

---

People are sometimes weary of mocking. You should think about certain things relating to this:

> If your mocking code is becoming complicated or you are having to mock out lots of things to test something, you should listen to that bad feeling and think about your code. Usually it is a sign of

>    The thing you are testing is having to do too many things (because it has too many dependencies to mock)

>>        Break the module apart so it does less

>    Its dependencies are too fine-grained

>>        Think about how you can consolidate some of these dependencies into one meaningful module

>    Your test is too concerned with implementation details

>>        Favour testing expected behaviour rather than the implementation

> Normally a lot of mocking points to bad abstraction in your code.