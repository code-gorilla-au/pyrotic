# pyrotic
code generator inspired by https://plopjs.com/ & https://github.com/jondot/hygen for golang.


## Motivation
Why not use hygen? great question! I would recommend [plop](https://plopjs.com/) over this, however [plop](https://plopjs.com/) is written in js.
This project is for people who want to use a code generator and not have to install node. Pyrotic also is specifically written to generate go code, all templates are parsed using go's builtin template parser and output is formatted using go's built in code formatter.



## Install

```
go install github.com/code-gorilla-au/pyrotic@latest

```

initial setup creates a `_templates` directory at the root of the project to hold the generators

```
pyrotic init
```

create your first generator

```
pyrotic new cmd
```

## Run

default template path is `_templates` and default file extension is `.tmpl`

```
pyrotic generate <name of generator> --name <name-to-pass-in>

eg: pyrotic generate cmd --name setup
```

### Use different directory

```
pyrotic --path example/_templates generate cmd --name setup
pyrotic -p example/_templates generate cmd --name setup
```

### Use different file extension

default file extension is `.tmpl`

```
pyrotic --extension ".template" generate cmd --name setup
pyrotic -x ".template" generate cmd --name setup
```

### Dry run mode

dry run will log to console rather than write to file

```
pyrotic -d generate cmd --name setup
pyrotic --dry-run generate cmd --name setup
```

### Different shared folder

default shared templates path is `_templates/shared`

```
pyrotic --shared foo/bar generate cmd --name setup
pyrotic -s foo/bar generate cmd --name setup
```

## Formatter properties

Formatter will pick up any of these variables within the `---` block and hydrate the metadata for the template. Any properties matching the signature will be added to the Meta property, for example `foo: bar` will be accessible by `{{ Meta.foo }}`. View more [examples](example/_templates).

| Property | Type          | Default | Example                 |
|----------|---------------|---------|-------------------------|
| to:      | string (path) | ""      | src/lib/utils/readme.md |
| append:  | bool          | false   | false                   |
| inject:  | bool          | false   | false                   |
| before:  | string        | ""      | type config struct      |
| after:   | string        | ""      | // commands             |


### Variables within formater properties

Tool's two-stage parser hydrates the data within the `---` blocks before parsing the template.

| Property | Type          | Default | Example                                      |
|----------|---------------|---------|----------------------------------------------|
| to:      | string (path) | ""      | src/lib/{{ .Name }}/{{ Meta.readmeName }}.md |
| append:  | bool          | false   | false                                        |
| inject:  | bool          | false   | false                                        |
| before:  | string        | ""      | type config struct                           |
| after:   | string        | ""      | // commands                                  |


### Using shared templates

In some instances you will want to reuse some templates across multiple generators. This can be done by having a `shared` directory within the `_templates` directory.
Any templates that are declared in the [shared](example/_templates/shared/config.tmpl) directory will be loading along with the generator. [Reference](example/_templates/fakr/shared_config.tmpl) the shared template within your generator directory to inject / append / create file.


## Built in template functions

ships with some already built in template funcs, some [examples](example/_templates/fakr/fakr_case.tmpl)

| func name           | description                         | code example                                | result                        |
|---------------------|-------------------------------------|---------------------------------------------|-------------------------------| 
| caseSnake           | convert to snake case               | {{ MetaData \| caseSnake }}                 | meta_data                     |
| caseKebab           | convert to kebab case               | {{ MetaData \| caseKebab }}                 | meta-data                     |
| casePascal          | convert to pascal case              | {{ meta_data \| casePascal }}               | MetaData                      |
| caseLower           | convert to lower case               | {{ MetaData \| caseLower }}                 | metadata                      |
| caseTitle           | convert to title case               | {{ MetaData \| caseTitle }}                 | METADATA                      |
| caseCamel           | convert to camel case               | {{ MetaData \| caseCamel }}                 | metaData                      |
| splitByDelimiter    | splits string by delimiter          | {{ splitByDelimiter "long,list" "," }}      | []string{"long" "list"}       |
| splitAfterDelimiter | splits string after delimiter       | {{ splitAfterDelimiter "a,long,list" "," }} | []string{"a," "long," "list"} |
| contains            | checks if string contains substring | {{ contains "foobarbin" "bar" }}            | true                          |
| hasPrefix           | checks if string has the prefix     | {{ contains "foobarbin" "foo" }}            | true                          |
| hasSuffix           | checks if string has the suffix     | {{ contains "foobarbin" "bin" }}            | true                          |


we also provide some Inflections using [flect](https://github.com/gobuffalo/flect)

- pluralise
- singularise
- ordinalize
- titleize
- humanize

## Pass in meta via cmd

you can pass in meta data via the `--meta` or `-m` flag, which takes in a comma (`,`) delimited list, and overrides the `{{ .Meta.<your-property> }}` within the template.

```

pyrotic generate fakr --meta foo=bar,bin=baz
pyrotic generate fakr -m foo=bar,bin=baz

```


## Dev mode
provides the short file name with logging

```bash

ENV=DEV ./pyrotic -p example/_templates generate fakr --meta foo=bar,bin=baz

```