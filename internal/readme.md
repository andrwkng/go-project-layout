# internal/
- The ```internal/``` directory gives an extra level of compiler protection. 
- The compiler will not let any other project import packages from an internal package. 
- Only packages from this project can import from internal

## Logging
- packages inside of internal can log

