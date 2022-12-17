# readmegen

Narzędzie do generowania zawartości pliku `README.md` w repozytorium `pastaandpizza/recipes`

## Użycie

> **Uwaga**: Narzędzie przeznaczone do użycia w GitHub Actions!

| Flaga | Opis | Wartość domyślna |
|-|-|-|
| `-template` | Plik z szablonem pliku `README.md`, składającym się ze zdefiniowanych sekcji `header`, `category` i `footer` | `readme.gotmpl` |
| `-out` | Nazwa pliku wynikowego | `README.md` |

```bash
./readmegen \
-template readme.gotmpl \
-out README.md \
pl-recipe-1.md pl-recipe-2.md ... pl-recipe-n.md
```

## Błędy

Jakby coś się działo to będzie pluł co go boli na `stdout`
