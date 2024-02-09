
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

---

# SOLID Principles

Write classes in a SOLID way. These concepts can help write cleaner code.

## S - Single Responsibility Principle

Classes should have a single responsibility - a class shouldn’t change for more than one reason.
- Restricting classes to one core responsibility leads to smaller classes, which tend to be easier to read.

## O - Open-Closed Principle

A class should be open for extension but closed for modification.
- Think Polymorphism.
- Extensibility ensures small classes (instead of growing classes) and can help prevent code duplication (DRY).

## L - Liskov Substitution Principle

Objects should be replaceable with instances of their subclasses without altering the behavior.
- Class `Bird` has `fly` method.
- Class `Hawk` extends `Bird`, has `fly` and `dive` methods.
- Class `Eagle` extends `Bird`, if need `Eagle` then you can use `fly` alone.

## I - Interface Segregation Principle

Many client-specific interfaces are better than one general-purpose interface.
- Interface `Database`.
- Class `SQL` implements `Database`.
- Class `InMemory` implements `Database`.
  - But `InMemory` doesn’t need `Database`, it's a general-purpose interface.
  - Instead, create a new interface `RemoteDatabase`.

## D - Dependency Inversion Principle

“You should depend upon abstractions, not concretions.”

### More on Dependency Inversion Principle (DIP)

Dependency Inversion Principle talks about the coupling between different classes or modules.
It focuses on the approach where the higher classes are not dependent on the lower classes instead depend upon the abstraction of the lower classes

- High-level classes should not depend on low-level classes. Both should depend on abstractions.
- Effectively, the DIP reduces coupling between different pieces of code.
- [More information](https://www.geeksforgeeks.org/dependecy-inversion-principle-solid/#)
