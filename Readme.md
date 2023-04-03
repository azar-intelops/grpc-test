# gRPC-Interceptors-using-go

#### prerequisites
- Go ([https://go.dev/doc/install](https://go.dev/doc/install))
- protoc ([https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/))
- Evans CLI ([https://github.com/ktr0731/evans#from-github-releases](https://github.com/ktr0731/evans#from-github-releases))
- code editor [I prefer [Visual Studio Code](https://code.visualstudio.com/)]

Demo for gRPC Interceptors, I have created logging middleware using interceptors using go.
- To run this code, please clone this repository using below commands.
    ```bash
    git clone https://github.com/intelops/go-interceptors-demo.git
    ```
- Navigate to ```go-interceptors-demo```
    ```bash
    cd go-interceptors-demo
    ```
- Update the go.mod file with each dependency that is required by your source files by downloading them all.
    ```bash
    go mod tidy
    ```
- Check ```Makefile``` for the common commands which I have used frequently.
    ```bash
    # to install grpc
    make install 
    
    # if path was not set, use
    make path
    
    # if proto gets updated, first remove pb folder using
    make clean
    
    # and run this command to recompile the proto file
    make protoc
    ```
- Run the server using
    ```bash
    go run main.go
    
    # if you see this in terminal, your server is working fine, else please check if anything is missing from above steps.
    
    #    Starting server in port :50051
    ```
- Open a new terminal and run the evans cli which works like a client:
    ```bash
    evans -r repl -p 50051
    
    # This command helps to choose the package 
    package demo
    
    # To check the services defined 
    show services
    
    :'
    # Output generated
    +-----------+------------+--------------+---------------+
    |  SERVICE  |    RPC     | REQUEST TYPE | RESPONSE TYPE |
    +-----------+------------+--------------+---------------+
    | MyService | DemoMethod | DemoRequest  | DemoResponse  |
    +-----------+------------+--------------+---------------+
    '
    # To call this method use this command
    call DemoMethod
    
    :'
    # Output generated
    # Example 1
    demo.MyService@127.0.0.1:50051> call DemoMethod
    message (TYPE_STRING) => Dr.Strange
    {
      "message": "Hello Dr.Strange"
    }
    # Example 2
    demo.MyService@127.0.0.1:50051> call DemoMethod
    message (TYPE_STRING) => Compage
    {
      "message": "Hello Compage"
    }
    
    #logs generated from server
    2023/03/31 23:37:40 Received request: message:"Dr.Strange"
    2023/03/31 23:39:16 Received request: message:"Compage"
    '
    # To stop CLI use this command
    exit
    ```