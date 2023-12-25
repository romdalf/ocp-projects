# Benchmarking languages and framework

Working with customers and partners on application modernization projects sparks discussions about governance. This system of rules and practices should influence the choice of language and related framework. However, from a governance and a bias perspective, I highlight the risks of randomly selecting these without performing relevant test case benchmarking. 

Regarding bias, Application teams have their favorite language and framework in mind and will research proof that this was the best option. The Internet is resourceful in supporting opinionated views with a few details about the methodology used. 

Let's take this [article](https://medium.com/deno-the-complete-reference/quarkus-java-vs-gin-go-hello-world-performance-0a2ec6d92078) as a reference. While the overall content provides a sufficient overview for an introduction, it does not help to perform a similar benchmarking within our environment, which will most likely significantly differ from the author. Let's address this missing part as a companion to the article. 

## The article benchmark
From a setup perspective, the author defines the following:  

- The hardware is based on a MacBook Pro M2 with 16GB of RAM
- [Bombardier](https://github.com/codesenberg/bombardier); a HTTP(S) bencharmking tool written in Go(lang)
- Go version 1.21.3; follow [this guide](https://go.dev/doc/install) to install Go on your system
- Quarkus 3.5.1 with Java v21; follow [this guide](https://quarkus.io/get-started/) to install Quarkus on your system

These are a good start; however, the Go code calls from [Gin](https://github.com/gin-gonic/gin), a Go HTTP framework for which we don't have any version reference. The same goes for Bombardier.

### Code
The below code is the as-is output from the article. These are the traditional ```hello world``` examples to showcase languages and easily compare them. Even with years of development experience, anyone needed more to run the code from an unknown language and benchmark it.

The code is also available within this repository within ```docs/sources/hello-world-article``` in the following subdirectories:   

- ```hello-world-go``` the ready to use Go code
- ```hello-world-quarkus``` the ready to use Quarkus code


#### Go/Gin   

```go
--8<-- "sources/hello-world-article/hello-world-go/main.go"
```

Here the steps to run this code:

- Open two terminal consoles
- Create a directory like ```hello-world-go```
- Create a file called ```main.go``` and copy the above code in
- Initialize the module(s) for the project:  

```
go mod init hello-world-go
```
This will create a file called ```go.mod``` with the following content: 

```
--8<-- "sources/hello-world-article/hello-world-go/go.mod"
```

- Check for missing modules:

```
go mod tidy
```
This will create a file called ```go.sum``` with the following content:

```
--8<-- "sources/hello-world-article/hello-world-go/go.sum"
```

- In one of the console, run the code:

```
go run main.go
```

Resulting in the following output:

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (1 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :3000
```

- In the second console, use ```curl``` to check the service:

```
curl localhost:3000
```

Resulting in the following output: 
```
Hello world!% 
```

***Note that there is no output logged into the first console where the service is running.***

- The code seems to be working, we can build a binary:

```
go build -o hello-world main.go
```

Resulting in the creating a binary called ```hello-world``` with the following size:

```bash
-rwxr-xr-x   1 romdalf  staff   9.5M Dec 25 10:32 hello-world
```

- Finally, to run the binary:

```
./hello-world
```
Resulting in the following output:

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (1 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :3000
```




#### Quarkus   

```java
--8<-- "sources/hello-world-article/hello-world-quarkus/src/main/java/org/acme/HelloWorldApplication.java"
```

To install Quarkus CLI, you can follow the official [Get Started](https://quarkus.io/get-started/) procedure by executing the following command in your favorite CLI:

```
curl -Ls https://sh.jbang.dev | bash -s - trust add https://repo1.maven.org/maven2/io/quarkus/quarkus-cli/
curl -Ls https://sh.jbang.dev | bash -s - app install --fresh --force quarkus@quarkusio
```
***Note: Remember to always check these scripts before executing them!***

Create the hello-world project with 
```
quarkus create app org.acme:hello-world --extension='resteasy-reactive'
```

Output
```
Looking for the newly published extensions in registry.quarkus.io
-----------
selected extensions: 
- io.quarkus:quarkus-resteasy-reactive


applying codestarts...
📚 java
🔨 maven
📦 quarkus
📝 config-properties
🔧 tooling-dockerfiles
🔧 tooling-maven-wrapper
🚀 resteasy-reactive-codestart

-----------
[SUCCESS] ✅  quarkus project has been successfully generated in:
--> /Users/romdalf/dev/opensource/ocp-projects/docs/sources/hello-world-article/quarkus/hello-world
-----------
Navigate into this directory and get started: quarkus dev
```

edit the 