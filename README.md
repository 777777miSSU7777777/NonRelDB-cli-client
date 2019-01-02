## NonRelDB-cli-client
CLI client for NonRelDB.

## NonRelDB
  
NonRelDB is an in-memory database that persists on disk. The data model is key-value. Written on pure golang.

## Installation
First of all you need to install [git](https://git-scm.com/).
Then you need to clone repository on your pc from github

    git clone https://github.com/777777miSSU7777777/NonRelDB-cli-client.git
    
In cloned repository you will find **Makefile** with following targets:

 - **clean** - removes client executable binary file.
 - **build** - builds client's executable binary file.
 - **run** - runs client with default configuration.
 
## Usage
### Flags
 - **-host -h** - defines host ip (default is 127.0.0.1)
 - **-port -p** - defines host port (default is 9090)
 - **--dump** - requests full database dump in json format on stdout.  
 Usage example
 
    ./client --dump > dump.json

 - **--restore** - restores database from stdin in json format.  
 Usage example 
 
    ./client --restore < dump.json
    
 ### Commands
Commands can be entered only in one register (**GET** and **get** but not **Get**).

 **List of supported commands**
 - **GET** - returns the value if existing, otherwise message "Value with this key not found".  
 Example
 

    GET 123
    

 - **SET** - set the value if existing, otherwise creates new. Also returns message "Value has changed".  
 **Value must be in double quotes.*  
 Example
 

    SET 123 "123"
    

 - **DEL** - deletes value from storage and returns it's value if existing, otherwise message "Value with this key not found".  
 Example
 

    DEL 123

- **KEYS** - returns all keys matching to entered regexp pattern, otherwise message "Keys with this pattern not found" or "Pattern is incorrect".  
**Regex pattern must be in double quotes.*  
Example

    KEYS "/*"

- **SUBSCRIBE** - subscribes the client on specified channel.  
Example 

    SUBSCRIBE redis

- **UNSUBSCRIBE** - unsubscribes the client from specified channel.  
**Cannot use from cli client directly.*  
Example

    UNSUBSCRIBE redis

- **PUBLISH** - sends the message to specified channel.  
*\*Message must be in double quotes.*   
Example

    PUBLISH redis "Hello world"
    
 ## Notice
 - Client will crash with non-existing flag's values.
 - Not implented -help flag.
