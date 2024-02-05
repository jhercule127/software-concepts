# Naming

## Why Good Names Matter?

Well-named things allow readers to understand your code without going through it in detail. You don't need to dive in to understand.

## Choosing Good Names

### Variables and Constants (Data Containers)

- Use nouns or short phrases with adjectives.
  
  Example: `const UserData`

### Functions/Methods (Commands or Calculated Values)

- Use verbs or short phrases with adjectives.
  
  Example: `SendData()`

### Classes (To Create Things)

- Use nouns or short phrases with nouns.

## Casing Conventions

- `snake_case` - Python
- `camelCase` - Java
- `PascalCase` - Python, Java
- `kebab-case` - HTML

## Variables, Constants, and Properties

- **Value is an Object:** Describe it.
- **Value is a Number or String:** Describe it.
- **Value is a Boolean:** Answer true/false. Use `is`
  
  Example: `isActive`

## Functions and Methods

- **Functions perform an operation:** Describe the operation.
  
  Example: `getUser`

- **Function computes a boolean:** Answer true/false.
  
  Example: `isValid()`


Certainly, here's the markdown file with the provided notes about classes, exceptions to be aware of, common errors, and pitfalls:

# Classes

## Describe the Object and Provide More Details Without Introducing Redundancy

Classes define blueprints for creating objects. They encapsulate data for the object's properties and define methods to operate on those properties. Provide descriptive names that convey the purpose and functionality of the class without redundancy.

## Exceptions to Be Aware Of

- Libraries may have methods and objects with predefined names that cannot be changed. Refer to specific library documentation to understand and adhere to naming conventions within that context. [Example](https://github.com/academind/clean-code-course-code/blob/naming-02-naming-exceptions/naming-exceptions.py).

## Common Errors and Pitfalls

- **Avoid Redundant Information in Names:** Refrain from including unnecessary or redundant details in variable, function, or class names. For instance, `userWithNameAndAge = User('Mac', 31)` is verbose; prefer `user = User('Mac', 31)`.

- **Avoid Unclear Abbreviations:** Steer clear of using ambiguous abbreviations, slang, or unclear language in names. Choose names that clearly convey their purpose and functionality.

- **Choose Distinctive Names:** When naming variables, functions, or classes, ensure that they are distinct and easily differentiable, especially if you have similar functions. Clear and distinct names enhance code readability and reduce the risk of confusion.

- **Be Consistent:** Maintain consistency in naming patterns across your codebase. Stick to a chosen naming convention for variables, functions, classes, etc., such as `getUsers()`, `getProducts()`, `getData()`.

## Examples

- Explore clean naming examples in Python code [here](https://github.com/academind/clean-code-course-code/blob/naming-03-demo/03-clean-name-examples.py).
- Additional examples and guidelines can be found in the [naming slides](https://github.com/academind/clean-code-course-code/blob/naming-extra-attachments/slides-naming.pdf).
