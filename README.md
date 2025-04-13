# Excel table in text

This project allows you to take data from an xlsx table and insert it into the template text instead of $NUM characters.

## How it works

1. You specify the path to the table (Excel file)
2. You specify the path to the template text (file with extensions .txt, .sql and others)
3. The project pulls data from the table and replaces the $NUM characters in the template text with the corresponding values.
4. $NUM corresponds to the order of the columns in the table. (First column is $1, fifth column is $5)

## Example of use

* Table:

| Name | Age | City |
| --- | --- | --- |
| Ivanov | 25 | Moscow |
| Petrov | 30 | Saint Petersburg |

* Text template:

"Hello, my name is $1, and I am $2 . I live in $3."

* Result:

"Hello, my name is Ivanov and I am 25 years old. I live in Moscow."

"Hello, my name is Petrov and I am 30 years old. I live in Saint Petersburg."

## Installation

1. Clone the repositories:

``` bash
git clone https://github.com/blacknoise228/excel-to-text-replacer.git
```

2. Go to the project directory:

``` bash
cd ./excel-to-text-replacer
```

3. Create a file template and a file to save the result and specify the path to these files in ./config.yaml.

4. Load as dependencies:

``` bash
go mod, carefully
```

5. Run:

``` bash
go run main.go
```
