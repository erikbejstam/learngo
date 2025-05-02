### Mocking

We'll now write a function that will use printing and the time lib to simulate a countdown. This is a little more advanced, which is why we should divide the steps up more. First, let's make the functionality of printing "3", and then all the numbers, and then implement timing.

---

After everything is implemented and timing is included, we realize that there are better ways to do things. We depend on the actual timer, that takes super long time to finish. We can't have tests running that long, especially not if the countdown is supposed to be longer for instance. So let's make Countdown have a Sleeper interface injected, and we can when testing inject a sleeper that only "fake sleeps". 

In this case we can use a ***sleeper*** which is a type of mock object. It records how many times a dependency has been called, which is handy in this case.

---

The thing now though is, what if the sleeps occurred out of order? we only know from our test that the mock sleeper was called three times, but when? Let's try to break this.