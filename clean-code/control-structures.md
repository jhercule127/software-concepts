# Control Structures

## Avoid Arrow Code

- **Arrow Code:** Avoid writing code that is overly condensed or "arrow code" as it is not clean and error-prone.

```javascript
const result = (condition1 && condition2) ? 'trueResult' : (condition3 || condition4) ? 'falseResult1' : 'falseResult2';
// can instead be 

let result;

if (condition1 && condition2) {
  result = 'trueResult';
} else if (condition3 || condition4) {
  result = 'falseResult1';
} else {
  result = 'falseResult2';
}


```

## Useful Concepts

- **Avoid Deep Nesting:** Refrain from deep nesting in control structures as it can decrease readability and maintainability.

- **Use Factory Functions & Polymorphism:** Utilize factory functions and polymorphism to improve code structure and maintainability.

- **Prefer Positive Checks:** Prefer positive checks over negative checks for better readability and understanding of code logic.

- **Utilize Errors:** Utilize error handling mechanisms to handle exceptional cases and improve code robustness.

## Using Guards

- **Guards:** Simplify control structures by using guards, which invert the if check and return early if the condition is not met. This approach avoids executing unnecessary code.

Example:
```javascript
if (!email.includes('@')) { 
    return;
}
// do stuff
```

>[!TIP]
> The guard fails fast, improving code clarity and efficiency. Look for opportunities to transform code into guards, especially if it contains lots of code and if/else statements.

## Extracting Control Structures

- **Extracting Control Structures:** Apply the same principles used for extracting parts of abstraction that don’t belong in functions by separating them out into their own methods.

- **Positive Phrasing:** Use positive phrasing, such as `isEmpty` instead of `isNotEmpty`, for improved readability and clarity.

For inverting conditional logic, refer to this [latest example](https://github.com/academind/clean-code-course-code/blob/control-06-inverting-logic/04-extract-functions.js#L60-L120).
vs [old example](https://github.com/academind/clean-code-course-code/blob/control-05-writing-clean-code/03-extract-functions.js)


# Utilizing Errors

- **Throwing and Handling Errors:** Throwing and handling errors can replace if statements and lead to more focused functions.

- **If SOMETHING is an Error, Make it an ERROR:** If something is considered an error condition, treat it as such by throwing an error.

- **Separating Error Handling:** Catch errors using try/catch blocks and separate error handling logic from the main function logic. 

- **Utilize Error Guards:** Utilize error guards to handle potential errors and maintain a clear flow of control. [Example](https://github.com/academind/clean-code-course-code/blob/control-07-creating-error-guards/use-errors.js#L70-L79).

- **Error Handling is "One Thing":** Error handling should be treated as a single responsibility within a function, separate from other concerns.
Check out [Error Handling is "One Thing"](https://github.com/academind/clean-code-course-code/blob/control-09-error-handling-is-one-thing/error-handling-is-one-thing.js#L69-L95)

# Using Factory Functions and Polymorphism

- **Factory Functions:** Factory functions are functions used to produce objects.

- **Polymorphism:** Polymorphism refers to objects or functions that can always be used in the same way. For more information, refer to [Polymorphism Wikipedia](https://en.wikipedia.org/wiki/Polymorphism_(computer_science)).

- **Example of Utilizing Factory Functions and Polymorphism:**
  - Factory functions and polymorphism can be used to refactor code that violates the DRY (Don't Repeat Yourself) principle. 
  - Example:, [Factory Functions](https://github.com/academind/clean-code-course-code/blob/control-10-factory-functions/factory-functions.js#L105-L121).

>[!TIP]
> **Using Default Parameters:** Utilize default parameters to eliminate some conditional statements and ensure parameters always have a value.

[Default Parameters](https://github.com/academind/clean-code-course-code/blob/control-11-default-parameters/more-improvements.js#L62-L66)


## Avoid "Magic" Numbers and Strings
There's one additional tweak you might want to make to the code. Sometimes you have hard-coded string identifiers like "PAYPAL" or "CREDIT_CARD".

You typically want to avoid this - not primarily for readability reasons but to avoid errors (e.g. due to typos).

Hence it is a better practice to use globally defined enums - if your programming language supports that - or globally defined constants.

```javascript
const TYPE_CREDIT_CARD = 'CREDIT_CARD';
// ...
if (transaction.type === TYPE_CREDIT_CARD) { ... }
```


## Final Product
This shows the final demo and how to split the work into their own separate directories
[Final Demo](https://github.com/academind/clean-code-course-code/tree/control-12-finished/final-demo)