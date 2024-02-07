# Control Structures

## Avoid Arrow Code

- **Arrow Code:** Avoid writing code that is overly condensed or "arrow code" as it is not clean and error-prone.

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
The guard fails fast, improving code clarity and efficiency. Look for opportunities to transform code into guards, especially if it contains lots of code and if/else statements.

## Extracting Control Structures

- **Extracting Control Structures:** Apply the same principles used for extracting parts of abstraction that don’t belong in functions by separating them out into their own methods.

- **Positive Phrasing:** Use positive phrasing, such as `isEmpty` instead of `isNotEmpty`, for improved readability and clarity.

For inverting conditional logic, refer to this [example](https://github.com/academind/clean-code-course-code/tree/control-06-inverting-logic).


# Utilizing Errors

- **Throwing and Handling Errors:** Throwing and handling errors can replace if statements and lead to more focused functions.

- **If SOMETHING is an Error, Make it an ERROR:** If something is considered an error condition, treat it as such by throwing an error.

- **Separating Error Handling:** Catch errors using try/catch blocks and separate error handling logic from the main function logic. 

- **Utilize Error Guards:** Utilize error guards to handle potential errors and maintain a clear flow of control. [Example](https://github.com/academind/clean-code-course-code/blob/control-07-creating-error-guards/use-errors.js#L70-L79).

- **Error Handling is "One Thing":** Error handling should be treated as a single responsibility within a function, separate from other concerns.

# Using Factory Functions and Polymorphism

- **Factory Functions:** Factory functions are functions used to produce objects.

- **Polymorphism:** Polymorphism refers to objects or functions that can always be used in the same way. For more information, refer to [Wikipedia](https://en.wikipedia.org/wiki/Polymorphism_(computer_science)).

- **Example of Utilizing Factory Functions and Polymorphism:**
  - Factory functions and polymorphism can be used to refactor code that violates the DRY (Don't Repeat Yourself) principle. 
  - Example: [Error Handling is "One Thing"](https://github.com/academind/clean-code-course-code/blob/control-09-error-handling-is-one-thing/error-handling-is-one-thing.js#L96-L146), [Factory Functions](https://github.com/academind/clean-code-course-code/blob/control-10-factory-functions/factory-functions.js#L105-L121).

- **Using Default Parameters:** Utilize default parameters to eliminate some conditional statements and ensure parameters always have a value.
