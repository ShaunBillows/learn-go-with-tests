	Pointers:
	
		In Go, when a function or method is called, arguments are copied.
		If you want to manipulate a value, but not expose it outside a package, you can use lowercase symbols to make it private.
		Pointers can be used to point to a value's memory address, so that methods and tests can refer to the same value in memory.
		Struct pointers do not need to be dereferenced when used with struct methods, as Go does this automatically.
	
	Methods:
	
        Go allows for methods to be created on type definitions as well as on structs.
        Implementing a Stringer interface and string method on a type, such as Bitcoin, can define how it is printed using %s.
	
	Errors:
	
        Errors are created as values with errors.new().
        To identify unchecked errors, the 'errcheck' command can be used.
        Global variables can be created for error messages to reuse in tests.
        Errors are the way to indicate failure when calling a function or method.
	
	Nil:
	
        Pointers can be nil.
        When a function returns a pointer, it is important to check if it is nil to avoid runtime exceptions.
	
	Creating new types:
	
        New types can be created from existing ones to add domain-specific meaning to values.
        This can also allow for the implementation of interfaces.
	
	Wrapping up:
	
        Pointers and errors are important concepts in writing Go code.
        The compiler will often help in identifying errors.
        Creating meaningful error messages can result in easier-to-test code.
        More sophisticated error handling strategies can be used.
