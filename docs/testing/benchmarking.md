# Properly benchmarking languages and framework






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