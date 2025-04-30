### Pointers and Errors

Vill nu kika på hur vi kan exponera viss del av state och structs. Här ska vi simulera en bank, och en plånbok. Användare ska ej kunna komma åt variabler i plånboken direkt, så vi gör dem privata.

Här tar de upp grejen med huruvida referensen är med pekare eller inte. Använder vi inte pekare `func (w Wallet)` så kopieras alla referenser bara. Så när funktionen kallas på är det ett helt nytt objekt man jobbar med, allt är helt nytt. Och eftersom man inte returnar något värde heller s.a.s. så bara "försvinner" detta objekt när man returnar funktionen.

`w *Wallet` är en pekare. Så egentligen borde man dereferera när man returnar, dvs returnera `return *w.balance` typ. Det betyder "dereferera pekaren dvs ta objektet, och sedan ta balance". Men Go abstraherar bort det, för det blir smidigare. Så man behöver bara returnera `w.balance`.

I fallet med `Balance`-funktionen behöver man egentligen inte returnera referens till det precisa värdet, det hade gått bra med en kopia. Men det är kutym att låta alla funktioner ha samma "method receiver type".

---

It's great to have more descriptive names for variables. For instance, here we are simulating the input of bitcoin into our wallet, but it does not appear in the code anywhere. How do we put it there? Just define a new type, like an alias for `int`. -> `type Bitcoin int`. För att skapa en Bitcoin variabel gör man då `Bitcoin(5)`, till exempel. 

---

Now we have to talk about errors. What if you try to withdraw more btc than what is available in your wallet? Let's make Withdraw return an error if this happens.

How do you return a new error? Like actually create the error? Here they do: `errors.New("error text...")`.

---

Note: Now we're writing a test in which we WANT the function to return an error. And it's pretty basic here: you just do the setup, return an error from the function that is supposed to return the error. In this case, we check that the balance is correct (no btc withdrawn from the wallet) and then do the `if err == nil` check. If err is nil, do the `t.Errorf()` thing.

If you want to compare errors you can do it with strings. If you have an error, you can "convert" it into a string with the Error() function. So if `want` is of type string, you can go `got.Error() != want`.

At first we compare if the error that the function returns is the same that we're looking for in the test function. However this is not good, you should have a ***single source of truth*** for this thing. So declare a error variable that both use instead. That way if you want to change it, you only have to do that in one place.

---

Now we are moving the helpers outside the main testing function, to make it clearer. I guess it's situational.

--- 

Installing a tool called `errcheck`, that is checking whether there are unchecked errors in my package. Very handy since it's easy to miss errors sometimes!

