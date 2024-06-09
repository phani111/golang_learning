correctUsageWithoutTempVar 

Key Points:
- Variable Declaration:
  - `var client *http.Client` declares a client variable at the beginning of the function.
  - `var err error` declares an err variable at the beginning of the function.
- Avoiding Shadowing:
  - Inside the if and else blocks, the line `client, err = ...` assigns values to the already declared client and err variables.
  - This avoids shadowing because it does not use the `:=` syntax to declare new variables.
- Logging:
  - The `log.Println(client)` inside the if and else blocks logs the client variable.
  - The `log.Println(client)` outside the blocks logs the client variable again, which now holds the value assigned within the blocks.
- Error Handling:
  - The `if err != nil` check is placed after the if-else block to handle any errors that might have occurred during the client creation.
  - This ensures that error handling is centralized and consistent.


### Conclusion

The shadowing function demonstrates the problem of variable shadowing, where the outer client variable remains nil.

The correctUsage function avoids shadowing by using a temporary variable c and then assigning it to the outer client variable.

The correctUsageWithoutTempVar function avoids shadowing by directly assigning the new value to the outer client variable and centralizing error handling, making the code cleaner and more maintainable.
