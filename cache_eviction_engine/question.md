### The Scenario: Memory-Safe Transaction Cache

Imagine your backend is receiving hundreds of transactions. To save database speed, you store these transactions in a Go map called a "Cache."

But memory is expensive. If the cache gets too full, your server will crash. You need to build an eviction machine that loops through a list of active transaction pointers and **destroys/wipes the data inside memory** for any transaction that has a status of `"failed"` or `"cancelled"`, freeing up space instantly.

Here is the twist: **You must modify the data directly in memory via pointers while looping through a slice of pointers.** If you make a mistake, you will either leak memory or cause a runtime crash.

---

### Your Strict Assignment

#### 1. The Setup (No code clues!)

Create a `Transaction` struct with three fields:

* `id` (string)
* `amount` (float64)
* `status` (string — e.g., `"pending"`, `"completed"`, `"failed"`, `"cancelled"`)

#### 2. The Core Machine: `EvictBadTransactions`

Write a function called `EvictBadTransactions`.

* It must accept a **slice of pointers to Transaction structs** (`[]*Transaction`).
* It should not return anything (`void`).

**Inside this function, you must (Think Deep Here):**

* Loop through the slice of pointers using a `for range` loop.
* For each transaction pointer, check its status.
* **The Brain Melt Rule:** If the status is `"failed"` or `"cancelled"`, you must completely clear that transaction's data from the computer's memory. To do this, reach into the pointer and set the struct fields to their absolute zero values:
* Set `id` to `""`
* Set `amount` to `0.0`
* Set `status` to `"evicted"`


* If the status is `"pending"` or `"completed"`, leave it completely alone.

#### 3. Write the `main` Function

To test this, you must build your input slice carefully because it holds **addresses**, not raw values.

Create 4 instances of your struct, then create a slice containing the **addresses** of those 4 instances:

1. `t1` $\rightarrow$ `{id: "tx_01", amount: 2500.0, status: "completed"}`
2. `t2` $\rightarrow$ `{id: "tx_02", amount: 15000.0, status: "failed"}`
3. `t3` $\rightarrow$ `{id: "tx_03", amount: 450.0, status: "pending"}`
4. `t4` $\rightarrow$ `{id: "tx_04", amount: 9000.0, status: "cancelled"}`

Pass that slice of pointers into `EvictBadTransactions`.

After the function runs, print out the original `t1`, `t2`, `t3`, and `t4` variables inside `main` individually.

---

### The Final Validation Test

If your pointer logic is 100% correct, printing `t2` and `t4` inside `main` should show they have been completely wiped out and reset to `{"", 0.0, "evicted"}`—even though your function returned absolutely nothing!

This is the real deal, bro. You are dealing with a slice of pointers (`[]*Transaction`), reading values through those pointers, and modifying the original memory blocks directly.

Let's see how you handle this memory engine. Fire down!