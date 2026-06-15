### The Scenario: The Ultimate Fintech Fraud Detector

Imagine you run a fintech company in Nigeria. Every second, hundreds of transfers are passing through your system. You need to build a single, reusable "Security Machine" function.

This machine itself **does not know what a fraudulent transaction looks like**. Instead, it just loops through a list of transactions, accepts a **Custom Guard Function** as data, and uses that guard function to block bad transactions.

Because fraud patterns change every single day, the security team needs to be able to write new guard rules on the fly and plug them into your machine without you rewriting your core loop.

---

### Your Strict Assignment

#### 1. The Input Data Structure

Create a `Transaction` struct with three fields:

* `username` (string)
* `amount` (float64)
* `location` (string — e.g., `"Lagos"`, `"London"`, `"Abuja"`)

#### 2. Create the Custom Function Type

Define a custom function type called `FraudGuard`.

* It must accept a single `Transaction` struct as an argument.
* It must return a `bool` (`true` if the transaction is fraudulent/bad, `false` if it is clean/safe).

#### 3. The Core Machine: `FilterTransactions`

Write a function called `FilterTransactions`.

* It must take two parameters:
1. A **slice of `Transaction` structs** (`[]Transaction`).
2. A **`FraudGuard` function** passed as data.


* It must return a **new slice of `Transaction` structs** containing *only* the transactions that were flagged as fraudulent (where the guard function returned `true`).

**Inside this function, you must:**

* Initialize an empty slice to hold the bad transactions. (Think about how to pre-allocate capacity if you can!).
* Use a `for range` loop to walk through the transactions.
* Pass the current transaction into your `FraudGuard` function. If it returns `true`, append that transaction to your bad slice.
* Return the bad slice.

#### 4. Write the `main` Function

Create a slice containing these 5 raw transactions:

1. `{username: "Tunde", amount: 5000.0, location: "Lagos"}`
2. `{username: "Chidi", amount: 6000000.0, location: "Abuja"}`
3. `{username: "Funmi", amount: 12000.0, location: "London"}`
4. `{username: "Tunde", amount: 4500.0, location: "Lagos"}`
5. `{username: "Osas", amount: 8000000.0, location: "Lagos"}`

Now, inside `main`, create **two separate anonymous functions** stored in variables of type `FraudGuard`:

* **Guard Rule 1 (`highAmountGuard`):** Flags any transaction where the amount is strictly greater than `5,000,000.0` (High-value alert).
* **Guard Rule 2 (`foreignLocationGuard`):** Flags any transaction where the location is NOT equal to `"Lagos"` and NOT equal to `"Abuja"` (International travel alert).

#### 5. Run the Machine

Call your `FilterTransactions` machine twice in `main`:

* Run it once passing the `highAmountGuard`. Print the results.
* Run it again passing the `foreignLocationGuard`. Print the results.

---

### The Brain Twist (Why this is beautiful)

Notice that your `FilterTransactions` function is written **only once**, but by passing different functions into it like raw data, it completely changes its behavior on the fly! That is the pure painkiller power of first-class functions.

No code snippets from me. Set up your struct, your function type, your filtering loop, and your two rule variables completely by yourself. Fire down, bro!