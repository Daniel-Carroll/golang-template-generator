Golang Template Generator

---
## Running the Application Locally

### Compile the Binary

```
go build ./cmd/templatr
```

### CLI Options

The options 

| Option      | Description |
| ----------- | ----------- |
| -templateName -t     | Name of the template you want to generate       |
| -output -o  | Directory where the generated template will be created.       |

### Generating a Template

After compiling the application generating a template is as simple as this:

```
./templatr -t="go" -o="golang-app"
```

In this case the template will be generated in a directory called "golang-app"