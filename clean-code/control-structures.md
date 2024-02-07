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
