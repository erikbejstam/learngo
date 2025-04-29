### Iteration

Bara repetition för mig själv: 
1. Skriv testet. Ha med want/got. `if want != got`...
2. Build kommer faila, för vi har inte funktionen `Repeat`. Lägg in den, mer eller mindre bara funktionssignatur + `return ""`. Detta failar, men inte för att det är inkorrekt Gokod.
3. Implementera minimal variant av funktionen. Grönt.
Refaktorera.

---

Nu provar vi lite benchmarking. Kommandot jag kör funkar, men ger inte exakt output som i tutorial. Nåval - man kan bara klicka på "benchmark"..

Anm. string concatenation i Go är långsamt pga strings är immutable - alltså är det massa kopiering som måste göras. Bättre då att använda `strings.Builder`. Om vi benchmarkar ser vi att det går många gånger fortare och effektivare.

---

TLDR: Mer TDD, skriva benchmarks, använda for-loops.