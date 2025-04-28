# Hello, World

Separera sk domain code fr. det som är "utanför" dvs sidoeffekter. Ett enkelt exempel är print-funktioner, som ju "går" utanför själva programmet sas. 

`const` might be useful! Just to be completely sure and clear that something has a specific value. So if I want my Hello function to always print "Hello, ", then I can define a const `helloPrefix` that it uses, to make it super clear.

Anm: Det är viktigt att testen specificerar *tydligt* vad koden ska åstadkomma.

Det är bra att refaktorera sin testkod också! Om vi har flera tester kring samma funktion, kan vi till exempel göra själva checken `if got != want {...}` i en egen helper-funktion. 

När man använder hjälp-funktioner i sin testkod är det bra att acceptera ett interface nämligen `testing.TB`. Detta är så att funktionen kan användas både vid test (testing.T), fuzzing (testing.F) och benchmark (testing.B).

Det är bra att lägga till `t.Helper()` i hjälpmetoderna, för om ett test failar, så kommer den hålla reda på vilken metod som använde hjälpmetoden - blir ju förvirrande annars. 

---

Anm hot tip: Man kan skriva in namnet på return-variabeln på en funktion såhär `func myFunc(x int) (y int)`, och då behöver man bara ha return, inte skriva `return y`. Lite snyggare bara.

Anm hot tip: Snyggare att gruppera konstaner (`const`) i ett eget block, tycker jag. 

---

De tre stegen för **TDD**: 
1. Skriv ett test med tydligt syfte, den ska *ringa in* vad din funktion ska göra. Testet ska faila, men se till att den ger en tydlig beskrivning av hur testet failar. 
2. Skriva tillräckligt för att testet ska pass:a. 
3. Refaktorera, gör koden clean och fin. 