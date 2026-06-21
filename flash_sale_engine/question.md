Omo, you are completely right! I went a bit too generous with the implementation steps. Let's strip away all the hints and architectural instructions.

Let's look at this purely from a high-level business requirement perspective. No more talking about buffered channels, tickers, or selectors. You design the engine yourself from scratch.

---

## The Rephrased Challenge: The High-Traffic Flash Sale Engine

### The Business Reality

You are the Lead Backend Engineer at an e-commerce startup. Tonight is Black Friday. You have exactly **5 units of a high-end phone** left in the warehouse. At midnight, **50 customers** are going to hit your purchase endpoint at the exact same microsecond.

### Your Requirements:

1. **Inventory Guard:** Exactly 5 phones must be sold. If your system sells 6, the company loses money. If it sells 4 because of a glitch, the company loses money.
2. **The 2-Second Checkout SLA:** Processing a payment is slow because you are calling an external API (like Paystack). If that external API takes longer than **2 seconds** for any customer, that customer's checkout session must immediately cancel/expire so that the stock can be freed up for the next person inline.
3. **The Live Dashboard:** While this chaos is happening, the business team needs a background monitoring system that logs how many phones are left in stock **every 500 milliseconds** until the sale is completely over.

### The Strict Constraints:

* **Constraint A:** You are **strictly forbidden** from using standard structural locks like `sync.Mutex` or `sync.RWMutex`.
* **Constraint B:** You must simulate the 50 concurrent buyers and the slow, erratic payment gateway using the tools we discussed.
