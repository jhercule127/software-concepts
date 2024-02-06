# Functions and Methods

## What makes up a Function?

- **Calling the Function Should be Readable:** The order and number of arguments matter for readability.
  
- **Working with the Function Should be Readable:** The length of the function body matters for readability.

- **Keep the Number of Parameters Low:**
  - Exercise caution with functions of 2 parameters.
  - Avoid having functions with 3 or more parameters.

## Refactoring Function Parameters

- **0 Parameter Functions:** Refactor functions without parameters for improved readability. [Example](https://github.com/academind/clean-code-course-code/blob/functions-01-refactoring-function-parameters/no-arg-functions.js).

## When to Use Parameters

- **1 Parameter Functions:** Some functions only need one argument, and that's acceptable. [Examples](https://github.com/academind/clean-code-course-code/tree/functions-02-when-one-parameter-is-just-right).

- **2 Parameter Functions:** Having 2 parameters is acceptable if it makes the function easy to read and understand. Try to minimize parameters, but it's fine if the function remains easy, readable, and clean. [Examples](https://github.com/academind/clean-code-course-code/tree/functions-03-two-parameters).

## Refactoring Multiple Parameters

- You can replace multiple parameters with a single parameter, such as passing in a mapping or another data type that contains the necessary data.

## Dynamic Number of Parameters

- It's acceptable to have functions with a dynamic number of parameters, but exercise caution if there are a lot of arguments. Despite this, such functions can still be readable and clean to look at. [Example](https://github.com/academind/clean-code-course-code/blob/functions-05-dynamic-parameters/infinite-arg-functions.js).
