    Maps:
    
        Maps are a key-value data structure, similar to dictionaries in other languages.
        The keys in a map must be a comparable type, while values can be of any type.
        To lookup a value in a map, you can use the syntax value, ok := mapType[key]. This will return the value associated with the key and a boolean indicating whether the key was found or not.
        Maps are passed by value, but the underlying data structure is a pointer  (to a runtime.hmap structure). Therefore, changes made to the map within a function or method will affect the original map.
        It is important to note that a nil map behaves like an empty map when reading, but writing to a nil map will cause a runtime panic. It is recommended to use make to initialise an empty map. Bad: var m map[string]string. Good:  var dictionary = map[string]string{} OR var dictionary = make(map[string]string)
        To add or update a value in a map, you can use the syntax map[key] = value. If the key already exists, the value will be overwritten, otherwise a new key-value pair will be added.
        To delete a key-value pair from a map, you can use the built-in function delete(dict, key). This will remove the key-value pair associated with the specified key.
    
    Errors:
    
        In Go, errors are values that implement the error interface. They are used to indicate failure when calling a function or method.
        Constants can be used to create errors that are reusable and immutable. To achieve this requires making an error type, eg DictionaryErr, which implements an error interface, eg Error() which returns the error value as a string.
        It is recommended to use error wrappers to provide additional context to the errors.
    
    Wrapping up:
    
        In this tutorial, we learned how to create, search, add, update, and delete values in a map.
        We also learned about the importance of error handling and how to create descriptive errors.
        Additionally, we discussed the concept of creating new types in Go.