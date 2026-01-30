# Go Monad: Maybe, Result, and Either

A lightweight, type-safe implementation of functional programming monads
for Go 1.18+. This library brings the safety of Haskell, Elm, and Gren to Go,
helping you reduce `if err != nil` boilerplate and handle optionality
without `nil` pointers.

---

## 🚀 Features

* **Maybe[T]**: Handle optional values without `nil`.
* **Result[T]**: Railway-oriented error handling using the standard Go `error` interface.
* **Either[A, B]**: Represent disjoint unions (Left or Right) for complex branching.
* **Type Safe**: Built from the ground up with Go Generics.
* **Zero Dependencies**: Pure Go logic.

---

## 📦 Installation

```bash
go get github.com/gilramir/gomonad

```

---

## 🛠 Usage

### 1. Maybe Monad

Use `Maybe` when a value might be missing, such as a database lookup or an optional configuration.

```go
func findUser(id int) Maybe[string] {
    if id == 42 {
        return Just("Alice")
    }
    return Nothing[string]()
}

// Chaining with Map
name := findUser(42).Map(func(s string) string {
    return strings.ToUpper(s)
})

fmt.Println(name.Unpack("Unknown")) // ALICE

```

### 2. Result Monad

Use `Result` to chain operations that might fail. This follows the
"Railway Oriented Programming" pattern, where an error at any step
bypasses the rest of the chain.

```go
func parseAge(input string) Result[int] {
    age, err := strconv.Atoi(input)
    if err != nil {
        return Err[int](err)
    }
    return Ok(age)
}

// Chain operations without checking errors at every step
res := BindResult(parseAge("25"), func(age int) Result[int] {
    if age < 18 {
        return Err[int](errors.New("too young"))
    }
    return Ok(age)
})

```

### 3. Either Monad

Use `Either` when you have two valid paths that aren't necessarily "success" or "failure."

```go
// Swap values or fold them into a single result
val := Right[string, int](100)
swapped := Swap(val) // Becomes Left[int, string](100)

result := Fold(swapped,
    func(i int) string { return "Got a number" },
    func(s string) string { return "Got a string" },
)

```

---

## 📑 Core Concepts

| Type | Intent | Success Path | Failure/Alt Path |
| --- | --- | --- | --- |
| **Maybe** | Presence | `Just(v)` | `Nothing()` |
| **Result** | Fallibility | `Ok(v)` | `Err(error)` |
| **Either** | Choice | `Right(b)` | `Left(a)` |

### Why Standalone Functions?

In Go, methods cannot introduce new type parameters. To allow transforming
a `Maybe[int]` into a `Maybe[string]`, we use package-level functions like
`Map` and `Bind` rather than struct methods.

---

