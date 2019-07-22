Implementation Approach
=======================

Memory
------
This sounds like an array[15] holding a bunch of array[16]s (possibly the array[15] should be a struct, not sure yet)
- Setter
  - Set an address
  - Set a specific byte of an address
  - Option to fail if address is not empty?
- Getter
  - Get an address
  - Get a specific byte of an address
  - Option to fail if address is not empty?


Registers
---------
Even more simple: same size, can only get and set
- Setter
  - Set a register, overwriting
- Getter
  - Get a register

Stack
-----
Probably an array
- Needs to resize itself
- Methods
  - Push
  - Pop
  - Peek?
  - Len?