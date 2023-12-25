# Benchmarking languages and framework

Working with customers and partners on application modernization projects sparks discussions about governance. This system of rules and practices should influence the choice of language and related framework. However, from a governance and a bias perspective, I highlight the risks of randomly selecting these without performing relevant test case benchmarking. 

Regarding bias, Application teams have their favorite language and framework in mind and will research proof that this was the best option. The Internet is resourceful in supporting opinionated views with a few details about the methodology used. 

Let's take this [article](https://medium.com/deno-the-complete-reference/quarkus-java-vs-gin-go-hello-world-performance-0a2ec6d92078) as a reference. While the overall content provides a sufficient overview for an introduction, it does not help to perform a similar benchmarking within our environment, which will most likely significantly differ from the author. Let's address this missing part as a companion to the article. 

## How to redo as-is the article benchmark
From a setup perspective, the author defines the following:
- The hardware is based on a MacBook Pro M2 with 16GB of RAM
- [Bombardier](https://github.com/codesenberg/bombardier); a HTTP(S) bencharmking tool written in Go(lang)
- Go version 1.21.3
- Quarkus 3.5.1 with Java v21

These are a good start; however, the Go code calls from [Gin](https://github.com/gin-gonic/gin), a Go HTTP framework for which we don't have any version reference. The same goes for Bombardier.

### Code

<div class="grid" markdown>

```yaml title="Go"
--8<-- "sources/hellow-world-article/hello-world-go/main.go" 
```


```java title="Quarkus"
--8<-- "sources/hello-world-article/hello-world-quarkus/src/main/java/org/acme/HelloWorldApplication.java" 
```

<div class="grid" markdown>




```
go mod init github.com/romdalf/ocp-projects/docs/sources/hello-world-article/go
go mod tidy
go run main.go
```

# Quarkus 
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