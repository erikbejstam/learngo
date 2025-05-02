### Dependency Injection

When we write code that *dependes* on something else, it is a good idea to *inject* (pass as a parameter) that thing into the code, instead of having the code create that dependency for itself. This makes stuff more modular and easier to test. 

We will try this now with a function that is going to do a greeting print. Now, fmt.Print is actually hard to test, so instead of using that directly, we will pass a writer into the function. This makes it so that when we want to test our greeting function we can pass a writer that is easy to test, and in prod we can pass stdout.

For writing the first test function, it is good to know that the `bytes.Buffer` implements the Writer interface, so let's use that.

What we do here, at least, is we use the `Fprintf` function which puts a string into the writer/buffer. What it does I guess is it converts the string into bytes, which the buffer wants.

Först definierar vi Greet så att den tar emot en pekare till en buffer. Buffer implementerar writer-interfacet, så Fprintf tar emot den. Men om vi nu faktiskt vill printa ut via stdout, så funkar det inte, för det är inte en Buffer. Här kommer principen *accept interface, return struct* in. Så det är bättre att acceptera en Writer, som både Buffer och Stdout implementerar.

Anm. Man behöver sällan skicka en pekare till ett interface som parameter. Så vi tar därför bort stjärnan bredvid argumentet i funktionen Greet. Har lite svårt att ingående förklara varför. Men såhär: man behöver inte definiera en funktion som att den ska ha en pekare till ett interface, vet inte ens när man skulle använda det. När man vill använda ett objekt så är det just ett objekt och inte ett interface. Därför är parametern bara ett interface. Sen när denna funktion ANVÄNDS, så skickar man in en referens till ett specifikt objekt. Det är typ så jag tänker nu. Men en *djup* förståelse har jag inte än. 

Then they give an example of why it is nice to write functions in this way. Namely, using it as a print function on the browser. For that you use `http.ResponseWriter` that implements Writer. So you have a handler func for a specific route and that handler runs Greet with the ResponseWriter and you get the greeting on the browser page.

Now the Go stdlib had code we could use for testing, the buffer, but further on we might want to *mock* stuff, that is create our own fake versions of a dependency, which can be really useful.