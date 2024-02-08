
# Section 6: Objects, Classes, and Data Containers/Structures

Patterns and Principles are about writing code which is maintainable and extensible. 
Clean code is about readability and easy to understand.

## Object vs. Data Container/Data Structure

| Object                                      | Data Container/Data Structure      |
|---------------------------------------------|------------------------------------|
| Private internals/properties                | Public internals, properties (almost no API methods) |
| Contain your business logic                 | Store and transport data           |
| Abstractions over concretions               | Concretions only                  |
| Example: Database to connect to             | User credentials in the database  |

### Why Differing These Matters?

There are times when a data container is just needed, but an object is needed where you want to hide the logic away.

## Polymorphism

Polymorphism is the ability of an object to take on many forms.

- [Example with polymorphism implemented](https://github.com/academind/clean-code-course-code/blob/obj-02-classes-polymorphism/polymorphism-start.ts)
- [With polymorphism implemented](https://github.com/academind/clean-code-course-code/blob/obj-02-classes-polymorphism/polymorphism.ts)

## How to Write Clean Classes:

- Classes should be small
- Classes should have a single responsibility: Single-Responsibility Principle

---

## Cohesion

Cohesion refers to how much your class methods are using the class properties.

| Max Cohesion            | No cohesion                                |
|-------------------------|--------------------------------------------|
| All methods each use all properties | All methods don’t use any class properties |
| Highly cohesive object  | Data structure/container with utility methods |

We want highly cohesive classes.

## “Law of Demeter” and Why You Should “Tell Not Ask”

The Law of Demeter is the **Principle of Least Knowledge**; don’t depend on the internals of “strangers” (other objects that you don’t know directly). This means no method chaining.

Basically, it recommends that objects should avoid accessing the internal data and methods of other objects. Instead, an object should only interact with its immediate dependencies.

It has 2 requirements: 
1. A single method can only operate on objects that are passed as arguments to the method.
2. Values of fields defined in this class. Globally values or objects created within methods can also be accessed.

For example, `this.customer.lastPurchase.data` is discouraged.

More information: 
- [The Law of Demeter by Example](https://medium.com/vattenfall-tech/the-law-of-demeter-by-example-fd7adbf0c324)
- https://www.baeldung.com/java-demeter-law


## Tell, Don’t Ask

Tell your code what to do, not ask.
