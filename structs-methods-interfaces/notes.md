    Formatting floating point values:
        .2f is used to format to 2 decimal places
        %g returns a higher precision value than %f
        %#v returns the value of a type (e.g. structs_methods_interfaces.Triangle{Base:12, Height:6})

    Methods:
        A function with a receiver
        Allows adding functionality to data types and implementing interfaces

    Interfaces:
        Allows creating functions that can be used with different types
        Provides decoupling for highly-decoupled code while maintaining type-safety
        Can be declared to define functions for different types (parametric polymorphism)

    Table driven tests:
        Builds a list of test cases that can be tested in the same manner
        Makes assertions clearer and test suites easier to extend and maintain

    Anonymous structs:
        Syntax for declaring a slice of structs with two fields: the shape and the expected result

    Importance of defining types:
        Essential for building easy-to-understand, piece together, and test software

    Interfaces for hiding complexity:
        Interfaces are a great tool for hiding complexity from other parts of the system, allowing test helper code to assert without knowing the exact shape