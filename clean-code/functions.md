# Functions and Methods

## What makes up a Function?

- **Calling the Function Should be Readable:** The order and number of arguments matter for readability.
  
- **Working with the Function Should be Readable:** The length of the function body matters for readability.

- **Keep the Number of Parameters Low:**
  - Exercise caution with functions of 2 parameters.
  - Avoid having functions with 3 or more parameters.

## Refactoring Function Parameters

- **0 Parameter Functions:** Refactor functions without parameters for improved readability. [Example](https://github.com/academind/clean-code-course-code/blob/functions-01-refactoring-function-parameters/no-arg-functions.js).

From lines 1-17 show how to save a user, but to refactor this better and become more readable

From lines 21 to 34 we use a function without parameters to make it more easier

## When to Use Parameters

- **1 Parameter Functions:** Some functions only need one argument, and that's acceptable. [Examples](https://github.com/academind/clean-code-course-code/tree/functions-02-when-one-parameter-is-just-right).

- **2 Parameter Functions:** Having 2 parameters is acceptable if it makes the function easy to read and understand. Try to minimize parameters, but it's fine if the function remains easy, readable, and clean. [Examples](https://github.com/academind/clean-code-course-code/tree/functions-03-two-parameters).

## Refactoring Multiple Parameters

- You can replace multiple parameters with a single parameter, such as passing in a mapping or another data type that contains the necessary data.

## Dynamic Number of Parameters

- It's acceptable to have functions with a dynamic number of parameters, but exercise caution if there are a lot of arguments. Despite this, such functions can still be readable and clean to look at. [Example](https://github.com/academind/clean-code-course-code/blob/functions-05-dynamic-parameters/infinite-arg-functions.js).


---

## Output Parameters: Try to Avoid Them

Output parameters are parameters passed into a function or method call that are modified during the function. **It's recommended to avoid them.**

Functions should either return a value or modify "this". No other mutations should be allowed. In true functional programming, only returning values is preferred. 

For more information, refer to this [article](https://medium.com/thinkster-io/code-smell-output-parameters-fcb90e0005aa) and this [code example](https://github.com/academind/clean-code-course-code/blob/functions-06-output-parameters/output-parameters.js).

## Functions Should Be Small and Do One Thing

- **"One Thing" Can Be Different Operations Done for One Thing:** Functions should be small and focused on a single responsibility, even if that responsibility involves multiple operations.

- **Levels of Abstraction:** Functions should operate at a consistent level of abstraction. Higher-level operations can be encapsulated in methods that are more readable and understandable, while lower-level operations can dive into specific API methods or operations.

- **Functions Should Do Work That's One Level of Abstraction Below Their Name:** For example:
  ```javascript
  function emailIsValid(email) {
      return email.includes('@');
  }
  ```
  The function `emailIsValid` operates at the same level of abstraction as its name suggests, providing a clear and concise indication of its purpose.

> [!TIP]
> **Avoid Mixing Levels of Abstraction:** Keep functions focused and avoid mixing different levels of abstraction within a single function to maintain clarity and readability.


---

# When to Split Long Functions

## Rule of Thumb:
1. **Extract Code That Works on the Same Functionality:** Split long functions by extracting code that performs a specific functionality.
2. **Extract Code That Requires More Interpretation Than the Surrounding Code:** Split long functions by extracting code that requires more interpretation or explanation than the surrounding code.

Bad Long Code: https://github.com/academind/clean-code-course-code/blob/functions-07-small-do-one-thing/bad-function.js

Better Looking Code: https://github.com/academind/clean-code-course-code/blob/functions-07-small-do-one-thing/good-function.js

## Staying DRY (Don't Repeat Yourself)

- **Don't Repeat Yourself (DRY):** Avoid writing the same code more than once. [Example](https://github.com/academind/clean-code-course-code/blob/functions-09-dry/dry.js#L10-L28).
  
[Article to Read](https://scientificallysound.org/2018/07/19/python-functions/)

- **Use Common Sense:** While following the DRY principle, always use common sense to determine the appropriate level of granularity. Being overly granular may not always improve readability.

### Situations to Avoid When Being Too Granular:
1. Renaming function.
2. Finding a new function will take longer.
3. Hard to find a new name for it.

## Try Keeping Functions Pure

- **Pure Functions:** Functions should ideally be pure, meaning they produce the same outputs for the same inputs and have no side effects. Pure functions are predictable, making them suitable for functional programming paradigms.

> [!NOTE]
> It's not bad to have impure functions, but impure ones are not predictable. Pure ones don't have side effects

### What are Side Effects?

- **Definition:** A side effect is an operation that changes the overall system/program state.

- **Example:**
  ```javascript
  createUser(email, password) {
    const user = new User();
    startSession(user);
    return user;
  }
  ```
  `startSession` is a side effect as it may change something in the `user` object.

  Another example: https://github.com/academind/clean-code-course-code/blob/functions-11-side-effects/pure-and-side-effects.js#L13-L17


> [!TIP]
> **Avoiding Side Effects:** Use meaningful names to indicate when side effects are expected to help avoid unexpected behavior.

So naming matters in avoiding side effects. Names should indicate when a side effect is expected.

> [!TIP]
> Choose function name that implies it or move it to another function/place


## Lastly, Unit Testing
It can help with writing good clean code.

One way to know if you need to split, can you easily test a function?
1. No -> consider splitting it
2. Yes - Good

[Summarization of notes](https://github.com/academind/clean-code-course-code/blob/functions-extra-attachments/Functions-Summary.pdf)