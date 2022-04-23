# fictional-engine

    Let build Task manager with grpc; GRPC support http2.0 and microservices. Another advantage grpc support many languages 


[Golang]

        [Export Path]

          $  export GOROOT=/usr/local/go
          $  export GOPATH=$HOME/go
          $  export GOBIN=$GOPATH/bin
          $  export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
        
        [Protoc command]

            protoc --go_out=paths=source_relative:. proto_file_path.proto

<!-- [Javascript]

        [Protoc command]

            protoc --js_out=import_style=commonjs,binary:. proto_file_path.proto -->

[Route]
   
    /          (get, post)        add task 
    /delete    (post)             delete task


[Run]

    [Docker] 
    
        [Redis]

            Open new terminal 
            $ docker run --rm  --name redis-test-instance -p 6379:6379  redis
        
        [Fictional-Engine Client]

           Fictional-Engine-client execute grpc client which provide add & delete task route
            Open new terminal
            $ docker pull ali2210/fictional-engine:latest
            $ docker run --rm -p 3000:3000 --net=host -it ali2210/fictional-engine