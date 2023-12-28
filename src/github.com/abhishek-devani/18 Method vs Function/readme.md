# 18 Method vs Function

### Key Differences:

- **Receiver Parameter:**
    - **Function:** Doesn't have a receiver parameter.
    - **Method:** Has a receiver parameter, indicating the type the method belongs to.

- **Invocation:**
    - **Function:** Called by the function name directly.
    - **Method:** Called on an instance of the type with the syntax instance.method().

- **Use Case:**
    - **Function:** Used for general-purpose tasks.
    - **Method:** Used when the behavior is closely associated with a particular type.
 
- **Visibility:**
    - **Function:** Can be declared at the package level, making them accessible from other files in the same package.
    - **Method:** Associated with a type, and their visibility depends on the visibility of the type.

### Notes

- We can write multiple method with the same name but different receiver.
- We can't write multiple function with the same but different arguments.

