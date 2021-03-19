# Checked GO

This library provides various checked number operations.

At this moment only checked type conversion is supported.

### Examples

```go
... := must.Int32ToUint8(200) // OK
... := must.Int32ToInt8(200) // panics
```

### Future plans

- Add conversion functions that return errors instead of panicking.
- Support other checked operations like add, sub, mul, etc.
