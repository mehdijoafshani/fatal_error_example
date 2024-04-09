## Fatala Error Example In Go
By running the main function, you will get the following error:
```
unexpected fault address [SOME_ADDRESS]
fatal error: fault
[signal SIGSEGV: segmentation violation code=0x2 addr=[SOME_ADDRESS] pc=[SOME_ADDRESS]]
```
Hyposesis: The garbage collector will remove the allocated space, while a local variable has a reference to it
