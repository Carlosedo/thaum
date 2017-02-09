# Thaum 🔮
Thaum is a boilerplate and scaffolding command line utility. It purposely requires little to no configuration and does not require you to have anything more than a [mustache](http://mustache.github.io/) template to make a boilerplate. (Read: No script.)

![screencapture](http://evanconrad.com/thaum.gif)

# Features
- No configuration.
- Global templates and git included project templates.
- Compiles mustache in paths for files like `{{name}}.spec.js`.

# Install

If you have [go](https://golang.org/), you can run:

```
$ go get github.com/flaque/thaum
```

or, you can run this long (but effective!) command: 

```
$ wget "https://github.com/Flaque/thaum/releases/download/v0.6.0-beta/thaum" -O "/usr/local/bin/thaum" && sudo chmod +x /usr/local/bin/thaum
```

to install the binary into your path.

**Note**: It is much easier to run that long and scary command than to install go. Seriously, if you don't have go, just copy paste that bad boy into your terminal. 

# Usage

Your templates go in a folder called `thaum_files`. Thaum will look up the file
tree for the nearest `thaum_files` folder and use that one for your command.

## Creating a Template
Create your thaum_files in your root project.

```
$ mkdir thaum_files
```

Then create your first template! This is just a folder in your `thaum_files`.

```
$ mkdir thaum_files/myTemplate
```

Then, we can create a file or a whole folder system if you want inside.

```
$ touch thaum_files/myTemplate/myWidget.js
```

In that file, you can put something like this:

```
import {{package}};

export class {{name}} {
  constructor(foo, bar) {
    //do something
  }
}
```

## Running Thaum

Once you have a template, you can run `thaum` like so:

```
$ thaum myTemplate
```

and thaum will ask you to fill in the details:

```
🔍  Using thaum_files at: "/Users/Flaque/thaum-test/thaum_files"

     package: foo
        name: bang

✍️  Created file: myWidget.js
```

You can also list all templates that exist by typing `thaum`.

```
$ thaum
Templates Available:
  component
  growler
  test
```

If you need help, you can type `thaum -h` to see the help screen.

```
$ thaum -h
```

## Other helpful things.

If you would not like your templates to named `.somethingOrAnother` so that test runners or other tools don't pick them up, you can use the extension `.thaum` at the end of your template and thaum will remove it for you when the file is created.

For example, if I have a template file with the name:
```
foo_test.go.thaum
```

and I run `thaum foo`, thaum will create the file:
```
foo_test.go
```
