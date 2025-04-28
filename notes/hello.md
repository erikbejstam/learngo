# Hello, World

Separera sk domain code fr. det som är "utanför" dvs sidoeffekter. Ett enkelt exempel är print-funktioner, som ju "går" utanför själva programmet sas. 

`const` might be useful! Just to be completely sure and clear that something has a specific value. So if I want my Hello function to always print "Hello, ", then I can define a const `helloPrefix` that it uses, to make it super clear.

Anm: Det är viktigt att testen specificerar *tydligt* vad koden ska åstadkomma.

Det är bra att refaktorera sin testkod också! Om vi har flera tester kring samma funktion, kan vi till exempel göra själva checken `if got != want {...}` i en egen helper-funktion. 

När man använder hjälp-funktioner i sin testkod är det bra att acceptera ett interface nämligen `testing.TB`. Detta är så att funktionen kan användas både vid test (testing.T), fuzzing (testing.F) och benchmark (testing.B).

Det är bra att lägga till `t.Helper()` i hjälpmetoderna, för om ett test failar, så kommer den hålla reda på vilken metod som använde hjälpmetoden - blir ju förvirrande annars. 

---

